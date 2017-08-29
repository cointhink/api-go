// Code generated by protoc-gen-go.
// source: proto/account.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Account struct {
	Id       string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email" json:"Email,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=Username" json:"Username,omitempty"`
	Fullname string `protobuf:"bytes,4,opt,name=Fullname" json:"Fullname,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto1.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetFullname() string {
	if m != nil {
		return m.Fullname
	}
	return ""
}

func init() {
	proto1.RegisterType((*Account)(nil), "proto.Account")
}

func init() { proto1.RegisterFile("proto/account.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x03, 0xf3, 0x84, 0x58, 0xc1, 0x94,
	0x52, 0x3a, 0x17, 0xbb, 0x23, 0x44, 0x5c, 0x88, 0x8f, 0x8b, 0xc9, 0x33, 0x45, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x33, 0x88, 0xc9, 0x33, 0x45, 0x48, 0x84, 0x8b, 0xd5, 0x35, 0x37, 0x31, 0x33, 0x47,
	0x82, 0x09, 0x2c, 0x04, 0xe1, 0x08, 0x49, 0x71, 0x71, 0x84, 0x16, 0xa7, 0x16, 0xe5, 0x25, 0xe6,
	0xa6, 0x4a, 0x30, 0x83, 0x25, 0xe0, 0x7c, 0x90, 0x9c, 0x5b, 0x69, 0x4e, 0x0e, 0x58, 0x8e, 0x05,
	0x22, 0x07, 0xe3, 0x27, 0xb1, 0x81, 0xed, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x92,
	0x43, 0xc7, 0x8d, 0x00, 0x00, 0x00,
}
