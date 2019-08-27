// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lingtong_fashion.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LingTongFashionInfo struct {
	FashionId        *int32 `protobuf:"varint,1,req,name=fashionId" json:"fashionId,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	LevelPro         *int32 `protobuf:"varint,3,req,name=levelPro" json:"levelPro,omitempty"`
	ActivteTime      *int64 `protobuf:"varint,4,req,name=activteTime,def=0" json:"activteTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LingTongFashionInfo) Reset()                    { *m = LingTongFashionInfo{} }
func (m *LingTongFashionInfo) String() string            { return proto.CompactTextString(m) }
func (*LingTongFashionInfo) ProtoMessage()               {}
func (*LingTongFashionInfo) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{0} }

const Default_LingTongFashionInfo_ActivteTime int64 = 0

func (m *LingTongFashionInfo) GetFashionId() int32 {
	if m != nil && m.FashionId != nil {
		return *m.FashionId
	}
	return 0
}

func (m *LingTongFashionInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *LingTongFashionInfo) GetLevelPro() int32 {
	if m != nil && m.LevelPro != nil {
		return *m.LevelPro
	}
	return 0
}

func (m *LingTongFashionInfo) GetActivteTime() int64 {
	if m != nil && m.ActivteTime != nil {
		return *m.ActivteTime
	}
	return Default_LingTongFashionInfo_ActivteTime
}

type CSLingTongFashionGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSLingTongFashionGet) Reset()                    { *m = CSLingTongFashionGet{} }
func (m *CSLingTongFashionGet) String() string            { return proto.CompactTextString(m) }
func (*CSLingTongFashionGet) ProtoMessage()               {}
func (*CSLingTongFashionGet) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{1} }

type SCLingTongFashionGet struct {
	LingTongId       *int32                 `protobuf:"varint,1,req,name=lingTongId" json:"lingTongId,omitempty"`
	FashionList      []*LingTongFashionInfo `protobuf:"bytes,2,rep,name=fashionList" json:"fashionList,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SCLingTongFashionGet) Reset()                    { *m = SCLingTongFashionGet{} }
func (m *SCLingTongFashionGet) String() string            { return proto.CompactTextString(m) }
func (*SCLingTongFashionGet) ProtoMessage()               {}
func (*SCLingTongFashionGet) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{2} }

func (m *SCLingTongFashionGet) GetLingTongId() int32 {
	if m != nil && m.LingTongId != nil {
		return *m.LingTongId
	}
	return 0
}

func (m *SCLingTongFashionGet) GetFashionList() []*LingTongFashionInfo {
	if m != nil {
		return m.FashionList
	}
	return nil
}

type CSLingTongFashionActive struct {
	FashionId        *int32 `protobuf:"varint,1,req,name=fashionId" json:"fashionId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSLingTongFashionActive) Reset()                    { *m = CSLingTongFashionActive{} }
func (m *CSLingTongFashionActive) String() string            { return proto.CompactTextString(m) }
func (*CSLingTongFashionActive) ProtoMessage()               {}
func (*CSLingTongFashionActive) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{3} }

func (m *CSLingTongFashionActive) GetFashionId() int32 {
	if m != nil && m.FashionId != nil {
		return *m.FashionId
	}
	return 0
}

type SCLingTongFashionActive struct {
	FashionInfo      *LingTongFashionInfo `protobuf:"bytes,1,req,name=fashionInfo" json:"fashionInfo,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SCLingTongFashionActive) Reset()                    { *m = SCLingTongFashionActive{} }
func (m *SCLingTongFashionActive) String() string            { return proto.CompactTextString(m) }
func (*SCLingTongFashionActive) ProtoMessage()               {}
func (*SCLingTongFashionActive) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{4} }

func (m *SCLingTongFashionActive) GetFashionInfo() *LingTongFashionInfo {
	if m != nil {
		return m.FashionInfo
	}
	return nil
}

type CSLingTongFashionWear struct {
	FashionId        *int32 `protobuf:"varint,1,req,name=fashionId" json:"fashionId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSLingTongFashionWear) Reset()                    { *m = CSLingTongFashionWear{} }
func (m *CSLingTongFashionWear) String() string            { return proto.CompactTextString(m) }
func (*CSLingTongFashionWear) ProtoMessage()               {}
func (*CSLingTongFashionWear) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{5} }

func (m *CSLingTongFashionWear) GetFashionId() int32 {
	if m != nil && m.FashionId != nil {
		return *m.FashionId
	}
	return 0
}

type SCLingTongFashionWear struct {
	FashionId        *int32 `protobuf:"varint,1,req,name=fashionId" json:"fashionId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLingTongFashionWear) Reset()                    { *m = SCLingTongFashionWear{} }
func (m *SCLingTongFashionWear) String() string            { return proto.CompactTextString(m) }
func (*SCLingTongFashionWear) ProtoMessage()               {}
func (*SCLingTongFashionWear) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{6} }

func (m *SCLingTongFashionWear) GetFashionId() int32 {
	if m != nil && m.FashionId != nil {
		return *m.FashionId
	}
	return 0
}

type CSLingTongFashionUnload struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSLingTongFashionUnload) Reset()                    { *m = CSLingTongFashionUnload{} }
func (m *CSLingTongFashionUnload) String() string            { return proto.CompactTextString(m) }
func (*CSLingTongFashionUnload) ProtoMessage()               {}
func (*CSLingTongFashionUnload) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{7} }

type SCLingTongFashionUnload struct {
	FashionId        *int32 `protobuf:"varint,1,req,name=fashionId" json:"fashionId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLingTongFashionUnload) Reset()                    { *m = SCLingTongFashionUnload{} }
func (m *SCLingTongFashionUnload) String() string            { return proto.CompactTextString(m) }
func (*SCLingTongFashionUnload) ProtoMessage()               {}
func (*SCLingTongFashionUnload) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{8} }

func (m *SCLingTongFashionUnload) GetFashionId() int32 {
	if m != nil && m.FashionId != nil {
		return *m.FashionId
	}
	return 0
}

type CSLingTongFashionUpstar struct {
	FashionId        *int32 `protobuf:"varint,1,req,name=fashionId" json:"fashionId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSLingTongFashionUpstar) Reset()                    { *m = CSLingTongFashionUpstar{} }
func (m *CSLingTongFashionUpstar) String() string            { return proto.CompactTextString(m) }
func (*CSLingTongFashionUpstar) ProtoMessage()               {}
func (*CSLingTongFashionUpstar) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{9} }

func (m *CSLingTongFashionUpstar) GetFashionId() int32 {
	if m != nil && m.FashionId != nil {
		return *m.FashionId
	}
	return 0
}

type SCLingTongFashionUpstar struct {
	FashionId        *int32 `protobuf:"varint,1,req,name=fashionId" json:"fashionId,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	LevelPro         *int32 `protobuf:"varint,3,req,name=levelPro" json:"levelPro,omitempty"`
	Result           *bool  `protobuf:"varint,4,req,name=result" json:"result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLingTongFashionUpstar) Reset()                    { *m = SCLingTongFashionUpstar{} }
func (m *SCLingTongFashionUpstar) String() string            { return proto.CompactTextString(m) }
func (*SCLingTongFashionUpstar) ProtoMessage()               {}
func (*SCLingTongFashionUpstar) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{10} }

func (m *SCLingTongFashionUpstar) GetFashionId() int32 {
	if m != nil && m.FashionId != nil {
		return *m.FashionId
	}
	return 0
}

func (m *SCLingTongFashionUpstar) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCLingTongFashionUpstar) GetLevelPro() int32 {
	if m != nil && m.LevelPro != nil {
		return *m.LevelPro
	}
	return 0
}

func (m *SCLingTongFashionUpstar) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

type SCLingTongFashionRemove struct {
	FashionId        *int32 `protobuf:"varint,1,req,name=fashionId" json:"fashionId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLingTongFashionRemove) Reset()                    { *m = SCLingTongFashionRemove{} }
func (m *SCLingTongFashionRemove) String() string            { return proto.CompactTextString(m) }
func (*SCLingTongFashionRemove) ProtoMessage()               {}
func (*SCLingTongFashionRemove) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{11} }

func (m *SCLingTongFashionRemove) GetFashionId() int32 {
	if m != nil && m.FashionId != nil {
		return *m.FashionId
	}
	return 0
}

type SCLingTongFashionTrialNotice struct {
	TrialFashionId   *int32 `protobuf:"varint,1,req,name=trialFashionId" json:"trialFashionId,omitempty"`
	ExpireTime       *int64 `protobuf:"varint,2,req,name=expireTime" json:"expireTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLingTongFashionTrialNotice) Reset()                    { *m = SCLingTongFashionTrialNotice{} }
func (m *SCLingTongFashionTrialNotice) String() string            { return proto.CompactTextString(m) }
func (*SCLingTongFashionTrialNotice) ProtoMessage()               {}
func (*SCLingTongFashionTrialNotice) Descriptor() ([]byte, []int) { return fileDescriptor61, []int{12} }

func (m *SCLingTongFashionTrialNotice) GetTrialFashionId() int32 {
	if m != nil && m.TrialFashionId != nil {
		return *m.TrialFashionId
	}
	return 0
}

func (m *SCLingTongFashionTrialNotice) GetExpireTime() int64 {
	if m != nil && m.ExpireTime != nil {
		return *m.ExpireTime
	}
	return 0
}

type SCLingTongFashionTrialOverdueNotice struct {
	TrialFashionId   *int32 `protobuf:"varint,1,req,name=trialFashionId" json:"trialFashionId,omitempty"`
	OverdueType      *int32 `protobuf:"varint,2,req,name=overdueType" json:"overdueType,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCLingTongFashionTrialOverdueNotice) Reset()         { *m = SCLingTongFashionTrialOverdueNotice{} }
func (m *SCLingTongFashionTrialOverdueNotice) String() string { return proto.CompactTextString(m) }
func (*SCLingTongFashionTrialOverdueNotice) ProtoMessage()    {}
func (*SCLingTongFashionTrialOverdueNotice) Descriptor() ([]byte, []int) {
	return fileDescriptor61, []int{13}
}

func (m *SCLingTongFashionTrialOverdueNotice) GetTrialFashionId() int32 {
	if m != nil && m.TrialFashionId != nil {
		return *m.TrialFashionId
	}
	return 0
}

func (m *SCLingTongFashionTrialOverdueNotice) GetOverdueType() int32 {
	if m != nil && m.OverdueType != nil {
		return *m.OverdueType
	}
	return 0
}

func init() {
	proto.RegisterType((*LingTongFashionInfo)(nil), "ui.LingTongFashionInfo")
	proto.RegisterType((*CSLingTongFashionGet)(nil), "ui.CSLingTongFashionGet")
	proto.RegisterType((*SCLingTongFashionGet)(nil), "ui.SCLingTongFashionGet")
	proto.RegisterType((*CSLingTongFashionActive)(nil), "ui.CSLingTongFashionActive")
	proto.RegisterType((*SCLingTongFashionActive)(nil), "ui.SCLingTongFashionActive")
	proto.RegisterType((*CSLingTongFashionWear)(nil), "ui.CSLingTongFashionWear")
	proto.RegisterType((*SCLingTongFashionWear)(nil), "ui.SCLingTongFashionWear")
	proto.RegisterType((*CSLingTongFashionUnload)(nil), "ui.CSLingTongFashionUnload")
	proto.RegisterType((*SCLingTongFashionUnload)(nil), "ui.SCLingTongFashionUnload")
	proto.RegisterType((*CSLingTongFashionUpstar)(nil), "ui.CSLingTongFashionUpstar")
	proto.RegisterType((*SCLingTongFashionUpstar)(nil), "ui.SCLingTongFashionUpstar")
	proto.RegisterType((*SCLingTongFashionRemove)(nil), "ui.SCLingTongFashionRemove")
	proto.RegisterType((*SCLingTongFashionTrialNotice)(nil), "ui.SCLingTongFashionTrialNotice")
	proto.RegisterType((*SCLingTongFashionTrialOverdueNotice)(nil), "ui.SCLingTongFashionTrialOverdueNotice")
}

func init() { proto.RegisterFile("lingtong_fashion.proto", fileDescriptor61) }

var fileDescriptor61 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x4b, 0x02, 0x41,
	0x10, 0xc7, 0xf1, 0x2e, 0xc3, 0xe6, 0x48, 0x6a, 0xd5, 0xf3, 0x82, 0x1e, 0xe4, 0x7a, 0x11, 0x11,
	0x89, 0x1e, 0x7b, 0x0b, 0x41, 0x31, 0xa4, 0x42, 0x4f, 0xea, 0x25, 0xe2, 0xd0, 0xf1, 0x5a, 0x58,
	0x77, 0x8f, 0xbd, 0xbd, 0xa3, 0xfe, 0xfb, 0xb8, 0x1f, 0x14, 0xb8, 0x6b, 0x3e, 0xce, 0xcc, 0x97,
	0xef, 0xe7, 0x3b, 0xbb, 0x03, 0x2e, 0xa3, 0x3c, 0x52, 0x82, 0x47, 0x1f, 0xdb, 0x30, 0xf9, 0xa4,
	0x82, 0x8f, 0x62, 0x29, 0x94, 0x20, 0x56, 0x4a, 0xfd, 0x10, 0x5a, 0x73, 0xca, 0xa3, 0x40, 0xf0,
	0x68, 0x52, 0x0e, 0x67, 0x7c, 0x2b, 0xc8, 0x25, 0x9c, 0x55, 0xda, 0xd9, 0xc6, 0xab, 0xf5, 0xac,
	0x7e, 0x9d, 0x9c, 0x43, 0x9d, 0x61, 0x86, 0xcc, 0xb3, 0x8a, 0xf2, 0x02, 0x1a, 0x45, 0xf9, 0x22,
	0x85, 0x67, 0x17, 0x1d, 0x17, 0x9c, 0x70, 0xad, 0x68, 0xa6, 0x30, 0xa0, 0x3b, 0xf4, 0x4e, 0x7a,
	0x56, 0xdf, 0xbe, 0xaf, 0xdd, 0xfa, 0x2e, 0xb4, 0xc7, 0xcb, 0x3d, 0xc8, 0x14, 0x95, 0xff, 0x06,
	0xed, 0xe5, 0x58, 0xef, 0x13, 0x02, 0xc0, 0xaa, 0xee, 0x2f, 0x7c, 0x08, 0x4e, 0x95, 0x67, 0x4e,
	0x13, 0xe5, 0x59, 0x3d, 0xbb, 0xef, 0xdc, 0x75, 0x47, 0x29, 0x1d, 0x19, 0xd2, 0xfb, 0x43, 0xe8,
	0x6a, 0xc4, 0x87, 0x3c, 0x1a, 0x1a, 0x16, 0xf3, 0xa7, 0xd0, 0xd5, 0x72, 0x54, 0xea, 0x3f, 0x6c,
	0xee, 0x5b, 0xe8, 0xff, 0xc1, 0x0e, 0xa0, 0xa3, 0x61, 0x5f, 0x31, 0x94, 0x26, 0xe8, 0x00, 0x3a,
	0x1a, 0xf4, 0x90, 0xf6, 0xca, 0xb0, 0xce, 0x8a, 0x33, 0x11, 0x6e, 0xf2, 0x4d, 0x35, 0x9b, 0x72,
	0x64, 0x32, 0x32, 0xbd, 0xcb, 0x2a, 0x4e, 0x94, 0x19, 0xfb, 0x6e, 0xf2, 0x3e, 0xa4, 0x3e, 0x7e,
	0x1e, 0x4d, 0x38, 0x95, 0x98, 0xa4, 0x4c, 0x15, 0x97, 0xd1, 0x30, 0x46, 0x5f, 0xe0, 0x4e, 0x98,
	0x3f, 0xe9, 0x11, 0xae, 0x35, 0x75, 0x20, 0x69, 0xc8, 0x9e, 0x84, 0xa2, 0x6b, 0x24, 0x2e, 0x34,
	0x55, 0x5e, 0x4e, 0xf6, 0x62, 0x11, 0x00, 0xfc, 0x8a, 0xa9, 0x2c, 0x6f, 0x32, 0xcf, 0x66, 0xfb,
	0x0b, 0xb8, 0x31, 0x7b, 0x3d, 0x67, 0x28, 0x37, 0x29, 0x1e, 0xb1, 0x6c, 0x81, 0x23, 0x4a, 0x61,
	0xf0, 0x1d, 0x97, 0x9e, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x5a, 0x6a, 0xdd, 0x64,
	0x03, 0x00, 0x00,
}
