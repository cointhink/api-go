// Code generated by protoc-gen-go.
// source: proto/schedule_detail.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleDetail struct {
	ScheduleId string `protobuf:"bytes,1,opt,name=schedule_id,json=scheduleId" json:"schedule_id,omitempty"`
}

func (m *ScheduleDetail) Reset()                    { *m = ScheduleDetail{} }
func (m *ScheduleDetail) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleDetail) ProtoMessage()               {}
func (*ScheduleDetail) Descriptor() ([]byte, []int) { return fileDescriptor23, []int{0} }

func (m *ScheduleDetail) GetScheduleId() string {
	if m != nil {
		return m.ScheduleId
	}
	return ""
}

func init() {
	proto1.RegisterType((*ScheduleDetail)(nil), "proto.ScheduleDetail")
}

func init() { proto1.RegisterFile("proto/schedule_detail.proto", fileDescriptor23) }

var fileDescriptor23 = []byte{
	// 88 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0x4f, 0x49, 0x2d, 0x49, 0xcc,
	0xcc, 0xd1, 0x03, 0x8b, 0x0a, 0xb1, 0x82, 0x29, 0x25, 0x43, 0x2e, 0xbe, 0x60, 0xa8, 0xbc, 0x0b,
	0x58, 0x5a, 0x48, 0x9e, 0x8b, 0x1b, 0xae, 0x23, 0x33, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x88, 0x0b, 0x26, 0xe4, 0x99, 0x92, 0xc4, 0x06, 0xd6, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0xbc, 0x44, 0x5b, 0x5f, 0x00, 0x00, 0x00,
}
