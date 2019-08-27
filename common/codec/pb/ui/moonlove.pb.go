// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moonlove.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MoonloveRank struct {
	PlayerName       *string `protobuf:"bytes,1,req,name=playerName" json:"playerName,omitempty"`
	Number           *int32  `protobuf:"varint,2,req,name=number" json:"number,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MoonloveRank) Reset()                    { *m = MoonloveRank{} }
func (m *MoonloveRank) String() string            { return proto.CompactTextString(m) }
func (*MoonloveRank) ProtoMessage()               {}
func (*MoonloveRank) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{0} }

func (m *MoonloveRank) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *MoonloveRank) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

type PlayerMoonloveInfo struct {
	GenerousRank     *int32 `protobuf:"varint,1,req,name=generousRank" json:"generousRank,omitempty"`
	CharmRank        *int32 `protobuf:"varint,2,req,name=charmRank" json:"charmRank,omitempty"`
	GenerousNum      *int32 `protobuf:"varint,3,req,name=generousNum" json:"generousNum,omitempty"`
	CharmNum         *int32 `protobuf:"varint,4,req,name=charmNum" json:"charmNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlayerMoonloveInfo) Reset()                    { *m = PlayerMoonloveInfo{} }
func (m *PlayerMoonloveInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerMoonloveInfo) ProtoMessage()               {}
func (*PlayerMoonloveInfo) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{1} }

func (m *PlayerMoonloveInfo) GetGenerousRank() int32 {
	if m != nil && m.GenerousRank != nil {
		return *m.GenerousRank
	}
	return 0
}

func (m *PlayerMoonloveInfo) GetCharmRank() int32 {
	if m != nil && m.CharmRank != nil {
		return *m.CharmRank
	}
	return 0
}

func (m *PlayerMoonloveInfo) GetGenerousNum() int32 {
	if m != nil && m.GenerousNum != nil {
		return *m.GenerousNum
	}
	return 0
}

func (m *PlayerMoonloveInfo) GetCharmNum() int32 {
	if m != nil && m.CharmNum != nil {
		return *m.CharmNum
	}
	return 0
}

type ScenePlayer struct {
	PlayerId         *int64  `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Sex              *int32  `protobuf:"varint,3,req,name=sex" json:"sex,omitempty"`
	Point            *int32  `protobuf:"varint,4,req,name=point" json:"point,omitempty"`
	Status           *int32  `protobuf:"varint,5,req,name=status" json:"status,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ScenePlayer) Reset()                    { *m = ScenePlayer{} }
func (m *ScenePlayer) String() string            { return proto.CompactTextString(m) }
func (*ScenePlayer) ProtoMessage()               {}
func (*ScenePlayer) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{2} }

func (m *ScenePlayer) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *ScenePlayer) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ScenePlayer) GetSex() int32 {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return 0
}

func (m *ScenePlayer) GetPoint() int32 {
	if m != nil && m.Point != nil {
		return *m.Point
	}
	return 0
}

func (m *ScenePlayer) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

type SCMoonloveGenerousChanged struct {
	PlayerId         *int64 `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	GenerousRank     *int32 `protobuf:"varint,2,req,name=generousRank" json:"generousRank,omitempty"`
	GenerousNum      *int32 `protobuf:"varint,3,req,name=generousNum" json:"generousNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCMoonloveGenerousChanged) Reset()                    { *m = SCMoonloveGenerousChanged{} }
func (m *SCMoonloveGenerousChanged) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveGenerousChanged) ProtoMessage()               {}
func (*SCMoonloveGenerousChanged) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{3} }

func (m *SCMoonloveGenerousChanged) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *SCMoonloveGenerousChanged) GetGenerousRank() int32 {
	if m != nil && m.GenerousRank != nil {
		return *m.GenerousRank
	}
	return 0
}

func (m *SCMoonloveGenerousChanged) GetGenerousNum() int32 {
	if m != nil && m.GenerousNum != nil {
		return *m.GenerousNum
	}
	return 0
}

type SCMoonloveCharmChanged struct {
	PlayerId         *int64 `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	CharmRank        *int32 `protobuf:"varint,2,req,name=charmRank" json:"charmRank,omitempty"`
	CharmNum         *int32 `protobuf:"varint,3,req,name=charmNum" json:"charmNum,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCMoonloveCharmChanged) Reset()                    { *m = SCMoonloveCharmChanged{} }
func (m *SCMoonloveCharmChanged) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveCharmChanged) ProtoMessage()               {}
func (*SCMoonloveCharmChanged) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{4} }

func (m *SCMoonloveCharmChanged) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *SCMoonloveCharmChanged) GetCharmRank() int32 {
	if m != nil && m.CharmRank != nil {
		return *m.CharmRank
	}
	return 0
}

func (m *SCMoonloveCharmChanged) GetCharmNum() int32 {
	if m != nil && m.CharmNum != nil {
		return *m.CharmNum
	}
	return 0
}

type SCMoonloveSceneInfo struct {
	CharmRank        []*MoonloveRank     `protobuf:"bytes,1,rep,name=charmRank" json:"charmRank,omitempty"`
	GenerousRank     []*MoonloveRank     `protobuf:"bytes,2,rep,name=generousRank" json:"generousRank,omitempty"`
	MoonloveInfo     *PlayerMoonloveInfo `protobuf:"bytes,3,req,name=moonloveInfo" json:"moonloveInfo,omitempty"`
	RewExp           *int64              `protobuf:"varint,5,req,name=rewExp" json:"rewExp,omitempty"`
	CreateTime       *int64              `protobuf:"varint,6,req,name=createTime" json:"createTime,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *SCMoonloveSceneInfo) Reset()                    { *m = SCMoonloveSceneInfo{} }
func (m *SCMoonloveSceneInfo) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveSceneInfo) ProtoMessage()               {}
func (*SCMoonloveSceneInfo) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{5} }

func (m *SCMoonloveSceneInfo) GetCharmRank() []*MoonloveRank {
	if m != nil {
		return m.CharmRank
	}
	return nil
}

func (m *SCMoonloveSceneInfo) GetGenerousRank() []*MoonloveRank {
	if m != nil {
		return m.GenerousRank
	}
	return nil
}

func (m *SCMoonloveSceneInfo) GetMoonloveInfo() *PlayerMoonloveInfo {
	if m != nil {
		return m.MoonloveInfo
	}
	return nil
}

func (m *SCMoonloveSceneInfo) GetRewExp() int64 {
	if m != nil && m.RewExp != nil {
		return *m.RewExp
	}
	return 0
}

func (m *SCMoonloveSceneInfo) GetCreateTime() int64 {
	if m != nil && m.CreateTime != nil {
		return *m.CreateTime
	}
	return 0
}

type SCMoonlovePushCharmRank struct {
	CharmRank        []*MoonloveRank `protobuf:"bytes,1,rep,name=charmRank" json:"charmRank,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SCMoonlovePushCharmRank) Reset()                    { *m = SCMoonlovePushCharmRank{} }
func (m *SCMoonlovePushCharmRank) String() string            { return proto.CompactTextString(m) }
func (*SCMoonlovePushCharmRank) ProtoMessage()               {}
func (*SCMoonlovePushCharmRank) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{6} }

func (m *SCMoonlovePushCharmRank) GetCharmRank() []*MoonloveRank {
	if m != nil {
		return m.CharmRank
	}
	return nil
}

type SCMoonlovePushGenerousRank struct {
	GenerousRank     []*MoonloveRank `protobuf:"bytes,2,rep,name=generousRank" json:"generousRank,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SCMoonlovePushGenerousRank) Reset()                    { *m = SCMoonlovePushGenerousRank{} }
func (m *SCMoonlovePushGenerousRank) String() string            { return proto.CompactTextString(m) }
func (*SCMoonlovePushGenerousRank) ProtoMessage()               {}
func (*SCMoonlovePushGenerousRank) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{7} }

func (m *SCMoonlovePushGenerousRank) GetGenerousRank() []*MoonloveRank {
	if m != nil {
		return m.GenerousRank
	}
	return nil
}

type SCMoonloveSceneResult struct {
	PlayerId         *int64 `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	RewExp           *int64 `protobuf:"varint,2,req,name=rewExp" json:"rewExp,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCMoonloveSceneResult) Reset()                    { *m = SCMoonloveSceneResult{} }
func (m *SCMoonloveSceneResult) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveSceneResult) ProtoMessage()               {}
func (*SCMoonloveSceneResult) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{8} }

func (m *SCMoonloveSceneResult) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *SCMoonloveSceneResult) GetRewExp() int64 {
	if m != nil && m.RewExp != nil {
		return *m.RewExp
	}
	return 0
}

type SCMoonloveRankRewards struct {
	RewProperty      *RewProperty `protobuf:"bytes,1,req,name=rewProperty" json:"rewProperty,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SCMoonloveRankRewards) Reset()                    { *m = SCMoonloveRankRewards{} }
func (m *SCMoonloveRankRewards) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveRankRewards) ProtoMessage()               {}
func (*SCMoonloveRankRewards) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{9} }

func (m *SCMoonloveRankRewards) GetRewProperty() *RewProperty {
	if m != nil {
		return m.RewProperty
	}
	return nil
}

type CSMoonloveViewDouble struct {
	PlayerId         *int64 `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSMoonloveViewDouble) Reset()                    { *m = CSMoonloveViewDouble{} }
func (m *CSMoonloveViewDouble) String() string            { return proto.CompactTextString(m) }
func (*CSMoonloveViewDouble) ProtoMessage()               {}
func (*CSMoonloveViewDouble) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{10} }

func (m *CSMoonloveViewDouble) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type SCMoonloveViewDouble struct {
	PlayerId         *int64    `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	TargetPlayerId   *int64    `protobuf:"varint,2,req,name=targetPlayerId" json:"targetPlayerId,omitempty"`
	TargetPosition   *Position `protobuf:"bytes,3,req,name=targetPosition" json:"targetPosition,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SCMoonloveViewDouble) Reset()                    { *m = SCMoonloveViewDouble{} }
func (m *SCMoonloveViewDouble) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveViewDouble) ProtoMessage()               {}
func (*SCMoonloveViewDouble) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{11} }

func (m *SCMoonloveViewDouble) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *SCMoonloveViewDouble) GetTargetPlayerId() int64 {
	if m != nil && m.TargetPlayerId != nil {
		return *m.TargetPlayerId
	}
	return 0
}

func (m *SCMoonloveViewDouble) GetTargetPosition() *Position {
	if m != nil {
		return m.TargetPosition
	}
	return nil
}

type CSMoonloveViewDoubleState struct {
	TargetPlayerId   *int64 `protobuf:"varint,1,req,name=targetPlayerId" json:"targetPlayerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSMoonloveViewDoubleState) Reset()                    { *m = CSMoonloveViewDoubleState{} }
func (m *CSMoonloveViewDoubleState) String() string            { return proto.CompactTextString(m) }
func (*CSMoonloveViewDoubleState) ProtoMessage()               {}
func (*CSMoonloveViewDoubleState) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{12} }

func (m *CSMoonloveViewDoubleState) GetTargetPlayerId() int64 {
	if m != nil && m.TargetPlayerId != nil {
		return *m.TargetPlayerId
	}
	return 0
}

type SCMoonloveViewDoubleState struct {
	TargetPlayerId   *int64 `protobuf:"varint,1,req,name=targetPlayerId" json:"targetPlayerId,omitempty"`
	IsSuccess        *bool  `protobuf:"varint,2,req,name=isSuccess" json:"isSuccess,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCMoonloveViewDoubleState) Reset()                    { *m = SCMoonloveViewDoubleState{} }
func (m *SCMoonloveViewDoubleState) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveViewDoubleState) ProtoMessage()               {}
func (*SCMoonloveViewDoubleState) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{13} }

func (m *SCMoonloveViewDoubleState) GetTargetPlayerId() int64 {
	if m != nil && m.TargetPlayerId != nil {
		return *m.TargetPlayerId
	}
	return 0
}

func (m *SCMoonloveViewDoubleState) GetIsSuccess() bool {
	if m != nil && m.IsSuccess != nil {
		return *m.IsSuccess
	}
	return false
}

type CSMoonloveViewDoubleRelease struct {
	TargetPlayerId   *int64 `protobuf:"varint,1,req,name=targetPlayerId" json:"targetPlayerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSMoonloveViewDoubleRelease) Reset()                    { *m = CSMoonloveViewDoubleRelease{} }
func (m *CSMoonloveViewDoubleRelease) String() string            { return proto.CompactTextString(m) }
func (*CSMoonloveViewDoubleRelease) ProtoMessage()               {}
func (*CSMoonloveViewDoubleRelease) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{14} }

func (m *CSMoonloveViewDoubleRelease) GetTargetPlayerId() int64 {
	if m != nil && m.TargetPlayerId != nil {
		return *m.TargetPlayerId
	}
	return 0
}

type SCMoonloveViewDoubleRelease struct {
	PlayerId         *int64 `protobuf:"varint,1,req,name=playerId" json:"playerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCMoonloveViewDoubleRelease) Reset()                    { *m = SCMoonloveViewDoubleRelease{} }
func (m *SCMoonloveViewDoubleRelease) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveViewDoubleRelease) ProtoMessage()               {}
func (*SCMoonloveViewDoubleRelease) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{15} }

func (m *SCMoonloveViewDoubleRelease) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type CSMoonlovePlayerList struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSMoonlovePlayerList) Reset()                    { *m = CSMoonlovePlayerList{} }
func (m *CSMoonlovePlayerList) String() string            { return proto.CompactTextString(m) }
func (*CSMoonlovePlayerList) ProtoMessage()               {}
func (*CSMoonlovePlayerList) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{16} }

type SCMoonlovePlayerList struct {
	PlayerList       []*ScenePlayer `protobuf:"bytes,1,rep,name=playerList" json:"playerList,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SCMoonlovePlayerList) Reset()                    { *m = SCMoonlovePlayerList{} }
func (m *SCMoonlovePlayerList) String() string            { return proto.CompactTextString(m) }
func (*SCMoonlovePlayerList) ProtoMessage()               {}
func (*SCMoonlovePlayerList) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{17} }

func (m *SCMoonlovePlayerList) GetPlayerList() []*ScenePlayer {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

type SCMoonloveGiftNotice struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	TargetName       *string `protobuf:"bytes,2,req,name=targetName" json:"targetName,omitempty"`
	Num              *int32  `protobuf:"varint,3,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCMoonloveGiftNotice) Reset()                    { *m = SCMoonloveGiftNotice{} }
func (m *SCMoonloveGiftNotice) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveGiftNotice) ProtoMessage()               {}
func (*SCMoonloveGiftNotice) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{18} }

func (m *SCMoonloveGiftNotice) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *SCMoonloveGiftNotice) GetTargetName() string {
	if m != nil && m.TargetName != nil {
		return *m.TargetName
	}
	return ""
}

func (m *SCMoonloveGiftNotice) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type SCMoonloveExpCountNotice struct {
	ExpCount         *int64 `protobuf:"varint,1,req,name=expCount" json:"expCount,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCMoonloveExpCountNotice) Reset()                    { *m = SCMoonloveExpCountNotice{} }
func (m *SCMoonloveExpCountNotice) String() string            { return proto.CompactTextString(m) }
func (*SCMoonloveExpCountNotice) ProtoMessage()               {}
func (*SCMoonloveExpCountNotice) Descriptor() ([]byte, []int) { return fileDescriptor74, []int{19} }

func (m *SCMoonloveExpCountNotice) GetExpCount() int64 {
	if m != nil && m.ExpCount != nil {
		return *m.ExpCount
	}
	return 0
}

func init() {
	proto.RegisterType((*MoonloveRank)(nil), "ui.MoonloveRank")
	proto.RegisterType((*PlayerMoonloveInfo)(nil), "ui.PlayerMoonloveInfo")
	proto.RegisterType((*ScenePlayer)(nil), "ui.ScenePlayer")
	proto.RegisterType((*SCMoonloveGenerousChanged)(nil), "ui.SCMoonloveGenerousChanged")
	proto.RegisterType((*SCMoonloveCharmChanged)(nil), "ui.SCMoonloveCharmChanged")
	proto.RegisterType((*SCMoonloveSceneInfo)(nil), "ui.SCMoonloveSceneInfo")
	proto.RegisterType((*SCMoonlovePushCharmRank)(nil), "ui.SCMoonlovePushCharmRank")
	proto.RegisterType((*SCMoonlovePushGenerousRank)(nil), "ui.SCMoonlovePushGenerousRank")
	proto.RegisterType((*SCMoonloveSceneResult)(nil), "ui.SCMoonloveSceneResult")
	proto.RegisterType((*SCMoonloveRankRewards)(nil), "ui.SCMoonloveRankRewards")
	proto.RegisterType((*CSMoonloveViewDouble)(nil), "ui.CSMoonloveViewDouble")
	proto.RegisterType((*SCMoonloveViewDouble)(nil), "ui.SCMoonloveViewDouble")
	proto.RegisterType((*CSMoonloveViewDoubleState)(nil), "ui.CSMoonloveViewDoubleState")
	proto.RegisterType((*SCMoonloveViewDoubleState)(nil), "ui.SCMoonloveViewDoubleState")
	proto.RegisterType((*CSMoonloveViewDoubleRelease)(nil), "ui.CSMoonloveViewDoubleRelease")
	proto.RegisterType((*SCMoonloveViewDoubleRelease)(nil), "ui.SCMoonloveViewDoubleRelease")
	proto.RegisterType((*CSMoonlovePlayerList)(nil), "ui.CSMoonlovePlayerList")
	proto.RegisterType((*SCMoonlovePlayerList)(nil), "ui.SCMoonlovePlayerList")
	proto.RegisterType((*SCMoonloveGiftNotice)(nil), "ui.SCMoonloveGiftNotice")
	proto.RegisterType((*SCMoonloveExpCountNotice)(nil), "ui.SCMoonloveExpCountNotice")
}

func init() { proto.RegisterFile("moonlove.proto", fileDescriptor74) }

var fileDescriptor74 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x41, 0x6f, 0xda, 0x4c,
	0x10, 0x55, 0x4c, 0x88, 0x92, 0xb1, 0x3f, 0xbe, 0x64, 0x93, 0x52, 0x27, 0xb9, 0xa0, 0x4d, 0x54,
	0x71, 0x88, 0xa8, 0x44, 0xd5, 0x43, 0x55, 0xb5, 0x17, 0x87, 0xa2, 0x48, 0x2d, 0x45, 0x50, 0x55,
	0xed, 0xd1, 0x31, 0x03, 0x58, 0xc5, 0xbb, 0xd6, 0xee, 0xba, 0x90, 0xff, 0xd5, 0x1f, 0x58, 0xed,
	0xda, 0xc6, 0x26, 0x98, 0x8a, 0xa3, 0x9f, 0xe6, 0xbd, 0x79, 0x6f, 0xc6, 0x3b, 0xd0, 0x88, 0x38,
	0x67, 0x0b, 0xfe, 0x1b, 0x3b, 0xb1, 0xe0, 0x8a, 0x13, 0x2b, 0x09, 0xaf, 0xce, 0x04, 0x2e, 0x63,
	0xc1, 0x63, 0x14, 0xea, 0x29, 0x85, 0xaf, 0x9c, 0x80, 0x47, 0x11, 0x67, 0xe9, 0x17, 0xed, 0x82,
	0xf3, 0x25, 0xa3, 0x8d, 0x7c, 0xf6, 0x8b, 0x10, 0x80, 0x78, 0xe1, 0x3f, 0xa1, 0x18, 0xf8, 0x11,
	0xba, 0x07, 0x2d, 0xab, 0x7d, 0x42, 0x1a, 0x70, 0xc4, 0x92, 0xe8, 0x11, 0x85, 0x6b, 0xb5, 0xac,
	0x76, 0x9d, 0x4e, 0x80, 0x0c, 0x4d, 0x4d, 0xce, 0x7c, 0x60, 0x53, 0x4e, 0x2e, 0xc0, 0x99, 0x21,
	0x43, 0xc1, 0x13, 0xa9, 0x95, 0x0c, 0xb7, 0x4e, 0xce, 0xe0, 0x24, 0x98, 0xfb, 0x22, 0x32, 0x90,
	0xa1, 0x93, 0x73, 0xb0, 0xf3, 0xc2, 0x41, 0x12, 0xb9, 0x35, 0x03, 0x9e, 0xc2, 0xb1, 0xa9, 0xd3,
	0xc8, 0xa1, 0xe9, 0xf2, 0x13, 0xec, 0x71, 0x80, 0x0c, 0xd3, 0x56, 0xba, 0x20, 0x35, 0xf6, 0x30,
	0x31, 0xd2, 0x35, 0xe2, 0xc0, 0x21, 0xd3, 0x26, 0x2d, 0x63, 0xd2, 0x86, 0x9a, 0xc4, 0x55, 0xa6,
	0xf6, 0x1f, 0xd4, 0x63, 0x1e, 0x32, 0x95, 0x4a, 0xe9, 0x00, 0x52, 0xf9, 0x2a, 0x91, 0x6e, 0xdd,
	0x48, 0xff, 0x80, 0xcb, 0xb1, 0x97, 0x9b, 0xef, 0x67, 0x5e, 0xbc, 0xb9, 0xcf, 0x66, 0x38, 0xa9,
	0x68, 0xf4, 0x3c, 0xd9, 0xee, 0x18, 0xf4, 0x2b, 0x34, 0x0b, 0x65, 0x4f, 0x07, 0xda, 0x2d, 0x5b,
	0x31, 0x9a, 0xf2, 0x14, 0x52, 0xc1, 0x3f, 0x07, 0x70, 0x5e, 0x28, 0x9a, 0x81, 0x98, 0x69, 0xdf,
	0x94, 0xc9, 0x07, 0xad, 0x5a, 0xdb, 0xee, 0x9e, 0x76, 0x92, 0xb0, 0xb3, 0xb1, 0xcc, 0x57, 0x5b,
	0xc6, 0xab, 0xeb, 0xee, 0xc0, 0x89, 0x4a, 0xab, 0x34, 0xad, 0xed, 0x6e, 0x53, 0xd7, 0x55, 0x2c,
	0xba, 0x01, 0x47, 0x02, 0x97, 0xbd, 0x55, 0x6c, 0xa6, 0x59, 0xd3, 0xbf, 0x4c, 0x20, 0xd0, 0x57,
	0xf8, 0x2d, 0x8c, 0xd0, 0x3d, 0xd2, 0x18, 0xfd, 0x08, 0x2f, 0x0b, 0xd7, 0xc3, 0x44, 0xce, 0xbd,
	0xdc, 0xec, 0x5e, 0xce, 0xe9, 0x3d, 0x5c, 0x6d, 0xf2, 0xfb, 0xa5, 0x1c, 0xfb, 0xe6, 0xa2, 0xef,
	0xe0, 0xc5, 0xb3, 0xd9, 0x8d, 0x50, 0x26, 0x0b, 0x55, 0xb1, 0x8c, 0x22, 0x94, 0x65, 0x02, 0x7c,
	0x28, 0x53, 0xb5, 0xd8, 0x08, 0x97, 0xbe, 0x98, 0x48, 0x72, 0x0b, 0xb6, 0xc0, 0xe5, 0x30, 0x7b,
	0x53, 0x86, 0x6d, 0x77, 0xff, 0xd7, 0xad, 0x47, 0x05, 0x4c, 0xdb, 0x70, 0xe1, 0x8d, 0x73, 0xfa,
	0xf7, 0x10, 0x97, 0xf7, 0x3c, 0x79, 0x5c, 0xe0, 0x76, 0x63, 0x3a, 0x85, 0x8b, 0xa2, 0xd1, 0xbf,
	0x2a, 0x49, 0x13, 0x1a, 0xca, 0x17, 0x33, 0x54, 0xc3, 0x1c, 0x37, 0x56, 0xc9, 0xed, 0x1a, 0xe7,
	0x32, 0x54, 0x21, 0x67, 0xd9, 0xfe, 0x1c, 0xb3, 0xbf, 0x0c, 0xa3, 0x6f, 0xe0, 0xb2, 0xca, 0xd1,
	0x58, 0xf9, 0x0a, 0x2b, 0xa4, 0x53, 0x73, 0x9f, 0xca, 0x0f, 0x65, 0x4f, 0x92, 0xfe, 0xaf, 0x43,
	0x39, 0x4e, 0x82, 0x00, 0xa5, 0x34, 0x16, 0x8f, 0xe9, 0x5b, 0xb8, 0xae, 0x6a, 0x3e, 0xc2, 0x05,
	0xfa, 0x72, 0x77, 0xfb, 0xd7, 0x70, 0x5d, 0xd5, 0x3e, 0xa7, 0x6d, 0x0f, 0xb3, 0x59, 0x1e, 0x7b,
	0x2a, 0xf6, 0x39, 0x94, 0x8a, 0xbe, 0x2f, 0x0f, 0xb9, 0xc0, 0xc9, 0x4d, 0x7e, 0xed, 0xf4, 0x57,
	0xf6, 0x33, 0x9a, 0x5d, 0x96, 0x2e, 0x0f, 0xed, 0x95, 0xc9, 0xfd, 0x70, 0xaa, 0x06, 0x5c, 0x85,
	0x01, 0xae, 0xef, 0x4f, 0x7a, 0x24, 0x09, 0x40, 0x9a, 0x61, 0xb0, 0x71, 0x93, 0xd8, 0xfa, 0x25,
	0xdf, 0x81, 0x5b, 0xc8, 0xf4, 0x56, 0xb1, 0xc7, 0x13, 0x96, 0x4b, 0x9d, 0xc2, 0x31, 0x66, 0x48,
	0x9a, 0xe4, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x4e, 0x45, 0xd7, 0xcd, 0x05, 0x00, 0x00,
}
