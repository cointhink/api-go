// Code generated by protoc-gen-go.
// source: proto/schedule_list_partial.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleListPartial struct {
	ListId      string       `protobuf:"bytes,1,opt,name=ListId" json:"ListId,omitempty"`
	ScheduleRun *ScheduleRun `protobuf:"bytes,2,opt,name=ScheduleRun" json:"ScheduleRun,omitempty"`
}

func (m *ScheduleListPartial) Reset()                    { *m = ScheduleListPartial{} }
func (m *ScheduleListPartial) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleListPartial) ProtoMessage()               {}
func (*ScheduleListPartial) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *ScheduleListPartial) GetListId() string {
	if m != nil {
		return m.ListId
	}
	return ""
}

func (m *ScheduleListPartial) GetScheduleRun() *ScheduleRun {
	if m != nil {
		return m.ScheduleRun
	}
	return nil
}

func init() {
	proto1.RegisterType((*ScheduleListPartial)(nil), "proto.ScheduleListPartial")
}

func init() { proto1.RegisterFile("proto/schedule_list_partial.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0xcf, 0xc9, 0x2c, 0x2e, 0x89,
	0x2f, 0x48, 0x2c, 0x2a, 0xc9, 0x4c, 0xcc, 0xd1, 0x03, 0xcb, 0x09, 0xb1, 0x82, 0x29, 0x29, 0x09,
	0x34, 0x95, 0x45, 0xa5, 0x79, 0x10, 0x05, 0x4a, 0xc9, 0x5c, 0xc2, 0xc1, 0x50, 0x51, 0x9f, 0xcc,
	0xe2, 0x92, 0x00, 0x88, 0x6e, 0x21, 0x31, 0x2e, 0x36, 0x10, 0xd7, 0x33, 0x45, 0x82, 0x51, 0x81,
	0x51, 0x83, 0x33, 0x08, 0xca, 0x13, 0x32, 0xe1, 0xe2, 0x86, 0x29, 0x0f, 0x2a, 0xcd, 0x93, 0x60,
	0x52, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0x82, 0x98, 0xa5, 0x87, 0x24, 0x13, 0x84, 0xac, 0x2c, 0x89,
	0x0d, 0x2c, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x3c, 0xd7, 0xcd, 0xb1, 0x00, 0x00,
	0x00,
}
