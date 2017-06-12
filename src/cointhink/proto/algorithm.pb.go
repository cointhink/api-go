// Code generated by protoc-gen-go.
// source: proto/algorithm.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Algorithm struct {
	Id        string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AccountId string `protobuf:"bytes,2,opt,name=AccountId" json:"AccountId,omitempty"`
	Status    string `protobuf:"bytes,3,opt,name=Status" json:"Status,omitempty"`
	Code      string `protobuf:"bytes,4,opt,name=Code" json:"Code,omitempty"`
}

func (m *Algorithm) Reset()                    { *m = Algorithm{} }
func (m *Algorithm) String() string            { return proto1.CompactTextString(m) }
func (*Algorithm) ProtoMessage()               {}
func (*Algorithm) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *Algorithm) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Algorithm) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *Algorithm) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Algorithm) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto1.RegisterType((*Algorithm)(nil), "proto.Algorithm")
}

func init() { proto1.RegisterFile("proto/algorithm.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcc, 0x49, 0xcf, 0x2f, 0xca, 0x2c, 0xc9, 0xc8, 0xd5, 0x03, 0xf3, 0x85, 0x58,
	0xc1, 0x94, 0x52, 0x2a, 0x17, 0xa7, 0x23, 0x4c, 0x46, 0x88, 0x8f, 0x8b, 0xc9, 0x33, 0x45, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x88, 0xc9, 0x33, 0x45, 0x48, 0x86, 0x8b, 0xd3, 0x31, 0x39, 0x39,
	0xbf, 0x34, 0xaf, 0xc4, 0x33, 0x45, 0x82, 0x09, 0x2c, 0x8c, 0x10, 0x10, 0x12, 0xe3, 0x62, 0x0b,
	0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x96, 0x60, 0x06, 0x4b, 0x41, 0x79, 0x42, 0x42, 0x5c, 0x2c, 0xce,
	0xf9, 0x29, 0xa9, 0x12, 0x2c, 0x60, 0x51, 0x30, 0x3b, 0x89, 0x0d, 0x6c, 0x9b, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x8d, 0xb2, 0x12, 0x50, 0x8d, 0x00, 0x00, 0x00,
}
