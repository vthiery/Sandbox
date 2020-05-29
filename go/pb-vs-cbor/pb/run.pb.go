// Code generated by protoc-gen-go. DO NOT EDIT.
// source: run.proto

package edgelab_adam_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Run struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Date                 string   `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Keyspace             string   `protobuf:"bytes,3,opt,name=keyspace,proto3" json:"keyspace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Run) Reset()         { *m = Run{} }
func (m *Run) String() string { return proto.CompactTextString(m) }
func (*Run) ProtoMessage()    {}
func (*Run) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3419bc3417bf873, []int{0}
}

func (m *Run) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Run.Unmarshal(m, b)
}
func (m *Run) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Run.Marshal(b, m, deterministic)
}
func (m *Run) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Run.Merge(m, src)
}
func (m *Run) XXX_Size() int {
	return xxx_messageInfo_Run.Size(m)
}
func (m *Run) XXX_DiscardUnknown() {
	xxx_messageInfo_Run.DiscardUnknown(m)
}

var xxx_messageInfo_Run proto.InternalMessageInfo

func (m *Run) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Run) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *Run) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

func init() {
	proto.RegisterType((*Run)(nil), "edgelab.adam.v1.Run")
}

func init() { proto.RegisterFile("run.proto", fileDescriptor_e3419bc3417bf873) }

var fileDescriptor_e3419bc3417bf873 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0xcd, 0xd3,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4f, 0x4d, 0x49, 0x4f, 0xcd, 0x49, 0x4c, 0xd2, 0x4b,
	0x4c, 0x49, 0xcc, 0xd5, 0x2b, 0x33, 0x54, 0x72, 0xe5, 0x62, 0x0e, 0x2a, 0xcd, 0x13, 0xe2, 0xe3,
	0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2,
	0x62, 0x49, 0x49, 0x2c, 0x49, 0x95, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xa4,
	0xb8, 0x38, 0xb2, 0x53, 0x2b, 0x8b, 0x0b, 0x12, 0x93, 0x53, 0x25, 0x98, 0xc1, 0xe2, 0x70, 0x7e,
	0x12, 0x1b, 0xd8, 0x78, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xba, 0xf9, 0xc6, 0x5e, 0x6b,
	0x00, 0x00, 0x00,
}
