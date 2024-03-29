// Code generated by protoc-gen-go.
// source: request.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Request struct {
	Rid uint64 `protobuf:"varint,1,opt,name=Rid,json=rid" json:"Rid,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Request) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
}

func init() { proto.RegisterFile("request.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 73 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe6,
	0x62, 0x0f, 0x82, 0x08, 0x0a, 0x09, 0x70, 0x31, 0x07, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x04, 0x31, 0x17, 0x65, 0xa6, 0x24, 0xb1, 0x81, 0xd5, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xd4, 0xc6, 0x75, 0x03, 0x38, 0x00, 0x00, 0x00,
}
