// Code generated by protoc-gen-go.
// source: proto/algorithm_detail.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AlgorithmDetail struct {
	AlgorithmId string `protobuf:"bytes,1,opt,name=algorithm_id,json=algorithmId" json:"algorithm_id,omitempty"`
}

func (m *AlgorithmDetail) Reset()                    { *m = AlgorithmDetail{} }
func (m *AlgorithmDetail) String() string            { return proto1.CompactTextString(m) }
func (*AlgorithmDetail) ProtoMessage()               {}
func (*AlgorithmDetail) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

func (m *AlgorithmDetail) GetAlgorithmId() string {
	if m != nil {
		return m.AlgorithmId
	}
	return ""
}

func init() {
	proto1.RegisterType((*AlgorithmDetail)(nil), "proto.AlgorithmDetail")
}

func init() { proto1.RegisterFile("proto/algorithm_detail.proto", fileDescriptor24) }

var fileDescriptor24 = []byte{
	// 89 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcc, 0x49, 0xcf, 0x2f, 0xca, 0x2c, 0xc9, 0xc8, 0x8d, 0x4f, 0x49, 0x2d, 0x49,
	0xcc, 0xcc, 0xd1, 0x03, 0x0b, 0x0b, 0xb1, 0x82, 0x29, 0x25, 0x13, 0x2e, 0x7e, 0x47, 0x98, 0x02,
	0x17, 0xb0, 0xbc, 0x90, 0x22, 0x17, 0x0f, 0x42, 0x4f, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x37, 0x5c, 0xcc, 0x33, 0x25, 0x89, 0x0d, 0xac, 0xd9, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x59, 0x53, 0x61, 0xeb, 0x63, 0x00, 0x00, 0x00,
}
