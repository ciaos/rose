package g

import (
	"fmt"
	"server/gamedata"
	"server/gamedata/cfg"
	"server/msg/proxymsg"
	"strings"
	//	"sync"
	"time"

	"github.com/ciaos/leaf/log"
)

const (
	MATCH_OK        = 1
	MATCH_CONTINUE  = 2
	MATCH_TIMEOUT   = 3
	MATCH_ALLOCROOM = 4
	MATCH_EMPTY     = 5
	MATCH_FINISH    = 6
	MATCH_ERROR     = 7
)

type Seat struct {
	charid     uint32
	jointime   int64
	serverid   int32
	servertype string
}

//for match server
type Table struct {
	seats         []*Seat
	createtime    int64
	matchmode     int32
	status        int32
	tableid       int32
	modeplayercnt int32
}

var TableManager = make(map[int32]*Table)
var PlayerTableIDMap = make(map[uint32]int32)
var tableid int32

//var mTableID *sync.Mutex

func InitTableManager() {
	tableid = 0

	//	mTableID = new(sync.Mutex)
}

func allocTableID() {
	//mTableID.Lock()
	//defer mTableID.Unlock()
	tableid += 1
	if tableid > MAX_TABLE_COUNT {
		tableid = 1
	}
}

func (table *Table) update(now *time.Time) int32 {
	r := gamedata.CSVMatchMode.Index((*table).matchmode)
	if r == nil {
		log.Error("CSVMatchMode ModeID %v Not Found", (*table).matchmode)
		return MATCH_ERROR
	}
	row := r.(*cfg.MatchMode)

	if (*now).Unix()-(*table).createtime > int64(row.TimeOutSec) {
		log.Debug("Tableid %v Timeout Createtime %v Now %v", (*table).tableid, (*table).createtime, (*now).Unix())
		return MATCH_TIMEOUT
	}

	if len((*table).seats) >= row.PlayerCnt {
		return MATCH_OK
	}

	if len((*table).seats) <= 0 {
		log.Debug("tableid %v empty", (*table).tableid)
		return MATCH_EMPTY
	}
	return MATCH_CONTINUE
}

func allocBattleRoom(tableid int32) {

	innerReq := &proxymsg.Proxy_MS_BS_AllocBattleRoom{
		Matchroomid: tableid,
		Matchmode:   TableManager[tableid].matchmode,
		Membercnt:   TableManager[tableid].modeplayercnt,
	}

	//todo 固定路由到指定的BattleServer
	go RandSendMessageTo("battleserver", uint32(tableid), uint32(proxymsg.ProxyMessageType_PMT_MS_BS_ALLOCBATTLEROOM), innerReq)
}

func UpdateTableManager(now *time.Time) {

	for i, table := range TableManager {
		if (*table).status != MATCH_ALLOCROOM {
			(*table).status = (*table).update(now)
		}

		if (*table).status == MATCH_OK || (*table).status == MATCH_TIMEOUT {

			(*table).status = MATCH_ALLOCROOM
			allocBattleRoom(i)
		}
		if (*table).status == MATCH_EMPTY {
			DeleteTable(i)
		}
		if (*table).status == MATCH_ERROR {

			//notify all member error
			deleteTableSeatInfo(i)
			DeleteTable(i)
		}
	}
}

func DeleteTable(tableid int32) {
	log.Debug("DeleteTable TableID %v", tableid)
	delete(TableManager, tableid)
}

func deleteTableSeatInfo(tableid int32) {
	table, ok := TableManager[tableid]
	if ok {
		for _, seat := range table.seats {
			delete(PlayerTableIDMap, seat.charid)
		}
	}
}

func JoinTable(charid uint32, matchmode int32, serverid int32, servertype string) {

	var createnew = true
	for i, table := range TableManager {
		if len((*table).seats) < int((*table).modeplayercnt) {
			seat := &Seat{
				charid:     charid,
				jointime:   time.Now().Unix(),
				serverid:   serverid,
				servertype: servertype,
			}
			TableManager[i].seats = append(TableManager[i].seats, seat)
			PlayerTableIDMap[charid] = i

			log.Debug("JoinTable TableID %v CharID %v", i, charid)

			createnew = false
			break
		}
	}
	if createnew {
		allocTableID()

		r := gamedata.CSVMatchMode.Index(matchmode)
		if r == nil {
			log.Error("JoinTable CSVMatchMode Not Found %v ", matchmode)
			return
		}
		row := r.(*cfg.MatchMode)

		table := &Table{
			tableid:    tableid,
			createtime: time.Now().Unix(),
			matchmode:  matchmode,
			seats: []*Seat{
				&Seat{
					charid:     charid,
					jointime:   time.Now().Unix(),
					serverid:   serverid,
					servertype: servertype,
				},
			},
			status:        MATCH_CONTINUE,
			modeplayercnt: int32(row.PlayerCnt),
		}
		TableManager[tableid] = table
		PlayerTableIDMap[charid] = tableid

		log.Debug("JoinTable CreateTableID %v CharID %v", tableid, charid)
	}
}

func LeaveTable(charid uint32, matchmode int32) {
	tableid, ok := PlayerTableIDMap[charid]
	if ok {
		table, ok := TableManager[tableid]
		if ok {
			for i, seat := range table.seats {
				if (*seat).charid == charid {
					TableManager[tableid].seats = append(table.seats[0:i], table.seats[i+1:]...)

					log.Debug("LeaveTable TableID %v CharID %v RestCount %v", tableid, charid, len(table.seats))
				}
			}
		} else {
			log.Error("LeaveTable TableID %v Not Exist CharID %v", tableid, charid)
		}

		delete(PlayerTableIDMap, charid)
	} else {
		log.Error("LeaveTable CharID %v Not Exist", charid)
	}
}

func ClearTable(tableid int32, battleroomid int32, battleserverid int32, battleservername string) {
	table, ok := TableManager[tableid]
	if ok {
		msg := &proxymsg.Proxy_MS_GS_MatchResult{
			Retcode:          0,
			Battleroomid:     battleroomid,
			Battleserverid:   battleserverid,
			Battleservername: battleservername,
		}

		for _, seat := range table.seats {
			log.Debug("NotifyConnectBS CharID %v BSID %v RoomID %v", (*seat).charid, battleserverid, battleroomid)

			go SendMessageTo((*seat).serverid, (*seat).servertype, (*seat).charid, uint32(proxymsg.ProxyMessageType_PMT_MS_GS_MATCHRESULT), msg)
		}

		deleteTableSeatInfo(tableid)
		table.seats = append([]*Seat{}) //clear seats
		DeleteTable(tableid)
	} else {
		log.Error("ClearTable TableID %v Not Found , TableCount %v", tableid, len(TableManager))
	}
}

func FormatTableInfo(tableid int32) string {
	table, ok := TableManager[tableid]
	if ok {
		return fmt.Sprintf("TableID:%v\tCTime:%v\tStatus:%v\tSeatCnt:%v", (*table).tableid, (*table).createtime, (*table).status, len((*table).seats))
	}
	return ""
}

func FormatSeatInfo(tableid int32) string {
	output := fmt.Sprintf("TableID:%v", tableid)
	table, ok := TableManager[tableid]
	if ok {
		for _, seat := range (*table).seats {
			output = strings.Join([]string{output, fmt.Sprintf("CharID:%v\tJoinTime:%v\tServerID:%v\tServerType:%v", (*seat).charid, (*seat).jointime, (*seat).serverid, (*seat).servertype)}, "\r\n")
		}
	}
	return output
}
