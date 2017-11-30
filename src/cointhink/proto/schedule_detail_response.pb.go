// Code generated by protoc-gen-go.
// source: proto/schedule_detail_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleDetailResponse struct {
	Ok       bool      `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Message  string    `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
	Schedule *Schedule `protobuf:"bytes,3,opt,name=Schedule" json:"Schedule,omitempty"`
}

func (m *ScheduleDetailResponse) Reset()                    { *m = ScheduleDetailResponse{} }
func (m *ScheduleDetailResponse) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleDetailResponse) ProtoMessage()               {}
func (*ScheduleDetailResponse) Descriptor() ([]byte, []int) { return fileDescriptor39, []int{0} }

func (m *ScheduleDetailResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ScheduleDetailResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ScheduleDetailResponse) GetSchedule() *Schedule {
	if m != nil {
		return m.Schedule
	}
	return nil
}

func init() {
	proto1.RegisterType((*ScheduleDetailResponse)(nil), "proto.ScheduleDetailResponse")
}

func init() { proto1.RegisterFile("proto/schedule_detail_response.proto", fileDescriptor39) }

var fileDescriptor39 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0x4f, 0x49, 0x2d, 0x49, 0xcc,
	0xcc, 0x89, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0x4b, 0x0b, 0xb1, 0x82,
	0x29, 0x29, 0x11, 0x54, 0xc5, 0x10, 0x49, 0xa5, 0x7c, 0x2e, 0xb1, 0x60, 0xa8, 0x88, 0x0b, 0x58,
	0x77, 0x10, 0x54, 0xb3, 0x10, 0x1f, 0x17, 0x93, 0x7f, 0xb6, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x47,
	0x10, 0x93, 0x7f, 0xb6, 0x90, 0x04, 0x17, 0xbb, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0xa4, 0xcd, 0xc5, 0x01, 0x33, 0x43, 0x82, 0x59,
	0x81, 0x51, 0x83, 0xdb, 0x88, 0x1f, 0x62, 0xba, 0x1e, 0x4c, 0x38, 0x08, 0xae, 0x20, 0x89, 0x0d,
	0x2c, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x20, 0x1d, 0xe3, 0xbc, 0x00, 0x00, 0x00,
}