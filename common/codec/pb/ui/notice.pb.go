// Code generated by protoc-gen-go. DO NOT EDIT.
// source: notice.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SCNotice struct {
	Content          []byte `protobuf:"bytes,1,req,name=content" json:"content,omitempty"`
	IntervalTime     *int64 `protobuf:"varint,2,req,name=intervalTime" json:"intervalTime,omitempty"`
	StartTime        *int64 `protobuf:"varint,3,opt,name=startTime,def=-1" json:"startTime,omitempty"`
	EndTime          *int64 `protobuf:"varint,4,opt,name=endTime,def=-1" json:"endTime,omitempty"`
	Num              *int32 `protobuf:"varint,5,opt,name=num,def=-1" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCNotice) Reset()                    { *m = SCNotice{} }
func (m *SCNotice) String() string            { return proto.CompactTextString(m) }
func (*SCNotice) ProtoMessage()               {}
func (*SCNotice) Descriptor() ([]byte, []int) { return fileDescriptor77, []int{0} }

const Default_SCNotice_StartTime int64 = -1
const Default_SCNotice_EndTime int64 = -1
const Default_SCNotice_Num int32 = -1

func (m *SCNotice) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *SCNotice) GetIntervalTime() int64 {
	if m != nil && m.IntervalTime != nil {
		return *m.IntervalTime
	}
	return 0
}

func (m *SCNotice) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return Default_SCNotice_StartTime
}

func (m *SCNotice) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return Default_SCNotice_EndTime
}

func (m *SCNotice) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return Default_SCNotice_Num
}

func init() {
	proto.RegisterType((*SCNotice)(nil), "ui.SCNotice")
}

func init() { proto.RegisterFile("notice.proto", fileDescriptor77) }

var fileDescriptor77 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xc8, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0x00, 0x50, 0x72, 0xb1, 0x54, 0x8f, 0x40, 0xe1, 0x54, 0xc8, 0x18, 0x9c, 0xb2, 0x28, 0xb8,
	0xba, 0xba, 0xbb, 0xe8, 0x0f, 0x94, 0x7a, 0x43, 0xc0, 0x5e, 0x24, 0x5e, 0xfc, 0x7e, 0x31, 0x43,
	0xd7, 0x87, 0x4e, 0xb2, 0xa6, 0x89, 0x4f, 0xef, 0x92, 0x35, 0x13, 0xd4, 0x74, 0x10, 0x5c, 0xdf,
	0xaf, 0xb7, 0xa6, 0x34, 0x60, 0x3f, 0x65, 0x51, 0x16, 0xf5, 0x26, 0x40, 0x74, 0xb4, 0x43, 0x97,
	0x44, 0xb9, 0x7c, 0xc7, 0xd7, 0x23, 0xcd, 0xec, 0x21, 0x40, 0xb4, 0xb4, 0xc7, 0xcd, 0x47, 0xc7,
	0xa2, 0x8d, 0x6c, 0x30, 0xd1, 0x5e, 0xe0, 0x78, 0xa6, 0x2d, 0xf6, 0x2c, 0xcf, 0x86, 0xab, 0x05,
	0x07, 0xb4, 0x52, 0x67, 0xdf, 0x05, 0x13, 0xbb, 0x3f, 0xfc, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf,
	0xd4, 0x39, 0xdb, 0x82, 0x00, 0x00, 0x00,
}