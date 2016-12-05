// Code generated by protoc-gen-go.
// source: syncInform.proto
// DO NOT EDIT!

package bean

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SyncInform struct {
	SessionId   int64 `protobuf:"varint,11,opt,name=sessionId" json:"sessionId,omitempty"`
	LatestIndex int64 `protobuf:"varint,12,opt,name=latestIndex" json:"latestIndex,omitempty"`
	ReadIndex   int64 `protobuf:"varint,13,opt,name=readIndex" json:"readIndex,omitempty"`
}

func (m *SyncInform) Reset()                    { *m = SyncInform{} }
func (m *SyncInform) String() string            { return proto.CompactTextString(m) }
func (*SyncInform) ProtoMessage()               {}
func (*SyncInform) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SyncInform) GetSessionId() int64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *SyncInform) GetLatestIndex() int64 {
	if m != nil {
		return m.LatestIndex
	}
	return 0
}

func (m *SyncInform) GetReadIndex() int64 {
	if m != nil {
		return m.ReadIndex
	}
	return 0
}

func init() {
	proto.RegisterType((*SyncInform)(nil), "bean.SyncInform")
}

func init() { proto.RegisterFile("syncInform.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 116 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xae, 0xcc, 0x4b,
	0xf6, 0xcc, 0x4b, 0xcb, 0x2f, 0xca, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x4a,
	0x4d, 0xcc, 0x53, 0xca, 0xe2, 0xe2, 0x0a, 0x86, 0xcb, 0x08, 0xc9, 0x70, 0x71, 0x16, 0xa7, 0x16,
	0x17, 0x67, 0xe6, 0xe7, 0x79, 0xa6, 0x48, 0x70, 0x2b, 0x30, 0x6a, 0x30, 0x07, 0x21, 0x04, 0x84,
	0x14, 0xb8, 0xb8, 0x73, 0x12, 0x4b, 0x52, 0x8b, 0x4b, 0x3c, 0xf3, 0x52, 0x52, 0x2b, 0x24, 0x78,
	0xc0, 0xf2, 0xc8, 0x42, 0x20, 0xfd, 0x45, 0xa9, 0x89, 0x29, 0x10, 0x79, 0x5e, 0x88, 0x7e, 0xb8,
	0x40, 0x12, 0x1b, 0xd8, 0x62, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6b, 0xbd, 0x63, 0xfc,
	0x8c, 0x00, 0x00, 0x00,
}