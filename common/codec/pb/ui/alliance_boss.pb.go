// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alliance_boss.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AllianceBossRank struct {
	PlayerId         *int64  `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	PlayerName       *string `protobuf:"bytes,2,req,name=playerName" json:"playerName,omitempty"`
	Damage           *int64  `protobuf:"varint,3,req,name=damage" json:"damage,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AllianceBossRank) Reset()                    { *m = AllianceBossRank{} }
func (m *AllianceBossRank) String() string            { return proto.CompactTextString(m) }
func (*AllianceBossRank) ProtoMessage()               {}
func (*AllianceBossRank) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *AllianceBossRank) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *AllianceBossRank) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *AllianceBossRank) GetDamage() int64 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

type AllianceBossInfo struct {
	Level            *int32 `protobuf:"varint,1,req,name=level" json:"level,omitempty"`
	Status           *int32 `protobuf:"varint,2,opt,name=status,def=-1" json:"status,omitempty"`
	MaxHp            *int64 `protobuf:"varint,3,opt,name=maxHp,def=-1" json:"maxHp,omitempty"`
	Hp               *int64 `protobuf:"varint,4,opt,name=hp,def=-1" json:"hp,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AllianceBossInfo) Reset()                    { *m = AllianceBossInfo{} }
func (m *AllianceBossInfo) String() string            { return proto.CompactTextString(m) }
func (*AllianceBossInfo) ProtoMessage()               {}
func (*AllianceBossInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

const Default_AllianceBossInfo_Status int32 = -1
const Default_AllianceBossInfo_MaxHp int64 = -1
const Default_AllianceBossInfo_Hp int64 = -1

func (m *AllianceBossInfo) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *AllianceBossInfo) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return Default_AllianceBossInfo_Status
}

func (m *AllianceBossInfo) GetMaxHp() int64 {
	if m != nil && m.MaxHp != nil {
		return *m.MaxHp
	}
	return Default_AllianceBossInfo_MaxHp
}

func (m *AllianceBossInfo) GetHp() int64 {
	if m != nil && m.Hp != nil {
		return *m.Hp
	}
	return Default_AllianceBossInfo_Hp
}

type AllianceBossSceneInfo struct {
	SummonTime       *int64              `protobuf:"varint,1,req,name=summonTime" json:"summonTime,omitempty"`
	BossInfo         *AllianceBossInfo   `protobuf:"bytes,2,req,name=bossInfo" json:"bossInfo,omitempty"`
	RankList         []*AllianceBossRank `protobuf:"bytes,3,rep,name=rankList" json:"rankList,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *AllianceBossSceneInfo) Reset()                    { *m = AllianceBossSceneInfo{} }
func (m *AllianceBossSceneInfo) String() string            { return proto.CompactTextString(m) }
func (*AllianceBossSceneInfo) ProtoMessage()               {}
func (*AllianceBossSceneInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *AllianceBossSceneInfo) GetSummonTime() int64 {
	if m != nil && m.SummonTime != nil {
		return *m.SummonTime
	}
	return 0
}

func (m *AllianceBossSceneInfo) GetBossInfo() *AllianceBossInfo {
	if m != nil {
		return m.BossInfo
	}
	return nil
}

func (m *AllianceBossSceneInfo) GetRankList() []*AllianceBossRank {
	if m != nil {
		return m.RankList
	}
	return nil
}

type CSAllianceBossSummon struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSAllianceBossSummon) Reset()                    { *m = CSAllianceBossSummon{} }
func (m *CSAllianceBossSummon) String() string            { return proto.CompactTextString(m) }
func (*CSAllianceBossSummon) ProtoMessage()               {}
func (*CSAllianceBossSummon) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

type SCAllianceBossSummon struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCAllianceBossSummon) Reset()                    { *m = SCAllianceBossSummon{} }
func (m *SCAllianceBossSummon) String() string            { return proto.CompactTextString(m) }
func (*SCAllianceBossSummon) ProtoMessage()               {}
func (*SCAllianceBossSummon) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

type CSAllianceBossEnter struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSAllianceBossEnter) Reset()                    { *m = CSAllianceBossEnter{} }
func (m *CSAllianceBossEnter) String() string            { return proto.CompactTextString(m) }
func (*CSAllianceBossEnter) ProtoMessage()               {}
func (*CSAllianceBossEnter) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

type SCAllianceBossEnter struct {
	SceneInfo        *AllianceBossSceneInfo `protobuf:"bytes,1,req,name=sceneInfo" json:"sceneInfo,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *SCAllianceBossEnter) Reset()                    { *m = SCAllianceBossEnter{} }
func (m *SCAllianceBossEnter) String() string            { return proto.CompactTextString(m) }
func (*SCAllianceBossEnter) ProtoMessage()               {}
func (*SCAllianceBossEnter) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *SCAllianceBossEnter) GetSceneInfo() *AllianceBossSceneInfo {
	if m != nil {
		return m.SceneInfo
	}
	return nil
}

type SCAllianceBossChanged struct {
	BossInfo         *AllianceBossInfo `protobuf:"bytes,1,req,name=bossInfo" json:"bossInfo,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *SCAllianceBossChanged) Reset()                    { *m = SCAllianceBossChanged{} }
func (m *SCAllianceBossChanged) String() string            { return proto.CompactTextString(m) }
func (*SCAllianceBossChanged) ProtoMessage()               {}
func (*SCAllianceBossChanged) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *SCAllianceBossChanged) GetBossInfo() *AllianceBossInfo {
	if m != nil {
		return m.BossInfo
	}
	return nil
}

type SCAllianceBossRank struct {
	RankList         []*AllianceBossRank `protobuf:"bytes,1,rep,name=rankList" json:"rankList,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *SCAllianceBossRank) Reset()                    { *m = SCAllianceBossRank{} }
func (m *SCAllianceBossRank) String() string            { return proto.CompactTextString(m) }
func (*SCAllianceBossRank) ProtoMessage()               {}
func (*SCAllianceBossRank) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *SCAllianceBossRank) GetRankList() []*AllianceBossRank {
	if m != nil {
		return m.RankList
	}
	return nil
}

type SCAllianceBossEnd struct {
	Sucess           *bool  `protobuf:"varint,1,req,name=sucess" json:"sucess,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCAllianceBossEnd) Reset()                    { *m = SCAllianceBossEnd{} }
func (m *SCAllianceBossEnd) String() string            { return proto.CompactTextString(m) }
func (*SCAllianceBossEnd) ProtoMessage()               {}
func (*SCAllianceBossEnd) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *SCAllianceBossEnd) GetSucess() bool {
	if m != nil && m.Sucess != nil {
		return *m.Sucess
	}
	return false
}

type CSAllianceBoss struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSAllianceBoss) Reset()                    { *m = CSAllianceBoss{} }
func (m *CSAllianceBoss) String() string            { return proto.CompactTextString(m) }
func (*CSAllianceBoss) ProtoMessage()               {}
func (*CSAllianceBoss) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

type SCAllianceBoss struct {
	BossStatus       *int32 `protobuf:"varint,1,req,name=bossStatus" json:"bossStatus,omitempty"`
	Level            *int32 `protobuf:"varint,2,req,name=level" json:"level,omitempty"`
	Exp              *int32 `protobuf:"varint,3,req,name=exp" json:"exp,omitempty"`
	SummonTime       *int64 `protobuf:"varint,4,req,name=summonTime" json:"summonTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCAllianceBoss) Reset()                    { *m = SCAllianceBoss{} }
func (m *SCAllianceBoss) String() string            { return proto.CompactTextString(m) }
func (*SCAllianceBoss) ProtoMessage()               {}
func (*SCAllianceBoss) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *SCAllianceBoss) GetBossStatus() int32 {
	if m != nil && m.BossStatus != nil {
		return *m.BossStatus
	}
	return 0
}

func (m *SCAllianceBoss) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SCAllianceBoss) GetExp() int32 {
	if m != nil && m.Exp != nil {
		return *m.Exp
	}
	return 0
}

func (m *SCAllianceBoss) GetSummonTime() int64 {
	if m != nil && m.SummonTime != nil {
		return *m.SummonTime
	}
	return 0
}

type SCAllianceBossSummonSucess struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCAllianceBossSummonSucess) Reset()                    { *m = SCAllianceBossSummonSucess{} }
func (m *SCAllianceBossSummonSucess) String() string            { return proto.CompactTextString(m) }
func (*SCAllianceBossSummonSucess) ProtoMessage()               {}
func (*SCAllianceBossSummonSucess) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func init() {
	proto.RegisterType((*AllianceBossRank)(nil), "ui.AllianceBossRank")
	proto.RegisterType((*AllianceBossInfo)(nil), "ui.AllianceBossInfo")
	proto.RegisterType((*AllianceBossSceneInfo)(nil), "ui.AllianceBossSceneInfo")
	proto.RegisterType((*CSAllianceBossSummon)(nil), "ui.CSAllianceBossSummon")
	proto.RegisterType((*SCAllianceBossSummon)(nil), "ui.SCAllianceBossSummon")
	proto.RegisterType((*CSAllianceBossEnter)(nil), "ui.CSAllianceBossEnter")
	proto.RegisterType((*SCAllianceBossEnter)(nil), "ui.SCAllianceBossEnter")
	proto.RegisterType((*SCAllianceBossChanged)(nil), "ui.SCAllianceBossChanged")
	proto.RegisterType((*SCAllianceBossRank)(nil), "ui.SCAllianceBossRank")
	proto.RegisterType((*SCAllianceBossEnd)(nil), "ui.SCAllianceBossEnd")
	proto.RegisterType((*CSAllianceBoss)(nil), "ui.CSAllianceBoss")
	proto.RegisterType((*SCAllianceBoss)(nil), "ui.SCAllianceBoss")
	proto.RegisterType((*SCAllianceBossSummonSucess)(nil), "ui.SCAllianceBossSummonSucess")
}

func init() { proto.RegisterFile("alliance_boss.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x24, 0x9b, 0xa6, 0xa4, 0xaf, 0x18, 0xda, 0x6d, 0x2b, 0xab, 0x78, 0x08, 0x2b, 0x94, 0x1c,
	0xb4, 0x60, 0x8f, 0x22, 0x88, 0x06, 0xa1, 0x05, 0xf1, 0x60, 0x7a, 0xf0, 0x26, 0xdb, 0x66, 0x6d,
	0x43, 0x93, 0x4d, 0xc8, 0x26, 0x52, 0xf1, 0xcf, 0x4b, 0xb6, 0x1f, 0x74, 0x6b, 0xf5, 0xb8, 0x93,
	0x97, 0x79, 0x33, 0xf3, 0x06, 0x3a, 0x2c, 0x8e, 0x23, 0x26, 0x66, 0xfc, 0x7d, 0x9a, 0x4a, 0x39,
	0xc8, 0xf2, 0xb4, 0x48, 0x31, 0x2a, 0x23, 0x3a, 0x82, 0xd6, 0xc3, 0xe6, 0xd3, 0x63, 0x2a, 0xe5,
	0x2b, 0x13, 0x4b, 0xdc, 0x02, 0x3b, 0x8b, 0xd9, 0x17, 0xcf, 0xc7, 0x21, 0x31, 0x5c, 0xe4, 0x99,
	0x18, 0x03, 0xac, 0x91, 0x17, 0x96, 0x70, 0x82, 0x5c, 0xe4, 0x35, 0xb0, 0x03, 0xf5, 0x90, 0x25,
	0x6c, 0xce, 0x89, 0x59, 0xcd, 0xd0, 0x37, 0x9d, 0x69, 0x2c, 0x3e, 0x52, 0x7c, 0x02, 0x56, 0xcc,
	0x3f, 0x79, 0xac, 0x68, 0x2c, 0x8c, 0xa1, 0x2e, 0x0b, 0x56, 0x94, 0x92, 0x20, 0xd7, 0xf0, 0xac,
	0x5b, 0x74, 0x7d, 0x83, 0xdb, 0x60, 0x25, 0x6c, 0x35, 0xca, 0x88, 0xe9, 0x1a, 0x9e, 0xa9, 0x20,
	0x07, 0xd0, 0x22, 0x23, 0xb5, 0xed, 0x9b, 0x7e, 0x43, 0x6f, 0x9f, 0x39, 0x98, 0x71, 0xc1, 0x15,
	0x3d, 0x06, 0x90, 0x65, 0x92, 0xa4, 0x62, 0x12, 0x25, 0x7c, 0x23, 0xb5, 0x0f, 0xf6, 0x74, 0xb3,
	0x5e, 0x09, 0x6d, 0x0e, 0xbb, 0x83, 0x32, 0x1a, 0xfc, 0x92, 0xd6, 0x07, 0x3b, 0x67, 0x62, 0xf9,
	0x1c, 0xc9, 0x82, 0x98, 0xae, 0x79, 0x6c, 0xae, 0x0a, 0x83, 0x9e, 0x42, 0xd7, 0x0f, 0xb4, 0xf5,
	0x6a, 0x63, 0x85, 0x07, 0xfe, 0x11, 0xbc, 0x07, 0x1d, 0x7d, 0xfe, 0x49, 0x14, 0x3c, 0xa7, 0x3e,
	0x74, 0xf4, 0x71, 0x05, 0xe3, 0x2b, 0x68, 0xc8, 0xad, 0x1d, 0x65, 0xa0, 0x39, 0x3c, 0x3b, 0x94,
	0xb1, 0xf3, 0x4b, 0xef, 0xa1, 0xa7, 0x93, 0xf8, 0x0b, 0x26, 0xe6, 0x3c, 0xd4, 0x4c, 0x1b, 0x7f,
	0x9b, 0xa6, 0x77, 0x80, 0x75, 0x02, 0x75, 0xef, 0xfd, 0x28, 0x8c, 0x7f, 0xa2, 0xb8, 0x84, 0xf6,
	0xa1, 0x87, 0xb0, 0xaa, 0x81, 0x2c, 0x67, 0x5c, 0x4a, 0xb5, 0xd8, 0xa6, 0x2d, 0x70, 0x74, 0xff,
	0x74, 0x02, 0x8e, 0xfe, 0x5b, 0x75, 0xb7, 0x4a, 0x6e, 0xb0, 0xee, 0xc2, 0xba, 0x1b, 0xbb, 0xaa,
	0x20, 0xf5, 0x6c, 0x82, 0xc9, 0x57, 0x99, 0xaa, 0x96, 0x75, 0x70, 0xe7, 0x9a, 0xaa, 0xdb, 0x05,
	0x9c, 0x1f, 0xcb, 0x3f, 0x50, 0x5a, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x1f, 0x07, 0x06,
	0xf0, 0x02, 0x00, 0x00,
}
