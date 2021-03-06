package main

import (
	"encoding/binary"
	"errors"
	"net"
	"server/msg/clientmsg"

	"github.com/golang/protobuf/proto"
)

const (
	LoginServerAddr = "127.0.0.1:8000"
	GameServerID    = 2
	ClientNum       = 5000
	OneBattleTime   = 20
)

func Send(conn *net.Conn, msgid clientmsg.MessageType, msgdata interface{}) {
	if *conn == nil {
		return
	}

	data, _ := proto.Marshal(msgdata.(proto.Message))
	reqbuf := make([]byte, 4+len(data))
	binary.BigEndian.PutUint16(reqbuf[0:], uint16(len(data)+2))
	binary.BigEndian.PutUint16(reqbuf[2:], uint16(msgid))

	copy(reqbuf[4:], data)
	(*conn).Write(reqbuf)
}

func Recv(conn *net.Conn) (error, clientmsg.MessageType, []byte) {
	if *conn == nil {
		return errors.New("nil conn"), 0, nil
	}

	headdata := make([]byte, 2)
	(*conn).Read(headdata[0:])

	msglen := binary.BigEndian.Uint16(headdata[0:])

	bodydata := make([]byte, msglen)
	bodylen, _ := (*conn).Read(bodydata[0:])
	if msglen == 0 || bodylen == 0 {
		return errors.New("Invalid msglen"), 0, nil
	}
	msgid := clientmsg.MessageType(binary.BigEndian.Uint16(bodydata[0:]))

	return nil, msgid, bodydata[2:bodylen]
}
