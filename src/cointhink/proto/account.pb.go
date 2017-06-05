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
	Email    string `protobuf:"bytes,1,opt,name=Email" json:"Email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=Username" json:"Username,omitempty"`
	Fullname string `protobuf:"bytes,3,opt,name=Fullname" json:"Fullname,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto1.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

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

func init() { proto1.RegisterFile("proto/account.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x03, 0xf3, 0x84, 0x58, 0xc1, 0x94,
	0x52, 0x38, 0x17, 0xbb, 0x23, 0x44, 0x5c, 0x48, 0x84, 0x8b, 0xd5, 0x35, 0x37, 0x31, 0x33, 0x47,
	0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc2, 0x11, 0x92, 0xe2, 0xe2, 0x08, 0x2d, 0x4e, 0x2d,
	0xca, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x4b, 0xc0, 0xf9, 0x20, 0x39, 0xb7, 0xd2, 0x9c, 0x1c,
	0xb0, 0x1c, 0x33, 0x44, 0x0e, 0xc6, 0x4f, 0x62, 0x03, 0x9b, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xe1, 0x8e, 0x93, 0xd5, 0x7d, 0x00, 0x00, 0x00,
}
