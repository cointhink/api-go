// Code generated by protoc-gen-go.
// source: proto/schedule_delete_response.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleDeleteResponse struct {
	Ok      bool   `protobuf:"varint,1,opt,name=Ok" json:"Ok,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message" json:"Message,omitempty"`
}

func (m *ScheduleDeleteResponse) Reset()                    { *m = ScheduleDeleteResponse{} }
func (m *ScheduleDeleteResponse) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleDeleteResponse) ProtoMessage()               {}
func (*ScheduleDeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor31, []int{0} }

func (m *ScheduleDeleteResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ScheduleDeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto1.RegisterType((*ScheduleDeleteResponse)(nil), "proto.ScheduleDeleteResponse")
}

func init() { proto1.RegisterFile("proto/schedule_delete_response.proto", fileDescriptor31) }

var fileDescriptor31 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0x4f, 0x49, 0xcd, 0x49, 0x2d,
	0x49, 0x8d, 0x2f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x03, 0x4b, 0x0b, 0xb1, 0x82,
	0x29, 0x25, 0x27, 0x2e, 0xb1, 0x60, 0xa8, 0x42, 0x17, 0xb0, 0xba, 0x20, 0xa8, 0x32, 0x21, 0x3e,
	0x2e, 0x26, 0xff, 0x6c, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x26, 0xff, 0x6c, 0x21, 0x09,
	0x2e, 0x76, 0xdf, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20,
	0x18, 0x37, 0x89, 0x0d, 0x6c, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x77, 0x7f, 0x2e, 0xcb,
	0x79, 0x00, 0x00, 0x00,
}
