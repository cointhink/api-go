// Code generated by protoc-gen-go.
// source: proto/signup_form_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SignupFormResponse_Reasons int32

const (
	SignupFormResponse_EMAIL_ALERT    SignupFormResponse_Reasons = 0
	SignupFormResponse_USERNAME_ALERT SignupFormResponse_Reasons = 1
)

var SignupFormResponse_Reasons_name = map[int32]string{
	0: "EMAIL_ALERT",
	1: "USERNAME_ALERT",
}
var SignupFormResponse_Reasons_value = map[string]int32{
	"EMAIL_ALERT":    0,
	"USERNAME_ALERT": 1,
}

func (x SignupFormResponse_Reasons) String() string {
	return proto1.EnumName(SignupFormResponse_Reasons_name, int32(x))
}
func (SignupFormResponse_Reasons) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor7, []int{0, 0}
}

type SignupFormResponse struct {
	Ok      bool                       `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Token   string                     `protobuf:"bytes,2,opt,name=Token" json:"Token,omitempty"`
	Reason  SignupFormResponse_Reasons `protobuf:"varint,3,opt,name=Reason,enum=proto.SignupFormResponse_Reasons" json:"Reason,omitempty"`
	Message string                     `protobuf:"bytes,4,opt,name=Message" json:"Message,omitempty"`
}

func (m *SignupFormResponse) Reset()                    { *m = SignupFormResponse{} }
func (m *SignupFormResponse) String() string            { return proto1.CompactTextString(m) }
func (*SignupFormResponse) ProtoMessage()               {}
func (*SignupFormResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *SignupFormResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *SignupFormResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SignupFormResponse) GetReason() SignupFormResponse_Reasons {
	if m != nil {
		return m.Reason
	}
	return SignupFormResponse_EMAIL_ALERT
}

func (m *SignupFormResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*SignupFormResponse)(nil), "proto.SignupFormResponse")
	proto1.RegisterEnum("proto.SignupFormResponse_Reasons", SignupFormResponse_Reasons_name, SignupFormResponse_Reasons_value)
}

func init() { proto1.RegisterFile("proto/signup_form_response.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0x4c, 0xcf, 0x2b, 0x2d, 0x88, 0x4f, 0xcb, 0x2f, 0xca, 0x8d, 0x2f, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0x4b, 0x09, 0xb1, 0x82, 0x29, 0xa5, 0xfd, 0x8c,
	0x5c, 0x42, 0xc1, 0x60, 0x55, 0x6e, 0xf9, 0x45, 0xb9, 0x41, 0x50, 0x35, 0x42, 0x7c, 0x5c, 0x4c,
	0xfe, 0xd9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x4c, 0xfe, 0xd9, 0x42, 0x22, 0x5c, 0xac,
	0x21, 0xf9, 0xd9, 0xa9, 0x79, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x25,
	0x17, 0x5b, 0x50, 0x6a, 0x62, 0x71, 0x7e, 0x9e, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x22,
	0xc4, 0x6c, 0x3d, 0x4c, 0x03, 0xf5, 0x20, 0xea, 0x8a, 0x83, 0xa0, 0x1a, 0x84, 0x24, 0xb8, 0xd8,
	0x7d, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x25, 0x58, 0xc0, 0x46, 0xc2, 0xb8, 0x4a, 0x7a, 0x5c,
	0xec, 0x50, 0xc5, 0x42, 0xfc, 0x5c, 0xdc, 0xae, 0xbe, 0x8e, 0x9e, 0x3e, 0xf1, 0x8e, 0x3e, 0xae,
	0x41, 0x21, 0x02, 0x0c, 0x42, 0x42, 0x5c, 0x7c, 0xa1, 0xc1, 0xae, 0x41, 0x7e, 0x8e, 0xbe, 0xae,
	0x50, 0x31, 0xc6, 0x24, 0x36, 0xb0, 0x9d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0xba,
	0x5e, 0x28, 0xf3, 0x00, 0x00, 0x00,
}
