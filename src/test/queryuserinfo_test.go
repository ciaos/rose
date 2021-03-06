package test

import (
	"fmt"
	"math/rand"
	"net"
	"server/msg/clientmsg"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	. "gopkg.in/check.v1"
)

func TestQueryUserInfo(t *testing.T) { TestingT(t) }

type QueryUserInfoSuite struct {
	conn   net.Conn
	err    error
	charid uint32
}

var _ = Suite(&QueryUserInfoSuite{})

func (s *QueryUserInfoSuite) SetUpSuite(c *C) {

}

func (s *QueryUserInfoSuite) TearDownSuite(c *C) {

}

func (s *QueryUserInfoSuite) SetUpTest(c *C) {
	s.conn, s.err = net.Dial("tcp", LoginServerAddr)
	if s.err != nil {
		c.Fatal("Connect Server Error ", s.err)
	}

	rand.Seed(time.Now().UnixNano())
	username := fmt.Sprintf("pengjing%d", rand.Intn(10000))
	password := "123456"

	s.charid = QuickLogin(c, &s.conn, username, password)
}

func (s *QueryUserInfoSuite) TearDownTest(c *C) {
	s.conn.Close()
}

func (s *QueryUserInfoSuite) TestQueryUserInfo(c *C) {
	req := &clientmsg.Req_QueryCharInfo{}
	req.CharIDs = append(req.CharIDs, s.charid)

	msgdata := SendAndRecvUtil(c, &s.conn, clientmsg.MessageType_MT_REQ_QUERY_CHARINFO, req, clientmsg.MessageType_MT_RLT_QUERY_CHARINFO)
	rspMsg := &clientmsg.Rlt_QueryCharInfo{}
	err := proto.Unmarshal(msgdata, rspMsg)
	if err != nil {
		c.Fatal("Rlt_QueryCharInfo Decode Error ", err)
	}

	c.Assert(rspMsg.RetCode, Equals, clientmsg.Type_GameRetCode_GRC_OK)
	c.Assert(len(rspMsg.UserInfo), Equals, 1)
	c.Assert(rspMsg.UserInfo[0].CharID, Equals, s.charid)
}
