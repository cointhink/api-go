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
	Email string `protobuf:"bytes,1,opt,name=Email" json:"Email,omitempty"`
}

func (m *SigninEmail) Reset()                    { *m = SigninEmail{} }
func (m *SigninEmail) String() string            { return proto1.CompactTextString(m) }
func (*SigninEmail) ProtoMessage()               {}
func (*SigninEmail) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *SigninEmail) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto1.RegisterType((*SigninEmail)(nil), "proto.SigninEmail")
}

func init() { proto1.RegisterFile("proto/signin_email.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 80 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0x4c, 0xcf, 0xcb, 0xcc, 0x8b, 0x4f, 0xcd, 0x4d, 0xcc, 0xcc, 0xd1, 0x03,
	0x0b, 0x09, 0xb1, 0x82, 0x29, 0x25, 0x65, 0x2e, 0xee, 0x60, 0xb0, 0xa4, 0x2b, 0x48, 0x4e, 0x48,
	0x84, 0x8b, 0x15, 0xcc, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x92, 0xd8, 0xc0,
	0x6a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x06, 0x68, 0x4b, 0x4e, 0x00, 0x00, 0x00,
}
