// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tulong.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TuLongRank struct {
	Pos              *int32  `protobuf:"varint,1,req,name=pos" json:"pos,omitempty"`
	ServerId         *int32  `protobuf:"varint,2,req,name=serverId" json:"serverId,omitempty"`
	Name             *string `protobuf:"bytes,3,req,name=name" json:"name,omitempty"`
	Num              *int32  `protobuf:"varint,4,req,name=num" json:"num,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TuLongRank) Reset()                    { *m = TuLongRank{} }
func (m *TuLongRank) String() string            { return proto.CompactTextString(m) }
func (*TuLongRank) ProtoMessage()               {}
func (*TuLongRank) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{0} }

func (m *TuLongRank) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

func (m *TuLongRank) GetServerId() int32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *TuLongRank) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TuLongRank) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

type CSTuLongRank struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSTuLongRank) Reset()                    { *m = CSTuLongRank{} }
func (m *CSTuLongRank) String() string            { return proto.CompactTextString(m) }
func (*CSTuLongRank) ProtoMessage()               {}
func (*CSTuLongRank) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{1} }

type SCTuLongRank struct {
	RankList         []*TuLongRank `protobuf:"bytes,1,rep,name=rankList" json:"rankList,omitempty"`
	Pos              *int32        `protobuf:"varint,2,req,name=pos" json:"pos,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *SCTuLongRank) Reset()                    { *m = SCTuLongRank{} }
func (m *SCTuLongRank) String() string            { return proto.CompactTextString(m) }
func (*SCTuLongRank) ProtoMessage()               {}
func (*SCTuLongRank) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{2} }

func (m *SCTuLongRank) GetRankList() []*TuLongRank {
	if m != nil {
		return m.RankList
	}
	return nil
}

func (m *SCTuLongRank) GetPos() int32 {
	if m != nil && m.Pos != nil {
		return *m.Pos
	}
	return 0
}

type CSTuLongCollect struct {
	NpcId            *int64 `protobuf:"varint,1,req,name=npcId" json:"npcId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSTuLongCollect) Reset()                    { *m = CSTuLongCollect{} }
func (m *CSTuLongCollect) String() string            { return proto.CompactTextString(m) }
func (*CSTuLongCollect) ProtoMessage()               {}
func (*CSTuLongCollect) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{3} }

func (m *CSTuLongCollect) GetNpcId() int64 {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return 0
}

type SCTuLongCollect struct {
	NpcId            *int64 `protobuf:"varint,1,req,name=npcId" json:"npcId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTuLongCollect) Reset()                    { *m = SCTuLongCollect{} }
func (m *SCTuLongCollect) String() string            { return proto.CompactTextString(m) }
func (*SCTuLongCollect) ProtoMessage()               {}
func (*SCTuLongCollect) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{4} }

func (m *SCTuLongCollect) GetNpcId() int64 {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return 0
}

type SCTuLongCollectStop struct {
	NpcId            *int64 `protobuf:"varint,1,req,name=npcId" json:"npcId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTuLongCollectStop) Reset()                    { *m = SCTuLongCollectStop{} }
func (m *SCTuLongCollectStop) String() string            { return proto.CompactTextString(m) }
func (*SCTuLongCollectStop) ProtoMessage()               {}
func (*SCTuLongCollectStop) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{5} }

func (m *SCTuLongCollectStop) GetNpcId() int64 {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return 0
}

type SCTuLongCollectFinish struct {
	NpcId            *int64 `protobuf:"varint,1,req,name=npcId" json:"npcId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTuLongCollectFinish) Reset()                    { *m = SCTuLongCollectFinish{} }
func (m *SCTuLongCollectFinish) String() string            { return proto.CompactTextString(m) }
func (*SCTuLongCollectFinish) ProtoMessage()               {}
func (*SCTuLongCollectFinish) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{6} }

func (m *SCTuLongCollectFinish) GetNpcId() int64 {
	if m != nil && m.NpcId != nil {
		return *m.NpcId
	}
	return 0
}

type SCTuLongBossStatus struct {
	Status           *int32 `protobuf:"varint,1,req,name=status" json:"status,omitempty"`
	BiaoShi          *int32 `protobuf:"varint,2,req,name=biaoShi" json:"biaoShi,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTuLongBossStatus) Reset()                    { *m = SCTuLongBossStatus{} }
func (m *SCTuLongBossStatus) String() string            { return proto.CompactTextString(m) }
func (*SCTuLongBossStatus) ProtoMessage()               {}
func (*SCTuLongBossStatus) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{7} }

func (m *SCTuLongBossStatus) GetStatus() int32 {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return 0
}

func (m *SCTuLongBossStatus) GetBiaoShi() int32 {
	if m != nil && m.BiaoShi != nil {
		return *m.BiaoShi
	}
	return 0
}

type SCTuLongStart struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTuLongStart) Reset()                    { *m = SCTuLongStart{} }
func (m *SCTuLongStart) String() string            { return proto.CompactTextString(m) }
func (*SCTuLongStart) ProtoMessage()               {}
func (*SCTuLongStart) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{8} }

type SCTuLongAllianceBiaoShi struct {
	BiaoShi          *int32 `protobuf:"varint,1,req,name=biaoShi" json:"biaoShi,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTuLongAllianceBiaoShi) Reset()                    { *m = SCTuLongAllianceBiaoShi{} }
func (m *SCTuLongAllianceBiaoShi) String() string            { return proto.CompactTextString(m) }
func (*SCTuLongAllianceBiaoShi) ProtoMessage()               {}
func (*SCTuLongAllianceBiaoShi) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{9} }

func (m *SCTuLongAllianceBiaoShi) GetBiaoShi() int32 {
	if m != nil && m.BiaoShi != nil {
		return *m.BiaoShi
	}
	return 0
}

type SCTuLongResult struct {
	Num              *int32      `protobuf:"varint,1,req,name=num" json:"num,omitempty"`
	ItemList         []*ItemInfo `protobuf:"bytes,2,rep,name=itemList" json:"itemList,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *SCTuLongResult) Reset()                    { *m = SCTuLongResult{} }
func (m *SCTuLongResult) String() string            { return proto.CompactTextString(m) }
func (*SCTuLongResult) ProtoMessage()               {}
func (*SCTuLongResult) Descriptor() ([]byte, []int) { return fileDescriptor121, []int{10} }

func (m *SCTuLongResult) GetNum() int32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *SCTuLongResult) GetItemList() []*ItemInfo {
	if m != nil {
		return m.ItemList
	}
	return nil
}

func init() {
	proto.RegisterType((*TuLongRank)(nil), "ui.TuLongRank")
	proto.RegisterType((*CSTuLongRank)(nil), "ui.CSTuLongRank")
	proto.RegisterType((*SCTuLongRank)(nil), "ui.SCTuLongRank")
	proto.RegisterType((*CSTuLongCollect)(nil), "ui.CSTuLongCollect")
	proto.RegisterType((*SCTuLongCollect)(nil), "ui.SCTuLongCollect")
	proto.RegisterType((*SCTuLongCollectStop)(nil), "ui.SCTuLongCollectStop")
	proto.RegisterType((*SCTuLongCollectFinish)(nil), "ui.SCTuLongCollectFinish")
	proto.RegisterType((*SCTuLongBossStatus)(nil), "ui.SCTuLongBossStatus")
	proto.RegisterType((*SCTuLongStart)(nil), "ui.SCTuLongStart")
	proto.RegisterType((*SCTuLongAllianceBiaoShi)(nil), "ui.SCTuLongAllianceBiaoShi")
	proto.RegisterType((*SCTuLongResult)(nil), "ui.SCTuLongResult")
}

func init() { proto.RegisterFile("tulong.proto", fileDescriptor121) }

var fileDescriptor121 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x59, 0xbb, 0xfd, 0xff, 0xf3, 0x5d, 0xd7, 0x49, 0x44, 0x2c, 0x1e, 0xa4, 0x04, 0x91,
	0xe1, 0x61, 0x07, 0xc1, 0xe3, 0x0e, 0x6e, 0x30, 0x28, 0xec, 0x64, 0xfc, 0x02, 0xb1, 0x8b, 0x5b,
	0x58, 0x9a, 0x94, 0xe6, 0x8d, 0x9f, 0x5f, 0x9a, 0x2e, 0x3a, 0x2a, 0xde, 0x92, 0x37, 0xbf, 0xe7,
	0x79, 0x92, 0x27, 0x90, 0xa0, 0x53, 0x46, 0xef, 0x17, 0x75, 0x63, 0xd0, 0x90, 0xc8, 0xc9, 0x5b,
	0x90, 0x28, 0xaa, 0x6e, 0x4f, 0x37, 0x00, 0x6f, 0x6e, 0x6b, 0xf4, 0xfe, 0x95, 0xeb, 0x23, 0x99,
	0x40, 0x5c, 0x1b, 0x9b, 0x0d, 0xf2, 0x68, 0x3e, 0x22, 0x97, 0x30, 0xb6, 0xa2, 0xf9, 0x14, 0x4d,
	0xb1, 0xcb, 0x22, 0x3f, 0x49, 0x60, 0xa8, 0x79, 0x25, 0xb2, 0x38, 0x8f, 0xe6, 0x17, 0x2d, 0xac,
	0x5d, 0x95, 0x0d, 0xdb, 0x23, 0x9a, 0x42, 0xb2, 0x66, 0x3f, 0x4e, 0x74, 0x09, 0x09, 0x5b, 0x9f,
	0x39, 0xe7, 0x30, 0x6e, 0xb8, 0x3e, 0x6e, 0xa5, 0xc5, 0x6c, 0x90, 0xc7, 0xf3, 0xc9, 0x53, 0xba,
	0x70, 0x72, 0xf1, 0x3b, 0xdb, 0x27, 0xd1, 0x1c, 0x66, 0xc1, 0x6e, 0x6d, 0x94, 0x12, 0x25, 0x92,
	0x29, 0x8c, 0x74, 0x5d, 0x16, 0x3b, 0x7f, 0xbb, 0xb8, 0x25, 0x42, 0xc0, 0x1f, 0xc4, 0x3d, 0x5c,
	0xf5, 0x08, 0x86, 0xa6, 0xee, 0x53, 0x0f, 0x70, 0xdd, 0xa3, 0x36, 0x52, 0x4b, 0x7b, 0xe8, 0x73,
	0xcf, 0x40, 0x02, 0xb7, 0x32, 0xd6, 0x32, 0xe4, 0xe8, 0x2c, 0x49, 0xe1, 0x9f, 0xf5, 0xab, 0x53,
	0x67, 0x33, 0xf8, 0xff, 0x2e, 0xb9, 0x61, 0x07, 0x79, 0x7a, 0xc8, 0x0c, 0xa6, 0x41, 0xc6, 0x90,
	0x37, 0x48, 0x1f, 0xe1, 0x26, 0x0c, 0x5e, 0x94, 0x92, 0x5c, 0x97, 0x62, 0xd5, 0x29, 0xce, 0xc5,
	0xde, 0x8d, 0x2e, 0x21, 0xfd, 0x2e, 0x51, 0x58, 0xa7, 0x30, 0x74, 0xde, 0x85, 0xdd, 0xc1, 0xb8,
	0xfd, 0x49, 0xdf, 0x69, 0xe4, 0x3b, 0x4d, 0xda, 0x4e, 0x0b, 0x14, 0x55, 0xa1, 0x3f, 0xcc, 0x57,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x74, 0x77, 0x86, 0xfa, 0x01, 0x00, 0x00,
}
