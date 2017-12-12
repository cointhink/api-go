// Code generated by protoc-gen-go.
// source: proto/market_prices.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MarketPrices struct {
	Prices []*MarketPrice `protobuf:"bytes,1,rep,name=Prices" json:"Prices,omitempty"`
}

func (m *MarketPrices) Reset()                    { *m = MarketPrices{} }
func (m *MarketPrices) String() string            { return proto1.CompactTextString(m) }
func (*MarketPrices) ProtoMessage()               {}
func (*MarketPrices) Descriptor() ([]byte, []int) { return fileDescriptor19, []int{0} }

func (m *MarketPrices) GetPrices() []*MarketPrice {
	if m != nil {
		return m.Prices
	}
	return nil
}

func init() {
	proto1.RegisterType((*MarketPrices)(nil), "proto.MarketPrices")
}

func init() { proto1.RegisterFile("proto/market_prices.proto", fileDescriptor19) }

var fileDescriptor19 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0x89, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0x2d, 0xd6,
	0x03, 0x8b, 0x09, 0xb1, 0x82, 0x29, 0x29, 0x09, 0x4c, 0x15, 0x10, 0x05, 0x4a, 0x56, 0x5c, 0x3c,
	0xbe, 0x60, 0xd1, 0x00, 0xb0, 0x36, 0x21, 0x2d, 0x2e, 0x36, 0x08, 0x4b, 0x82, 0x51, 0x81, 0x59,
	0x83, 0xdb, 0x48, 0x08, 0xa2, 0x4e, 0x0f, 0x49, 0x51, 0x10, 0x54, 0x45, 0x12, 0x1b, 0x58, 0xca,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xe9, 0x12, 0xab, 0x2e, 0x80, 0x00, 0x00, 0x00,
}
