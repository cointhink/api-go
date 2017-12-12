// Code generated by protoc-gen-go.
// source: proto/algolog_list.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AlgologList struct {
	AlgorunId string `protobuf:"bytes,1,opt,name=AlgorunId" json:"AlgorunId,omitempty"`
}

func (m *AlgologList) Reset()                    { *m = AlgologList{} }
func (m *AlgologList) String() string            { return proto1.CompactTextString(m) }
func (*AlgologList) ProtoMessage()               {}
func (*AlgologList) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *AlgologList) GetAlgorunId() string {
	if m != nil {
		return m.AlgorunId
	}
	return ""
}

func init() {
	proto1.RegisterType((*AlgologList)(nil), "proto.AlgologList")
}

func init() { proto1.RegisterFile("proto/algolog_list.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 86 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcc, 0x49, 0xcf, 0xcf, 0xc9, 0x4f, 0x8f, 0xcf, 0xc9, 0x2c, 0x2e, 0xd1, 0x03,
	0x0b, 0x09, 0xb1, 0x82, 0x29, 0x25, 0x6d, 0x2e, 0x6e, 0x47, 0x88, 0xa4, 0x4f, 0x66, 0x71, 0x89,
	0x90, 0x0c, 0x17, 0x27, 0x88, 0x5b, 0x54, 0x9a, 0xe7, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x84, 0x10, 0x48, 0x62, 0x03, 0xeb, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x27,
	0xab, 0x90, 0x56, 0x00, 0x00, 0x00,
}
