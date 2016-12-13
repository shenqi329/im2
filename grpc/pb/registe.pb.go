// Code generated by protoc-gen-go.
// source: registe.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeviceRegisteRequest struct {
	Rid      uint64 `protobuf:"varint,1,opt,name=Rid,json=rid" json:"Rid,omitempty"`
	SsoToken string `protobuf:"bytes,11,opt,name=SsoToken,json=ssoToken" json:"SsoToken,omitempty"`
	AppId    string `protobuf:"bytes,12,opt,name=AppId,json=appId" json:"AppId,omitempty"`
	DeviceId string `protobuf:"bytes,13,opt,name=DeviceId,json=deviceId" json:"DeviceId,omitempty"`
	Platform string `protobuf:"bytes,14,opt,name=Platform,json=platform" json:"Platform,omitempty"`
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

type DeviceRegisteResponse struct {
	Rid   uint64 `protobuf:"varint,1,opt,name=Rid,json=rid" json:"Rid,omitempty"`
	Code  string `protobuf:"bytes,2,opt,name=Code,json=code" json:"Code,omitempty"`
	Desc  string `protobuf:"bytes,3,opt,name=Desc,json=desc" json:"Desc,omitempty"`
	Token string `protobuf:"bytes,11,opt,name=Token,json=token" json:"Token,omitempty"`
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
	proto.RegisterType((*DeviceRegisteRequest)(nil), "pb.DeviceRegisteRequest")
	proto.RegisterType((*DeviceRegisteResponse)(nil), "pb.DeviceRegisteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Registe service

type RegisteClient interface {
	// Sends a greeting
	Login(ctx context.Context, in *DeviceRegisteRequest, opts ...grpc.CallOption) (*DeviceRegisteResponse, error)
}

type registeClient struct {
	cc *grpc.ClientConn
}

func NewRegisteClient(cc *grpc.ClientConn) RegisteClient {
	return &registeClient{cc}
}

func (c *registeClient) Login(ctx context.Context, in *DeviceRegisteRequest, opts ...grpc.CallOption) (*DeviceRegisteResponse, error) {
	out := new(DeviceRegisteResponse)
	err := grpc.Invoke(ctx, "/pb.Registe/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Registe service

type RegisteServer interface {
	// Sends a greeting
	Login(context.Context, *DeviceRegisteRequest) (*DeviceRegisteResponse, error)
}

func RegisterRegisteServer(s *grpc.Server, srv RegisteServer) {
	s.RegisterService(&_Registe_serviceDesc, srv)
}

func _Registe_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRegisteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisteServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Registe/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisteServer).Login(ctx, req.(*DeviceRegisteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Registe_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Registe",
	HandlerType: (*RegisteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Registe_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registe.proto",
}

func init() { proto.RegisterFile("registe.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x40, 0xed, 0xb6, 0xd5, 0x75, 0x74, 0x45, 0xc2, 0x0a, 0x71, 0x4f, 0x4b, 0x4f, 0x7b, 0xea,
	0x41, 0xef, 0x82, 0xb8, 0x97, 0x82, 0x07, 0x89, 0xfe, 0xc0, 0xb6, 0x33, 0x96, 0xa0, 0x76, 0x62,
	0x27, 0xfa, 0x21, 0x7e, 0xb1, 0xa4, 0xc9, 0x41, 0xa4, 0xc7, 0x97, 0xc7, 0x4c, 0x1e, 0x03, 0xab,
	0x91, 0x7a, 0x2b, 0x9e, 0x6a, 0x37, 0xb2, 0x67, 0xb5, 0x70, 0x6d, 0xf5, 0x93, 0xc1, 0x7a, 0x4f,
	0xdf, 0xb6, 0x23, 0x13, 0x9d, 0xa1, 0xcf, 0x2f, 0x12, 0xaf, 0x2e, 0x21, 0x37, 0x16, 0x75, 0xb6,
	0xcd, 0x76, 0x85, 0xc9, 0x47, 0x8b, 0x6a, 0x03, 0xcb, 0x67, 0xe1, 0x17, 0x7e, 0xa3, 0x41, 0x9f,
	0x6d, 0xb3, 0xdd, 0xa9, 0x59, 0x4a, 0x62, 0xb5, 0x86, 0xf2, 0xde, 0xb9, 0x06, 0xf5, 0xf9, 0x24,
	0xca, 0x43, 0x80, 0x30, 0x11, 0x77, 0x37, 0xa8, 0x57, 0x71, 0x02, 0x13, 0x07, 0xf7, 0xf4, 0x7e,
	0xf0, 0xaf, 0x3c, 0x7e, 0xe8, 0x8b, 0xe8, 0x5c, 0xe2, 0xaa, 0x87, 0xab, 0x7f, 0x4d, 0xe2, 0x78,
	0x10, 0x9a, 0x89, 0x52, 0x50, 0x3c, 0x30, 0x92, 0x5e, 0x4c, 0x2b, 0x8a, 0x8e, 0x91, 0xc2, 0xdb,
	0x9e, 0xa4, 0xd3, 0x79, 0x7c, 0x43, 0x92, 0x2e, 0x04, 0xfe, 0x2d, 0x2f, 0x7d, 0x80, 0x9b, 0x06,
	0x4e, 0xd2, 0x17, 0xea, 0x0e, 0xca, 0x47, 0xee, 0xed, 0xa0, 0x74, 0xed, 0xda, 0x7a, 0xee, 0x24,
	0x9b, 0xeb, 0x19, 0x13, 0xc3, 0xaa, 0xa3, 0xf6, 0x78, 0xba, 0xe9, 0xed, 0x6f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0c, 0x0d, 0xdd, 0xb3, 0x64, 0x01, 0x00, 0x00,
}
