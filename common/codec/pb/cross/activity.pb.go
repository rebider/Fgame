// Code generated by protoc-gen-go. DO NOT EDIT.
// source: activity.proto

/*
Package cross is a generated protocol buffer package.

It is generated from these files:
	activity.proto
	alliance.proto
	arena.proto
	arenapvp.proto
	battle.proto
	buff.proto
	chuangshi.proto
	collect.proto
	common.proto
	dense_wat.proto
	drop.proto
	godsiege.proto
	item.proto
	jieyi.proto
	lianyu.proto
	lineup.proto
	lingtong.proto
	login.proto
	massacre.proto
	message_type.proto
	pk.proto
	player.proto
	property.proto
	qixue.proto
	relive.proto
	scene.proto
	shenmo.proto
	show.proto
	skill.proto
	team.proto
	teamcopy.proto
	tulong.proto
	xuechi.proto

It has these top-level messages:
	ActivityPkData
	ISPlayerActivityPkDataChanged
	SIPlayerActivityPkDataChanged
	PlayerActivityRankData
	ActivityRankData
	ISPlayerActivityRankDataChanged
	ISPlayerActivityTickRewDataChanged
	AllianceData
	SIPlayerAllianceSync
	ArenaPlayer
	SIArenaMatch
	ISArenaMatch
	ISArenaMatchResult
	SIArenaStopMatch
	ISArenaStopMatch
	ISArenaWin
	SIArenaWin
	PlayerArenaData
	SIPlayerArenaDataChanged
	ISArenaRelive
	SIArenaRelive
	ISArenaCollectExpTree
	SIArenaCollectExpTree
	ISArenaCollectBox
	SIArenaCollectBox
	ISArenaGiveUp
	SIArenaGiveUp
	ISArenaResetReliveTimes
	SIArenapvpAttend
	ISArenapvpAttend
	ISArenapvpRelive
	SIArenapvpRelive
	PlayerArenapvpData
	SIPlayerArenapvpDataChanged
	ISArenapvpResetReliveTimes
	ISArenapvpAttendSuccess
	ISArenapvpResultBattle
	ISArenapvpResultElection
	ISAreanapvpElectionLuckyRew
	PlayerBattleData
	SIPlayerBattleDataChanged
	BuffData
	SIBuffAdd
	SIBuffRemove
	SIBuffUpdate
	ChuangShiData
	SIChuangShiEnterCity
	ISChuangShiEnterCity
	ISChuangShiKillPlayer
	SIChuangShiSceneFinish
	ISChuangShiSceneFinish
	ISCollectFinish
	SICollectFinish
	ISCollectMiZangFinish
	SICollectMiZangFinish
	SIHeartBeat
	ISHeartBeat
	DenseWatData
	ISDenseWatSync
	SIDenseWatSync
	ISPlayerGetDropItem
	SIPlayerGetDropItem
	SIGodSiegeAttend
	ISGodSiegeAttend
	ISGodSiegeLineUpSuccess
	SIGodSiegeLineUpSuccess
	SIGodSiegeCancleLineUp
	ISGodSiegeCancleLineUp
	ISGodSiegeFinishLineUpCancle
	SIGodSiegeFinishLineUpCancle
	ItemInfo
	JieYiData
	SIPlayerJieYiSync
	ISShengWeiDrop
	SIShengWeiDrop
	SILianYuAttend
	ISLianYuAttend
	ISLianYuLineUpSuccess
	SILianYuLineUpSuccess
	SILianYuCancleLineUp
	ISLianYuCancleLineUp
	ISLianYuFinishLineUpCancle
	SILianYuFinishLineUpCancle
	SILineupAttend
	ISLineupAttend
	SILineupCancle
	ISLineupCancle
	ISLineupSuccess
	SILineupSuccess
	ISLineupSceneFinishToCancel
	LingTongShowData
	LingTongData
	SILingTongDataInit
	SILingTongDataChanged
	SILingTongDataRemove
	SILogin
	ISLogin
	ISMassacreDrop
	SIMassacreDrop
	PkData
	PlayerBasicData
	Property
	PropertyData
	ISQiXueDrop
	SIQiXueDrop
	PlayerReliveData
	ISPlayerReliveSync
	SIPlayerReliveSync
	ISPlayerRelive
	SIPlayerRelive
	CrossData
	SIPlayerData
	SIPlayerSystemBattlePropertyChanged
	SIPlayerBasicPropertyChanged
	SIPlayerExitCross
	ISPlayerExitCross
	ISPlayerKillBiology
	SIPlayerKillBiology
	ISPlayerMountSync
	PlayerBossReliveData
	ISPlayerBossReliveSync
	ShenMoData
	ISShenMoSync
	SIShenMoSync
	SIShenMoAttend
	ISShenMoAttend
	ISShenMoLineUpSuccess
	SIShenMoLineUpSuccess
	SIShenMoCancleLineUp
	ISShenMoCancleLineUp
	ISShenMoFinishLineUpCancle
	SIShenMoFinishLineUpCancle
	ISShenMoKillNumChanged
	SIShenMoKillNumChanged
	ISPlayerGongXunAdd
	SIPlayerGongXunAdd
	ISPlayerGongXunSub
	SIPlayerGongXunSub
	SIPlayerGongXunChanged
	ISPlayerGongXunChanged
	PlayerShowData
	SIPlayerShowDataChanged
	SkillTianFu
	SkillData
	TeShuSkillData
	SIPlayerTeshuSkillReset
	TeamData
	SIPlayerTeamSync
	TeamPlayer
	SITeamCopyStartBattle
	ISTeamCopyStartBattle
	ISTeamCopyBattleResult
	SITeamCopyBattleResult
	SITuLongAttend
	ISTuLongAttend
	ISTuLongKillBoss
	SITuLongKillBoss
	XueChiData
	ISXueChiSync
	SIXueChiSync
	SIXueChiAdd
*/
package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ActivityPkData struct {
	ActivityType     *int32 `protobuf:"varint,1,req,name=activityType" json:"activityType,omitempty"`
	KilledNum        *int32 `protobuf:"varint,2,req,name=killedNum" json:"killedNum,omitempty"`
	LastKillTime     *int64 `protobuf:"varint,3,req,name=lastKillTime" json:"lastKillTime,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ActivityPkData) Reset()                    { *m = ActivityPkData{} }
func (m *ActivityPkData) String() string            { return proto.CompactTextString(m) }
func (*ActivityPkData) ProtoMessage()               {}
func (*ActivityPkData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActivityPkData) GetActivityType() int32 {
	if m != nil && m.ActivityType != nil {
		return *m.ActivityType
	}
	return 0
}

func (m *ActivityPkData) GetKilledNum() int32 {
	if m != nil && m.KilledNum != nil {
		return *m.KilledNum
	}
	return 0
}

func (m *ActivityPkData) GetLastKillTime() int64 {
	if m != nil && m.LastKillTime != nil {
		return *m.LastKillTime
	}
	return 0
}

type ISPlayerActivityPkDataChanged struct {
	ActivityPkData   *ActivityPkData `protobuf:"bytes,1,req,name=activityPkData" json:"activityPkData,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ISPlayerActivityPkDataChanged) Reset()                    { *m = ISPlayerActivityPkDataChanged{} }
func (m *ISPlayerActivityPkDataChanged) String() string            { return proto.CompactTextString(m) }
func (*ISPlayerActivityPkDataChanged) ProtoMessage()               {}
func (*ISPlayerActivityPkDataChanged) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ISPlayerActivityPkDataChanged) GetActivityPkData() *ActivityPkData {
	if m != nil {
		return m.ActivityPkData
	}
	return nil
}

type SIPlayerActivityPkDataChanged struct {
	ActivityPkData   *ActivityPkData `protobuf:"bytes,1,req,name=activityPkData" json:"activityPkData,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SIPlayerActivityPkDataChanged) Reset()                    { *m = SIPlayerActivityPkDataChanged{} }
func (m *SIPlayerActivityPkDataChanged) String() string            { return proto.CompactTextString(m) }
func (*SIPlayerActivityPkDataChanged) ProtoMessage()               {}
func (*SIPlayerActivityPkDataChanged) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SIPlayerActivityPkDataChanged) GetActivityPkData() *ActivityPkData {
	if m != nil {
		return m.ActivityPkData
	}
	return nil
}

type PlayerActivityRankData struct {
	ActivityType     *int32 `protobuf:"varint,1,req,name=activityType" json:"activityType,omitempty"`
	RankType         *int32 `protobuf:"varint,2,req,name=rankType" json:"rankType,omitempty"`
	Val              *int64 `protobuf:"varint,3,req,name=val" json:"val,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *PlayerActivityRankData) Reset()                    { *m = PlayerActivityRankData{} }
func (m *PlayerActivityRankData) String() string            { return proto.CompactTextString(m) }
func (*PlayerActivityRankData) ProtoMessage()               {}
func (*PlayerActivityRankData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PlayerActivityRankData) GetActivityType() int32 {
	if m != nil && m.ActivityType != nil {
		return *m.ActivityType
	}
	return 0
}

func (m *PlayerActivityRankData) GetRankType() int32 {
	if m != nil && m.RankType != nil {
		return *m.RankType
	}
	return 0
}

func (m *PlayerActivityRankData) GetVal() int64 {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return 0
}

type ActivityRankData struct {
	ActivityType       *int32                    `protobuf:"varint,1,req,name=activityType" json:"activityType,omitempty"`
	PlayerRankDataList []*PlayerActivityRankData `protobuf:"bytes,2,rep,name=playerRankDataList" json:"playerRankDataList,omitempty"`
	EndTime            *int64                    `protobuf:"varint,3,req,name=endTime" json:"endTime,omitempty"`
	XXX_unrecognized   []byte                    `json:"-"`
}

func (m *ActivityRankData) Reset()                    { *m = ActivityRankData{} }
func (m *ActivityRankData) String() string            { return proto.CompactTextString(m) }
func (*ActivityRankData) ProtoMessage()               {}
func (*ActivityRankData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ActivityRankData) GetActivityType() int32 {
	if m != nil && m.ActivityType != nil {
		return *m.ActivityType
	}
	return 0
}

func (m *ActivityRankData) GetPlayerRankDataList() []*PlayerActivityRankData {
	if m != nil {
		return m.PlayerRankDataList
	}
	return nil
}

func (m *ActivityRankData) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

type ISPlayerActivityRankDataChanged struct {
	ActivityType     *int32 `protobuf:"varint,1,req,name=activityType" json:"activityType,omitempty"`
	RankType         *int32 `protobuf:"varint,2,req,name=rankType" json:"rankType,omitempty"`
	Val              *int64 `protobuf:"varint,3,req,name=val" json:"val,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ISPlayerActivityRankDataChanged) Reset()                    { *m = ISPlayerActivityRankDataChanged{} }
func (m *ISPlayerActivityRankDataChanged) String() string            { return proto.CompactTextString(m) }
func (*ISPlayerActivityRankDataChanged) ProtoMessage()               {}
func (*ISPlayerActivityRankDataChanged) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ISPlayerActivityRankDataChanged) GetActivityType() int32 {
	if m != nil && m.ActivityType != nil {
		return *m.ActivityType
	}
	return 0
}

func (m *ISPlayerActivityRankDataChanged) GetRankType() int32 {
	if m != nil && m.RankType != nil {
		return *m.RankType
	}
	return 0
}

func (m *ISPlayerActivityRankDataChanged) GetVal() int64 {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return 0
}

type ISPlayerActivityTickRewDataChanged struct {
	PropertyList        []*Property `protobuf:"bytes,1,rep,name=propertyList" json:"propertyList,omitempty"`
	SpecialPropertyList []*Property `protobuf:"bytes,2,rep,name=specialPropertyList" json:"specialPropertyList,omitempty"`
	XXX_unrecognized    []byte      `json:"-"`
}

func (m *ISPlayerActivityTickRewDataChanged) Reset()         { *m = ISPlayerActivityTickRewDataChanged{} }
func (m *ISPlayerActivityTickRewDataChanged) String() string { return proto.CompactTextString(m) }
func (*ISPlayerActivityTickRewDataChanged) ProtoMessage()    {}
func (*ISPlayerActivityTickRewDataChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6}
}

func (m *ISPlayerActivityTickRewDataChanged) GetPropertyList() []*Property {
	if m != nil {
		return m.PropertyList
	}
	return nil
}

func (m *ISPlayerActivityTickRewDataChanged) GetSpecialPropertyList() []*Property {
	if m != nil {
		return m.SpecialPropertyList
	}
	return nil
}

func init() {
	proto.RegisterType((*ActivityPkData)(nil), "cross.ActivityPkData")
	proto.RegisterType((*ISPlayerActivityPkDataChanged)(nil), "cross.ISPlayerActivityPkDataChanged")
	proto.RegisterType((*SIPlayerActivityPkDataChanged)(nil), "cross.SIPlayerActivityPkDataChanged")
	proto.RegisterType((*PlayerActivityRankData)(nil), "cross.PlayerActivityRankData")
	proto.RegisterType((*ActivityRankData)(nil), "cross.ActivityRankData")
	proto.RegisterType((*ISPlayerActivityRankDataChanged)(nil), "cross.ISPlayerActivityRankDataChanged")
	proto.RegisterType((*ISPlayerActivityTickRewDataChanged)(nil), "cross.ISPlayerActivityTickRewDataChanged")
}

func init() { proto.RegisterFile("activity.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x59, 0xcb, 0x50, 0x5f, 0x4b, 0x37, 0xe3, 0x94, 0x22, 0x0c, 0x4b, 0x41, 0xe8, 0x41,
	0x7b, 0xd8, 0xcd, 0xa3, 0xe8, 0x65, 0xa8, 0xb3, 0x6c, 0xc5, 0xfb, 0xa3, 0x0d, 0x1a, 0x9a, 0xb5,
	0x21, 0x8d, 0x93, 0xfe, 0xf7, 0x62, 0xd6, 0x60, 0x3b, 0x54, 0x44, 0xbc, 0xbe, 0x7c, 0xf9, 0xe5,
	0xfb, 0x3d, 0x02, 0x1e, 0x66, 0x8a, 0x6d, 0x98, 0x6a, 0x62, 0x21, 0x2b, 0x55, 0x91, 0x61, 0x26,
	0xab, 0xba, 0x3e, 0xf5, 0x84, 0xac, 0x04, 0x95, 0x66, 0x1c, 0x3e, 0x82, 0x77, 0xdd, 0x06, 0x93,
	0xe2, 0x16, 0x15, 0x92, 0x09, 0xb8, 0xe6, 0x6a, 0xda, 0x08, 0xea, 0x0f, 0x02, 0x2b, 0x1a, 0x92,
	0x43, 0x38, 0x28, 0x18, 0xe7, 0x34, 0x5f, 0xbc, 0xae, 0x7d, 0x4b, 0x8f, 0x26, 0xe0, 0x72, 0xac,
	0xd5, 0x1d, 0xe3, 0x3c, 0x65, 0x6b, 0xea, 0xdb, 0x81, 0x15, 0xd9, 0xe1, 0x02, 0xa6, 0xf3, 0x55,
	0xc2, 0xb1, 0xa1, 0xb2, 0x0f, 0xbe, 0x79, 0xc1, 0xf2, 0x99, 0xe6, 0xe4, 0xf2, 0xb3, 0xda, 0xf6,
	0x40, 0xbf, 0xe0, 0xcc, 0x8e, 0x63, 0xdd, 0x30, 0xee, 0xdf, 0xfa, 0xe0, 0xad, 0xe6, 0xff, 0xc8,
	0x7b, 0x80, 0x93, 0x3e, 0x6d, 0x89, 0xe5, 0x4f, 0xe2, 0x63, 0xd8, 0x97, 0x58, 0x16, 0x7a, 0xb2,
	0xf5, 0x76, 0xc0, 0xde, 0x20, 0x6f, 0x75, 0x15, 0x8c, 0x7f, 0x09, 0xba, 0x02, 0x22, 0xf4, 0xc3,
	0x26, 0x77, 0xcf, 0x6a, 0xe5, 0x5b, 0x81, 0x1d, 0x39, 0xb3, 0x69, 0xdb, 0xf5, 0x9b, 0x66, 0x23,
	0xd8, 0xa3, 0x65, 0xde, 0x59, 0xf2, 0x13, 0x9c, 0xed, 0x2e, 0xd9, 0x84, 0xcd, 0x5a, 0xfe, 0x64,
	0xd3, 0x40, 0xb8, 0xcb, 0x4d, 0x59, 0x56, 0x2c, 0xe9, 0x5b, 0x17, 0x7d, 0x0e, 0xae, 0xf9, 0x45,
	0xda, 0x61, 0xa0, 0x1d, 0x46, 0xc6, 0xa1, 0x3d, 0x22, 0x17, 0x70, 0x54, 0x0b, 0x9a, 0x31, 0xe4,
	0x49, 0x37, 0x6d, 0x7d, 0x99, 0x7e, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x73, 0x31, 0x88, 0xb0,
	0x02, 0x00, 0x00,
}