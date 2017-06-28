// Code generated by protoc-gen-go.
// source: proto/signin_email_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SigninEmailResponse struct {
	Ok      bool   `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *SigninEmailResponse) Reset()                    { *m = SigninEmailResponse{} }
func (m *SigninEmailResponse) String() string            { return proto1.CompactTextString(m) }
func (*SigninEmailResponse) ProtoMessage()               {}
func (*SigninEmailResponse) Descriptor() ([]byte, []int) { return fileDescriptor20, []int{0} }

func (m *SigninEmailResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *SigninEmailResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*SigninEmailResponse)(nil), "proto.SigninEmailResponse")
}

func init() { proto1.RegisterFile("proto/signin_email_response.proto", fileDescriptor20) }

var fileDescriptor20 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0x4c, 0xcf, 0xcb, 0xcc, 0x8b, 0x4f, 0xcd, 0x4d, 0xcc, 0xcc, 0x89, 0x2f,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0xcb, 0x09, 0xb1, 0x82, 0x29, 0x25, 0x7b,
	0x2e, 0xe1, 0x60, 0xb0, 0x2a, 0x57, 0x90, 0xa2, 0x20, 0xa8, 0x1a, 0x21, 0x3e, 0x2e, 0x26, 0xff,
	0x6c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x26, 0xff, 0x6c, 0x21, 0x09, 0x2e, 0x76, 0xdf,
	0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x37, 0x89,
	0x0d, 0x6c, 0x8e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x04, 0x19, 0xc7, 0x73, 0x00, 0x00,
	0x00,
}
