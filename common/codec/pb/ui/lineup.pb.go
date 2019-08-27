// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lineup.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CSLineupCancel struct {
	CrossType        *int32 `protobuf:"varint,1,req,name=crossType" json:"crossType,omitempty"`
	SceneId          *int64 `protobuf:"varint,2,req,name=sceneId" json:"sceneId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSLineupCancel) Reset()                    { *m = CSLineupCancel{} }
func (m *CSLineupCancel) String() string            { return proto.CompactTextString(m) }
func (*CSLineupCancel) ProtoMessage()               {}
func (*CSLineupCancel) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{0} }

func (m *CSLineupCancel) GetCrossType() int32 {
	if m != nil && m.CrossType != nil {
		return *m.CrossType
	}
	return 0
}

func (m *CSLineupCancel) GetSceneId() int64 {
	if m != nil && m.SceneId != nil {
		return *m.SceneId
	}
	return 0
}

type SCLineupCancel struct {
	CrossType        *int32 `protobuf:"varint,1,req,name=crossType" json:"crossType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLineupCancel) Reset()                    { *m = SCLineupCancel{} }
func (m *SCLineupCancel) String() string            { return proto.CompactTextString(m) }
func (*SCLineupCancel) ProtoMessage()               {}
func (*SCLineupCancel) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{1} }

func (m *SCLineupCancel) GetCrossType() int32 {
	if m != nil && m.CrossType != nil {
		return *m.CrossType
	}
	return 0
}

type SCLineupNotice struct {
	BeforeNum        *int32 `protobuf:"varint,1,req,name=beforeNum" json:"beforeNum,omitempty"`
	CrossType        *int32 `protobuf:"varint,2,req,name=crossType" json:"crossType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLineupNotice) Reset()                    { *m = SCLineupNotice{} }
func (m *SCLineupNotice) String() string            { return proto.CompactTextString(m) }
func (*SCLineupNotice) ProtoMessage()               {}
func (*SCLineupNotice) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{2} }

func (m *SCLineupNotice) GetBeforeNum() int32 {
	if m != nil && m.BeforeNum != nil {
		return *m.BeforeNum
	}
	return 0
}

func (m *SCLineupNotice) GetCrossType() int32 {
	if m != nil && m.CrossType != nil {
		return *m.CrossType
	}
	return 0
}

type SCLineupSuccess struct {
	CrossType        *int32 `protobuf:"varint,1,req,name=crossType" json:"crossType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLineupSuccess) Reset()                    { *m = SCLineupSuccess{} }
func (m *SCLineupSuccess) String() string            { return proto.CompactTextString(m) }
func (*SCLineupSuccess) ProtoMessage()               {}
func (*SCLineupSuccess) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{3} }

func (m *SCLineupSuccess) GetCrossType() int32 {
	if m != nil && m.CrossType != nil {
		return *m.CrossType
	}
	return 0
}

type SCLineupSceneFinishToCancel struct {
	CrossType        *int32 `protobuf:"varint,1,req,name=crossType" json:"crossType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLineupSceneFinishToCancel) Reset()                    { *m = SCLineupSceneFinishToCancel{} }
func (m *SCLineupSceneFinishToCancel) String() string            { return proto.CompactTextString(m) }
func (*SCLineupSceneFinishToCancel) ProtoMessage()               {}
func (*SCLineupSceneFinishToCancel) Descriptor() ([]byte, []int) { return fileDescriptor58, []int{4} }

func (m *SCLineupSceneFinishToCancel) GetCrossType() int32 {
	if m != nil && m.CrossType != nil {
		return *m.CrossType
	}
	return 0
}

func init() {
	proto.RegisterType((*CSLineupCancel)(nil), "ui.CSLineupCancel")
	proto.RegisterType((*SCLineupCancel)(nil), "ui.SCLineupCancel")
	proto.RegisterType((*SCLineupNotice)(nil), "ui.SCLineupNotice")
	proto.RegisterType((*SCLineupSuccess)(nil), "ui.SCLineupSuccess")
	proto.RegisterType((*SCLineupSceneFinishToCancel)(nil), "ui.SCLineupSceneFinishToCancel")
}

func init() { proto.RegisterFile("lineup.proto", fileDescriptor58) }

var fileDescriptor58 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0xc9, 0xcc, 0x4b,
	0x2d, 0x2d, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0xcd, 0x94, 0xe2, 0xca, 0x2c,
	0x49, 0xcd, 0x85, 0xf0, 0xa5, 0x78, 0x92, 0xf3, 0x73, 0x73, 0xf3, 0xf3, 0x20, 0x3c, 0x25, 0x13,
	0x2e, 0x3e, 0xe7, 0x60, 0x1f, 0xb0, 0x7a, 0xe7, 0xc4, 0xbc, 0xe4, 0xd4, 0x1c, 0x21, 0x41, 0x2e,
	0xce, 0xe4, 0xa2, 0xfc, 0xe2, 0xe2, 0x90, 0xca, 0x82, 0x54, 0x09, 0x46, 0x05, 0x26, 0x0d, 0x56,
	0x21, 0x7e, 0x2e, 0xf6, 0xe2, 0xe4, 0xd4, 0xbc, 0x54, 0xcf, 0x14, 0x09, 0x26, 0x05, 0x26, 0x0d,
	0x66, 0x25, 0x65, 0x2e, 0xbe, 0x60, 0x67, 0x02, 0xba, 0x94, 0xcc, 0x10, 0x8a, 0xfc, 0xf2, 0x4b,
	0x32, 0x93, 0x53, 0x41, 0x8a, 0x92, 0x52, 0xd3, 0xf2, 0x8b, 0x52, 0xfd, 0x4a, 0x73, 0xa1, 0x46,
	0xa3, 0xe8, 0x63, 0x02, 0xeb, 0x53, 0xe1, 0xe2, 0x87, 0xe9, 0x0b, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d,
	0x2e, 0xc6, 0x66, 0xba, 0x01, 0x97, 0x34, 0x5c, 0x15, 0xc8, 0x6d, 0x6e, 0x99, 0x79, 0x99, 0xc5,
	0x19, 0x21, 0xf9, 0x38, 0xdd, 0x03, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xdf, 0x71, 0x71, 0x17,
	0x01, 0x00, 0x00,
}