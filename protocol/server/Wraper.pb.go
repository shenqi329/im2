// Code generated by protoc-gen-go.
// source: Wraper.proto
// DO NOT EDIT!

/*
Package bean is a generated protocol buffer package.

It is generated from these files:
	Wraper.proto

It has these top-level messages:
	WraperMessage
*/
package bean

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WraperMessage struct {
	ConnId     uint64 `protobuf:"varint,1,opt,name=ConnId" json:"ConnId,omitempty"`
	Message    []byte `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	IsLoginIn  bool   `protobuf:"varint,3,opt,name=IsLoginIn" json:"IsLoginIn,omitempty"`
	IsLoginOut bool   `protobuf:"varint,4,opt,name=IsLoginOut" json:"IsLoginOut,omitempty"`
	Token      string `protobuf:"bytes,5,opt,name=Token" json:"Token,omitempty"`
	UserId     string `protobuf:"bytes,6,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *WraperMessage) Reset()                    { *m = WraperMessage{} }
func (m *WraperMessage) String() string            { return proto.CompactTextString(m) }
func (*WraperMessage) ProtoMessage()               {}
func (*WraperMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WraperMessage) GetConnId() uint64 {
	if m != nil {
		return m.ConnId
	}
	return 0
}

func (m *WraperMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *WraperMessage) GetIsLoginIn() bool {
	if m != nil {
		return m.IsLoginIn
	}
	return false
}

func (m *WraperMessage) GetIsLoginOut() bool {
	if m != nil {
		return m.IsLoginOut
	}
	return false
}

func (m *WraperMessage) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *WraperMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*WraperMessage)(nil), "bean.WraperMessage")
}

func init() { proto.RegisterFile("Wraper.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x09, 0x2f, 0x4a, 0x2c,
	0x48, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x4a, 0x4d, 0xcc, 0x53, 0x5a,
	0xcb, 0xc8, 0xc5, 0x0b, 0x11, 0xf6, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x12, 0xe3, 0x62,
	0x73, 0xce, 0xcf, 0xcb, 0xf3, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x82, 0xf2, 0x84,
	0x24, 0xb8, 0xd8, 0xa1, 0x4a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x21, 0x19,
	0x2e, 0x4e, 0xcf, 0x62, 0x9f, 0xfc, 0xf4, 0xcc, 0x3c, 0xcf, 0x3c, 0x09, 0x66, 0x05, 0x46, 0x0d,
	0x8e, 0x20, 0x84, 0x80, 0x90, 0x1c, 0x17, 0x17, 0x94, 0xe3, 0x5f, 0x5a, 0x22, 0xc1, 0x02, 0x96,
	0x46, 0x12, 0x11, 0x12, 0xe1, 0x62, 0x0d, 0xc9, 0xcf, 0x4e, 0xcd, 0x93, 0x60, 0x55, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x70, 0x40, 0xae, 0x08, 0x2d, 0x4e, 0x2d, 0xf2, 0x4c, 0x91, 0x60, 0x03, 0x0b,
	0x43, 0x79, 0x49, 0x6c, 0x60, 0xc7, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x98, 0x62, 0x09,
	0x20, 0xcc, 0x00, 0x00, 0x00,
}
