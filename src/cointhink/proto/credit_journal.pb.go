// Code generated by protoc-gen-go.
// source: proto/credit_journal.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreditJournal struct {
	Id               string  `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	AccountId        string  `protobuf:"bytes,2,opt,name=AccountId" json:"AccountId,omitempty"`
	ScheduleId       string  `protobuf:"bytes,3,opt,name=ScheduleId" json:"ScheduleId,omitempty"`
	Status           string  `protobuf:"bytes,4,opt,name=Status" json:"Status,omitempty"`
	StripeTx         string  `protobuf:"bytes,5,opt,name=StripeTx" json:"StripeTx,omitempty"`
	CreditAdjustment int32   `protobuf:"varint,6,opt,name=CreditAdjustment" json:"CreditAdjustment,omitempty"`
	TotalUsd         float32 `protobuf:"fixed32,7,opt,name=TotalUsd" json:"TotalUsd,omitempty"`
}

func (m *CreditJournal) Reset()                    { *m = CreditJournal{} }
func (m *CreditJournal) String() string            { return proto1.CompactTextString(m) }
func (*CreditJournal) ProtoMessage()               {}
func (*CreditJournal) Descriptor() ([]byte, []int) { return fileDescriptor33, []int{0} }

func (m *CreditJournal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreditJournal) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *CreditJournal) GetScheduleId() string {
	if m != nil {
		return m.ScheduleId
	}
	return ""
}

func (m *CreditJournal) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CreditJournal) GetStripeTx() string {
	if m != nil {
		return m.StripeTx
	}
	return ""
}

func (m *CreditJournal) GetCreditAdjustment() int32 {
	if m != nil {
		return m.CreditAdjustment
	}
	return 0
}

func (m *CreditJournal) GetTotalUsd() float32 {
	if m != nil {
		return m.TotalUsd
	}
	return 0
}

func init() {
	proto1.RegisterType((*CreditJournal)(nil), "proto.CreditJournal")
}

func init() { proto1.RegisterFile("proto/credit_journal.proto", fileDescriptor33) }

var fileDescriptor33 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x2e, 0x4a, 0x4d, 0xc9, 0x2c, 0x89, 0xcf, 0xca, 0x2f, 0x2d, 0xca, 0x4b, 0xcc,
	0xd1, 0x03, 0x0b, 0x0a, 0xb1, 0x82, 0x29, 0xa5, 0x9b, 0x8c, 0x5c, 0xbc, 0xce, 0x60, 0x79, 0x2f,
	0x88, 0xb4, 0x10, 0x1f, 0x17, 0x93, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x93,
	0x67, 0x8a, 0x90, 0x0c, 0x17, 0xa7, 0x63, 0x72, 0x72, 0x7e, 0x69, 0x5e, 0x89, 0x67, 0x8a, 0x04,
	0x13, 0x58, 0x18, 0x21, 0x20, 0x24, 0xc7, 0xc5, 0x15, 0x9c, 0x9c, 0x91, 0x9a, 0x52, 0x9a, 0x93,
	0xea, 0x99, 0x22, 0xc1, 0x0c, 0x96, 0x46, 0x12, 0x11, 0x12, 0xe3, 0x62, 0x0b, 0x2e, 0x49, 0x2c,
	0x29, 0x2d, 0x96, 0x60, 0x01, 0xcb, 0x41, 0x79, 0x42, 0x52, 0x5c, 0x1c, 0xc1, 0x25, 0x45, 0x99,
	0x05, 0xa9, 0x21, 0x15, 0x12, 0xac, 0x60, 0x19, 0x38, 0x5f, 0x48, 0x8b, 0x4b, 0x00, 0xe2, 0x24,
	0xc7, 0x94, 0xac, 0xd2, 0xe2, 0x92, 0xdc, 0xd4, 0xbc, 0x12, 0x09, 0x36, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0x0c, 0x71, 0x90, 0x39, 0x21, 0xf9, 0x25, 0x89, 0x39, 0xa1, 0xc5, 0x29, 0x12, 0xec, 0x0a,
	0x8c, 0x1a, 0x4c, 0x41, 0x70, 0x7e, 0x12, 0x1b, 0xd8, 0x8b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2a, 0x33, 0xb4, 0x3b, 0x07, 0x01, 0x00, 0x00,
}
