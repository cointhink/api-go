// Code generated by protoc-gen-go.
// source: proto/schedule_list_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleListResponse struct {
	Ok        bool           `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Message   string         `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Schedules []*ScheduleRun `protobuf:"bytes,3,rep,name=Schedules" json:"Schedules,omitempty"`
}

func (m *ScheduleListResponse) Reset()                    { *m = ScheduleListResponse{} }
func (m *ScheduleListResponse) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleListResponse) ProtoMessage()               {}
func (*ScheduleListResponse) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *ScheduleListResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ScheduleListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ScheduleListResponse) GetSchedules() []*ScheduleRun {
	if m != nil {
		return m.Schedules
	}
	return nil
}

func init() {
	proto1.RegisterType((*ScheduleListResponse)(nil), "proto.ScheduleListResponse")
}

func init() { proto1.RegisterFile("proto/schedule_list_response.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0xcf, 0xc9, 0x2c, 0x2e, 0x89,
	0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0x4b, 0x0a, 0xb1, 0x82, 0x29, 0x29,
	0x09, 0x34, 0xa5, 0x45, 0xa5, 0x79, 0x10, 0x05, 0x4a, 0x45, 0x5c, 0x22, 0xc1, 0x50, 0x51, 0x9f,
	0xcc, 0xe2, 0x92, 0x20, 0xa8, 0x76, 0x21, 0x3e, 0x2e, 0x26, 0xff, 0x6c, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x8e, 0x20, 0x26, 0xff, 0x6c, 0x21, 0x09, 0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4,
	0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0xc8, 0x80, 0x8b, 0x13, 0x66, 0x42,
	0xb1, 0x04, 0xb3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x10, 0xc4, 0x70, 0x3d, 0x98, 0x78, 0x50, 0x69,
	0x5e, 0x10, 0x42, 0x51, 0x12, 0x1b, 0x58, 0xd6, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xd5,
	0x28, 0x84, 0xc1, 0x00, 0x00, 0x00,
}
