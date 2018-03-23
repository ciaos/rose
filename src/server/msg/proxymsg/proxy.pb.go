// Code generated by protoc-gen-go.
// source: proxy.proto
// DO NOT EDIT!

/*
Package proxymsg is a generated protocol buffer package.

It is generated from these files:
	proxy.proto

It has these top-level messages:
	InternalMessage
	Proxy_GS_MS_Match
	Proxy_MS_BS_AllocBattleRoom
	Proxy_MS_GS_Begin_Battle
	Proxy_BS_MS_AllocBattleRoom
*/
package proxymsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProxyMessageType int32

const (
	ProxyMessageType_PMT_GS_MS_MATCH           ProxyMessageType = 0
	ProxyMessageType_PMT_MS_BS_ALLOCBATTLEROOM ProxyMessageType = 1
	ProxyMessageType_PMT_MS_GS_MATCH_RESULT    ProxyMessageType = 2
	ProxyMessageType_PMT_BS_MS_ALLOCBATTLEROOM ProxyMessageType = 3
	ProxyMessageType_PMT_MS_GS_BEGIN_BATTLE    ProxyMessageType = 4
	ProxyMessageType_PMT_GS_MS_TEAM_OPERATE    ProxyMessageType = 5
	ProxyMessageType_PMT_MS_GS_TEAM_OPERATE    ProxyMessageType = 6
)

var ProxyMessageType_name = map[int32]string{
	0: "PMT_GS_MS_MATCH",
	1: "PMT_MS_BS_ALLOCBATTLEROOM",
	2: "PMT_MS_GS_MATCH_RESULT",
	3: "PMT_BS_MS_ALLOCBATTLEROOM",
	4: "PMT_MS_GS_BEGIN_BATTLE",
	5: "PMT_GS_MS_TEAM_OPERATE",
	6: "PMT_MS_GS_TEAM_OPERATE",
}
var ProxyMessageType_value = map[string]int32{
	"PMT_GS_MS_MATCH":           0,
	"PMT_MS_BS_ALLOCBATTLEROOM": 1,
	"PMT_MS_GS_MATCH_RESULT":    2,
	"PMT_BS_MS_ALLOCBATTLEROOM": 3,
	"PMT_MS_GS_BEGIN_BATTLE":    4,
	"PMT_GS_MS_TEAM_OPERATE":    5,
	"PMT_MS_GS_TEAM_OPERATE":    6,
}

func (x ProxyMessageType) String() string {
	return proto.EnumName(ProxyMessageType_name, int32(x))
}
func (ProxyMessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// proxy message protocol
type InternalMessage struct {
	Fromid   int32  `protobuf:"varint,1,opt,name=fromid" json:"fromid,omitempty"`
	Fromtype string `protobuf:"bytes,2,opt,name=fromtype" json:"fromtype,omitempty"`
	Toid     int32  `protobuf:"varint,3,opt,name=toid" json:"toid,omitempty"`
	Totype   string `protobuf:"bytes,4,opt,name=totype" json:"totype,omitempty"`
	Charid   uint32 `protobuf:"varint,5,opt,name=charid" json:"charid,omitempty"`
	Msgid    uint32 `protobuf:"varint,6,opt,name=msgid" json:"msgid,omitempty"`
	Msgdata  []byte `protobuf:"bytes,7,opt,name=msgdata,proto3" json:"msgdata,omitempty"`
}

func (m *InternalMessage) Reset()                    { *m = InternalMessage{} }
func (m *InternalMessage) String() string            { return proto.CompactTextString(m) }
func (*InternalMessage) ProtoMessage()               {}
func (*InternalMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Proxy_GS_MS_Match struct {
	Charid    uint32 `protobuf:"varint,1,opt,name=charid" json:"charid,omitempty"`
	Charname  string `protobuf:"bytes,2,opt,name=charname" json:"charname,omitempty"`
	Matchmode int32  `protobuf:"varint,3,opt,name=matchmode" json:"matchmode,omitempty"`
	Mapid     int32  `protobuf:"varint,4,opt,name=mapid" json:"mapid,omitempty"`
	Action    int32  `protobuf:"varint,5,opt,name=action" json:"action,omitempty"`
}

func (m *Proxy_GS_MS_Match) Reset()                    { *m = Proxy_GS_MS_Match{} }
func (m *Proxy_GS_MS_Match) String() string            { return proto.CompactTextString(m) }
func (*Proxy_GS_MS_Match) ProtoMessage()               {}
func (*Proxy_GS_MS_Match) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Proxy_MS_BS_AllocBattleRoom struct {
	Matchmode    int32                                     `protobuf:"varint,1,opt,name=matchmode" json:"matchmode,omitempty"`
	Mapid        int32                                     `protobuf:"varint,2,opt,name=mapid" json:"mapid,omitempty"`
	Matchtableid int32                                     `protobuf:"varint,3,opt,name=matchtableid" json:"matchtableid,omitempty"`
	Members      []*Proxy_MS_BS_AllocBattleRoom_MemberInfo `protobuf:"bytes,4,rep,name=members" json:"members,omitempty"`
}

func (m *Proxy_MS_BS_AllocBattleRoom) Reset()                    { *m = Proxy_MS_BS_AllocBattleRoom{} }
func (m *Proxy_MS_BS_AllocBattleRoom) String() string            { return proto.CompactTextString(m) }
func (*Proxy_MS_BS_AllocBattleRoom) ProtoMessage()               {}
func (*Proxy_MS_BS_AllocBattleRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Proxy_MS_BS_AllocBattleRoom) GetMembers() []*Proxy_MS_BS_AllocBattleRoom_MemberInfo {
	if m != nil {
		return m.Members
	}
	return nil
}

type Proxy_MS_BS_AllocBattleRoom_MemberInfo struct {
	CharID       uint32 `protobuf:"varint,1,opt,name=CharID" json:"CharID,omitempty"`
	CharName     string `protobuf:"bytes,2,opt,name=CharName" json:"CharName,omitempty"`
	CharType     int32  `protobuf:"varint,3,opt,name=CharType" json:"CharType,omitempty"`
	TeamID       int32  `protobuf:"varint,4,opt,name=TeamID" json:"TeamID,omitempty"`
	OwnerID      uint32 `protobuf:"varint,5,opt,name=OwnerID" json:"OwnerID,omitempty"`
	GameServerID int32  `protobuf:"varint,6,opt,name=GameServerID" json:"GameServerID,omitempty"`
}

func (m *Proxy_MS_BS_AllocBattleRoom_MemberInfo) Reset() {
	*m = Proxy_MS_BS_AllocBattleRoom_MemberInfo{}
}
func (m *Proxy_MS_BS_AllocBattleRoom_MemberInfo) String() string { return proto.CompactTextString(m) }
func (*Proxy_MS_BS_AllocBattleRoom_MemberInfo) ProtoMessage()    {}
func (*Proxy_MS_BS_AllocBattleRoom_MemberInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

type Proxy_MS_GS_Begin_Battle struct {
	BattleRoomID int32  `protobuf:"varint,1,opt,name=BattleRoomID" json:"BattleRoomID,omitempty"`
	BattleAddr   string `protobuf:"bytes,2,opt,name=BattleAddr" json:"BattleAddr,omitempty"`
	BattleKey    []byte `protobuf:"bytes,3,opt,name=BattleKey,proto3" json:"BattleKey,omitempty"`
}

func (m *Proxy_MS_GS_Begin_Battle) Reset()                    { *m = Proxy_MS_GS_Begin_Battle{} }
func (m *Proxy_MS_GS_Begin_Battle) String() string            { return proto.CompactTextString(m) }
func (*Proxy_MS_GS_Begin_Battle) ProtoMessage()               {}
func (*Proxy_MS_GS_Begin_Battle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Proxy_BS_MS_AllocBattleRoom struct {
	Retcode        int32  `protobuf:"varint,1,opt,name=retcode" json:"retcode,omitempty"`
	Matchtableid   int32  `protobuf:"varint,2,opt,name=matchtableid" json:"matchtableid,omitempty"`
	Battleroomid   int32  `protobuf:"varint,3,opt,name=battleroomid" json:"battleroomid,omitempty"`
	Battleroomkey  []byte `protobuf:"bytes,4,opt,name=battleroomkey,proto3" json:"battleroomkey,omitempty"`
	Battleserverid int32  `protobuf:"varint,5,opt,name=battleserverid" json:"battleserverid,omitempty"`
	Connectaddr    string `protobuf:"bytes,6,opt,name=connectaddr" json:"connectaddr,omitempty"`
}

func (m *Proxy_BS_MS_AllocBattleRoom) Reset()                    { *m = Proxy_BS_MS_AllocBattleRoom{} }
func (m *Proxy_BS_MS_AllocBattleRoom) String() string            { return proto.CompactTextString(m) }
func (*Proxy_BS_MS_AllocBattleRoom) ProtoMessage()               {}
func (*Proxy_BS_MS_AllocBattleRoom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*InternalMessage)(nil), "proxymsg.InternalMessage")
	proto.RegisterType((*Proxy_GS_MS_Match)(nil), "proxymsg.Proxy_GS_MS_Match")
	proto.RegisterType((*Proxy_MS_BS_AllocBattleRoom)(nil), "proxymsg.Proxy_MS_BS_AllocBattleRoom")
	proto.RegisterType((*Proxy_MS_BS_AllocBattleRoom_MemberInfo)(nil), "proxymsg.Proxy_MS_BS_AllocBattleRoom.MemberInfo")
	proto.RegisterType((*Proxy_MS_GS_Begin_Battle)(nil), "proxymsg.Proxy_MS_GS_Begin_Battle")
	proto.RegisterType((*Proxy_BS_MS_AllocBattleRoom)(nil), "proxymsg.Proxy_BS_MS_AllocBattleRoom")
	proto.RegisterEnum("proxymsg.ProxyMessageType", ProxyMessageType_name, ProxyMessageType_value)
}

var fileDescriptor0 = []byte{
	// 622 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0xdb, 0x6e, 0xd3, 0x4c,
	0x10, 0xfe, 0x9d, 0x83, 0xdb, 0x6e, 0xd2, 0xbf, 0x66, 0x41, 0x55, 0x28, 0x07, 0x55, 0x16, 0x42,
	0x15, 0x17, 0x11, 0x82, 0x27, 0xb0, 0x83, 0x15, 0x02, 0x71, 0x13, 0x39, 0xe6, 0xda, 0xda, 0x38,
	0xdb, 0x34, 0x22, 0xf6, 0x46, 0xce, 0x0a, 0x88, 0xc4, 0x5b, 0xf0, 0x26, 0xdc, 0xf1, 0x20, 0x3c,
	0x07, 0x77, 0x5c, 0x33, 0x7b, 0x8a, 0xed, 0x70, 0x90, 0xa2, 0x68, 0xbf, 0x6f, 0x76, 0x66, 0xe7,
	0xfb, 0x66, 0xd7, 0xa8, 0xb3, 0x29, 0xd8, 0xa7, 0x5d, 0x1f, 0xfe, 0x39, 0xc3, 0xc7, 0x12, 0x64,
	0xdb, 0xa5, 0xfb, 0xcd, 0x42, 0x67, 0xa3, 0x9c, 0xd3, 0x22, 0x27, 0xeb, 0x90, 0x6e, 0xb7, 0x64,
	0x49, 0xf1, 0x39, 0xb2, 0x6f, 0x0a, 0x96, 0xad, 0x16, 0x3d, 0xeb, 0xd2, 0xba, 0x6a, 0x47, 0x1a,
	0xe1, 0x0b, 0x74, 0x2c, 0x56, 0x7c, 0xb7, 0xa1, 0xbd, 0x06, 0x44, 0x4e, 0xa2, 0x3d, 0xc6, 0x18,
	0xb5, 0x38, 0x83, 0x8c, 0xa6, 0xcc, 0x90, 0x6b, 0x51, 0x87, 0x33, 0xb9, 0xbb, 0x25, 0x77, 0x6b,
	0x24, 0xf8, 0xf4, 0x96, 0x14, 0xb0, 0xbb, 0x0d, 0xfc, 0x69, 0xa4, 0x11, 0xbe, 0x87, 0xda, 0xd0,
	0x12, 0xd0, 0xb6, 0xa4, 0x15, 0xc0, 0x3d, 0x74, 0x04, 0x8b, 0x05, 0xe1, 0xa4, 0x77, 0x04, 0x7c,
	0x37, 0x32, 0xd0, 0xfd, 0x62, 0xa1, 0x3b, 0x53, 0x21, 0x24, 0x19, 0xce, 0x92, 0x10, 0x7e, 0x84,
	0xa7, 0xb7, 0x95, 0xea, 0x56, 0xad, 0x3a, 0x74, 0x2f, 0x56, 0x39, 0xc9, 0xf6, 0xdd, 0x1b, 0x8c,
	0x1f, 0xa2, 0x93, 0x4c, 0x24, 0x67, 0x6c, 0x41, 0xb5, 0x84, 0x92, 0x90, 0x7d, 0x91, 0x0d, 0x14,
	0x6c, 0xc9, 0x88, 0x02, 0xe2, 0x1c, 0x92, 0xf2, 0x15, 0xcb, 0xa5, 0x0a, 0x70, 0x49, 0x21, 0xf7,
	0x67, 0x03, 0x3d, 0x50, 0x5d, 0x41, 0x4b, 0xfe, 0x2c, 0xf1, 0xd6, 0x6b, 0x96, 0xfa, 0x84, 0xf3,
	0x35, 0x8d, 0x18, 0xcb, 0xea, 0x67, 0x59, 0x7f, 0x3d, 0xab, 0x51, 0x3d, 0xcb, 0x45, 0x5d, 0xb9,
	0x85, 0x93, 0xf9, 0x9a, 0xee, 0x5d, 0xae, 0x71, 0xf8, 0x0d, 0xf8, 0x44, 0xb3, 0x39, 0x2d, 0xb6,
	0xd0, 0x67, 0xf3, 0xaa, 0xf3, 0xe2, 0x79, 0xdf, 0x4c, 0xb9, 0xff, 0x8f, 0x7e, 0xfa, 0xa1, 0x4c,
	0x1a, 0xe5, 0x37, 0x2c, 0x32, 0x05, 0x2e, 0xbe, 0x5a, 0x08, 0x95, 0xbc, 0x90, 0x3a, 0x00, 0xab,
	0x46, 0xaf, 0x8c, 0xa5, 0x0a, 0x09, 0x4b, 0xc5, 0xea, 0xba, 0x62, 0xa9, 0xc1, 0x26, 0x16, 0x8b,
	0xf1, 0xab, 0x76, 0xf7, 0x58, 0xd4, 0x8b, 0x29, 0xc9, 0xa0, 0x9e, 0x72, 0x54, 0x23, 0x31, 0xea,
	0xc9, 0xc7, 0x9c, 0x8a, 0x83, 0xd4, 0xcd, 0x30, 0x50, 0x18, 0x30, 0x84, 0xaa, 0x33, 0x5a, 0x7c,
	0x90, 0x61, 0x5b, 0x19, 0x50, 0xe5, 0xdc, 0xcf, 0xa8, 0xb7, 0xd7, 0x09, 0x17, 0xc2, 0xa7, 0xcb,
	0x55, 0x9e, 0x28, 0xa1, 0x22, 0xbf, 0x94, 0xac, 0x75, 0x40, 0x7e, 0x95, 0xc3, 0x8f, 0x11, 0x52,
	0xd8, 0x5b, 0x2c, 0x0a, 0xad, 0xa7, 0xc2, 0x88, 0xc1, 0x29, 0xf4, 0x96, 0xee, 0xa4, 0xa4, 0x6e,
	0x54, 0x12, 0xee, 0x0f, 0xcb, 0x8c, 0xdd, 0x97, 0x97, 0xf1, 0x70, 0xec, 0xa0, 0xad, 0xa0, 0x3c,
	0x2d, 0x87, 0x6e, 0xe0, 0x6f, 0xc3, 0x6d, 0xfc, 0x61, 0xb8, 0xb0, 0x67, 0x2e, 0x6b, 0x15, 0x4c,
	0x3e, 0x4c, 0x7d, 0x01, 0xaa, 0x1c, 0x7e, 0x82, 0x4e, 0x4b, 0xfc, 0x1e, 0x7a, 0x6c, 0xc9, 0x1e,
	0xeb, 0x24, 0x7e, 0x8a, 0xfe, 0x57, 0xc4, 0x56, 0xfa, 0xa6, 0x1f, 0x61, 0x3b, 0x3a, 0x60, 0xf1,
	0x25, 0xea, 0xa4, 0x2c, 0xcf, 0x69, 0xca, 0x89, 0xb0, 0xc3, 0x96, 0x76, 0x54, 0xa9, 0x67, 0xdf,
	0x2d, 0xe4, 0x48, 0xc5, 0xfa, 0xbb, 0x21, 0x47, 0x7b, 0x17, 0x9d, 0x4d, 0xc3, 0xd8, 0x3c, 0x48,
	0x2f, 0x1e, 0xbc, 0x76, 0xfe, 0xc3, 0x8f, 0xd0, 0x7d, 0x41, 0xea, 0xfb, 0x37, 0x1e, 0x4f, 0x06,
	0xbe, 0x17, 0xc7, 0xe3, 0x20, 0x9a, 0x4c, 0x42, 0xc7, 0x82, 0xab, 0x72, 0xae, 0xc3, 0x43, 0x9d,
	0x93, 0x44, 0xc1, 0xec, 0xdd, 0x38, 0x76, 0x1a, 0x26, 0x55, 0x7b, 0x7a, 0x90, 0xda, 0xac, 0xa7,
	0xfa, 0xc1, 0x70, 0x74, 0x9d, 0xa8, 0xb8, 0xd3, 0x32, 0x31, 0xd5, 0x4a, 0x1c, 0x78, 0x61, 0x32,
	0x99, 0x06, 0x91, 0x17, 0x07, 0x4e, 0xbb, 0x9e, 0x57, 0x8b, 0xd9, 0x73, 0x5b, 0x7e, 0x23, 0x5f,
	0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x28, 0x08, 0xa2, 0x8c, 0x32, 0x05, 0x00, 0x00,
}
