// Code generated by protoc-gen-go.
// source: proto/schedule_create.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleCreate struct {
	Schedule *Schedule `protobuf:"bytes,2,opt,name=Schedule" json:"Schedule,omitempty"`
}

func (m *ScheduleCreate) Reset()                    { *m = ScheduleCreate{} }
func (m *ScheduleCreate) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleCreate) ProtoMessage()               {}
func (*ScheduleCreate) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *ScheduleCreate) GetSchedule() *Schedule {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func init() {
	proto1.RegisterType((*ScheduleCreate)(nil), "proto.ScheduleCreate")
}

func init() { proto1.RegisterFile("proto/schedule_create.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 96 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0x4f, 0x2e, 0x4a, 0x4d, 0x2c,
	0x49, 0xd5, 0x03, 0x8b, 0x0a, 0xb1, 0x82, 0x29, 0x29, 0x11, 0x54, 0x35, 0x10, 0x49, 0x25, 0x5b,
	0x2e, 0xbe, 0x60, 0xa8, 0x88, 0x33, 0x58, 0x93, 0x90, 0x36, 0x17, 0x07, 0x4c, 0x44, 0x82, 0x49,
	0x81, 0x51, 0x83, 0xdb, 0x88, 0x1f, 0xa2, 0x56, 0x0f, 0x26, 0x1c, 0x04, 0x57, 0x90, 0xc4, 0x06,
	0x96, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x4c, 0x57, 0xe2, 0x81, 0x00, 0x00, 0x00,
}
