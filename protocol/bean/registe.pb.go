// Code generated by protoc-gen-go.
// source: registe.proto
// DO NOT EDIT!

package bean

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// messageType = 11
type DeviceRegisteRequest struct {
	Rid      uint64 `protobuf:"varint,1,opt,name=rid" json:"rid,omitempty"`
	SsoToken string `protobuf:"bytes,11,opt,name=ssoToken" json:"ssoToken,omitempty"`
	AppId    string `protobuf:"bytes,12,opt,name=appId" json:"appId,omitempty"`
	DeviceId string `protobuf:"bytes,13,opt,name=deviceId" json:"deviceId,omitempty"`
	Platform string `protobuf:"bytes,14,opt,name=platform" json:"platform,omitempty"`
}

func (m *DeviceRegisteRequest) Reset()                    { *m = DeviceRegisteRequest{} }
func (m *DeviceRegisteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceRegisteRequest) ProtoMessage()               {}
func (*DeviceRegisteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *DeviceRegisteRequest) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *DeviceRegisteRequest) GetSsoToken() string {
	if m != nil {
		return m.SsoToken
	}
	return ""
}

func (m *DeviceRegisteRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeviceRegisteRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DeviceRegisteRequest) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

// messageType = 12
type DeviceRegisteResponse struct {
	Rid   uint64 `protobuf:"varint,1,opt,name=rid" json:"rid,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Desc  string `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	Token string `protobuf:"bytes,11,opt,name=token" json:"token,omitempty"`
}

func (m *DeviceRegisteResponse) Reset()                    { *m = DeviceRegisteResponse{} }
func (m *DeviceRegisteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeviceRegisteResponse) ProtoMessage()               {}
func (*DeviceRegisteResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *DeviceRegisteResponse) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *DeviceRegisteResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DeviceRegisteResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *DeviceRegisteResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceRegisteRequest)(nil), "bean.DeviceRegisteRequest")
	proto.RegisterType((*DeviceRegisteResponse)(nil), "bean.DeviceRegisteResponse")
}

func init() { proto.RegisterFile("registe.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x41, 0xaa, 0xc2, 0x30,
	0x10, 0x86, 0xc9, 0x6b, 0xdf, 0xe3, 0x19, 0xad, 0x48, 0xa8, 0x10, 0x5c, 0x95, 0xae, 0xba, 0x72,
	0xe3, 0x15, 0xdc, 0x74, 0x1b, 0xbc, 0x40, 0xdb, 0x8c, 0xa5, 0xa8, 0x9d, 0x98, 0x89, 0x1e, 0xc4,
	0x13, 0x4b, 0x12, 0x2a, 0x22, 0xee, 0xfe, 0x6f, 0x7e, 0x66, 0xf8, 0x18, 0x9e, 0x59, 0xe8, 0x07,
	0x72, 0xb0, 0x35, 0x16, 0x1d, 0x8a, 0xb4, 0x85, 0x66, 0x2c, 0x1f, 0x8c, 0xe7, 0x7b, 0xb8, 0x0f,
	0x1d, 0xa8, 0xd8, 0x2a, 0xb8, 0xde, 0x80, 0x9c, 0x58, 0xf1, 0xc4, 0x0e, 0x5a, 0xb2, 0x82, 0x55,
	0xa9, 0xf2, 0x51, 0x6c, 0xf8, 0x3f, 0x11, 0x1e, 0xf0, 0x04, 0xa3, 0x9c, 0x17, 0xac, 0x9a, 0xa9,
	0x17, 0x8b, 0x9c, 0xff, 0x36, 0xc6, 0xd4, 0x5a, 0x2e, 0x42, 0x11, 0xc1, 0x6f, 0xe8, 0x70, 0xbb,
	0xd6, 0x32, 0x8b, 0x1b, 0x13, 0xfb, 0xce, 0x9c, 0x1b, 0x77, 0x44, 0x7b, 0x91, 0xcb, 0xd8, 0x4d,
	0x5c, 0xf6, 0x7c, 0xfd, 0xe1, 0x44, 0x06, 0x47, 0x82, 0x2f, 0x52, 0x82, 0xa7, 0x1d, 0x6a, 0x90,
	0x3f, 0xe1, 0x44, 0xc8, 0x7e, 0xa6, 0x81, 0x3a, 0x99, 0xc4, 0x99, 0xcf, 0x5e, 0xd0, 0xbd, 0x99,
	0x47, 0x68, 0xff, 0xc2, 0x2b, 0x76, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x0f, 0x91, 0x33,
	0x1b, 0x01, 0x00, 0x00,
}
