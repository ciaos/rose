package conf

import (
	"log"
	"time"
)

var (
	// log conf
	LogFlag = log.LstdFlags | log.Lshortfile

	// gate conf
	PendingWriteNum        = 10000
	MaxMsgLen       uint32 = 32768
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 2
	LittleEndian           = false

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 20000
)
