// Code generated by protoc-gen-go.
// source: proto/schedule.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Schedule_States int32

const (
	Schedule_unknown  Schedule_States = 0
	Schedule_disabled Schedule_States = 1
	Schedule_enabled  Schedule_States = 2
	Schedule_deleted  Schedule_States = 3
)

var Schedule_States_name = map[int32]string{
	0: "unknown",
	1: "disabled",
	2: "enabled",
	3: "deleted",
}
var Schedule_States_value = map[string]int32{
	"unknown":  0,
	"disabled": 1,
	"enabled":  2,
	"deleted":  3,
}

func (x Schedule_States) String() string {
	return proto1.EnumName(Schedule_States_name, int32(x))
}
func (Schedule_States) EnumDescriptor() ([]byte, []int) { return fileDescriptor21, []int{0, 0} }

type Schedule struct {
	Id           string          `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AccountId    string          `protobuf:"bytes,2,opt,name=AccountId" json:"AccountId,omitempty"`
	AlgorithmId  string          `protobuf:"bytes,3,opt,name=AlgorithmId" json:"AlgorithmId,omitempty"`
	Status       Schedule_States `protobuf:"varint,4,opt,name=Status,enum=proto.Schedule_States" json:"Status,omitempty"`
	InitialState string          `protobuf:"bytes,5,opt,name=InitialState" json:"InitialState,omitempty"`
}

func (m *Schedule) Reset()                    { *m = Schedule{} }
func (m *Schedule) String() string            { return proto1.CompactTextString(m) }
func (*Schedule) ProtoMessage()               {}
func (*Schedule) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *Schedule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Schedule) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Schedule) GetAlgorithmId() string {
	if m != nil {
		return m.AlgorithmId
	}
	return ""
}

func (m *Schedule) GetStatus() Schedule_States {
	if m != nil {
		return m.Status
	}
	return Schedule_unknown
}

func (m *Schedule) GetInitialState() string {
	if m != nil {
		return m.InitialState
	}
	return ""
}

func init() {
	proto1.RegisterType((*Schedule)(nil), "proto.Schedule")
	proto1.RegisterEnum("proto.Schedule_States", Schedule_States_name, Schedule_States_value)
}

func init() { proto1.RegisterFile("proto/schedule.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0xb1, 0x4b, 0x4b, 0x7b, 0xad, 0xaa, 0xe8, 0x84, 0x90, 0x07, 0x86, 0x28, 0x53, 0x27,
	0x23, 0xc1, 0xcc, 0xd0, 0xd1, 0x6b, 0xfb, 0x04, 0x69, 0xee, 0x44, 0x2c, 0x8c, 0x8d, 0x62, 0x5b,
	0x3c, 0x35, 0xef, 0x80, 0xe2, 0x04, 0x41, 0x27, 0xeb, 0xff, 0x3e, 0xeb, 0xff, 0x0f, 0xee, 0x3f,
	0x87, 0x90, 0xc2, 0x53, 0xec, 0x7a, 0xa6, 0xec, 0x58, 0x97, 0x88, 0xcb, 0xf2, 0x34, 0xdf, 0x02,
	0xd6, 0xe7, 0xd9, 0xe0, 0x1e, 0xa4, 0x21, 0x25, 0x6a, 0x71, 0xd8, 0x9c, 0xa4, 0x21, 0x7c, 0x84,
	0xcd, 0xb1, 0xeb, 0x42, 0xf6, 0xc9, 0x90, 0x92, 0x05, 0xff, 0x01, 0xac, 0x61, 0x7b, 0x74, 0x6f,
	0x61, 0xb0, 0xa9, 0xff, 0x30, 0xa4, 0x16, 0xc5, 0xff, 0x47, 0xa8, 0x61, 0x75, 0x4e, 0x6d, 0xca,
	0x51, 0xdd, 0xd6, 0xe2, 0xb0, 0x7f, 0x7e, 0x98, 0xb6, 0xf5, 0xef, 0xa0, 0x1e, 0x2d, 0xc7, 0xd3,
	0xfc, 0x0b, 0x1b, 0xd8, 0x19, 0x6f, 0x93, 0x6d, 0x5d, 0x11, 0x6a, 0x59, 0x2a, 0xaf, 0x58, 0xf3,
	0x3a, 0x75, 0x72, 0xc4, 0x2d, 0xdc, 0x65, 0xff, 0xee, 0xc3, 0x97, 0xaf, 0x6e, 0x70, 0x07, 0x6b,
	0xb2, 0xb1, 0xbd, 0x38, 0xa6, 0x4a, 0x8c, 0x8a, 0xfd, 0x14, 0xe4, 0x18, 0x88, 0x1d, 0x27, 0xa6,
	0x6a, 0x71, 0x59, 0x95, 0x0b, 0x5e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x66, 0x71, 0x42,
	0x15, 0x01, 0x00, 0x00,
}
