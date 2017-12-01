// Code generated by protoc-gen-go.
// source: proto/algorun.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Algorun_States int32

const (
	Algorun_unknown    Algorun_States = 0
	Algorun_building   Algorun_States = 1
	Algorun_starting   Algorun_States = 2
	Algorun_running    Algorun_States = 3
	Algorun_stopped    Algorun_States = 4
	Algorun_destroying Algorun_States = 5
	Algorun_deleted    Algorun_States = 6
)

var Algorun_States_name = map[int32]string{
	0: "unknown",
	1: "building",
	2: "starting",
	3: "running",
	4: "stopped",
	5: "destroying",
	6: "deleted",
}
var Algorun_States_value = map[string]int32{
	"unknown":    0,
	"building":   1,
	"starting":   2,
	"running":    3,
	"stopped":    4,
	"destroying": 5,
	"deleted":    6,
}

func (x Algorun_States) String() string {
	return proto1.EnumName(Algorun_States_name, int32(x))
}
func (Algorun_States) EnumDescriptor() ([]byte, []int) { return fileDescriptor18, []int{0, 0} }

type Algorun struct {
	Id          string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AlgorithmId string `protobuf:"bytes,2,opt,name=AlgorithmId" json:"AlgorithmId,omitempty"`
	AccountId   string `protobuf:"bytes,3,opt,name=AccountId" json:"AccountId,omitempty"`
	ScheduleId  string `protobuf:"bytes,4,opt,name=ScheduleId" json:"ScheduleId,omitempty"`
	Status      string `protobuf:"bytes,5,opt,name=Status" json:"Status,omitempty"`
	Code        string `protobuf:"bytes,6,opt,name=Code" json:"Code,omitempty"`
	Image       string `protobuf:"bytes,7,opt,name=Image" json:"Image,omitempty"`
	State       string `protobuf:"bytes,8,opt,name=State" json:"State,omitempty"`
}

func (m *Algorun) Reset()                    { *m = Algorun{} }
func (m *Algorun) String() string            { return proto1.CompactTextString(m) }
func (*Algorun) ProtoMessage()               {}
func (*Algorun) Descriptor() ([]byte, []int) { return fileDescriptor18, []int{0} }

func (m *Algorun) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Algorun) GetAlgorithmId() string {
	if m != nil {
		return m.AlgorithmId
	}
	return ""
}

func (m *Algorun) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Algorun) GetScheduleId() string {
	if m != nil {
		return m.ScheduleId
	}
	return ""
}

func (m *Algorun) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Algorun) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Algorun) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Algorun) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto1.RegisterType((*Algorun)(nil), "proto.Algorun")
	proto1.RegisterEnum("proto.Algorun_States", Algorun_States_name, Algorun_States_value)
}

func init() { proto1.RegisterFile("proto/algorun.proto", fileDescriptor18) }

var fileDescriptor18 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x80, 0x49, 0x9a, 0x9f, 0xf6, 0x8a, 0xaa, 0xe8, 0x40, 0xc8, 0x03, 0x42, 0x55, 0x27, 0x26,
	0x18, 0x78, 0x82, 0x8a, 0xc9, 0x2b, 0x7d, 0x82, 0x34, 0x77, 0x4a, 0x22, 0x52, 0x3b, 0x4a, 0xce,
	0x42, 0x3c, 0x1b, 0x2f, 0x87, 0x7c, 0x41, 0xa2, 0x93, 0xfd, 0x7d, 0x9f, 0x7d, 0xc3, 0xc1, 0xdd,
	0x38, 0x79, 0xf1, 0xaf, 0xf5, 0xd0, 0xfa, 0x29, 0xb8, 0x17, 0x25, 0xcc, 0xf5, 0x38, 0xfc, 0xa4,
	0x50, 0x1e, 0x97, 0x80, 0x3b, 0x48, 0x2d, 0x99, 0x64, 0x9f, 0x3c, 0x6f, 0x3e, 0x52, 0x4b, 0xb8,
	0x87, 0xad, 0xa6, 0x5e, 0xba, 0x8b, 0x25, 0x93, 0x6a, 0xb8, 0x56, 0xf8, 0x08, 0x9b, 0x63, 0xd3,
	0xf8, 0xe0, 0xc4, 0x92, 0x59, 0x69, 0xff, 0x17, 0xf8, 0x04, 0x70, 0x6a, 0x3a, 0xa6, 0x30, 0xb0,
	0x25, 0x93, 0x69, 0xbe, 0x32, 0xf8, 0x00, 0xc5, 0x49, 0x6a, 0x09, 0xb3, 0xc9, 0xb5, 0xfd, 0x11,
	0x22, 0x64, 0xef, 0x9e, 0xd8, 0x14, 0x6a, 0xf5, 0x8e, 0xf7, 0x90, 0xdb, 0x4b, 0xdd, 0xb2, 0x29,
	0x55, 0x2e, 0x10, 0x6d, 0xfc, 0xc3, 0x66, 0xbd, 0x58, 0x85, 0x43, 0xb7, 0xcc, 0xe5, 0x19, 0xb7,
	0x50, 0x06, 0xf7, 0xe9, 0xfc, 0x97, 0xab, 0x6e, 0xf0, 0x16, 0xd6, 0xe7, 0xd0, 0x0f, 0xd4, 0xbb,
	0xb6, 0x4a, 0x22, 0xcd, 0x52, 0x4f, 0x12, 0x29, 0x8d, 0x0f, 0xa7, 0xe0, 0x5c, 0x84, 0x55, 0x84,
	0x59, 0xfc, 0x38, 0x32, 0x55, 0x19, 0xee, 0x00, 0x88, 0x67, 0x99, 0xfc, 0x77, 0x8c, 0x79, 0x8c,
	0xc4, 0x03, 0x0b, 0x53, 0x55, 0x9c, 0x0b, 0x5d, 0xe2, 0xdb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x90, 0xdf, 0x81, 0x42, 0x62, 0x01, 0x00, 0x00,
}
