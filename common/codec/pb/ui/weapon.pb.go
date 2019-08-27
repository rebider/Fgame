// Code generated by protoc-gen-go. DO NOT EDIT.
// source: weapon.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WeaponInfo struct {
	WeaponId         *int32 `protobuf:"varint,1,req,name=weaponId" json:"weaponId,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	CulLevel         *int32 `protobuf:"varint,3,req,name=culLevel" json:"culLevel,omitempty"`
	CulPro           *int32 `protobuf:"varint,4,req,name=culPro" json:"culPro,omitempty"`
	State            *int32 `protobuf:"varint,5,req,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *WeaponInfo) Reset()                    { *m = WeaponInfo{} }
func (m *WeaponInfo) String() string            { return proto.CompactTextString(m) }
func (*WeaponInfo) ProtoMessage()               {}
func (*WeaponInfo) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{0} }

func (m *WeaponInfo) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

func (m *WeaponInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *WeaponInfo) GetCulLevel() int32 {
	if m != nil && m.CulLevel != nil {
		return *m.CulLevel
	}
	return 0
}

func (m *WeaponInfo) GetCulPro() int32 {
	if m != nil && m.CulPro != nil {
		return *m.CulPro
	}
	return 0
}

func (m *WeaponInfo) GetState() int32 {
	if m != nil && m.State != nil {
		return *m.State
	}
	return 0
}

type AllWeaponInfo struct {
	WeaponWear       *int32        `protobuf:"varint,1,req,name=weaponWear" json:"weaponWear,omitempty"`
	WeaponList       []*WeaponInfo `protobuf:"bytes,2,rep,name=weaponList" json:"weaponList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *AllWeaponInfo) Reset()                    { *m = AllWeaponInfo{} }
func (m *AllWeaponInfo) String() string            { return proto.CompactTextString(m) }
func (*AllWeaponInfo) ProtoMessage()               {}
func (*AllWeaponInfo) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{1} }

func (m *AllWeaponInfo) GetWeaponWear() int32 {
	if m != nil && m.WeaponWear != nil {
		return *m.WeaponWear
	}
	return 0
}

func (m *AllWeaponInfo) GetWeaponList() []*WeaponInfo {
	if m != nil {
		return m.WeaponList
	}
	return nil
}

type CSWeaponGet struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWeaponGet) Reset()                    { *m = CSWeaponGet{} }
func (m *CSWeaponGet) String() string            { return proto.CompactTextString(m) }
func (*CSWeaponGet) ProtoMessage()               {}
func (*CSWeaponGet) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{2} }

type SCWeaponGet struct {
	WeaponWear       *int32        `protobuf:"varint,1,req,name=weaponWear" json:"weaponWear,omitempty"`
	WeaponList       []*WeaponInfo `protobuf:"bytes,2,rep,name=weaponList" json:"weaponList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SCWeaponGet) Reset()                    { *m = SCWeaponGet{} }
func (m *SCWeaponGet) String() string            { return proto.CompactTextString(m) }
func (*SCWeaponGet) ProtoMessage()               {}
func (*SCWeaponGet) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{3} }

func (m *SCWeaponGet) GetWeaponWear() int32 {
	if m != nil && m.WeaponWear != nil {
		return *m.WeaponWear
	}
	return 0
}

func (m *SCWeaponGet) GetWeaponList() []*WeaponInfo {
	if m != nil {
		return m.WeaponList
	}
	return nil
}

type CSWeaponActive struct {
	WeaponId         *int32 `protobuf:"varint,1,req,name=weaponId" json:"weaponId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWeaponActive) Reset()                    { *m = CSWeaponActive{} }
func (m *CSWeaponActive) String() string            { return proto.CompactTextString(m) }
func (*CSWeaponActive) ProtoMessage()               {}
func (*CSWeaponActive) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{4} }

func (m *CSWeaponActive) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

type SCWeaponActive struct {
	WeaponId         *int32 `protobuf:"varint,1,req,name=weaponId" json:"weaponId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCWeaponActive) Reset()                    { *m = SCWeaponActive{} }
func (m *SCWeaponActive) String() string            { return proto.CompactTextString(m) }
func (*SCWeaponActive) ProtoMessage()               {}
func (*SCWeaponActive) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{5} }

func (m *SCWeaponActive) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

type CSWeaponEatDan struct {
	WeaponId         *int32 `protobuf:"varint,1,req,name=weaponId" json:"weaponId,omitempty"`
	Num              *int32 `protobuf:"varint,2,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWeaponEatDan) Reset()                    { *m = CSWeaponEatDan{} }
func (m *CSWeaponEatDan) String() string            { return proto.CompactTextString(m) }
func (*CSWeaponEatDan) ProtoMessage()               {}
func (*CSWeaponEatDan) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{6} }

func (m *CSWeaponEatDan) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

func (m *CSWeaponEatDan) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type SCWeaponEatDan struct {
	WeaponId         *int32 `protobuf:"varint,1,req,name=weaponId" json:"weaponId,omitempty"`
	CulLevel         *int32 `protobuf:"varint,2,req,name=culLevel" json:"culLevel,omitempty"`
	CulPro           *int32 `protobuf:"varint,3,req,name=culPro" json:"culPro,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCWeaponEatDan) Reset()                    { *m = SCWeaponEatDan{} }
func (m *SCWeaponEatDan) String() string            { return proto.CompactTextString(m) }
func (*SCWeaponEatDan) ProtoMessage()               {}
func (*SCWeaponEatDan) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{7} }

func (m *SCWeaponEatDan) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

func (m *SCWeaponEatDan) GetCulLevel() int32 {
	if m != nil && m.CulLevel != nil {
		return *m.CulLevel
	}
	return 0
}

func (m *SCWeaponEatDan) GetCulPro() int32 {
	if m != nil && m.CulPro != nil {
		return *m.CulPro
	}
	return 0
}

type CSWeaponUpstar struct {
	WeaponId         *int32 `protobuf:"varint,1,req,name=weaponId" json:"weaponId,omitempty"`
	AutoFlag         *bool  `protobuf:"varint,2,req,name=autoFlag" json:"autoFlag,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWeaponUpstar) Reset()                    { *m = CSWeaponUpstar{} }
func (m *CSWeaponUpstar) String() string            { return proto.CompactTextString(m) }
func (*CSWeaponUpstar) ProtoMessage()               {}
func (*CSWeaponUpstar) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{8} }

func (m *CSWeaponUpstar) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

func (m *CSWeaponUpstar) GetAutoFlag() bool {
	if m != nil && m.AutoFlag != nil {
		return *m.AutoFlag
	}
	return false
}

type SCWeaponUpstar struct {
	WeaponId         *int32 `protobuf:"varint,1,req,name=weaponId" json:"weaponId,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	UpPro            *int32 `protobuf:"varint,3,req,name=upPro" json:"upPro,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCWeaponUpstar) Reset()                    { *m = SCWeaponUpstar{} }
func (m *SCWeaponUpstar) String() string            { return proto.CompactTextString(m) }
func (*SCWeaponUpstar) ProtoMessage()               {}
func (*SCWeaponUpstar) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{9} }

func (m *SCWeaponUpstar) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

func (m *SCWeaponUpstar) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCWeaponUpstar) GetUpPro() int32 {
	if m != nil && m.UpPro != nil {
		return *m.UpPro
	}
	return 0
}

type CSWeaponAwaken struct {
	WeaponId         *int32 `protobuf:"varint,1,req,name=weaponId" json:"weaponId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWeaponAwaken) Reset()                    { *m = CSWeaponAwaken{} }
func (m *CSWeaponAwaken) String() string            { return proto.CompactTextString(m) }
func (*CSWeaponAwaken) ProtoMessage()               {}
func (*CSWeaponAwaken) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{10} }

func (m *CSWeaponAwaken) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

type SCWeaponAwaken struct {
	Result           *int32 `protobuf:"varint,1,req,name=result" json:"result,omitempty"`
	WeaponId         *int32 `protobuf:"varint,2,req,name=weaponId" json:"weaponId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCWeaponAwaken) Reset()                    { *m = SCWeaponAwaken{} }
func (m *SCWeaponAwaken) String() string            { return proto.CompactTextString(m) }
func (*SCWeaponAwaken) ProtoMessage()               {}
func (*SCWeaponAwaken) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{11} }

func (m *SCWeaponAwaken) GetResult() int32 {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return 0
}

func (m *SCWeaponAwaken) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

type CSWeaponWear struct {
	WeaponWear       *int32 `protobuf:"varint,1,req,name=weaponWear" json:"weaponWear,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWeaponWear) Reset()                    { *m = CSWeaponWear{} }
func (m *CSWeaponWear) String() string            { return proto.CompactTextString(m) }
func (*CSWeaponWear) ProtoMessage()               {}
func (*CSWeaponWear) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{12} }

func (m *CSWeaponWear) GetWeaponWear() int32 {
	if m != nil && m.WeaponWear != nil {
		return *m.WeaponWear
	}
	return 0
}

type SCWeaponWear struct {
	WeaponWear       *int32 `protobuf:"varint,1,req,name=weaponWear" json:"weaponWear,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCWeaponWear) Reset()                    { *m = SCWeaponWear{} }
func (m *SCWeaponWear) String() string            { return proto.CompactTextString(m) }
func (*SCWeaponWear) ProtoMessage()               {}
func (*SCWeaponWear) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{13} }

func (m *SCWeaponWear) GetWeaponWear() int32 {
	if m != nil && m.WeaponWear != nil {
		return *m.WeaponWear
	}
	return 0
}

type CSWeaponUnLoad struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWeaponUnLoad) Reset()                    { *m = CSWeaponUnLoad{} }
func (m *CSWeaponUnLoad) String() string            { return proto.CompactTextString(m) }
func (*CSWeaponUnLoad) ProtoMessage()               {}
func (*CSWeaponUnLoad) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{14} }

type SCWeaponUnLoad struct {
	WeaponWear       *int32 `protobuf:"varint,1,req,name=weaponWear" json:"weaponWear,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCWeaponUnLoad) Reset()                    { *m = SCWeaponUnLoad{} }
func (m *SCWeaponUnLoad) String() string            { return proto.CompactTextString(m) }
func (*SCWeaponUnLoad) ProtoMessage()               {}
func (*SCWeaponUnLoad) Descriptor() ([]byte, []int) { return fileDescriptor126, []int{15} }

func (m *SCWeaponUnLoad) GetWeaponWear() int32 {
	if m != nil && m.WeaponWear != nil {
		return *m.WeaponWear
	}
	return 0
}

func init() {
	proto.RegisterType((*WeaponInfo)(nil), "ui.WeaponInfo")
	proto.RegisterType((*AllWeaponInfo)(nil), "ui.AllWeaponInfo")
	proto.RegisterType((*CSWeaponGet)(nil), "ui.CSWeaponGet")
	proto.RegisterType((*SCWeaponGet)(nil), "ui.SCWeaponGet")
	proto.RegisterType((*CSWeaponActive)(nil), "ui.CSWeaponActive")
	proto.RegisterType((*SCWeaponActive)(nil), "ui.SCWeaponActive")
	proto.RegisterType((*CSWeaponEatDan)(nil), "ui.CSWeaponEatDan")
	proto.RegisterType((*SCWeaponEatDan)(nil), "ui.SCWeaponEatDan")
	proto.RegisterType((*CSWeaponUpstar)(nil), "ui.CSWeaponUpstar")
	proto.RegisterType((*SCWeaponUpstar)(nil), "ui.SCWeaponUpstar")
	proto.RegisterType((*CSWeaponAwaken)(nil), "ui.CSWeaponAwaken")
	proto.RegisterType((*SCWeaponAwaken)(nil), "ui.SCWeaponAwaken")
	proto.RegisterType((*CSWeaponWear)(nil), "ui.CSWeaponWear")
	proto.RegisterType((*SCWeaponWear)(nil), "ui.SCWeaponWear")
	proto.RegisterType((*CSWeaponUnLoad)(nil), "ui.CSWeaponUnLoad")
	proto.RegisterType((*SCWeaponUnLoad)(nil), "ui.SCWeaponUnLoad")
}

func init() { proto.RegisterFile("weapon.proto", fileDescriptor126) }

var fileDescriptor126 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x3d, 0x6b, 0xf3, 0x30,
	0x10, 0xc7, 0x79, 0x94, 0xc7, 0xc1, 0x9c, 0x63, 0x63, 0x3c, 0x79, 0x0c, 0x47, 0x87, 0x4c, 0x2e,
	0x84, 0x7e, 0x80, 0x86, 0x24, 0x0d, 0x01, 0x0f, 0x85, 0x50, 0xb2, 0x15, 0x84, 0xa3, 0x16, 0x53,
	0xd5, 0x32, 0xb6, 0x94, 0x7c, 0xfd, 0x62, 0x39, 0x8a, 0x44, 0x6a, 0x2f, 0x1d, 0xef, 0xf8, 0xbf,
	0x48, 0xbf, 0x83, 0xd9, 0x85, 0xd1, 0x5a, 0x54, 0x59, 0xdd, 0x08, 0x29, 0x12, 0xa2, 0x4a, 0x7c,
	0x07, 0x38, 0xea, 0xdd, 0xbe, 0xfa, 0x10, 0x49, 0x0c, 0x7e, 0xaf, 0xd8, 0x9f, 0xd2, 0x7f, 0x73,
	0xb2, 0xf0, 0x92, 0x10, 0x3c, 0xce, 0xce, 0x8c, 0xa7, 0x44, 0x8f, 0x31, 0xf8, 0x85, 0xe2, 0xb9,
	0xde, 0x4c, 0xf4, 0x26, 0x82, 0x69, 0xa1, 0xf8, 0x6b, 0x23, 0xd2, 0xff, 0xc6, 0xd0, 0x4a, 0x2a,
	0x59, 0xea, 0x75, 0x23, 0xee, 0x20, 0x5c, 0x71, 0xee, 0x54, 0x24, 0x00, 0x7d, 0xc5, 0x91, 0xd1,
	0xe6, 0x5a, 0x82, 0x66, 0x97, 0x97, 0xad, 0x4c, 0xc9, 0x7c, 0xb2, 0x08, 0x96, 0x51, 0xa6, 0xca,
	0xcc, 0xfa, 0x30, 0x84, 0x60, 0x7d, 0xe8, 0xe7, 0x1d, 0x93, 0xb8, 0x85, 0xe0, 0xb0, 0xbe, 0x8d,
	0x7f, 0x4e, 0x45, 0x88, 0x4c, 0xea, 0xaa, 0x90, 0xe5, 0x99, 0xfd, 0x46, 0xd0, 0x69, 0x4c, 0xd5,
	0xa8, 0xe6, 0xd1, 0xe6, 0x6c, 0xa9, 0xdc, 0xd0, 0x6a, 0x00, 0x65, 0x00, 0x93, 0x4a, 0x7d, 0xf7,
	0x20, 0x71, 0x63, 0x43, 0x47, 0x0d, 0x2e, 0x6c, 0x72, 0x07, 0x5b, 0xc3, 0xc7, 0x27, 0x5b, 0xfb,
	0x56, 0xb7, 0x92, 0x36, 0xc3, 0x29, 0x54, 0x49, 0xf1, 0xc2, 0xe9, 0xa7, 0x4e, 0xf1, 0xf1, 0xd9,
	0x76, 0x8f, 0xba, 0xee, 0xee, 0x1e, 0x82, 0xa7, 0x6a, 0xdb, 0xeb, 0x62, 0xbb, 0xd0, 0x2f, 0x36,
	0xf0, 0x7a, 0x5c, 0x3a, 0xd8, 0x7a, 0x4d, 0x04, 0xd3, 0x86, 0xb5, 0x8a, 0x4b, 0xfb, 0xb2, 0x9b,
	0x87, 0x5c, 0x73, 0x67, 0x26, 0xb7, 0x3b, 0xe4, 0xd0, 0x59, 0x3b, 0x8d, 0xc9, 0x1d, 0xd5, 0xc4,
	0x0e, 0x97, 0x2a, 0x17, 0xf4, 0x84, 0x0f, 0xce, 0x9f, 0xf5, 0x66, 0xc8, 0xf7, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x4e, 0x1f, 0x20, 0x2f, 0x20, 0x03, 0x00, 0x00,
}
