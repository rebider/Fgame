// Code generated by protoc-gen-go. DO NOT EDIT.
// source: item.proto

package scene

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SceneItemData struct {
	Uid              *int64    `protobuf:"varint,1,req,name=uid" json:"uid,omitempty"`
	ItemId           *int32    `protobuf:"varint,2,req,name=itemId" json:"itemId,omitempty"`
	ItemNum          *int32    `protobuf:"varint,3,req,name=itemNum" json:"itemNum,omitempty"`
	CurPos           *Position `protobuf:"bytes,4,req,name=curPos" json:"curPos,omitempty"`
	Level            *int32    `protobuf:"varint,5,opt,name=level" json:"level,omitempty"`
	State            *bool     `protobuf:"varint,6,req,name=state" json:"state,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *SceneItemData) Reset()                    { *m = SceneItemData{} }
func (m *SceneItemData) String() string            { return proto.CompactTextString(m) }
func (*SceneItemData) ProtoMessage()               {}
func (*SceneItemData) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *SceneItemData) GetUid() int64 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *SceneItemData) GetItemId() int32 {
	if m != nil && m.ItemId != nil {
		return *m.ItemId
	}
	return 0
}

func (m *SceneItemData) GetItemNum() int32 {
	if m != nil && m.ItemNum != nil {
		return *m.ItemNum
	}
	return 0
}

func (m *SceneItemData) GetCurPos() *Position {
	if m != nil {
		return m.CurPos
	}
	return nil
}

func (m *SceneItemData) GetLevel() int32 {
	if m != nil && m.Level != nil {
		return *m.Level
	}
	return 0
}

func (m *SceneItemData) GetState() bool {
	if m != nil && m.State != nil {
		return *m.State
	}
	return false
}

func init() {
	proto.RegisterType((*SceneItemData)(nil), "scene.SceneItemData")
}

func init() { proto.RegisterFile("item.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x24, 0xcd, 0x31, 0xcf, 0x82, 0x30,
	0x10, 0xc6, 0xf1, 0x50, 0xde, 0xf2, 0x9a, 0x43, 0x24, 0xe9, 0x74, 0x71, 0xb1, 0x71, 0xea, 0xc4,
	0xe0, 0x67, 0x70, 0x61, 0x31, 0x24, 0x7e, 0x02, 0x02, 0x37, 0x34, 0xa1, 0x9c, 0xa1, 0x57, 0x27,
	0x3f, 0xbc, 0x29, 0x8e, 0xff, 0xdf, 0xf0, 0x3c, 0x00, 0x5e, 0x28, 0x74, 0xaf, 0x8d, 0x85, 0x8d,
	0x8e, 0x13, 0xad, 0x74, 0x3e, 0x4e, 0x1c, 0x02, 0xaf, 0x3f, 0xbc, 0x7e, 0xa0, 0x79, 0x66, 0xee,
	0x85, 0xc2, 0x7d, 0x94, 0xd1, 0xd4, 0x50, 0x26, 0x3f, 0x63, 0x61, 0x95, 0x2b, 0xcd, 0x09, 0xaa,
	0x3c, 0xd0, 0xcf, 0xa8, 0xac, 0x72, 0xda, 0xb4, 0xf0, 0x9f, 0xfb, 0x91, 0x02, 0x96, 0x3b, 0x5c,
	0xa0, 0x9a, 0xd2, 0x36, 0x70, 0xc4, 0x3f, 0xab, 0x5c, 0x7d, 0x6b, 0xbb, 0xfd, 0xa4, 0x1b, 0x38,
	0x7a, 0xf1, 0xbc, 0x9a, 0x06, 0xf4, 0x42, 0x6f, 0x5a, 0x50, 0xdb, 0xc2, 0xe9, 0x9c, 0x51, 0x46,
	0x21, 0xac, 0xac, 0x72, 0x87, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xa3, 0x9a, 0xba, 0x9f,
	0x00, 0x00, 0x00,
}
