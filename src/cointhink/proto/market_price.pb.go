// Code generated by protoc-gen-go.
// source: proto/market_price.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MarketPrice struct {
	Exchange   string `protobuf:"bytes,1,opt,name=Exchange" json:"Exchange,omitempty"`
	Market     string `protobuf:"bytes,2,opt,name=Market" json:"Market,omitempty"`
	Amount     string `protobuf:"bytes,3,opt,name=Amount" json:"Amount,omitempty"`
	Currency   string `protobuf:"bytes,4,opt,name=Currency" json:"Currency,omitempty"`
	ReceivedAt string `protobuf:"bytes,5,opt,name=ReceivedAt" json:"ReceivedAt,omitempty"`
}

func (m *MarketPrice) Reset()                    { *m = MarketPrice{} }
func (m *MarketPrice) String() string            { return proto1.CompactTextString(m) }
func (*MarketPrice) ProtoMessage()               {}
func (*MarketPrice) Descriptor() ([]byte, []int) { return fileDescriptor29, []int{0} }

func (m *MarketPrice) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *MarketPrice) GetMarket() string {
	if m != nil {
		return m.Market
	}
	return ""
}

func (m *MarketPrice) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *MarketPrice) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *MarketPrice) GetReceivedAt() string {
	if m != nil {
		return m.ReceivedAt
	}
	return ""
}

func init() {
	proto1.RegisterType((*MarketPrice)(nil), "proto.MarketPrice")
}

func init() { proto1.RegisterFile("proto/market_price.proto", fileDescriptor29) }

var fileDescriptor29 = []byte{
	// 148 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0xca, 0x4e, 0x2d, 0x89, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0xd5, 0x03,
	0x0b, 0x09, 0xb1, 0x82, 0x29, 0xa5, 0xa9, 0x8c, 0x5c, 0xdc, 0xbe, 0x60, 0xd9, 0x00, 0x90, 0xa4,
	0x90, 0x14, 0x17, 0x87, 0x6b, 0x45, 0x72, 0x46, 0x62, 0x5e, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x9c, 0x2f, 0x24, 0xc6, 0xc5, 0x06, 0x51, 0x2a, 0xc1, 0x04, 0x96, 0x81, 0xf2,
	0x40, 0xe2, 0x8e, 0xb9, 0xf9, 0xa5, 0x79, 0x25, 0x12, 0xcc, 0x10, 0x71, 0x08, 0x0f, 0x64, 0x96,
	0x73, 0x69, 0x51, 0x51, 0x6a, 0x5e, 0x72, 0xa5, 0x04, 0x0b, 0xc4, 0x2c, 0x18, 0x5f, 0x48, 0x8e,
	0x8b, 0x2b, 0x28, 0x35, 0x39, 0x35, 0xb3, 0x2c, 0x35, 0xc5, 0xb1, 0x44, 0x82, 0x15, 0x2c, 0x8b,
	0x24, 0x92, 0xc4, 0x06, 0x76, 0x9e, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xcf, 0x4d, 0x35,
	0xc1, 0x00, 0x00, 0x00,
}
