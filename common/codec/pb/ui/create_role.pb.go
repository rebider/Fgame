// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create_role.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SCEnterSelectJob struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCEnterSelectJob) Reset()                    { *m = SCEnterSelectJob{} }
func (m *SCEnterSelectJob) String() string            { return proto.CompactTextString(m) }
func (*SCEnterSelectJob) ProtoMessage()               {}
func (*SCEnterSelectJob) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{0} }

type CSSelectJob struct {
	Job              *int32  `protobuf:"varint,1,req,name=job" json:"job,omitempty"`
	Sex              *int32  `protobuf:"varint,2,req,name=sex" json:"sex,omitempty"`
	Name             *string `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CSSelectJob) Reset()                    { *m = CSSelectJob{} }
func (m *CSSelectJob) String() string            { return proto.CompactTextString(m) }
func (*CSSelectJob) ProtoMessage()               {}
func (*CSSelectJob) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{1} }

func (m *CSSelectJob) GetJob() int32 {
	if m != nil && m.Job != nil {
		return *m.Job
	}
	return 0
}

func (m *CSSelectJob) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *CSSelectJob) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type SCSelectJob struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCSelectJob) Reset()                    { *m = SCSelectJob{} }
func (m *SCSelectJob) String() string            { return proto.CompactTextString(m) }
func (*SCSelectJob) ProtoMessage()               {}
func (*SCSelectJob) Descriptor() ([]byte, []int) { return fileDescriptor22, []int{2} }

func init() {
	proto.RegisterType((*SCEnterSelectJob)(nil), "ui.SCEnterSelectJob")
	proto.RegisterType((*CSSelectJob)(nil), "ui.CSSelectJob")
	proto.RegisterType((*SCSelectJob)(nil), "ui.SCSelectJob")
}

func init() { proto.RegisterFile("create_role.proto", fileDescriptor22) }

var fileDescriptor22 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x2e, 0x4a, 0x4d,
	0x2c, 0x49, 0x8d, 0x2f, 0xca, 0xcf, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a,
	0xcd, 0x54, 0x12, 0xe2, 0x12, 0x08, 0x76, 0x76, 0xcd, 0x2b, 0x49, 0x2d, 0x0a, 0x4e, 0xcd, 0x49,
	0x4d, 0x2e, 0xf1, 0xca, 0x4f, 0x52, 0x32, 0xe5, 0xe2, 0x76, 0x0e, 0x86, 0x73, 0x85, 0xb8, 0xb9,
	0x98, 0xb3, 0xf2, 0x93, 0x24, 0x18, 0x15, 0x98, 0x34, 0x58, 0x41, 0x9c, 0xe2, 0xd4, 0x0a, 0x09,
	0x26, 0x30, 0x87, 0x87, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x59, 0x81, 0x49, 0x83, 0x53,
	0x89, 0x97, 0x8b, 0x3b, 0xd8, 0x19, 0xae, 0x0d, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x38, 0xb9,
	0x69, 0x71, 0x00, 0x00, 0x00,
}
