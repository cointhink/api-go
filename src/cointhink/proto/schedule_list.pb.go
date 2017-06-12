// Code generated by protoc-gen-go.
// source: proto/schedule_list.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScheduleList struct {
	FilterAccountId string `protobuf:"bytes,1,opt,name=filter_account_id,json=filterAccountId" json:"filter_account_id,omitempty"`
}

func (m *ScheduleList) Reset()                    { *m = ScheduleList{} }
func (m *ScheduleList) String() string            { return proto1.CompactTextString(m) }
func (*ScheduleList) ProtoMessage()               {}
func (*ScheduleList) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *ScheduleList) GetFilterAccountId() string {
	if m != nil {
		return m.FilterAccountId
	}
	return ""
}

func init() {
	proto1.RegisterType((*ScheduleList)(nil), "proto.ScheduleList")
}

func init() { proto1.RegisterFile("proto/schedule_list.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 103 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x8d, 0xcf, 0xc9, 0x2c, 0x2e, 0xd1,
	0x03, 0x8b, 0x09, 0xb1, 0x82, 0x29, 0x25, 0x2b, 0x2e, 0x9e, 0x60, 0xa8, 0xac, 0x4f, 0x66, 0x71,
	0x89, 0x90, 0x16, 0x97, 0x60, 0x5a, 0x66, 0x4e, 0x49, 0x6a, 0x51, 0x7c, 0x62, 0x72, 0x72, 0x7e,
	0x69, 0x5e, 0x49, 0x7c, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x3f, 0x44, 0xc2,
	0x11, 0x22, 0xee, 0x99, 0x92, 0xc4, 0x06, 0x36, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x97,
	0x4c, 0xa7, 0x1f, 0x66, 0x00, 0x00, 0x00,
}
