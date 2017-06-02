// Code generated by protoc-gen-go.
// source: proto/session_create_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SessionCreateResponse struct {
	Ok      bool     `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Account *Account `protobuf:"bytes,2,opt,name=Account" json:"Account,omitempty"`
}

func (m *SessionCreateResponse) Reset()                    { *m = SessionCreateResponse{} }
func (m *SessionCreateResponse) String() string            { return proto1.CompactTextString(m) }
func (*SessionCreateResponse) ProtoMessage()               {}
func (*SessionCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor9, []int{0} }

func (m *SessionCreateResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *SessionCreateResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func init() {
	proto1.RegisterType((*SessionCreateResponse)(nil), "proto.SessionCreateResponse")
}

func init() { proto1.RegisterFile("proto/session_create_response.proto", fileDescriptor9) }

var fileDescriptor9 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x8b, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c, 0x49,
	0x8d, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0xcb, 0x0a, 0xb1, 0x82, 0x29,
	0x29, 0x61, 0x88, 0xda, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0x88, 0x9c, 0x52, 0x20, 0x97,
	0x68, 0x30, 0x44, 0xb3, 0x33, 0x58, 0x6f, 0x10, 0x54, 0xab, 0x10, 0x1f, 0x17, 0x93, 0x7f, 0xb6,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47, 0x10, 0x93, 0x7f, 0xb6, 0x90, 0x06, 0x17, 0xbb, 0x23, 0x44,
	0xa7, 0x04, 0x93, 0x02, 0xa3, 0x06, 0xb7, 0x11, 0x1f, 0xc4, 0x04, 0x3d, 0xa8, 0x68, 0x10, 0x4c,
	0x3a, 0x89, 0x0d, 0x2c, 0x6e, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xf9, 0x79, 0xbc, 0x9c,
	0x00, 0x00, 0x00,
}
