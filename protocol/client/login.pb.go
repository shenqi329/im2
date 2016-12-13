// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

package client

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeviceLoginRequest struct {
	Rid      uint64 `protobuf:"varint,1,opt,name=Rid" json:"Rid,omitempty"`
	Token    string `protobuf:"bytes,11,opt,name=Token" json:"Token,omitempty"`
	AppId    string `protobuf:"bytes,12,opt,name=AppId" json:"AppId,omitempty"`
	DeviceId string `protobuf:"bytes,13,opt,name=DeviceId" json:"DeviceId,omitempty"`
	Platform string `protobuf:"bytes,14,opt,name=Platform" json:"Platform,omitempty"`
	UserId   string `protobuf:"bytes,15,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *DeviceLoginRequest) Reset()                    { *m = DeviceLoginRequest{} }
func (m *DeviceLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceLoginRequest) ProtoMessage()               {}
func (*DeviceLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DeviceLoginRequest) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *DeviceLoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeviceLoginRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeviceLoginRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DeviceLoginRequest) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *DeviceLoginRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

// messageType = 14
type DeviceLoginResponse struct {
	Rid  uint64 `protobuf:"varint,1,opt,name=rid" json:"rid,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
}

func (m *DeviceLoginResponse) Reset()                    { *m = DeviceLoginResponse{} }
func (m *DeviceLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*DeviceLoginResponse) ProtoMessage()               {}
func (*DeviceLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *DeviceLoginResponse) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *DeviceLoginResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DeviceLoginResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceLoginRequest)(nil), "client.DeviceLoginRequest")
	proto.RegisterType((*DeviceLoginResponse)(nil), "client.DeviceLoginResponse")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0x89, 0xad, 0x45, 0xa7, 0xfe, 0x31, 0x8a, 0x0c, 0xae, 0x4a, 0x57, 0x5d, 0xb9, 0xf1,
	0x09, 0x04, 0x37, 0x05, 0x41, 0x09, 0xfa, 0x00, 0x9a, 0x8c, 0x12, 0xac, 0x49, 0x4c, 0xa2, 0x0f,
	0xe4, 0x93, 0x5e, 0x92, 0xdc, 0x1f, 0xee, 0xee, 0x7c, 0xe7, 0x23, 0xe1, 0x30, 0xd0, 0x2f, 0xee,
	0xd3, 0xd8, 0x5b, 0x1f, 0x5c, 0x72, 0xd8, 0xa9, 0xc5, 0xb0, 0x4d, 0xe3, 0xbf, 0x00, 0x7c, 0xe0,
	0x3f, 0xa3, 0xf8, 0x31, 0x5b, 0xc9, 0x3f, 0xbf, 0x1c, 0x13, 0x5e, 0x40, 0x23, 0x8d, 0x26, 0x31,
	0x88, 0xa9, 0x95, 0x39, 0xe2, 0x15, 0x1c, 0xbe, 0xb8, 0x2f, 0xb6, 0xd4, 0x0f, 0x62, 0x3a, 0x96,
	0x15, 0x72, 0x7b, 0xef, 0xfd, 0xac, 0xe9, 0xa4, 0xb6, 0x05, 0xf0, 0x06, 0x8e, 0xea, 0x9f, 0xb3,
	0xa6, 0xd3, 0x22, 0xb6, 0x9c, 0xdd, 0xf3, 0xf2, 0x96, 0x3e, 0x5c, 0xf8, 0xa6, 0xb3, 0xea, 0x36,
	0x8c, 0xd7, 0xd0, 0xbd, 0x46, 0x0e, 0xb3, 0xa6, 0xf3, 0x62, 0xd6, 0x34, 0x3e, 0xc1, 0xe5, 0xde,
	0xc6, 0xe8, 0x9d, 0x8d, 0x9c, 0x47, 0x86, 0xdd, 0xc8, 0x60, 0x34, 0x22, 0xb4, 0xca, 0x69, 0xa6,
	0x83, 0xf2, 0xbc, 0xe4, 0xdc, 0x69, 0x8e, 0x8a, 0x9a, 0xda, 0xe5, 0xfc, 0xde, 0x95, 0x23, 0xdc,
	0xad, 0x02, 0x00, 0x00, 0xff, 0xff, 0x76, 0x4b, 0x76, 0x3a, 0x13, 0x01, 0x00, 0x00,
}
