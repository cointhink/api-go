// Code generated by protoc-gen-go.
// source: proto/signin_email.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SigninEmail struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *SigninEmail) Reset()                    { *m = SigninEmail{} }
func (m *SigninEmail) String() string            { return proto1.CompactTextString(m) }
func (*SigninEmail) ProtoMessage()               {}
func (*SigninEmail) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SigninEmail) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto1.RegisterType((*SigninEmail)(nil), "proto.SigninEmail")
}

func init() { proto1.RegisterFile("proto/signin_email.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0x4c, 0xcf, 0xcb, 0xcc, 0x8b, 0x4f, 0xcd, 0x4d, 0xcc, 0xcc, 0xd1, 0x03,
	0x0b, 0x09, 0xb1, 0x82, 0x29, 0x29, 0x61, 0x88, 0x82, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12,
	0x88, 0x9c, 0x92, 0x32, 0x17, 0x77, 0x30, 0x58, 0x87, 0x2b, 0x48, 0x83, 0x90, 0x08, 0x17, 0x2b,
	0x58, 0xa7, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x93, 0xc4, 0x06, 0x56, 0x6b, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x4f, 0xb7, 0x0a, 0x63, 0x00, 0x00, 0x00,
}
