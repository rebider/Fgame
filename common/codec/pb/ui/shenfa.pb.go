// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shenfa.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ShenfaInfo struct {
	AdvancedId       *int32            `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	ShenfaId         *int32            `protobuf:"varint,2,req,name=shenfaId" json:"shenfaId,omitempty"`
	UnrealLevel      *int32            `protobuf:"varint,3,req,name=unrealLevel" json:"unrealLevel,omitempty"`
	UnrealPro        *int32            `protobuf:"varint,4,req,name=unrealPro" json:"unrealPro,omitempty"`
	SkinList         []*ShenFaSkinInfo `protobuf:"bytes,5,rep,name=skinList" json:"skinList,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *ShenfaInfo) Reset()                    { *m = ShenfaInfo{} }
func (m *ShenfaInfo) String() string            { return proto.CompactTextString(m) }
func (*ShenfaInfo) ProtoMessage()               {}
func (*ShenfaInfo) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{0} }

func (m *ShenfaInfo) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *ShenfaInfo) GetShenfaId() int32 {
	if m != nil && m.ShenfaId != nil {
		return *m.ShenfaId
	}
	return 0
}

func (m *ShenfaInfo) GetUnrealLevel() int32 {
	if m != nil && m.UnrealLevel != nil {
		return *m.UnrealLevel
	}
	return 0
}

func (m *ShenfaInfo) GetUnrealPro() int32 {
	if m != nil && m.UnrealPro != nil {
		return *m.UnrealPro
	}
	return 0
}

func (m *ShenfaInfo) GetSkinList() []*ShenFaSkinInfo {
	if m != nil {
		return m.SkinList
	}
	return nil
}

type ShenFaSkinInfo struct {
	ShenFaId         *int32 `protobuf:"varint,1,req,name=shenFaId" json:"shenFaId,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Pro              *int32 `protobuf:"varint,3,req,name=pro" json:"pro,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShenFaSkinInfo) Reset()                    { *m = ShenFaSkinInfo{} }
func (m *ShenFaSkinInfo) String() string            { return proto.CompactTextString(m) }
func (*ShenFaSkinInfo) ProtoMessage()               {}
func (*ShenFaSkinInfo) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{1} }

func (m *ShenFaSkinInfo) GetShenFaId() int32 {
	if m != nil && m.ShenFaId != nil {
		return *m.ShenFaId
	}
	return 0
}

func (m *ShenFaSkinInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *ShenFaSkinInfo) GetPro() int32 {
	if m != nil && m.Pro != nil {
		return *m.Pro
	}
	return 0
}

type CSShenfaGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShenfaGet) Reset()                    { *m = CSShenfaGet{} }
func (m *CSShenfaGet) String() string            { return proto.CompactTextString(m) }
func (*CSShenfaGet) ProtoMessage()               {}
func (*CSShenfaGet) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{2} }

type SCShenfaGet struct {
	AdvancedId       *int32            `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	ShenfaId         *int32            `protobuf:"varint,2,req,name=shenfaId" json:"shenfaId,omitempty"`
	UnrealLevel      *int32            `protobuf:"varint,3,req,name=unrealLevel" json:"unrealLevel,omitempty"`
	UnrealPro        *int32            `protobuf:"varint,4,req,name=unrealPro" json:"unrealPro,omitempty"`
	UnrealList       []int32           `protobuf:"varint,5,rep,name=unrealList" json:"unrealList,omitempty"`
	Bless            *int32            `protobuf:"varint,6,opt,name=bless" json:"bless,omitempty"`
	BlessTime        *int64            `protobuf:"varint,7,opt,name=blessTime" json:"blessTime,omitempty"`
	Hidden           *bool             `protobuf:"varint,8,req,name=hidden" json:"hidden,omitempty"`
	ShenFaSkinList   []*ShenFaSkinInfo `protobuf:"bytes,9,rep,name=shenFaSkinList" json:"shenFaSkinList,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SCShenfaGet) Reset()                    { *m = SCShenfaGet{} }
func (m *SCShenfaGet) String() string            { return proto.CompactTextString(m) }
func (*SCShenfaGet) ProtoMessage()               {}
func (*SCShenfaGet) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{3} }

func (m *SCShenfaGet) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *SCShenfaGet) GetShenfaId() int32 {
	if m != nil && m.ShenfaId != nil {
		return *m.ShenfaId
	}
	return 0
}

func (m *SCShenfaGet) GetUnrealLevel() int32 {
	if m != nil && m.UnrealLevel != nil {
		return *m.UnrealLevel
	}
	return 0
}

func (m *SCShenfaGet) GetUnrealPro() int32 {
	if m != nil && m.UnrealPro != nil {
		return *m.UnrealPro
	}
	return 0
}

func (m *SCShenfaGet) GetUnrealList() []int32 {
	if m != nil {
		return m.UnrealList
	}
	return nil
}

func (m *SCShenfaGet) GetBless() int32 {
	if m != nil && m.Bless != nil {
		return *m.Bless
	}
	return 0
}

func (m *SCShenfaGet) GetBlessTime() int64 {
	if m != nil && m.BlessTime != nil {
		return *m.BlessTime
	}
	return 0
}

func (m *SCShenfaGet) GetHidden() bool {
	if m != nil && m.Hidden != nil {
		return *m.Hidden
	}
	return false
}

func (m *SCShenfaGet) GetShenFaSkinList() []*ShenFaSkinInfo {
	if m != nil {
		return m.ShenFaSkinList
	}
	return nil
}

type CSShenfaUnrealDan struct {
	Num              *int32 `protobuf:"varint,1,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShenfaUnrealDan) Reset()                    { *m = CSShenfaUnrealDan{} }
func (m *CSShenfaUnrealDan) String() string            { return proto.CompactTextString(m) }
func (*CSShenfaUnrealDan) ProtoMessage()               {}
func (*CSShenfaUnrealDan) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{4} }

func (m *CSShenfaUnrealDan) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type SCShenfaUnrealDan struct {
	Level            *int32 `protobuf:"varint,1,req,name=level" json:"level,omitempty"`
	Progress         *int32 `protobuf:"varint,2,req,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShenfaUnrealDan) Reset()                    { *m = SCShenfaUnrealDan{} }
func (m *SCShenfaUnrealDan) String() string            { return proto.CompactTextString(m) }
func (*SCShenfaUnrealDan) ProtoMessage()               {}
func (*SCShenfaUnrealDan) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{5} }

func (m *SCShenfaUnrealDan) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCShenfaUnrealDan) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSShenfaUnreal struct {
	ShenfaId         *int32 `protobuf:"varint,1,req,name=shenfaId" json:"shenfaId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShenfaUnreal) Reset()                    { *m = CSShenfaUnreal{} }
func (m *CSShenfaUnreal) String() string            { return proto.CompactTextString(m) }
func (*CSShenfaUnreal) ProtoMessage()               {}
func (*CSShenfaUnreal) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{6} }

func (m *CSShenfaUnreal) GetShenfaId() int32 {
	if m != nil && m.ShenfaId != nil {
		return *m.ShenfaId
	}
	return 0
}

type SCShenfaUnreal struct {
	ShenfaId         *int32 `protobuf:"varint,1,opt,name=shenfaId" json:"shenfaId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShenfaUnreal) Reset()                    { *m = SCShenfaUnreal{} }
func (m *SCShenfaUnreal) String() string            { return proto.CompactTextString(m) }
func (*SCShenfaUnreal) ProtoMessage()               {}
func (*SCShenfaUnreal) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{7} }

func (m *SCShenfaUnreal) GetShenfaId() int32 {
	if m != nil && m.ShenfaId != nil {
		return *m.ShenfaId
	}
	return 0
}

type CSShenfaUnload struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShenfaUnload) Reset()                    { *m = CSShenfaUnload{} }
func (m *CSShenfaUnload) String() string            { return proto.CompactTextString(m) }
func (*CSShenfaUnload) ProtoMessage()               {}
func (*CSShenfaUnload) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{8} }

type SCShenfaUnload struct {
	ShenfaId         *int32 `protobuf:"varint,1,req,name=shenfaId" json:"shenfaId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShenfaUnload) Reset()                    { *m = SCShenfaUnload{} }
func (m *SCShenfaUnload) String() string            { return proto.CompactTextString(m) }
func (*SCShenfaUnload) ProtoMessage()               {}
func (*SCShenfaUnload) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{9} }

func (m *SCShenfaUnload) GetShenfaId() int32 {
	if m != nil && m.ShenfaId != nil {
		return *m.ShenfaId
	}
	return 0
}

type CSShenfaAdvanced struct {
	AutoFlag         *bool  `protobuf:"varint,1,req,name=autoFlag" json:"autoFlag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShenfaAdvanced) Reset()                    { *m = CSShenfaAdvanced{} }
func (m *CSShenfaAdvanced) String() string            { return proto.CompactTextString(m) }
func (*CSShenfaAdvanced) ProtoMessage()               {}
func (*CSShenfaAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{10} }

func (m *CSShenfaAdvanced) GetAutoFlag() bool {
	if m != nil && m.AutoFlag != nil {
		return *m.AutoFlag
	}
	return false
}

type SCShenfaAdvanced struct {
	AdvancedId       *int32 `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	ShenfaId         *int32 `protobuf:"varint,2,req,name=shenfaId" json:"shenfaId,omitempty"`
	Bless            *int32 `protobuf:"varint,3,opt,name=bless" json:"bless,omitempty"`
	BlessTime        *int64 `protobuf:"varint,4,opt,name=blessTime" json:"blessTime,omitempty"`
	AdvancedType     *int32 `protobuf:"varint,5,req,name=advancedType" json:"advancedType,omitempty"`
	IsDouble         *bool  `protobuf:"varint,6,opt,name=isDouble,def=0" json:"isDouble,omitempty"`
	TotalBless       *int32 `protobuf:"varint,7,opt,name=totalBless" json:"totalBless,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShenfaAdvanced) Reset()                    { *m = SCShenfaAdvanced{} }
func (m *SCShenfaAdvanced) String() string            { return proto.CompactTextString(m) }
func (*SCShenfaAdvanced) ProtoMessage()               {}
func (*SCShenfaAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{11} }

const Default_SCShenfaAdvanced_IsDouble bool = false

func (m *SCShenfaAdvanced) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *SCShenfaAdvanced) GetShenfaId() int32 {
	if m != nil && m.ShenfaId != nil {
		return *m.ShenfaId
	}
	return 0
}

func (m *SCShenfaAdvanced) GetBless() int32 {
	if m != nil && m.Bless != nil {
		return *m.Bless
	}
	return 0
}

func (m *SCShenfaAdvanced) GetBlessTime() int64 {
	if m != nil && m.BlessTime != nil {
		return *m.BlessTime
	}
	return 0
}

func (m *SCShenfaAdvanced) GetAdvancedType() int32 {
	if m != nil && m.AdvancedType != nil {
		return *m.AdvancedType
	}
	return 0
}

func (m *SCShenfaAdvanced) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return Default_SCShenfaAdvanced_IsDouble
}

func (m *SCShenfaAdvanced) GetTotalBless() int32 {
	if m != nil && m.TotalBless != nil {
		return *m.TotalBless
	}
	return 0
}

type CSShenfaHidden struct {
	Hidden           *bool  `protobuf:"varint,1,req,name=hidden" json:"hidden,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShenfaHidden) Reset()                    { *m = CSShenfaHidden{} }
func (m *CSShenfaHidden) String() string            { return proto.CompactTextString(m) }
func (*CSShenfaHidden) ProtoMessage()               {}
func (*CSShenfaHidden) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{12} }

func (m *CSShenfaHidden) GetHidden() bool {
	if m != nil && m.Hidden != nil {
		return *m.Hidden
	}
	return false
}

type SCShenfaHidden struct {
	Hidden           *bool  `protobuf:"varint,1,req,name=hidden" json:"hidden,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShenfaHidden) Reset()                    { *m = SCShenfaHidden{} }
func (m *SCShenfaHidden) String() string            { return proto.CompactTextString(m) }
func (*SCShenfaHidden) ProtoMessage()               {}
func (*SCShenfaHidden) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{13} }

func (m *SCShenfaHidden) GetHidden() bool {
	if m != nil && m.Hidden != nil {
		return *m.Hidden
	}
	return false
}

type CSShenFaUpstar struct {
	ShenFaId         *int32 `protobuf:"varint,1,req,name=shenFaId" json:"shenFaId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShenFaUpstar) Reset()                    { *m = CSShenFaUpstar{} }
func (m *CSShenFaUpstar) String() string            { return proto.CompactTextString(m) }
func (*CSShenFaUpstar) ProtoMessage()               {}
func (*CSShenFaUpstar) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{14} }

func (m *CSShenFaUpstar) GetShenFaId() int32 {
	if m != nil && m.ShenFaId != nil {
		return *m.ShenFaId
	}
	return 0
}

type SCShenFaUpstar struct {
	ShenFaId         *int32 `protobuf:"varint,1,req,name=shenFaId" json:"shenFaId,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	UpPro            *int32 `protobuf:"varint,3,req,name=upPro" json:"upPro,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShenFaUpstar) Reset()                    { *m = SCShenFaUpstar{} }
func (m *SCShenFaUpstar) String() string            { return proto.CompactTextString(m) }
func (*SCShenFaUpstar) ProtoMessage()               {}
func (*SCShenFaUpstar) Descriptor() ([]byte, []int) { return fileDescriptor96, []int{15} }

func (m *SCShenFaUpstar) GetShenFaId() int32 {
	if m != nil && m.ShenFaId != nil {
		return *m.ShenFaId
	}
	return 0
}

func (m *SCShenFaUpstar) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCShenFaUpstar) GetUpPro() int32 {
	if m != nil && m.UpPro != nil {
		return *m.UpPro
	}
	return 0
}

func init() {
	proto.RegisterType((*ShenfaInfo)(nil), "ui.ShenfaInfo")
	proto.RegisterType((*ShenFaSkinInfo)(nil), "ui.ShenFaSkinInfo")
	proto.RegisterType((*CSShenfaGet)(nil), "ui.CSShenfaGet")
	proto.RegisterType((*SCShenfaGet)(nil), "ui.SCShenfaGet")
	proto.RegisterType((*CSShenfaUnrealDan)(nil), "ui.CSShenfaUnrealDan")
	proto.RegisterType((*SCShenfaUnrealDan)(nil), "ui.SCShenfaUnrealDan")
	proto.RegisterType((*CSShenfaUnreal)(nil), "ui.CSShenfaUnreal")
	proto.RegisterType((*SCShenfaUnreal)(nil), "ui.SCShenfaUnreal")
	proto.RegisterType((*CSShenfaUnload)(nil), "ui.CSShenfaUnload")
	proto.RegisterType((*SCShenfaUnload)(nil), "ui.SCShenfaUnload")
	proto.RegisterType((*CSShenfaAdvanced)(nil), "ui.CSShenfaAdvanced")
	proto.RegisterType((*SCShenfaAdvanced)(nil), "ui.SCShenfaAdvanced")
	proto.RegisterType((*CSShenfaHidden)(nil), "ui.CSShenfaHidden")
	proto.RegisterType((*SCShenfaHidden)(nil), "ui.SCShenfaHidden")
	proto.RegisterType((*CSShenFaUpstar)(nil), "ui.CSShenFaUpstar")
	proto.RegisterType((*SCShenFaUpstar)(nil), "ui.SCShenFaUpstar")
}

func init() { proto.RegisterFile("shenfa.proto", fileDescriptor96) }

var fileDescriptor96 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x14, 0x54, 0x36, 0x4d, 0x9b, 0xbe, 0x74, 0x57, 0xbb, 0x06, 0x09, 0x1f, 0x2d, 0xab, 0x87, 0x88,
	0xc3, 0x1e, 0x10, 0x27, 0x0e, 0x08, 0x68, 0xb5, 0xb0, 0x52, 0x0f, 0x95, 0xd2, 0x7e, 0x80, 0x4b,
	0xbc, 0xad, 0x55, 0xd7, 0x8e, 0xe2, 0xa4, 0x12, 0x37, 0x7e, 0x84, 0x4f, 0xe2, 0x9f, 0x50, 0x9e,
	0x6b, 0x6f, 0xd3, 0x65, 0x25, 0x0e, 0x1c, 0x33, 0x9a, 0x37, 0xf3, 0xe6, 0x4d, 0x0c, 0x27, 0xee,
	0x4e, 0x9a, 0x8d, 0x58, 0x36, 0xad, 0xed, 0x2c, 0x99, 0xf4, 0x8a, 0xff, 0x4c, 0x00, 0x2a, 0x04,
	0xd7, 0x66, 0x63, 0x09, 0x01, 0x10, 0xf5, 0xa3, 0x30, 0xdf, 0x65, 0xbd, 0xae, 0x69, 0xc2, 0x26,
	0x65, 0x46, 0xe6, 0x90, 0xfb, 0xb1, 0x75, 0x4d, 0x27, 0x88, 0xbc, 0x82, 0xa2, 0x37, 0xad, 0x14,
	0xfa, 0x42, 0x3e, 0x4a, 0x4d, 0x53, 0x04, 0x17, 0x70, 0xec, 0xc1, 0xcb, 0xd6, 0xd2, 0x03, 0x84,
	0x4e, 0x21, 0x77, 0xf7, 0xca, 0x5c, 0x28, 0xd7, 0xd1, 0x8c, 0xa5, 0x65, 0xf1, 0x8e, 0x2c, 0x7b,
	0xb5, 0x1c, 0xfc, 0x56, 0xa2, 0xba, 0x57, 0x66, 0xf0, 0xe4, 0x1f, 0x61, 0x36, 0x46, 0x82, 0xe3,
	0x4a, 0xc4, 0x1d, 0xa6, 0x90, 0x69, 0xf4, 0xf2, 0x0b, 0x14, 0x90, 0x36, 0xad, 0xf5, 0xc6, 0x7c,
	0x0a, 0xc5, 0x59, 0xe5, 0x33, 0x7c, 0x95, 0x1d, 0xff, 0x9d, 0x40, 0x51, 0x9d, 0xc5, 0xef, 0xff,
	0x1c, 0x89, 0x00, 0x3c, 0xf1, 0x42, 0x28, 0x5c, 0xee, 0x46, 0x4b, 0xe7, 0xe8, 0x21, 0x4b, 0xfc,
	0x14, 0x7e, 0x5e, 0xa9, 0x07, 0x49, 0x8f, 0x58, 0x52, 0xa6, 0x64, 0x06, 0x87, 0x77, 0xaa, 0xae,
	0xa5, 0xa1, 0x39, 0x9b, 0x94, 0x39, 0x79, 0x0b, 0x33, 0x17, 0x23, 0xa3, 0xd2, 0xf1, 0xde, 0xf3,
	0x30, 0x58, 0x84, 0x78, 0xd7, 0xe8, 0x7c, 0x2e, 0xcc, 0x70, 0x00, 0xd3, 0x3f, 0xf8, 0x34, 0xfc,
	0x3d, 0x2c, 0x42, 0xe0, 0x2d, 0x23, 0x5e, 0x2c, 0x26, 0x6e, 0x5a, 0x7b, 0xdb, 0x0e, 0x6b, 0x62,
	0x62, 0xce, 0x61, 0x36, 0xd6, 0x1d, 0x5d, 0x25, 0x09, 0x9c, 0xb1, 0xf2, 0x0b, 0x4e, 0x52, 0x66,
	0x7c, 0xfe, 0x5c, 0x47, 0x5b, 0x51, 0x8f, 0xa7, 0x06, 0xe4, 0x2f, 0xca, 0xa7, 0x30, 0x0f, 0x53,
	0x9f, 0x9f, 0xda, 0x19, 0x58, 0xa2, 0xef, 0xec, 0x4a, 0x8b, 0x5b, 0x64, 0xe5, 0xfc, 0x57, 0x02,
	0xf3, 0x20, 0x15, 0x69, 0xff, 0x56, 0x68, 0x2c, 0x25, 0xdd, 0x2d, 0xe5, 0x00, 0x4b, 0x79, 0x0d,
	0x27, 0x41, 0xe7, 0xea, 0x47, 0x23, 0x69, 0x86, 0x73, 0x6f, 0x20, 0x57, 0xee, 0xdc, 0xf6, 0x37,
	0x5a, 0x62, 0x9f, 0xf9, 0x87, 0x6c, 0x23, 0xb4, 0x93, 0x83, 0x6d, 0x67, 0x3b, 0xa1, 0xbf, 0xa0,
	0xea, 0x11, 0x66, 0x67, 0xdb, 0xec, 0xdf, 0xb0, 0xdf, 0x67, 0x4d, 0xfb, 0x04, 0x6c, 0x7b, 0x8b,
	0x3d, 0x8c, 0xd8, 0xc3, 0x4a, 0x5c, 0x37, 0xae, 0x13, 0xed, 0xee, 0xef, 0xcf, 0x3f, 0x05, 0x95,
	0xfd, 0x9c, 0x97, 0x4f, 0x64, 0x0a, 0x59, 0xdf, 0x5c, 0x86, 0x47, 0xf2, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x72, 0xd4, 0x94, 0x58, 0xfa, 0x03, 0x00, 0x00,
}