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

func TestPing(t *testing.T) { TestingT(t) }

type PingSuite struct {
	conn net.Conn
	err  error

	charid uint32
}

var _ = Suite(&PingSuite{})

func (s *PingSuite) SetUpSuite(c *C) {
}

func (s *PingSuite) TearDownSuite(c *C) {
}

func (s *PingSuite) SetUpTest(c *C) {
	s.conn, s.err = net.Dial("tcp", LoginServerAddr)
	if s.err != nil {
		c.Fatal("Connect Server Error ", s.err)
	}

	rand.Seed(time.Now().UnixNano())
	username := fmt.Sprintf("pengjing%d", rand.Intn(10000))
	password := "123456"

	s.charid = QuickLogin(c, &s.conn, username, password)
}

func (s *PingSuite) TearDownTest(c *C) {
	s.conn.Close()
}

func (s *PingSuite) TestPing(c *C) {

	rand.Seed(time.Now().UnixNano())
	reqMsg := &clientmsg.Ping{
		ID: uint32(rand.Intn(10000)),
	}

	msgdata := SendAndRecvUtil(c, &s.conn, clientmsg.MessageType_MT_PING, reqMsg, clientmsg.MessageType_MT_PONG)
	rspMsg := &clientmsg.Pong{}
	proto.Unmarshal(msgdata, rspMsg)
	c.Assert(rspMsg.ID, Equals, reqMsg.ID)
}
