// Code generated by protoc-gen-go. DO NOT EDIT.
// source: show.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PlayerShowData struct {
	TitleId          *int32  `protobuf:"varint,1,opt,name=titleId" json:"titleId,omitempty"`
	WeaponId         *int32  `protobuf:"varint,2,opt,name=weaponId" json:"weaponId,omitempty"`
	ClothesId        *int32  `protobuf:"varint,3,opt,name=clothesId" json:"clothesId,omitempty"`
	RideId           *int32  `protobuf:"varint,4,opt,name=rideId" json:"rideId,omitempty"`
	WingId           *int32  `protobuf:"varint,5,opt,name=wingId" json:"wingId,omitempty"`
	ShenFaId         *int32  `protobuf:"varint,6,opt,name=shenFaId" json:"shenFaId,omitempty"`
	LingYuId         *int32  `protobuf:"varint,7,opt,name=lingYuId" json:"lingYuId,omitempty"`
	FourGodKey       *int32  `protobuf:"varint,8,opt,name=fourGodKey" json:"fourGodKey,omitempty"`
	Realm            *int32  `protobuf:"varint,9,opt,name=realm" json:"realm,omitempty"`
	Spouse           *string `protobuf:"bytes,10,opt,name=spouse" json:"spouse,omitempty"`
	WeddingStatus    *int32  `protobuf:"varint,11,opt,name=weddingStatus" json:"weddingStatus,omitempty"`
	Model            *int32  `protobuf:"varint,12,opt,name=model" json:"model,omitempty"`
	WeaponState      *int32  `protobuf:"varint,13,opt,name=weaponState" json:"weaponState,omitempty"`
	RingType         *int32  `protobuf:"varint,14,opt,name=ringType" json:"ringType,omitempty"`
	FaBaoId          *int32  `protobuf:"varint,15,opt,name=faBaoId" json:"faBaoId,omitempty"`
	PetId            *int32  `protobuf:"varint,16,opt,name=petId" json:"petId,omitempty"`
	XianTiId         *int32  `protobuf:"varint,17,opt,name=xianTiId" json:"xianTiId,omitempty"`
	BaGua            *int32  `protobuf:"varint,18,opt,name=baGua" json:"baGua,omitempty"`
	MountHidden      *bool   `protobuf:"varint,19,opt,name=mountHidden" json:"mountHidden,omitempty"`
	MountAdvanceId   *int32  `protobuf:"varint,20,opt,name=mountAdvanceId" json:"mountAdvanceId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *PlayerShowData) Reset()                    { *m = PlayerShowData{} }
func (m *PlayerShowData) String() string            { return proto.CompactTextString(m) }
func (*PlayerShowData) ProtoMessage()               {}
func (*PlayerShowData) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{0} }

func (m *PlayerShowData) GetTitleId() int32 {
	if m != nil && m.TitleId != nil {
		return *m.TitleId
	}
	return 0
}

func (m *PlayerShowData) GetWeaponId() int32 {
	if m != nil && m.WeaponId != nil {
		return *m.WeaponId
	}
	return 0
}

func (m *PlayerShowData) GetClothesId() int32 {
	if m != nil && m.ClothesId != nil {
		return *m.ClothesId
	}
	return 0
}

func (m *PlayerShowData) GetRideId() int32 {
	if m != nil && m.RideId != nil {
		return *m.RideId
	}
	return 0
}

func (m *PlayerShowData) GetWingId() int32 {
	if m != nil && m.WingId != nil {
		return *m.WingId
	}
	return 0
}

func (m *PlayerShowData) GetShenFaId() int32 {
	if m != nil && m.ShenFaId != nil {
		return *m.ShenFaId
	}
	return 0
}

func (m *PlayerShowData) GetLingYuId() int32 {
	if m != nil && m.LingYuId != nil {
		return *m.LingYuId
	}
	return 0
}

func (m *PlayerShowData) GetFourGodKey() int32 {
	if m != nil && m.FourGodKey != nil {
		return *m.FourGodKey
	}
	return 0
}

func (m *PlayerShowData) GetRealm() int32 {
	if m != nil && m.Realm != nil {
		return *m.Realm
	}
	return 0
}

func (m *PlayerShowData) GetSpouse() string {
	if m != nil && m.Spouse != nil {
		return *m.Spouse
	}
	return ""
}

func (m *PlayerShowData) GetWeddingStatus() int32 {
	if m != nil && m.WeddingStatus != nil {
		return *m.WeddingStatus
	}
	return 0
}

func (m *PlayerShowData) GetModel() int32 {
	if m != nil && m.Model != nil {
		return *m.Model
	}
	return 0
}

func (m *PlayerShowData) GetWeaponState() int32 {
	if m != nil && m.WeaponState != nil {
		return *m.WeaponState
	}
	return 0
}

func (m *PlayerShowData) GetRingType() int32 {
	if m != nil && m.RingType != nil {
		return *m.RingType
	}
	return 0
}

func (m *PlayerShowData) GetFaBaoId() int32 {
	if m != nil && m.FaBaoId != nil {
		return *m.FaBaoId
	}
	return 0
}

func (m *PlayerShowData) GetPetId() int32 {
	if m != nil && m.PetId != nil {
		return *m.PetId
	}
	return 0
}

func (m *PlayerShowData) GetXianTiId() int32 {
	if m != nil && m.XianTiId != nil {
		return *m.XianTiId
	}
	return 0
}

func (m *PlayerShowData) GetBaGua() int32 {
	if m != nil && m.BaGua != nil {
		return *m.BaGua
	}
	return 0
}

func (m *PlayerShowData) GetMountHidden() bool {
	if m != nil && m.MountHidden != nil {
		return *m.MountHidden
	}
	return false
}

func (m *PlayerShowData) GetMountAdvanceId() int32 {
	if m != nil && m.MountAdvanceId != nil {
		return *m.MountAdvanceId
	}
	return 0
}

type SIPlayerShowDataChanged struct {
	PlayerShowData   *PlayerShowData `protobuf:"bytes,1,req,name=playerShowData" json:"playerShowData,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SIPlayerShowDataChanged) Reset()                    { *m = SIPlayerShowDataChanged{} }
func (m *SIPlayerShowDataChanged) String() string            { return proto.CompactTextString(m) }
func (*SIPlayerShowDataChanged) ProtoMessage()               {}
func (*SIPlayerShowDataChanged) Descriptor() ([]byte, []int) { return fileDescriptor27, []int{1} }

func (m *SIPlayerShowDataChanged) GetPlayerShowData() *PlayerShowData {
	if m != nil {
		return m.PlayerShowData
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayerShowData)(nil), "cross.PlayerShowData")
	proto.RegisterType((*SIPlayerShowDataChanged)(nil), "cross.SIPlayerShowDataChanged")
}

func init() { proto.RegisterFile("show.proto", fileDescriptor27) }

var fileDescriptor27 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x3f, 0x6f, 0xc2, 0x30,
	0x10, 0xc5, 0x15, 0xda, 0xf0, 0xe7, 0x28, 0x01, 0x4c, 0x69, 0x6f, 0x8c, 0x98, 0xb2, 0x94, 0xa1,
	0xdf, 0xa0, 0x7f, 0x54, 0xb0, 0xba, 0x54, 0x82, 0xa5, 0xa3, 0xcb, 0x99, 0xc4, 0x52, 0xb0, 0xa3,
	0xd8, 0x69, 0xca, 0x37, 0xef, 0x58, 0xc5, 0x95, 0x07, 0xc6, 0xfb, 0xf9, 0xbd, 0x77, 0xf2, 0x3b,
	0x00, 0x5b, 0x98, 0x76, 0x5d, 0xd5, 0xc6, 0x19, 0x16, 0x1f, 0x6a, 0x63, 0xed, 0xea, 0xb7, 0x07,
	0xc9, 0x47, 0x29, 0xce, 0xb2, 0xde, 0x15, 0xa6, 0x7d, 0x15, 0x4e, 0xb0, 0x29, 0x0c, 0x9c, 0x72,
	0xa5, 0xe4, 0x84, 0x51, 0x1a, 0x65, 0x31, 0x9b, 0xc1, 0xb0, 0x95, 0xa2, 0x32, 0x9a, 0x13, 0xf6,
	0x3c, 0x99, 0xc3, 0xe8, 0x50, 0x1a, 0x57, 0x48, 0xcb, 0x09, 0xaf, 0x3c, 0x4a, 0xa0, 0x5f, 0x2b,
	0xea, 0x4c, 0xd7, 0x61, 0x6e, 0x95, 0xce, 0x39, 0x61, 0x1c, 0x42, 0x6c, 0x21, 0xf5, 0x9b, 0xe0,
	0x84, 0xfd, 0x40, 0x4a, 0xa5, 0xf3, 0xcf, 0x86, 0x13, 0x0e, 0x3c, 0x61, 0x00, 0x47, 0xd3, 0xd4,
	0x1b, 0x43, 0xef, 0xf2, 0x8c, 0x43, 0xcf, 0x26, 0x10, 0xd7, 0x52, 0x94, 0x27, 0x1c, 0x85, 0x58,
	0x5b, 0x99, 0xc6, 0x4a, 0x84, 0x34, 0xca, 0x46, 0x6c, 0x09, 0x93, 0x56, 0x12, 0x29, 0x9d, 0xef,
	0x9c, 0x70, 0x8d, 0xc5, 0x71, 0x70, 0x9d, 0x0c, 0xc9, 0x12, 0x6f, 0xfc, 0xb8, 0x80, 0xf1, 0xff,
	0x0f, 0x3a, 0x91, 0xc4, 0x49, 0xd8, 0x5f, 0x2b, 0x9d, 0xef, 0xcf, 0x95, 0xc4, 0xc4, 0x93, 0x29,
	0x0c, 0x8e, 0xe2, 0x59, 0x18, 0x4e, 0x38, 0x0d, 0x31, 0x95, 0x74, 0x9c, 0x70, 0x16, 0x1c, 0x3f,
	0x4a, 0xe8, 0xbd, 0xe2, 0x84, 0xf3, 0x20, 0xf8, 0x12, 0x9b, 0x46, 0x20, 0x0b, 0x7b, 0x4e, 0xa6,
	0xd1, 0x6e, 0xab, 0x88, 0xa4, 0xc6, 0x45, 0x1a, 0x65, 0x43, 0x76, 0x07, 0x89, 0x87, 0x4f, 0xf4,
	0x2d, 0xf4, 0xa1, 0x6b, 0xe8, 0xb6, 0x13, 0xaf, 0xb6, 0x70, 0xbf, 0xe3, 0x97, 0xdd, 0xbf, 0x14,
	0x42, 0xe7, 0x92, 0xd8, 0x03, 0x24, 0xd5, 0xc5, 0x03, 0x46, 0x69, 0x2f, 0x1b, 0x3f, 0x2e, 0xd7,
	0xfe, 0x6a, 0xeb, 0x4b, 0xd7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xf6, 0x4a, 0xb9, 0xd8,
	0x01, 0x00, 0x00,
}
