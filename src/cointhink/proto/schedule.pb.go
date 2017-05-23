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

type Schedule struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	AccountId   string `protobuf:"bytes,2,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	AlgorithmId string `protobuf:"bytes,3,opt,name=algorithm_id,json=algorithmId" json:"algorithm_id,omitempty"`
	Status      string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
}

func (m *Schedule) Reset()                    { *m = Schedule{} }
func (m *Schedule) String() string            { return proto1.CompactTextString(m) }
func (*Schedule) ProtoMessage()               {}
func (*Schedule) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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

func (m *Schedule) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto1.RegisterType((*Schedule)(nil), "proto.Schedule")
}

func init() { proto1.RegisterFile("proto/schedule.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0xd5, 0x03, 0x73, 0x85, 0x58, 0xc1,
	0x94, 0x52, 0x09, 0x17, 0x47, 0x30, 0x54, 0x42, 0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51,
	0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x96, 0x8b, 0x2b, 0x31, 0x39, 0x39, 0xbf,
	0x34, 0xaf, 0x24, 0x3e, 0x33, 0x45, 0x82, 0x09, 0x2c, 0xce, 0x09, 0x15, 0xf1, 0x4c, 0x11, 0x52,
	0xe4, 0xe2, 0x49, 0xcc, 0x49, 0xcf, 0x2f, 0xca, 0x2c, 0xc9, 0xc8, 0x05, 0x29, 0x60, 0x06, 0x2b,
	0xe0, 0x86, 0x8b, 0x79, 0xa6, 0x08, 0x89, 0x71, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x4b,
	0xb0, 0x80, 0x25, 0xa1, 0xbc, 0x24, 0x36, 0xb0, 0xe5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x21, 0x50, 0xc9, 0x73, 0x9b, 0x00, 0x00, 0x00,
}