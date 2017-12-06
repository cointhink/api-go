// Code generated by protoc-gen-go.
// source: proto/lambda.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Lambda struct {
	Token   string               `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
	Method  string               `protobuf:"bytes,2,opt,name=Method" json:"Method,omitempty"`
	Object  *google_protobuf.Any `protobuf:"bytes,3,opt,name=Object" json:"Object,omitempty"`
	StateIn string               `protobuf:"bytes,4,opt,name=StateIn" json:"StateIn,omitempty"`
}

func (m *Lambda) Reset()                    { *m = Lambda{} }
func (m *Lambda) String() string            { return proto1.CompactTextString(m) }
func (*Lambda) ProtoMessage()               {}
func (*Lambda) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *Lambda) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Lambda) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Lambda) GetObject() *google_protobuf.Any {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *Lambda) GetStateIn() string {
	if m != nil {
		return m.StateIn
	}
	return ""
}

func init() {
	proto1.RegisterType((*Lambda)(nil), "proto.Lambda")
}

func init() { proto1.RegisterFile("proto/lambda.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x49, 0xcc, 0x4d, 0x4a, 0x49, 0xd4, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x94, 0x94,
	0x64, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x98, 0x97, 0x54, 0x9a, 0xa6, 0x9f, 0x98, 0x57,
	0x09, 0x51, 0xa1, 0x54, 0xc7, 0xc5, 0xe6, 0x03, 0xd6, 0x21, 0x24, 0xc2, 0xc5, 0x1a, 0x92, 0x9f,
	0x9d, 0x9a, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x89, 0x71, 0xb1, 0xf9,
	0xa6, 0x96, 0x64, 0xe4, 0xa7, 0x48, 0x30, 0x81, 0x85, 0xa1, 0x3c, 0x21, 0x1d, 0x2e, 0x36, 0xff,
	0xa4, 0xac, 0xd4, 0xe4, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x11, 0x3d, 0x88, 0x1d,
	0x7a, 0x30, 0x3b, 0xf4, 0x1c, 0xf3, 0x2a, 0x83, 0xa0, 0x6a, 0x84, 0x24, 0xb8, 0xd8, 0x83, 0x4b,
	0x12, 0x4b, 0x52, 0x3d, 0xf3, 0x24, 0x58, 0xc0, 0xc6, 0xc0, 0xb8, 0x49, 0x6c, 0x60, 0xf5, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x45, 0x16, 0x64, 0xbe, 0x00, 0x00, 0x00,
}
