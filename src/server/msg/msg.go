package msg

import (
	"server/msg/clientmsg"

	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&clientmsg.Ping{})
	Processor.Register(&clientmsg.Pong{})
	Processor.Register(&clientmsg.Req_ServerTime{})
	Processor.Register(&clientmsg.Rlt_ServerTime{})
	Processor.Register(&clientmsg.Req_Register{})
	Processor.Register(&clientmsg.Rlt_Register{})
	Processor.Register(&clientmsg.Req_ServerList{})
	Processor.Register(&clientmsg.Rlt_ServerList{})
	Processor.Register(&clientmsg.Req_Login{})
	Processor.Register(&clientmsg.Rlt_Login{})
	Processor.Register(&clientmsg.Req_Match{})
	Processor.Register(&clientmsg.Rlt_NotifyBattleAddress{})
	Processor.Register(&clientmsg.Req_ConnectBS{})
	Processor.Register(&clientmsg.Rlt_ConnectBS{})
	Processor.Register(&clientmsg.Rlt_StartBattle{})
}
