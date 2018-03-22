package gate

import (
	"server/game"
	"server/login"
	"server/msg"
	"server/msg/clientmsg"
)

func init() {

	//login server
	msg.Processor.SetRouter(&clientmsg.Req_Register{}, login.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Req_ServerList{}, login.ChanRPC)

	//game server
	msg.Processor.SetRouter(&clientmsg.Ping{}, game.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Req_ServerTime{}, game.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Req_Login{}, game.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Req_Match{}, game.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Transfer_Team_Operate{}, game.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Req_SetCharName{}, game.ChanRPC)

	//battle server
	msg.Processor.SetRouter(&clientmsg.Req_ConnectBS{}, game.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Req_EndBattle{}, game.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Transfer_Command{}, game.ChanRPC)
	msg.Processor.SetRouter(&clientmsg.Transfer_Loading_Progress{}, game.ChanRPC)
}
