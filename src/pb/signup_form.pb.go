// Code generated by protoc-gen-go.
// source: signup_form.proto
// DO NOT EDIT!

/*
Package signup_form is a generated protocol buffer package.

It is generated from these files:
	signup_form.proto

It has these top-level messages:
	SignupForm
*/
package signup_form

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SignupForm struct {
	Email    string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *SignupForm) Reset()                    { *m = SignupForm{} }
func (m *SignupForm) String() string            { return proto.CompactTextString(m) }
func (*SignupForm) ProtoMessage()               {}
func (*SignupForm) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignupForm) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignupForm) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignupForm) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*SignupForm)(nil), "SignupForm")
}

func init() { proto.RegisterFile("signup_form.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xce, 0x4c, 0xcf,
	0x2b, 0x2d, 0x88, 0x4f, 0xcb, 0x2f, 0xca, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe2,
	0xe2, 0x0a, 0x06, 0x0b, 0xba, 0xe5, 0x17, 0xe5, 0x0a, 0x89, 0x70, 0xb1, 0xa6, 0xe6, 0x26, 0x66,
	0xe6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x52, 0x5c, 0x1c, 0xa5, 0xc5,
	0xa9, 0x45, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x09, 0x38, 0x1f, 0x24, 0x57, 0x90, 0x58,
	0x5c, 0x5c, 0x9e, 0x5f, 0x94, 0x22, 0xc1, 0x0c, 0x91, 0x83, 0xf1, 0x93, 0xd8, 0xc0, 0x56, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xbc, 0x37, 0x04, 0x77, 0x00, 0x00, 0x00,
}
