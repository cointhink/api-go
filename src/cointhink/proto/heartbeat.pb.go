// Code generated by protoc-gen-go.
// source: proto/heartbeat.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Heartbeat struct {
	Challenge string `protobuf:"bytes,1,opt,name=Challenge" json:"Challenge,omitempty"`
}

func (m *Heartbeat) Reset()                    { *m = Heartbeat{} }
func (m *Heartbeat) String() string            { return proto1.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()               {}
func (*Heartbeat) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Heartbeat) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

func init() {
	proto1.RegisterType((*Heartbeat)(nil), "proto.Heartbeat")
}

func init() { proto1.RegisterFile("proto/heartbeat.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 81 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0x4d, 0x2c, 0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0xd1, 0x03, 0xf3, 0x85, 0x58,
	0xc1, 0x94, 0x92, 0x26, 0x17, 0xa7, 0x07, 0x4c, 0x46, 0x48, 0x86, 0x8b, 0xd3, 0x39, 0x23, 0x31,
	0x27, 0x27, 0x35, 0x2f, 0x3d, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0x21, 0x90, 0xc4,
	0x06, 0xd6, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x88, 0xcb, 0x05, 0x51, 0x00, 0x00,
	0x00,
}