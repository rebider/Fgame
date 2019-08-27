// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jieyi.proto

package cross

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type JieYiData struct {
	JieYiName        *string `protobuf:"bytes,1,req,name=jieYiName" json:"jieYiName,omitempty"`
	Rank             *int32  `protobuf:"varint,2,req,name=rank" json:"rank,omitempty"`
	JieYiId          *int64  `protobuf:"varint,3,req,name=jieYiId" json:"jieYiId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *JieYiData) Reset()                    { *m = JieYiData{} }
func (m *JieYiData) String() string            { return proto.CompactTextString(m) }
func (*JieYiData) ProtoMessage()               {}
func (*JieYiData) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{0} }

func (m *JieYiData) GetJieYiName() string {
	if m != nil && m.JieYiName != nil {
		return *m.JieYiName
	}
	return ""
}

func (m *JieYiData) GetRank() int32 {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return 0
}

func (m *JieYiData) GetJieYiId() int64 {
	if m != nil && m.JieYiId != nil {
		return *m.JieYiId
	}
	return 0
}

type SIPlayerJieYiSync struct {
	JieYiData        *JieYiData `protobuf:"bytes,1,req,name=jieYiData" json:"jieYiData,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *SIPlayerJieYiSync) Reset()                    { *m = SIPlayerJieYiSync{} }
func (m *SIPlayerJieYiSync) String() string            { return proto.CompactTextString(m) }
func (*SIPlayerJieYiSync) ProtoMessage()               {}
func (*SIPlayerJieYiSync) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{1} }

func (m *SIPlayerJieYiSync) GetJieYiData() *JieYiData {
	if m != nil {
		return m.JieYiData
	}
	return nil
}

type ISShengWeiDrop struct {
	AttackerId       *int64  `protobuf:"varint,1,req,name=attackerId" json:"attackerId,omitempty"`
	AttackerName     *string `protobuf:"bytes,2,req,name=attackerName" json:"attackerName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ISShengWeiDrop) Reset()                    { *m = ISShengWeiDrop{} }
func (m *ISShengWeiDrop) String() string            { return proto.CompactTextString(m) }
func (*ISShengWeiDrop) ProtoMessage()               {}
func (*ISShengWeiDrop) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{2} }

func (m *ISShengWeiDrop) GetAttackerId() int64 {
	if m != nil && m.AttackerId != nil {
		return *m.AttackerId
	}
	return 0
}

func (m *ISShengWeiDrop) GetAttackerName() string {
	if m != nil && m.AttackerName != nil {
		return *m.AttackerName
	}
	return ""
}

type SIShengWeiDrop struct {
	ItemId           *int32 `protobuf:"varint,1,req,name=itemId" json:"itemId,omitempty"`
	ItemNum          *int64 `protobuf:"varint,2,req,name=itemNum" json:"itemNum,omitempty"`
	AttackerId       *int64 `protobuf:"varint,3,req,name=attackerId" json:"attackerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SIShengWeiDrop) Reset()                    { *m = SIShengWeiDrop{} }
func (m *SIShengWeiDrop) String() string            { return proto.CompactTextString(m) }
func (*SIShengWeiDrop) ProtoMessage()               {}
func (*SIShengWeiDrop) Descriptor() ([]byte, []int) { return fileDescriptor13, []int{3} }

func (m *SIShengWeiDrop) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *SIShengWeiDrop) GetItemNum() int64 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return 0
}

func (m *SIShengWeiDrop) GetAttackerId() int64 {
	if m != nil && m.AttackerId != nil {
		return *m.AttackerId
	}
	return 0
}

func init() {
	proto.RegisterType((*JieYiData)(nil), "cross.JieYiData")
	proto.RegisterType((*SIPlayerJieYiSync)(nil), "cross.SIPlayerJieYiSync")
	proto.RegisterType((*ISShengWeiDrop)(nil), "cross.ISShengWeiDrop")
	proto.RegisterType((*SIShengWeiDrop)(nil), "cross.SIShengWeiDrop")
}

func init() { proto.RegisterFile("jieyi.proto", fileDescriptor13) }

var fileDescriptor13 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0x87, 0x31, 0xf1, 0x94, 0xbc, 0x3b, 0xaa, 0xf7, 0x70, 0xe8, 0x58, 0xe2, 0xd2, 0xa9, 0x83,
	0x93, 0x08, 0x6e, 0xe7, 0x10, 0x87, 0x43, 0xc8, 0x20, 0x8e, 0xa1, 0xf7, 0xd0, 0xb4, 0xb6, 0x29,
	0x69, 0x1c, 0xfa, 0xdf, 0x4b, 0x1e, 0x55, 0xd0, 0xf1, 0x85, 0x7c, 0xdf, 0xef, 0x83, 0x6d, 0xe7,
	0x69, 0xf1, 0xcd, 0x14, 0x43, 0x0a, 0xb8, 0x69, 0x63, 0x98, 0x67, 0xfd, 0x08, 0xea, 0xd9, 0xd3,
	0x9b, 0x3f, 0xb8, 0xe4, 0x70, 0x0f, 0xaa, 0xcb, 0xc7, 0xd1, 0x0d, 0x54, 0x9e, 0x55, 0xa2, 0x56,
	0xb8, 0x83, 0xf3, 0xe8, 0xc6, 0xbe, 0x14, 0x95, 0xa8, 0x37, 0x78, 0x05, 0x97, 0xfc, 0xc1, 0x9c,
	0x4a, 0x59, 0x89, 0x5a, 0xea, 0x7b, 0xd8, 0x5b, 0xf3, 0xf2, 0xe9, 0x16, 0x8a, 0xac, 0xb1, 0xcb,
	0xd8, 0xe2, 0xed, 0xaa, 0xc9, 0x4e, 0xd6, 0x6c, 0xef, 0xae, 0x1b, 0x9e, 0x6b, 0x7e, 0xb7, 0xf4,
	0x03, 0x14, 0xc6, 0xda, 0x0f, 0x1a, 0xdf, 0x5f, 0xc9, 0x1f, 0x62, 0x98, 0x10, 0x01, 0x5c, 0x4a,
	0xae, 0xed, 0x29, 0x9a, 0x13, 0x73, 0x12, 0x6f, 0x60, 0xf7, 0xf3, 0xc6, 0x51, 0x39, 0x43, 0xe9,
	0x27, 0x28, 0xac, 0xf9, 0xc3, 0x16, 0x70, 0xe1, 0x13, 0x0d, 0x2b, 0xc7, 0xa1, 0xf9, 0x3e, 0x7e,
	0x0d, 0x8c, 0xc8, 0x7f, 0x72, 0x8e, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x58, 0x3c, 0x38, 0x68,
	0x10, 0x01, 0x00, 0x00,
}