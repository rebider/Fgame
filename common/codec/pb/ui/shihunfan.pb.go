// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shihunfan.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ShihunfanInfo struct {
	AdvancedId       *int32 `protobuf:"varint,1,req,name=advancedId" json:"advancedId,omitempty"`
	Bless            *int32 `protobuf:"varint,2,opt,name=bless" json:"bless,omitempty"`
	BlessTime        *int64 `protobuf:"varint,3,opt,name=blessTime" json:"blessTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShihunfanInfo) Reset()                    { *m = ShihunfanInfo{} }
func (m *ShihunfanInfo) String() string            { return proto.CompactTextString(m) }
func (*ShihunfanInfo) ProtoMessage()               {}
func (*ShihunfanInfo) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{0} }

func (m *ShihunfanInfo) GetAdvancedId() int32 {
	if m != nil && m.AdvancedId != nil {
		return *m.AdvancedId
	}
	return 0
}

func (m *ShihunfanInfo) GetBless() int32 {
	if m != nil && m.Bless != nil {
		return *m.Bless
	}
	return 0
}

func (m *ShihunfanInfo) GetBlessTime() int64 {
	if m != nil && m.BlessTime != nil {
		return *m.BlessTime
	}
	return 0
}

type ShihunfanDanInfo struct {
	Level            *int32 `protobuf:"varint,1,req,name=level" json:"level,omitempty"`
	Progress         *int32 `protobuf:"varint,2,opt,name=progress" json:"progress,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShihunfanDanInfo) Reset()                    { *m = ShihunfanDanInfo{} }
func (m *ShihunfanDanInfo) String() string            { return proto.CompactTextString(m) }
func (*ShihunfanDanInfo) ProtoMessage()               {}
func (*ShihunfanDanInfo) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{1} }

func (m *ShihunfanDanInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *ShihunfanDanInfo) GetProgress() int32 {
	if m != nil && m.Progress != nil {
		return *m.Progress
	}
	return 0
}

type CSShihunfanGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShihunfanGet) Reset()                    { *m = CSShihunfanGet{} }
func (m *CSShihunfanGet) String() string            { return proto.CompactTextString(m) }
func (*CSShihunfanGet) ProtoMessage()               {}
func (*CSShihunfanGet) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{2} }

type SCShihunfanGet struct {
	ShihunfanInfo    *ShihunfanInfo    `protobuf:"bytes,1,req,name=shihunfanInfo" json:"shihunfanInfo,omitempty"`
	ShihunfandanInfo *ShihunfanDanInfo `protobuf:"bytes,2,req,name=shihunfandanInfo" json:"shihunfandanInfo,omitempty"`
	ChargeVal        *int32            `protobuf:"varint,3,req,name=chargeVal" json:"chargeVal,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SCShihunfanGet) Reset()                    { *m = SCShihunfanGet{} }
func (m *SCShihunfanGet) String() string            { return proto.CompactTextString(m) }
func (*SCShihunfanGet) ProtoMessage()               {}
func (*SCShihunfanGet) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{3} }

func (m *SCShihunfanGet) GetShihunfanInfo() *ShihunfanInfo {
	if m != nil {
		return m.ShihunfanInfo
	}
	return nil
}

func (m *SCShihunfanGet) GetShihunfandanInfo() *ShihunfanDanInfo {
	if m != nil {
		return m.ShihunfandanInfo
	}
	return nil
}

func (m *SCShihunfanGet) GetChargeVal() int32 {
	if m != nil && m.ChargeVal != nil {
		return *m.ChargeVal
	}
	return 0
}

type SCShihunfanChargeVary struct {
	ChargeVal        *int32 `protobuf:"varint,1,req,name=chargeVal" json:"chargeVal,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCShihunfanChargeVary) Reset()                    { *m = SCShihunfanChargeVary{} }
func (m *SCShihunfanChargeVary) String() string            { return proto.CompactTextString(m) }
func (*SCShihunfanChargeVary) ProtoMessage()               {}
func (*SCShihunfanChargeVary) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{4} }

func (m *SCShihunfanChargeVary) GetChargeVal() int32 {
	if m != nil && m.ChargeVal != nil {
		return *m.ChargeVal
	}
	return 0
}

type CSShihunfanAdvanced struct {
	BuyFlag          *bool  `protobuf:"varint,1,req,name=buyFlag" json:"buyFlag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShihunfanAdvanced) Reset()                    { *m = CSShihunfanAdvanced{} }
func (m *CSShihunfanAdvanced) String() string            { return proto.CompactTextString(m) }
func (*CSShihunfanAdvanced) ProtoMessage()               {}
func (*CSShihunfanAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{5} }

func (m *CSShihunfanAdvanced) GetBuyFlag() bool {
	if m != nil && m.BuyFlag != nil {
		return *m.BuyFlag
	}
	return false
}

type SCShihunfanAdvanced struct {
	ShihunfanInfo    *ShihunfanInfo `protobuf:"bytes,1,req,name=shihunfanInfo" json:"shihunfanInfo,omitempty"`
	IsDouble         *bool          `protobuf:"varint,2,opt,name=isDouble,def=0" json:"isDouble,omitempty"`
	AdvancedType     *int32         `protobuf:"varint,3,req,name=advancedType" json:"advancedType,omitempty"`
	ChargeVal        *int32         `protobuf:"varint,4,req,name=chargeVal" json:"chargeVal,omitempty"`
	BuyFlag          *bool          `protobuf:"varint,5,opt,name=buyFlag,def=0" json:"buyFlag,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SCShihunfanAdvanced) Reset()                    { *m = SCShihunfanAdvanced{} }
func (m *SCShihunfanAdvanced) String() string            { return proto.CompactTextString(m) }
func (*SCShihunfanAdvanced) ProtoMessage()               {}
func (*SCShihunfanAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{6} }

const Default_SCShihunfanAdvanced_IsDouble bool = false
const Default_SCShihunfanAdvanced_BuyFlag bool = false

func (m *SCShihunfanAdvanced) GetShihunfanInfo() *ShihunfanInfo {
	if m != nil {
		return m.ShihunfanInfo
	}
	return nil
}

func (m *SCShihunfanAdvanced) GetIsDouble() bool {
	if m != nil && m.IsDouble != nil {
		return *m.IsDouble
	}
	return Default_SCShihunfanAdvanced_IsDouble
}

func (m *SCShihunfanAdvanced) GetAdvancedType() int32 {
	if m != nil && m.AdvancedType != nil {
		return *m.AdvancedType
	}
	return 0
}

func (m *SCShihunfanAdvanced) GetChargeVal() int32 {
	if m != nil && m.ChargeVal != nil {
		return *m.ChargeVal
	}
	return 0
}

func (m *SCShihunfanAdvanced) GetBuyFlag() bool {
	if m != nil && m.BuyFlag != nil {
		return *m.BuyFlag
	}
	return Default_SCShihunfanAdvanced_BuyFlag
}

type CSShihunfanDanAdvanced struct {
	Num              *int32 `protobuf:"varint,1,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSShihunfanDanAdvanced) Reset()                    { *m = CSShihunfanDanAdvanced{} }
func (m *CSShihunfanDanAdvanced) String() string            { return proto.CompactTextString(m) }
func (*CSShihunfanDanAdvanced) ProtoMessage()               {}
func (*CSShihunfanDanAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{7} }

func (m *CSShihunfanDanAdvanced) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type SCShihunfanDanAdvanced struct {
	ShihunfandanInfo *ShihunfanDanInfo `protobuf:"bytes,1,req,name=shihunfandanInfo" json:"shihunfandanInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SCShihunfanDanAdvanced) Reset()                    { *m = SCShihunfanDanAdvanced{} }
func (m *SCShihunfanDanAdvanced) String() string            { return proto.CompactTextString(m) }
func (*SCShihunfanDanAdvanced) ProtoMessage()               {}
func (*SCShihunfanDanAdvanced) Descriptor() ([]byte, []int) { return fileDescriptor102, []int{8} }

func (m *SCShihunfanDanAdvanced) GetShihunfandanInfo() *ShihunfanDanInfo {
	if m != nil {
		return m.ShihunfandanInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ShihunfanInfo)(nil), "ui.ShihunfanInfo")
	proto.RegisterType((*ShihunfanDanInfo)(nil), "ui.ShihunfanDanInfo")
	proto.RegisterType((*CSShihunfanGet)(nil), "ui.CSShihunfanGet")
	proto.RegisterType((*SCShihunfanGet)(nil), "ui.SCShihunfanGet")
	proto.RegisterType((*SCShihunfanChargeVary)(nil), "ui.SCShihunfanChargeVary")
	proto.RegisterType((*CSShihunfanAdvanced)(nil), "ui.CSShihunfanAdvanced")
	proto.RegisterType((*SCShihunfanAdvanced)(nil), "ui.SCShihunfanAdvanced")
	proto.RegisterType((*CSShihunfanDanAdvanced)(nil), "ui.CSShihunfanDanAdvanced")
	proto.RegisterType((*SCShihunfanDanAdvanced)(nil), "ui.SCShihunfanDanAdvanced")
}

func init() { proto.RegisterFile("shihunfan.proto", fileDescriptor102) }

var fileDescriptor102 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4f, 0x83, 0x30,
	0x18, 0xc7, 0x03, 0xd8, 0xc8, 0x9e, 0xc9, 0xc6, 0xba, 0x39, 0x39, 0x2e, 0x4d, 0x34, 0xc4, 0x03,
	0x07, 0xbd, 0x79, 0x33, 0x9b, 0x2f, 0x3b, 0x6f, 0xf1, 0xde, 0x8d, 0xf2, 0x92, 0x74, 0x40, 0xe8,
	0xba, 0x84, 0x83, 0x1f, 0xc4, 0x6f, 0x6b, 0xac, 0x05, 0x4b, 0xa2, 0x89, 0x47, 0x1e, 0x7e, 0xfd,
	0xf7, 0xf7, 0x7f, 0x0a, 0x63, 0x91, 0xe5, 0x99, 0x2c, 0x12, 0x5a, 0x44, 0x55, 0x5d, 0x1e, 0x4b,
	0x6c, 0xcb, 0x9c, 0x3c, 0x81, 0xb7, 0x69, 0xc7, 0xeb, 0x22, 0x29, 0x31, 0x06, 0xa0, 0xf1, 0x89,
	0x16, 0x7b, 0x16, 0xaf, 0xe3, 0xc0, 0x5a, 0xd8, 0x21, 0xc2, 0x1e, 0xa0, 0x1d, 0x67, 0x42, 0x04,
	0xf6, 0xc2, 0x0a, 0x11, 0x9e, 0xc0, 0x40, 0x7d, 0x6e, 0xf3, 0x03, 0x0b, 0x9c, 0x85, 0x15, 0x3a,
	0xe4, 0x1e, 0xfc, 0x2e, 0x66, 0xa5, 0x93, 0x3c, 0x40, 0x9c, 0x9d, 0x18, 0xd7, 0x21, 0x3e, 0xb8,
	0x55, 0x5d, 0xa6, 0x75, 0x97, 0x43, 0x7c, 0x18, 0x2d, 0x37, 0xdd, 0xb1, 0x17, 0x76, 0x24, 0xef,
	0x30, 0xda, 0x2c, 0xcd, 0x09, 0x0e, 0xc1, 0x13, 0xa6, 0x9f, 0x0a, 0x1b, 0xde, 0x4d, 0x22, 0x99,
	0x47, 0x7d, 0xf1, 0x08, 0xfc, 0x8e, 0x8c, 0x35, 0x6c, 0x2b, 0x78, 0xd6, 0x83, 0x5b, 0xbd, 0x09,
	0x0c, 0xf6, 0x19, 0xad, 0x53, 0xf6, 0x46, 0x79, 0xe0, 0x7c, 0x29, 0x92, 0x5b, 0xb8, 0x34, 0xae,
	0x5f, 0xea, 0xbf, 0x75, 0xd3, 0x67, 0x55, 0x1d, 0x72, 0x03, 0x53, 0x43, 0xfe, 0x51, 0xaf, 0x0c,
	0x8f, 0xe1, 0x7c, 0x27, 0x9b, 0x67, 0x4e, 0x53, 0xc5, 0xb9, 0xe4, 0xc3, 0x82, 0xa9, 0x11, 0xda,
	0x81, 0xff, 0x2f, 0x76, 0x05, 0x6e, 0x2e, 0x56, 0xa5, 0xdc, 0x71, 0xa6, 0x16, 0xe7, 0x3e, 0xa0,
	0x84, 0x72, 0xc1, 0xf0, 0x0c, 0x2e, 0xda, 0xa7, 0xda, 0x36, 0x15, 0xfb, 0x2e, 0xd1, 0x77, 0x3d,
	0x53, 0xa3, 0xf9, 0x8f, 0x14, 0x32, 0x02, 0xc8, 0x35, 0xcc, 0x8d, 0x0e, 0x2b, 0xc3, 0x6e, 0x08,
	0x4e, 0x21, 0x0f, 0xba, 0xea, 0x2b, 0xcc, 0x8d, 0x06, 0x26, 0xf6, 0xdb, 0xce, 0xad, 0xbf, 0x77,
	0xfe, 0x19, 0x00, 0x00, 0xff, 0xff, 0x53, 0xaa, 0xfa, 0x46, 0x83, 0x02, 0x00, 0x00,
}