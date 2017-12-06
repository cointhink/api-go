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
func (Schedule_States) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{0, 0} }

type Schedule_Executors int32

const (
	Schedule_container     Schedule_Executors = 0
	Schedule_lambda        Schedule_Executors = 1
	Schedule_lambda_master Schedule_Executors = 2
)

var Schedule_Executors_name = map[int32]string{
	0: "container",
	1: "lambda",
	2: "lambda_master",
}
var Schedule_Executors_value = map[string]int32{
	"container":     0,
	"lambda":        1,
	"lambda_master": 2,
}

func (x Schedule_Executors) String() string {
	return proto1.EnumName(Schedule_Executors_name, int32(x))
}
func (Schedule_Executors) EnumDescriptor() ([]byte, []int) { return fileDescriptor26, []int{0, 1} }

type Schedule struct {
	Id           string             `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AccountId    string             `protobuf:"bytes,2,opt,name=AccountId" json:"AccountId,omitempty"`
	AlgorithmId  string             `protobuf:"bytes,3,opt,name=AlgorithmId" json:"AlgorithmId,omitempty"`
	Status       Schedule_States    `protobuf:"varint,4,opt,name=Status,enum=proto.Schedule_States" json:"Status,omitempty"`
	InitialState string             `protobuf:"bytes,5,opt,name=InitialState" json:"InitialState,omitempty"`
	EnabledUntil string             `protobuf:"bytes,6,opt,name=EnabledUntil" json:"EnabledUntil,omitempty"`
	Executor     Schedule_Executors `protobuf:"varint,7,opt,name=Executor,enum=proto.Schedule_Executors" json:"Executor,omitempty"`
}

func (m *Schedule) Reset()                    { *m = Schedule{} }
func (m *Schedule) String() string            { return proto1.CompactTextString(m) }
func (*Schedule) ProtoMessage()               {}
func (*Schedule) Descriptor() ([]byte, []int) { return fileDescriptor26, []int{0} }

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

func (m *Schedule) GetEnabledUntil() string {
	if m != nil {
		return m.EnabledUntil
	}
	return ""
}

func (m *Schedule) GetExecutor() Schedule_Executors {
	if m != nil {
		return m.Executor
	}
	return Schedule_container
}

func init() {
	proto1.RegisterType((*Schedule)(nil), "proto.Schedule")
	proto1.RegisterEnum("proto.Schedule_States", Schedule_States_name, Schedule_States_value)
	proto1.RegisterEnum("proto.Schedule_Executors", Schedule_Executors_name, Schedule_Executors_value)
}

func init() { proto1.RegisterFile("proto/schedule.proto", fileDescriptor26) }

var fileDescriptor26 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xbf, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x9b, 0x94, 0xa6, 0xcd, 0xf5, 0x87, 0xcc, 0x09, 0x21, 0x23, 0x31, 0x54, 0x99, 0x3a,
	0x05, 0x09, 0xc4, 0xc0, 0xc0, 0xd0, 0xa1, 0x43, 0xd6, 0x56, 0xcc, 0xc8, 0x89, 0x4f, 0xd4, 0xc2,
	0xb1, 0x51, 0xec, 0x08, 0x76, 0xfe, 0x71, 0x14, 0xa7, 0x2d, 0x85, 0xc9, 0xbe, 0xf7, 0x3d, 0xdd,
	0x77, 0x70, 0xf5, 0xd1, 0x58, 0x6f, 0xef, 0x5c, 0xb5, 0x27, 0xd9, 0x6a, 0xca, 0xc3, 0x88, 0xa3,
	0xf0, 0x64, 0xdf, 0x43, 0x98, 0xec, 0x0e, 0x04, 0x17, 0x10, 0x17, 0x92, 0x47, 0xcb, 0x68, 0x95,
	0x6e, 0xe3, 0x42, 0xe2, 0x2d, 0xa4, 0xeb, 0xaa, 0xb2, 0xad, 0xf1, 0x85, 0xe4, 0x71, 0x88, 0x7f,
	0x03, 0x5c, 0xc2, 0x74, 0xad, 0xdf, 0x6c, 0xa3, 0xfc, 0xbe, 0x2e, 0x24, 0x1f, 0x06, 0x7e, 0x1e,
	0x61, 0x0e, 0xc9, 0xce, 0x0b, 0xdf, 0x3a, 0x7e, 0xb1, 0x8c, 0x56, 0x8b, 0xfb, 0xeb, 0xde, 0x9d,
	0x1f, 0x85, 0x79, 0x47, 0xc9, 0x6d, 0x0f, 0x2d, 0xcc, 0x60, 0x56, 0x18, 0xe5, 0x95, 0xd0, 0x01,
	0xf0, 0x51, 0x58, 0xf9, 0x27, 0xeb, 0x3a, 0x1b, 0x23, 0x4a, 0x4d, 0xf2, 0xc5, 0x78, 0xa5, 0x79,
	0xd2, 0x77, 0xce, 0x33, 0x7c, 0x84, 0xc9, 0xe6, 0x8b, 0xaa, 0xd6, 0xdb, 0x86, 0x8f, 0x83, 0xf9,
	0xe6, 0xbf, 0xf9, 0xc8, 0xdd, 0xf6, 0x54, 0xcd, 0x9e, 0xfb, 0x73, 0xc9, 0xe1, 0x14, 0xc6, 0xad,
	0x79, 0x37, 0xf6, 0xd3, 0xb0, 0x01, 0xce, 0x60, 0x22, 0x95, 0x0b, 0xeb, 0x59, 0xd4, 0x21, 0xea,
	0x5d, 0x2c, 0xee, 0x06, 0x49, 0x9a, 0x3c, 0x49, 0x36, 0xcc, 0x9e, 0x20, 0x3d, 0x6d, 0xc5, 0x39,
	0xa4, 0x95, 0x35, 0x5e, 0x28, 0x43, 0x0d, 0x1b, 0x20, 0x40, 0xa2, 0x45, 0x5d, 0x4a, 0xc1, 0x22,
	0xbc, 0x84, 0x79, 0xff, 0x7f, 0xad, 0x85, 0xf3, 0xd4, 0xb0, 0xb8, 0x4c, 0xc2, 0x75, 0x0f, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x43, 0xa6, 0xaf, 0xab, 0x01, 0x00, 0x00,
}
