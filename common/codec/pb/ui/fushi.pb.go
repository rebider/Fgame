// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fushi.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FushiInfo struct {
	Typ              *int32 `protobuf:"varint,1,req,name=typ" json:"typ,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FushiInfo) Reset()                    { *m = FushiInfo{} }
func (m *FushiInfo) String() string            { return proto.CompactTextString(m) }
func (*FushiInfo) ProtoMessage()               {}
func (*FushiInfo) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{0} }

func (m *FushiInfo) GetTyp() int32 {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return 0
}

func (m *FushiInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

type CSFushiInfo struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSFushiInfo) Reset()                    { *m = CSFushiInfo{} }
func (m *CSFushiInfo) String() string            { return proto.CompactTextString(m) }
func (*CSFushiInfo) ProtoMessage()               {}
func (*CSFushiInfo) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{1} }

type SCFushiInfo struct {
	FushiList        []*FushiInfo `protobuf:"bytes,1,rep,name=fushiList" json:"fushiList,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SCFushiInfo) Reset()                    { *m = SCFushiInfo{} }
func (m *SCFushiInfo) String() string            { return proto.CompactTextString(m) }
func (*SCFushiInfo) ProtoMessage()               {}
func (*SCFushiInfo) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{2} }

func (m *SCFushiInfo) GetFushiList() []*FushiInfo {
	if m != nil {
		return m.FushiList
	}
	return nil
}

type CSFuShiActivite struct {
	Typ              *int32 `protobuf:"varint,1,req,name=typ" json:"typ,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSFuShiActivite) Reset()                    { *m = CSFuShiActivite{} }
func (m *CSFuShiActivite) String() string            { return proto.CompactTextString(m) }
func (*CSFuShiActivite) ProtoMessage()               {}
func (*CSFuShiActivite) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{3} }

func (m *CSFuShiActivite) GetTyp() int32 {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return 0
}

type SCFuShiActivite struct {
	Typ              *int32 `protobuf:"varint,1,req,name=typ" json:"typ,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCFuShiActivite) Reset()                    { *m = SCFuShiActivite{} }
func (m *SCFuShiActivite) String() string            { return proto.CompactTextString(m) }
func (*SCFuShiActivite) ProtoMessage()               {}
func (*SCFuShiActivite) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{4} }

func (m *SCFuShiActivite) GetTyp() int32 {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return 0
}

type CSFuShiUplevel struct {
	Typ              *int32 `protobuf:"varint,1,req,name=typ" json:"typ,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSFuShiUplevel) Reset()                    { *m = CSFuShiUplevel{} }
func (m *CSFuShiUplevel) String() string            { return proto.CompactTextString(m) }
func (*CSFuShiUplevel) ProtoMessage()               {}
func (*CSFuShiUplevel) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{5} }

func (m *CSFuShiUplevel) GetTyp() int32 {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return 0
}

type SCFuShiUplevel struct {
	Typ              *int32 `protobuf:"varint,1,req,name=typ" json:"typ,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCFuShiUplevel) Reset()                    { *m = SCFuShiUplevel{} }
func (m *SCFuShiUplevel) String() string            { return proto.CompactTextString(m) }
func (*SCFuShiUplevel) ProtoMessage()               {}
func (*SCFuShiUplevel) Descriptor() ([]byte, []int) { return fileDescriptor41, []int{6} }

func (m *SCFuShiUplevel) GetTyp() int32 {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return 0
}

func (m *SCFuShiUplevel) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func init() {
	proto.RegisterType((*FushiInfo)(nil), "ui.FushiInfo")
	proto.RegisterType((*CSFushiInfo)(nil), "ui.CSFushiInfo")
	proto.RegisterType((*SCFushiInfo)(nil), "ui.SCFushiInfo")
	proto.RegisterType((*CSFuShiActivite)(nil), "ui.CSFuShiActivite")
	proto.RegisterType((*SCFuShiActivite)(nil), "ui.SCFuShiActivite")
	proto.RegisterType((*CSFuShiUplevel)(nil), "ui.CSFuShiUplevel")
	proto.RegisterType((*SCFuShiUplevel)(nil), "ui.SCFuShiUplevel")
}

func init() { proto.RegisterFile("fushi.proto", fileDescriptor41) }

var fileDescriptor41 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2b, 0x2d, 0xce,
	0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0xcd, 0x54, 0x52, 0xe7, 0xe2, 0x74,
	0x03, 0x09, 0x79, 0xe6, 0xa5, 0xe5, 0x0b, 0x71, 0x73, 0x31, 0x97, 0x54, 0x16, 0x48, 0x30, 0x2a,
	0x30, 0x69, 0xb0, 0x0a, 0xf1, 0x72, 0xb1, 0xe6, 0xa4, 0x96, 0xa5, 0xe6, 0x48, 0x30, 0x81, 0xb8,
	0x4a, 0xbc, 0x5c, 0xdc, 0xce, 0xc1, 0x70, 0xa5, 0x4a, 0xfa, 0x5c, 0xdc, 0xc1, 0xce, 0x08, 0x9d,
	0x0a, 0x5c, 0x9c, 0x60, 0x93, 0x7d, 0x32, 0x8b, 0x4b, 0x24, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d,
	0x78, 0xf5, 0x4a, 0x33, 0xf5, 0x10, 0x1a, 0xe4, 0xb8, 0xf8, 0x41, 0xfa, 0x83, 0x33, 0x32, 0x1d,
	0x93, 0x4b, 0x32, 0xcb, 0x32, 0x4b, 0x52, 0x51, 0xac, 0x03, 0xc9, 0x83, 0x0c, 0xc4, 0x29, 0x2f,
	0xcb, 0xc5, 0x07, 0xd5, 0x1f, 0x5a, 0x00, 0x76, 0x17, 0xaa, 0xb4, 0x0e, 0x17, 0x1f, 0x54, 0x3b,
	0x36, 0x69, 0x34, 0xcf, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x10, 0x69, 0xbb, 0x07, 0x01,
	0x00, 0x00,
}
