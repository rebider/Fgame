// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tower.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type FloorInfo struct {
	BiologyId        *int32 `protobuf:"varint,1,req,name=biologyId" json:"biologyId,omitempty"`
	IsDead           *bool  `protobuf:"varint,2,req,name=isDead" json:"isDead,omitempty"`
	DeadTime         *int64 `protobuf:"varint,3,req,name=deadTime" json:"deadTime,omitempty"`
	Floor            *int32 `protobuf:"varint,4,req,name=floor" json:"floor,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FloorInfo) Reset()                    { *m = FloorInfo{} }
func (m *FloorInfo) String() string            { return proto.CompactTextString(m) }
func (*FloorInfo) ProtoMessage()               {}
func (*FloorInfo) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{0} }

func (m *FloorInfo) GetBiologyId() int32 {
	if m != nil && m.BiologyId != nil {
		return *m.BiologyId
	}
	return 0
}

func (m *FloorInfo) GetIsDead() bool {
	if m != nil && m.IsDead != nil {
		return *m.IsDead
	}
	return false
}

func (m *FloorInfo) GetDeadTime() int64 {
	if m != nil && m.DeadTime != nil {
		return *m.DeadTime
	}
	return 0
}

func (m *FloorInfo) GetFloor() int32 {
	if m != nil && m.Floor != nil {
		return *m.Floor
	}
	return 0
}

type TowerLog struct {
	CreateTime       *int64  `protobuf:"varint,1,req,name=createTime" json:"createTime,omitempty"`
	PlayerName       *string `protobuf:"bytes,2,req,name=playerName" json:"playerName,omitempty"`
	BiologyName      *string `protobuf:"bytes,3,req,name=biologyName" json:"biologyName,omitempty"`
	ItemId           *int32  `protobuf:"varint,4,req,name=itemId" json:"itemId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TowerLog) Reset()                    { *m = TowerLog{} }
func (m *TowerLog) String() string            { return proto.CompactTextString(m) }
func (*TowerLog) ProtoMessage()               {}
func (*TowerLog) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{1} }

func (m *TowerLog) GetCreateTime() int64 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

func (m *TowerLog) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *TowerLog) GetBiologyName() string {
	if m != nil && m.BiologyName != nil {
		return *m.BiologyName
	}
	return ""
}

func (m *TowerLog) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

type CSTowerEnter struct {
	Floor            *int32 `protobuf:"varint,1,req,name=floor" json:"floor,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSTowerEnter) Reset()                    { *m = CSTowerEnter{} }
func (m *CSTowerEnter) String() string            { return proto.CompactTextString(m) }
func (*CSTowerEnter) ProtoMessage()               {}
func (*CSTowerEnter) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{2} }

func (m *CSTowerEnter) GetFloor() int32 {
	if m != nil && m.Floor != nil {
		return *m.Floor
	}
	return 0
}

type SCTowerEnter struct {
	RemainTime       *int64      `protobuf:"varint,1,req,name=remainTime" json:"remainTime,omitempty"`
	LogList          []*TowerLog `protobuf:"bytes,2,rep,name=logList" json:"logList,omitempty"`
	Floor            *int32      `protobuf:"varint,3,req,name=floor" json:"floor,omitempty"`
	IsDead           *bool       `protobuf:"varint,4,opt,name=isDead" json:"isDead,omitempty"`
	DeadTime         *int64      `protobuf:"varint,5,opt,name=deadTime" json:"deadTime,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCTowerEnter) Reset()                    { *m = SCTowerEnter{} }
func (m *SCTowerEnter) String() string            { return proto.CompactTextString(m) }
func (*SCTowerEnter) ProtoMessage()               {}
func (*SCTowerEnter) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{3} }

func (m *SCTowerEnter) GetRemainTime() int64 {
	if m != nil && m.RemainTime != nil {
		return *m.RemainTime
	}
	return 0
}

func (m *SCTowerEnter) GetLogList() []*TowerLog {
	if m != nil {
		return m.LogList
	}
	return nil
}

func (m *SCTowerEnter) GetFloor() int32 {
	if m != nil && m.Floor != nil {
		return *m.Floor
	}
	return 0
}

func (m *SCTowerEnter) GetIsDead() bool {
	if m != nil && m.IsDead != nil {
		return *m.IsDead
	}
	return false
}

func (m *SCTowerEnter) GetDeadTime() int64 {
	if m != nil && m.DeadTime != nil {
		return *m.DeadTime
	}
	return 0
}

type SCTowerBossInfoNotice struct {
	FloorInfo        *FloorInfo `protobuf:"bytes,1,req,name=floorInfo" json:"floorInfo,omitempty"`
	PlayerName       *string    `protobuf:"bytes,2,opt,name=playerName" json:"playerName,omitempty"`
	BossName         *string    `protobuf:"bytes,3,opt,name=bossName" json:"bossName,omitempty"`
	MapName          *string    `protobuf:"bytes,4,opt,name=mapName" json:"mapName,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *SCTowerBossInfoNotice) Reset()                    { *m = SCTowerBossInfoNotice{} }
func (m *SCTowerBossInfoNotice) String() string            { return proto.CompactTextString(m) }
func (*SCTowerBossInfoNotice) ProtoMessage()               {}
func (*SCTowerBossInfoNotice) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{4} }

func (m *SCTowerBossInfoNotice) GetFloorInfo() *FloorInfo {
	if m != nil {
		return m.FloorInfo
	}
	return nil
}

func (m *SCTowerBossInfoNotice) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *SCTowerBossInfoNotice) GetBossName() string {
	if m != nil && m.BossName != nil {
		return *m.BossName
	}
	return ""
}

func (m *SCTowerBossInfoNotice) GetMapName() string {
	if m != nil && m.MapName != nil {
		return *m.MapName
	}
	return ""
}

type SCTowerTimeNotice struct {
	RemainTime       *int64 `protobuf:"varint,1,req,name=remainTime" json:"remainTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTowerTimeNotice) Reset()                    { *m = SCTowerTimeNotice{} }
func (m *SCTowerTimeNotice) String() string            { return proto.CompactTextString(m) }
func (*SCTowerTimeNotice) ProtoMessage()               {}
func (*SCTowerTimeNotice) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{5} }

func (m *SCTowerTimeNotice) GetRemainTime() int64 {
	if m != nil && m.RemainTime != nil {
		return *m.RemainTime
	}
	return 0
}

type SCTowerResultNotice struct {
	TotalExp         *int64      `protobuf:"varint,1,req,name=totalExp" json:"totalExp,omitempty"`
	DropList         []*DropInfo `protobuf:"bytes,2,rep,name=dropList" json:"dropList,omitempty"`
	RemainTime       *int64      `protobuf:"varint,3,req,name=remainTime" json:"remainTime,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCTowerResultNotice) Reset()                    { *m = SCTowerResultNotice{} }
func (m *SCTowerResultNotice) String() string            { return proto.CompactTextString(m) }
func (*SCTowerResultNotice) ProtoMessage()               {}
func (*SCTowerResultNotice) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{6} }

func (m *SCTowerResultNotice) GetTotalExp() int64 {
	if m != nil && m.TotalExp != nil {
		return *m.TotalExp
	}
	return 0
}

func (m *SCTowerResultNotice) GetDropList() []*DropInfo {
	if m != nil {
		return m.DropList
	}
	return nil
}

func (m *SCTowerResultNotice) GetRemainTime() int64 {
	if m != nil && m.RemainTime != nil {
		return *m.RemainTime
	}
	return 0
}

type CSTowerFloorList struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSTowerFloorList) Reset()                    { *m = CSTowerFloorList{} }
func (m *CSTowerFloorList) String() string            { return proto.CompactTextString(m) }
func (*CSTowerFloorList) ProtoMessage()               {}
func (*CSTowerFloorList) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{7} }

type SCTowerFloorList struct {
	RemainTime       *int64       `protobuf:"varint,1,req,name=remainTime" json:"remainTime,omitempty"`
	FloorList        []*FloorInfo `protobuf:"bytes,2,rep,name=floorList" json:"floorList,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SCTowerFloorList) Reset()                    { *m = SCTowerFloorList{} }
func (m *SCTowerFloorList) String() string            { return proto.CompactTextString(m) }
func (*SCTowerFloorList) ProtoMessage()               {}
func (*SCTowerFloorList) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{8} }

func (m *SCTowerFloorList) GetRemainTime() int64 {
	if m != nil && m.RemainTime != nil {
		return *m.RemainTime
	}
	return 0
}

func (m *SCTowerFloorList) GetFloorList() []*FloorInfo {
	if m != nil {
		return m.FloorList
	}
	return nil
}

type CSTowerLogIncr struct {
	LogTime          *int64 `protobuf:"varint,1,req,name=logTime" json:"logTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSTowerLogIncr) Reset()                    { *m = CSTowerLogIncr{} }
func (m *CSTowerLogIncr) String() string            { return proto.CompactTextString(m) }
func (*CSTowerLogIncr) ProtoMessage()               {}
func (*CSTowerLogIncr) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{9} }

func (m *CSTowerLogIncr) GetLogTime() int64 {
	if m != nil && m.LogTime != nil {
		return *m.LogTime
	}
	return 0
}

type SCTowerLogIncr struct {
	LogList          []*TowerLog `protobuf:"bytes,1,rep,name=logList" json:"logList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCTowerLogIncr) Reset()                    { *m = SCTowerLogIncr{} }
func (m *SCTowerLogIncr) String() string            { return proto.CompactTextString(m) }
func (*SCTowerLogIncr) ProtoMessage()               {}
func (*SCTowerLogIncr) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{10} }

func (m *SCTowerLogIncr) GetLogList() []*TowerLog {
	if m != nil {
		return m.LogList
	}
	return nil
}

type CSTowerDaBao struct {
	Typ              *int32 `protobuf:"varint,1,req,name=typ" json:"typ,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSTowerDaBao) Reset()                    { *m = CSTowerDaBao{} }
func (m *CSTowerDaBao) String() string            { return proto.CompactTextString(m) }
func (*CSTowerDaBao) ProtoMessage()               {}
func (*CSTowerDaBao) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{11} }

func (m *CSTowerDaBao) GetTyp() int32 {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return 0
}

type SCTowerDaBao struct {
	RemainTime       *int64 `protobuf:"varint,1,req,name=remainTime" json:"remainTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTowerDaBao) Reset()                    { *m = SCTowerDaBao{} }
func (m *SCTowerDaBao) String() string            { return proto.CompactTextString(m) }
func (*SCTowerDaBao) ProtoMessage()               {}
func (*SCTowerDaBao) Descriptor() ([]byte, []int) { return fileDescriptor117, []int{12} }

func (m *SCTowerDaBao) GetRemainTime() int64 {
	if m != nil && m.RemainTime != nil {
		return *m.RemainTime
	}
	return 0
}

func init() {
	proto.RegisterType((*FloorInfo)(nil), "ui.FloorInfo")
	proto.RegisterType((*TowerLog)(nil), "ui.TowerLog")
	proto.RegisterType((*CSTowerEnter)(nil), "ui.CSTowerEnter")
	proto.RegisterType((*SCTowerEnter)(nil), "ui.SCTowerEnter")
	proto.RegisterType((*SCTowerBossInfoNotice)(nil), "ui.SCTowerBossInfoNotice")
	proto.RegisterType((*SCTowerTimeNotice)(nil), "ui.SCTowerTimeNotice")
	proto.RegisterType((*SCTowerResultNotice)(nil), "ui.SCTowerResultNotice")
	proto.RegisterType((*CSTowerFloorList)(nil), "ui.CSTowerFloorList")
	proto.RegisterType((*SCTowerFloorList)(nil), "ui.SCTowerFloorList")
	proto.RegisterType((*CSTowerLogIncr)(nil), "ui.CSTowerLogIncr")
	proto.RegisterType((*SCTowerLogIncr)(nil), "ui.SCTowerLogIncr")
	proto.RegisterType((*CSTowerDaBao)(nil), "ui.CSTowerDaBao")
	proto.RegisterType((*SCTowerDaBao)(nil), "ui.SCTowerDaBao")
}

func init() { proto.RegisterFile("tower.proto", fileDescriptor117) }

var fileDescriptor117 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xd1, 0x6b, 0x9c, 0x40,
	0x10, 0xc6, 0x51, 0x73, 0x8d, 0x8e, 0x97, 0xab, 0xd9, 0x50, 0x90, 0x94, 0x14, 0xbb, 0x2f, 0xf5,
	0xe9, 0x0a, 0xf9, 0x13, 0x92, 0x4b, 0xa9, 0x70, 0xa4, 0xd0, 0x04, 0xfa, 0xd0, 0xa7, 0x8d, 0xee,
	0xc9, 0x82, 0x3a, 0xcb, 0xba, 0x47, 0x7b, 0xff, 0x7d, 0xd9, 0xbd, 0xd5, 0x58, 0x73, 0x6f, 0x32,
	0xce, 0x7e, 0xbf, 0x6f, 0xbe, 0x19, 0x88, 0x35, 0xfe, 0xe1, 0x6a, 0x2d, 0x15, 0x6a, 0x24, 0xfe,
	0x5e, 0x5c, 0x2f, 0x4b, 0x6c, 0x5b, 0xec, 0x8e, 0x95, 0x6b, 0xa8, 0x14, 0xca, 0xe3, 0x37, 0xfd,
	0x01, 0xd1, 0xb7, 0x06, 0x51, 0x15, 0xdd, 0x0e, 0xc9, 0x25, 0x44, 0x2f, 0x02, 0x1b, 0xac, 0x0f,
	0x45, 0x95, 0x7a, 0x99, 0x9f, 0x2f, 0xc8, 0x0a, 0xde, 0x89, 0x7e, 0xc3, 0x59, 0x95, 0xfa, 0x99,
	0x9f, 0x87, 0x24, 0x81, 0xb0, 0xe2, 0xac, 0x7a, 0x16, 0x2d, 0x4f, 0x83, 0xcc, 0xcf, 0x03, 0x72,
	0x01, 0x8b, 0x9d, 0x51, 0x48, 0xcf, 0xcc, 0x03, 0xfa, 0x0b, 0xc2, 0x67, 0x43, 0xdf, 0x62, 0x4d,
	0x08, 0x40, 0xa9, 0x38, 0xd3, 0xdc, 0xb6, 0x7b, 0xb6, 0x9d, 0x00, 0xc8, 0x86, 0x1d, 0xb8, 0x7a,
	0x64, 0x2d, 0xb7, 0xa2, 0x11, 0xb9, 0x82, 0xd8, 0x71, 0x6d, 0x31, 0xb0, 0x45, 0x43, 0xd6, 0xbc,
	0x2d, 0x2a, 0x27, 0x7c, 0x03, 0xcb, 0xfb, 0x27, 0x2b, 0xfd, 0xd0, 0x69, 0xae, 0x5e, 0xb9, 0xd6,
	0x28, 0x55, 0xb0, 0x7c, 0xba, 0x9f, 0xfc, 0x26, 0x00, 0x8a, 0xb7, 0x4c, 0x74, 0x13, 0xf6, 0x0d,
	0x9c, 0x37, 0x58, 0x6f, 0x45, 0xaf, 0x53, 0x3f, 0x0b, 0xf2, 0xf8, 0x76, 0xb9, 0xde, 0x8b, 0xf5,
	0x68, 0x77, 0x54, 0x0c, 0x66, 0xa3, 0x9f, 0x65, 0xde, 0x6c, 0xf4, 0x45, 0xe6, 0xe5, 0x01, 0x95,
	0xf0, 0xc1, 0x31, 0xef, 0xb0, 0xef, 0x4d, 0x84, 0x8f, 0xa8, 0x45, 0xc9, 0x49, 0x06, 0xd1, 0x6e,
	0x48, 0xd5, 0xb2, 0xe3, 0xdb, 0x0b, 0x83, 0x7a, 0x8d, 0x7a, 0x1e, 0x83, 0x97, 0x47, 0x06, 0xf0,
	0x82, 0x7d, 0xef, 0x32, 0x30, 0x95, 0xf7, 0x70, 0xde, 0x32, 0x69, 0x0b, 0xc6, 0x43, 0x44, 0xbf,
	0xc0, 0xa5, 0x23, 0x1a, 0x1b, 0x8e, 0x76, 0x62, 0x54, 0xfa, 0x1b, 0xae, 0x5c, 0xe3, 0x4f, 0xde,
	0xef, 0x1b, 0xed, 0x5a, 0x13, 0x08, 0x35, 0x6a, 0xd6, 0x3c, 0xfc, 0x95, 0x2e, 0x93, 0x4f, 0x10,
	0x9a, 0x73, 0x98, 0x87, 0xb2, 0x51, 0x28, 0x07, 0xa3, 0x13, 0x71, 0xbb, 0x72, 0x4a, 0x20, 0x71,
	0xab, 0xb0, 0x03, 0x99, 0xb7, 0xf4, 0x3b, 0x24, 0x0e, 0x38, 0xd6, 0x4e, 0xee, 0x60, 0x88, 0x66,
	0x02, 0xfc, 0x3f, 0x1a, 0xfa, 0x19, 0x56, 0x4e, 0x7d, 0x8b, 0x75, 0xd1, 0x95, 0xca, 0xc4, 0xd0,
	0x60, 0x3d, 0x99, 0xee, 0x2b, 0xac, 0x1c, 0x6c, 0x68, 0x99, 0xac, 0xd6, 0x7b, 0xbb, 0x5a, 0xfa,
	0x71, 0x3c, 0x9e, 0x0d, 0xbb, 0x63, 0x48, 0x62, 0x08, 0xf4, 0x41, 0xba, 0xd3, 0xa1, 0xe3, 0xe9,
	0x1c, 0x7f, 0x9e, 0xb0, 0xfd, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x73, 0x95, 0x3b, 0x34, 0x53, 0x03,
	0x00, 0x00,
}
