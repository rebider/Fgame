// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wushuangweapon.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WushuangWeaponBodyPos int32

const (
	WushuangWeaponBodyPos_Weapon   WushuangWeaponBodyPos = 1
	WushuangWeaponBodyPos_Clothes  WushuangWeaponBodyPos = 2
	WushuangWeaponBodyPos_Head     WushuangWeaponBodyPos = 3
	WushuangWeaponBodyPos_Shoes    WushuangWeaponBodyPos = 4
	WushuangWeaponBodyPos_Necklace WushuangWeaponBodyPos = 5
	WushuangWeaponBodyPos_Pendant  WushuangWeaponBodyPos = 6
)

var WushuangWeaponBodyPos_name = map[int32]string{
	1: "Weapon",
	2: "Clothes",
	3: "Head",
	4: "Shoes",
	5: "Necklace",
	6: "Pendant",
}
var WushuangWeaponBodyPos_value = map[string]int32{
	"Weapon":   1,
	"Clothes":  2,
	"Head":     3,
	"Shoes":    4,
	"Necklace": 5,
	"Pendant":  6,
}

func (x WushuangWeaponBodyPos) Enum() *WushuangWeaponBodyPos {
	p := new(WushuangWeaponBodyPos)
	*p = x
	return p
}
func (x WushuangWeaponBodyPos) String() string {
	return proto.EnumName(WushuangWeaponBodyPos_name, int32(x))
}
func (x *WushuangWeaponBodyPos) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WushuangWeaponBodyPos_value, data, "WushuangWeaponBodyPos")
	if err != nil {
		return err
	}
	*x = WushuangWeaponBodyPos(value)
	return nil
}
func (WushuangWeaponBodyPos) EnumDescriptor() ([]byte, []int) { return fileDescriptor132, []int{0} }

type CSWushuangWeaponDevouring struct {
	BodyPos          *WushuangWeaponBodyPos `protobuf:"varint,1,req,name=bodyPos,enum=ui.WushuangWeaponBodyPos" json:"bodyPos,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *CSWushuangWeaponDevouring) Reset()                    { *m = CSWushuangWeaponDevouring{} }
func (m *CSWushuangWeaponDevouring) String() string            { return proto.CompactTextString(m) }
func (*CSWushuangWeaponDevouring) ProtoMessage()               {}
func (*CSWushuangWeaponDevouring) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{0} }

func (m *CSWushuangWeaponDevouring) GetBodyPos() WushuangWeaponBodyPos {
	if m != nil && m.BodyPos != nil {
		return *m.BodyPos
	}
	return WushuangWeaponBodyPos_Weapon
}

type SCWushuangWeaponDevouring struct {
	BodyPos          *WushuangWeaponBodyPos `protobuf:"varint,1,req,name=bodyPos,enum=ui.WushuangWeaponBodyPos" json:"bodyPos,omitempty"`
	Level            *int32                 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Experience       *int64                 `protobuf:"varint,3,req,name=experience" json:"experience,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SCWushuangWeaponDevouring) Reset()                    { *m = SCWushuangWeaponDevouring{} }
func (m *SCWushuangWeaponDevouring) String() string            { return proto.CompactTextString(m) }
func (*SCWushuangWeaponDevouring) ProtoMessage()               {}
func (*SCWushuangWeaponDevouring) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{1} }

func (m *SCWushuangWeaponDevouring) GetBodyPos() WushuangWeaponBodyPos {
	if m != nil && m.BodyPos != nil {
		return *m.BodyPos
	}
	return WushuangWeaponBodyPos_Weapon
}

func (m *SCWushuangWeaponDevouring) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCWushuangWeaponDevouring) GetExperience() int64 {
	if m != nil && m.Experience != nil {
		return *m.Experience
	}
	return 0
}

type CSWushuangWeaponBreakthrough struct {
	BodyPos          *WushuangWeaponBodyPos `protobuf:"varint,1,req,name=bodyPos,enum=ui.WushuangWeaponBodyPos" json:"bodyPos,omitempty"`
	IndexList        []int32                `protobuf:"varint,2,rep,name=indexList" json:"indexList,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *CSWushuangWeaponBreakthrough) Reset()                    { *m = CSWushuangWeaponBreakthrough{} }
func (m *CSWushuangWeaponBreakthrough) String() string            { return proto.CompactTextString(m) }
func (*CSWushuangWeaponBreakthrough) ProtoMessage()               {}
func (*CSWushuangWeaponBreakthrough) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{2} }

func (m *CSWushuangWeaponBreakthrough) GetBodyPos() WushuangWeaponBodyPos {
	if m != nil && m.BodyPos != nil {
		return *m.BodyPos
	}
	return WushuangWeaponBodyPos_Weapon
}

func (m *CSWushuangWeaponBreakthrough) GetIndexList() []int32 {
	if m != nil {
		return m.IndexList
	}
	return nil
}

type SCWushuangWeaponBreakthrough struct {
	BodyPos          *WushuangWeaponBodyPos `protobuf:"varint,1,req,name=bodyPos,enum=ui.WushuangWeaponBodyPos" json:"bodyPos,omitempty"`
	Level            *int32                 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SCWushuangWeaponBreakthrough) Reset()                    { *m = SCWushuangWeaponBreakthrough{} }
func (m *SCWushuangWeaponBreakthrough) String() string            { return proto.CompactTextString(m) }
func (*SCWushuangWeaponBreakthrough) ProtoMessage()               {}
func (*SCWushuangWeaponBreakthrough) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{3} }

func (m *SCWushuangWeaponBreakthrough) GetBodyPos() WushuangWeaponBodyPos {
	if m != nil && m.BodyPos != nil {
		return *m.BodyPos
	}
	return WushuangWeaponBodyPos_Weapon
}

func (m *SCWushuangWeaponBreakthrough) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

type CSWushuangWeaponPutOn struct {
	Index            *int32 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWushuangWeaponPutOn) Reset()                    { *m = CSWushuangWeaponPutOn{} }
func (m *CSWushuangWeaponPutOn) String() string            { return proto.CompactTextString(m) }
func (*CSWushuangWeaponPutOn) ProtoMessage()               {}
func (*CSWushuangWeaponPutOn) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{4} }

func (m *CSWushuangWeaponPutOn) GetIndex() int32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

type SCWushuangWeaponPutOn struct {
	BodyPos          *WushuangWeaponBodyPos `protobuf:"varint,1,req,name=bodyPos,enum=ui.WushuangWeaponBodyPos" json:"bodyPos,omitempty"`
	Level            *int32                 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Experience       *int64                 `protobuf:"varint,3,req,name=experience" json:"experience,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SCWushuangWeaponPutOn) Reset()                    { *m = SCWushuangWeaponPutOn{} }
func (m *SCWushuangWeaponPutOn) String() string            { return proto.CompactTextString(m) }
func (*SCWushuangWeaponPutOn) ProtoMessage()               {}
func (*SCWushuangWeaponPutOn) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{5} }

func (m *SCWushuangWeaponPutOn) GetBodyPos() WushuangWeaponBodyPos {
	if m != nil && m.BodyPos != nil {
		return *m.BodyPos
	}
	return WushuangWeaponBodyPos_Weapon
}

func (m *SCWushuangWeaponPutOn) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCWushuangWeaponPutOn) GetExperience() int64 {
	if m != nil && m.Experience != nil {
		return *m.Experience
	}
	return 0
}

type CSWushuangWeaponTakeOff struct {
	BodyPos          *WushuangWeaponBodyPos `protobuf:"varint,1,req,name=bodyPos,enum=ui.WushuangWeaponBodyPos" json:"bodyPos,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *CSWushuangWeaponTakeOff) Reset()                    { *m = CSWushuangWeaponTakeOff{} }
func (m *CSWushuangWeaponTakeOff) String() string            { return proto.CompactTextString(m) }
func (*CSWushuangWeaponTakeOff) ProtoMessage()               {}
func (*CSWushuangWeaponTakeOff) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{6} }

func (m *CSWushuangWeaponTakeOff) GetBodyPos() WushuangWeaponBodyPos {
	if m != nil && m.BodyPos != nil {
		return *m.BodyPos
	}
	return WushuangWeaponBodyPos_Weapon
}

type SCWushuangWeaponTakeOff struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCWushuangWeaponTakeOff) Reset()                    { *m = SCWushuangWeaponTakeOff{} }
func (m *SCWushuangWeaponTakeOff) String() string            { return proto.CompactTextString(m) }
func (*SCWushuangWeaponTakeOff) ProtoMessage()               {}
func (*SCWushuangWeaponTakeOff) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{7} }

type CSWushuangWeaponInfo struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSWushuangWeaponInfo) Reset()                    { *m = CSWushuangWeaponInfo{} }
func (m *CSWushuangWeaponInfo) String() string            { return proto.CompactTextString(m) }
func (*CSWushuangWeaponInfo) ProtoMessage()               {}
func (*CSWushuangWeaponInfo) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{8} }

type SCWushuangWeaponInfo struct {
	AllBodyPosInfo   []*WushuangWeaponBodyPosInfo `protobuf:"bytes,1,rep,name=allBodyPosInfo" json:"allBodyPosInfo,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *SCWushuangWeaponInfo) Reset()                    { *m = SCWushuangWeaponInfo{} }
func (m *SCWushuangWeaponInfo) String() string            { return proto.CompactTextString(m) }
func (*SCWushuangWeaponInfo) ProtoMessage()               {}
func (*SCWushuangWeaponInfo) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{9} }

func (m *SCWushuangWeaponInfo) GetAllBodyPosInfo() []*WushuangWeaponBodyPosInfo {
	if m != nil {
		return m.AllBodyPosInfo
	}
	return nil
}

type WushuangWeaponBodyPosInfo struct {
	BodyPos          *WushuangWeaponBodyPos `protobuf:"varint,1,req,name=bodyPos,enum=ui.WushuangWeaponBodyPos" json:"bodyPos,omitempty"`
	ItemId           *int32                 `protobuf:"varint,2,req,name=itemId" json:"itemId,omitempty"`
	Level            *int32                 `protobuf:"varint,3,req,name=level" json:"level,omitempty"`
	Experience       *int64                 `protobuf:"varint,4,req,name=experience" json:"experience,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *WushuangWeaponBodyPosInfo) Reset()                    { *m = WushuangWeaponBodyPosInfo{} }
func (m *WushuangWeaponBodyPosInfo) String() string            { return proto.CompactTextString(m) }
func (*WushuangWeaponBodyPosInfo) ProtoMessage()               {}
func (*WushuangWeaponBodyPosInfo) Descriptor() ([]byte, []int) { return fileDescriptor132, []int{10} }

func (m *WushuangWeaponBodyPosInfo) GetBodyPos() WushuangWeaponBodyPos {
	if m != nil && m.BodyPos != nil {
		return *m.BodyPos
	}
	return WushuangWeaponBodyPos_Weapon
}

func (m *WushuangWeaponBodyPosInfo) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *WushuangWeaponBodyPosInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *WushuangWeaponBodyPosInfo) GetExperience() int64 {
	if m != nil && m.Experience != nil {
		return *m.Experience
	}
	return 0
}

func init() {
	proto.RegisterType((*CSWushuangWeaponDevouring)(nil), "ui.CSWushuangWeaponDevouring")
	proto.RegisterType((*SCWushuangWeaponDevouring)(nil), "ui.SCWushuangWeaponDevouring")
	proto.RegisterType((*CSWushuangWeaponBreakthrough)(nil), "ui.CSWushuangWeaponBreakthrough")
	proto.RegisterType((*SCWushuangWeaponBreakthrough)(nil), "ui.SCWushuangWeaponBreakthrough")
	proto.RegisterType((*CSWushuangWeaponPutOn)(nil), "ui.CSWushuangWeaponPutOn")
	proto.RegisterType((*SCWushuangWeaponPutOn)(nil), "ui.SCWushuangWeaponPutOn")
	proto.RegisterType((*CSWushuangWeaponTakeOff)(nil), "ui.CSWushuangWeaponTakeOff")
	proto.RegisterType((*SCWushuangWeaponTakeOff)(nil), "ui.SCWushuangWeaponTakeOff")
	proto.RegisterType((*CSWushuangWeaponInfo)(nil), "ui.CSWushuangWeaponInfo")
	proto.RegisterType((*SCWushuangWeaponInfo)(nil), "ui.SCWushuangWeaponInfo")
	proto.RegisterType((*WushuangWeaponBodyPosInfo)(nil), "ui.WushuangWeaponBodyPosInfo")
	proto.RegisterEnum("ui.WushuangWeaponBodyPos", WushuangWeaponBodyPos_name, WushuangWeaponBodyPos_value)
}

func init() { proto.RegisterFile("wushuangweapon.proto", fileDescriptor132) }

var fileDescriptor132 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0x4f, 0x6b, 0xc2, 0x40,
	0x10, 0xc5, 0x31, 0x31, 0xfe, 0x19, 0x5b, 0x49, 0x17, 0xad, 0x09, 0x58, 0x08, 0x7b, 0x28, 0xc1,
	0x83, 0x07, 0xa1, 0x5f, 0x40, 0x5b, 0x5a, 0xa1, 0xad, 0x82, 0x05, 0xe9, 0xa1, 0xd0, 0xad, 0x99,
	0x98, 0xd4, 0x74, 0x57, 0xf2, 0x47, 0x6d, 0x3f, 0x7d, 0x49, 0x0c, 0x05, 0x17, 0x3d, 0x58, 0xbc,
	0xce, 0xbc, 0xfd, 0xbd, 0x37, 0x6f, 0xa1, 0xb1, 0x4e, 0x22, 0x2f, 0x61, 0x7c, 0xbe, 0x46, 0xb6,
	0x14, 0xbc, 0xbb, 0x0c, 0x45, 0x2c, 0x88, 0x92, 0xf8, 0xf4, 0x1e, 0xcc, 0xc1, 0x64, 0x9a, 0x6f,
	0xa7, 0xd9, 0xf6, 0x16, 0x57, 0x22, 0x09, 0x7d, 0x3e, 0x27, 0x1d, 0x28, 0x7f, 0x08, 0xe7, 0x7b,
	0x2c, 0x22, 0xa3, 0x60, 0x29, 0x76, 0xbd, 0x67, 0x76, 0x13, 0xbf, 0xbb, 0xab, 0xee, 0x6f, 0x05,
	0xf4, 0x13, 0xcc, 0xc9, 0xe0, 0x04, 0x20, 0x72, 0x0e, 0x5a, 0x80, 0x2b, 0x0c, 0x0c, 0xc5, 0x52,
	0x6c, 0x8d, 0x10, 0x00, 0xdc, 0x2c, 0x31, 0xf4, 0x91, 0xcf, 0xd0, 0x50, 0x2d, 0xc5, 0x56, 0xe9,
	0x1b, 0xb4, 0xe5, 0xd0, 0xfd, 0x10, 0xd9, 0x22, 0xf6, 0x42, 0x91, 0xcc, 0xbd, 0xa3, 0xec, 0x2e,
	0xa0, 0xea, 0x73, 0x07, 0x37, 0x8f, 0x7e, 0x14, 0x1b, 0x8a, 0xa5, 0xda, 0x1a, 0x7d, 0x85, 0xb6,
	0x7c, 0xca, 0xbf, 0xf1, 0xbb, 0xd7, 0xd0, 0x6b, 0x68, 0xca, 0xc9, 0xc7, 0x49, 0x3c, 0xe2, 0xa9,
	0x2e, 0x8b, 0x91, 0x11, 0x35, 0xea, 0x42, 0x53, 0x8e, 0xb0, 0xd5, 0x9d, 0xb8, 0xc9, 0x3b, 0x68,
	0xc9, 0x79, 0x5e, 0xd8, 0x02, 0x47, 0xae, 0x7b, 0xd4, 0xe7, 0x9b, 0xd0, 0x92, 0xe3, 0xe6, 0x18,
	0x7a, 0x09, 0x0d, 0xd9, 0x61, 0xc8, 0x5d, 0x41, 0x9f, 0xa0, 0x21, 0x3f, 0x49, 0xe7, 0xe4, 0x06,
	0xea, 0x2c, 0x08, 0x72, 0x70, 0x3a, 0x31, 0x0a, 0x96, 0x6a, 0xd7, 0x7a, 0x57, 0x07, 0xdd, 0x33,
	0xdc, 0x0f, 0x98, 0x07, 0x97, 0x47, 0x95, 0x56, 0x87, 0x92, 0x1f, 0xe3, 0xd7, 0xd0, 0xc9, 0x5b,
	0xfb, 0x2b, 0x51, 0xdd, 0x53, 0x62, 0x31, 0x2d, 0xb1, 0xf3, 0x0e, 0xcd, 0xfd, 0x2c, 0x80, 0xd2,
	0x76, 0xa0, 0x17, 0x48, 0x0d, 0xca, 0x83, 0x40, 0xc4, 0x1e, 0x46, 0xba, 0x42, 0x2a, 0x50, 0x7c,
	0x40, 0xe6, 0xe8, 0x2a, 0xa9, 0x82, 0x36, 0xf1, 0x04, 0x46, 0x7a, 0x91, 0x9c, 0x41, 0xe5, 0x19,
	0x67, 0x8b, 0x80, 0xcd, 0x50, 0xd7, 0x52, 0xfd, 0x18, 0xb9, 0xc3, 0x78, 0xac, 0x97, 0x7e, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x5c, 0x09, 0xf6, 0x39, 0xc0, 0x03, 0x00, 0x00,
}
