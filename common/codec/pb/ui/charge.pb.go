// Code generated by protoc-gen-go. DO NOT EDIT.
// source: charge.proto

package ui

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CSCharge struct {
	ChargeId         *int32 `protobuf:"varint,1,req,name=chargeId" json:"chargeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSCharge) Reset()                    { *m = CSCharge{} }
func (m *CSCharge) String() string            { return proto.CompactTextString(m) }
func (*CSCharge) ProtoMessage()               {}
func (*CSCharge) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{0} }

func (m *CSCharge) GetChargeId() int32 {
	if m != nil && m.ChargeId != nil {
		return *m.ChargeId
	}
	return 0
}

type SCCharge struct {
	ChargeId         *int32 `protobuf:"varint,1,req,name=chargeId" json:"chargeId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCCharge) Reset()                    { *m = SCCharge{} }
func (m *SCCharge) String() string            { return proto.CompactTextString(m) }
func (*SCCharge) ProtoMessage()               {}
func (*SCCharge) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{1} }

func (m *SCCharge) GetChargeId() int32 {
	if m != nil && m.ChargeId != nil {
		return *m.ChargeId
	}
	return 0
}

type SCChargeOrder struct {
	ChargeId         *int32  `protobuf:"varint,1,req,name=chargeId" json:"chargeId,omitempty"`
	ChargeType       *int32  `protobuf:"varint,2,req,name=chargeType" json:"chargeType,omitempty"`
	OrderId          *string `protobuf:"bytes,3,req,name=orderId" json:"orderId,omitempty"`
	NotifyUrl        *string `protobuf:"bytes,4,req,name=notifyUrl" json:"notifyUrl,omitempty"`
	SdkOrderId       *string `protobuf:"bytes,5,req,name=sdkOrderId" json:"sdkOrderId,omitempty"`
	PlatformUserId   *string `protobuf:"bytes,6,req,name=platformUserId" json:"platformUserId,omitempty"`
	Money            *int32  `protobuf:"varint,7,req,name=money" json:"money,omitempty"`
	PlayerId         *int64  `protobuf:"varint,8,req,name=playerId" json:"playerId,omitempty"`
	PlayerName       *string `protobuf:"bytes,9,req,name=playerName" json:"playerName,omitempty"`
	ServerId         *int32  `protobuf:"varint,10,req,name=serverId" json:"serverId,omitempty"`
	ServerName       *string `protobuf:"bytes,11,req,name=serverName" json:"serverName,omitempty"`
	Extension        *string `protobuf:"bytes,12,req,name=extension" json:"extension,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCChargeOrder) Reset()                    { *m = SCChargeOrder{} }
func (m *SCChargeOrder) String() string            { return proto.CompactTextString(m) }
func (*SCChargeOrder) ProtoMessage()               {}
func (*SCChargeOrder) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{2} }

func (m *SCChargeOrder) GetChargeId() int32 {
	if m != nil && m.ChargeId != nil {
		return *m.ChargeId
	}
	return 0
}

func (m *SCChargeOrder) GetChargeType() int32 {
	if m != nil && m.ChargeType != nil {
		return *m.ChargeType
	}
	return 0
}

func (m *SCChargeOrder) GetOrderId() string {
	if m != nil && m.OrderId != nil {
		return *m.OrderId
	}
	return ""
}

func (m *SCChargeOrder) GetNotifyUrl() string {
	if m != nil && m.NotifyUrl != nil {
		return *m.NotifyUrl
	}
	return ""
}

func (m *SCChargeOrder) GetSdkOrderId() string {
	if m != nil && m.SdkOrderId != nil {
		return *m.SdkOrderId
	}
	return ""
}

func (m *SCChargeOrder) GetPlatformUserId() string {
	if m != nil && m.PlatformUserId != nil {
		return *m.PlatformUserId
	}
	return ""
}

func (m *SCChargeOrder) GetMoney() int32 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func (m *SCChargeOrder) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *SCChargeOrder) GetPlayerName() string {
	if m != nil && m.PlayerName != nil {
		return *m.PlayerName
	}
	return ""
}

func (m *SCChargeOrder) GetServerId() int32 {
	if m != nil && m.ServerId != nil {
		return *m.ServerId
	}
	return 0
}

func (m *SCChargeOrder) GetServerName() string {
	if m != nil && m.ServerName != nil {
		return *m.ServerName
	}
	return ""
}

func (m *SCChargeOrder) GetExtension() string {
	if m != nil && m.Extension != nil {
		return *m.Extension
	}
	return ""
}

type SCFirstChargeRecordNotice struct {
	ChargeIdList     []int32 `protobuf:"varint,1,rep,name=chargeIdList" json:"chargeIdList,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCFirstChargeRecordNotice) Reset()                    { *m = SCFirstChargeRecordNotice{} }
func (m *SCFirstChargeRecordNotice) String() string            { return proto.CompactTextString(m) }
func (*SCFirstChargeRecordNotice) ProtoMessage()               {}
func (*SCFirstChargeRecordNotice) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{3} }

func (m *SCFirstChargeRecordNotice) GetChargeIdList() []int32 {
	if m != nil {
		return m.ChargeIdList
	}
	return nil
}

type SCChargeGoldNotice struct {
	ChargeGold       *int64 `protobuf:"varint,1,req,name=chargeGold" json:"chargeGold,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCChargeGoldNotice) Reset()                    { *m = SCChargeGoldNotice{} }
func (m *SCChargeGoldNotice) String() string            { return proto.CompactTextString(m) }
func (*SCChargeGoldNotice) ProtoMessage()               {}
func (*SCChargeGoldNotice) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{4} }

func (m *SCChargeGoldNotice) GetChargeGold() int64 {
	if m != nil && m.ChargeGold != nil {
		return *m.ChargeGold
	}
	return 0
}

type SCTodayChargeGold struct {
	ChargeGold       *int64 `protobuf:"varint,1,req,name=chargeGold" json:"chargeGold,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SCTodayChargeGold) Reset()                    { *m = SCTodayChargeGold{} }
func (m *SCTodayChargeGold) String() string            { return proto.CompactTextString(m) }
func (*SCTodayChargeGold) ProtoMessage()               {}
func (*SCTodayChargeGold) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{5} }

func (m *SCTodayChargeGold) GetChargeGold() int64 {
	if m != nil && m.ChargeGold != nil {
		return *m.ChargeGold
	}
	return 0
}

type CSNewFirstChargeRecord struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CSNewFirstChargeRecord) Reset()                    { *m = CSNewFirstChargeRecord{} }
func (m *CSNewFirstChargeRecord) String() string            { return proto.CompactTextString(m) }
func (*CSNewFirstChargeRecord) ProtoMessage()               {}
func (*CSNewFirstChargeRecord) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{6} }

type SCNewFirstChargeRecord struct {
	StartTime        *int64  `protobuf:"varint,1,req,name=startTime" json:"startTime,omitempty"`
	Duration         *int64  `protobuf:"varint,2,req,name=duration" json:"duration,omitempty"`
	ChargeIdList     []int32 `protobuf:"varint,3,rep,name=chargeIdList" json:"chargeIdList,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCNewFirstChargeRecord) Reset()                    { *m = SCNewFirstChargeRecord{} }
func (m *SCNewFirstChargeRecord) String() string            { return proto.CompactTextString(m) }
func (*SCNewFirstChargeRecord) ProtoMessage()               {}
func (*SCNewFirstChargeRecord) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{7} }

func (m *SCNewFirstChargeRecord) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *SCNewFirstChargeRecord) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *SCNewFirstChargeRecord) GetChargeIdList() []int32 {
	if m != nil {
		return m.ChargeIdList
	}
	return nil
}

type SCNewFirstChargeRecordNotice struct {
	ChargeIdList     []int32 `protobuf:"varint,1,rep,name=chargeIdList" json:"chargeIdList,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SCNewFirstChargeRecordNotice) Reset()                    { *m = SCNewFirstChargeRecordNotice{} }
func (m *SCNewFirstChargeRecordNotice) String() string            { return proto.CompactTextString(m) }
func (*SCNewFirstChargeRecordNotice) ProtoMessage()               {}
func (*SCNewFirstChargeRecordNotice) Descriptor() ([]byte, []int) { return fileDescriptor14, []int{8} }

func (m *SCNewFirstChargeRecordNotice) GetChargeIdList() []int32 {
	if m != nil {
		return m.ChargeIdList
	}
	return nil
}

func init() {
	proto.RegisterType((*CSCharge)(nil), "ui.CSCharge")
	proto.RegisterType((*SCCharge)(nil), "ui.SCCharge")
	proto.RegisterType((*SCChargeOrder)(nil), "ui.SCChargeOrder")
	proto.RegisterType((*SCFirstChargeRecordNotice)(nil), "ui.SCFirstChargeRecordNotice")
	proto.RegisterType((*SCChargeGoldNotice)(nil), "ui.SCChargeGoldNotice")
	proto.RegisterType((*SCTodayChargeGold)(nil), "ui.SCTodayChargeGold")
	proto.RegisterType((*CSNewFirstChargeRecord)(nil), "ui.CSNewFirstChargeRecord")
	proto.RegisterType((*SCNewFirstChargeRecord)(nil), "ui.SCNewFirstChargeRecord")
	proto.RegisterType((*SCNewFirstChargeRecordNotice)(nil), "ui.SCNewFirstChargeRecordNotice")
}

func init() { proto.RegisterFile("charge.proto", fileDescriptor14) }

var fileDescriptor14 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xe1, 0x4a, 0xeb, 0x30,
	0x18, 0x86, 0x59, 0x7b, 0x7a, 0xd6, 0x7d, 0x67, 0x3b, 0x67, 0x0b, 0x87, 0x11, 0x61, 0x3f, 0x4a,
	0xff, 0xd8, 0x5f, 0x82, 0xe0, 0x1d, 0x04, 0x94, 0x81, 0x6c, 0x60, 0xb6, 0x0b, 0x08, 0x4b, 0xa6,
	0xc1, 0xb6, 0x29, 0x49, 0xa6, 0xf6, 0xba, 0xbd, 0x01, 0xc9, 0x57, 0xab, 0xe0, 0xa6, 0x7f, 0x9f,
	0xf7, 0xc9, 0x47, 0xde, 0x17, 0xc6, 0xbb, 0x07, 0x61, 0xef, 0xd5, 0x45, 0x63, 0x8d, 0x37, 0x24,
	0x3a, 0xe8, 0x7c, 0x01, 0x29, 0xe3, 0x0c, 0x29, 0x99, 0x42, 0xda, 0xe5, 0x4b, 0x49, 0x07, 0x59,
	0x54, 0x24, 0x21, 0xe5, 0xec, 0xdb, 0xf4, 0x75, 0x00, 0x93, 0x3e, 0x5e, 0x5b, 0xa9, 0xec, 0xb1,
	0x43, 0x08, 0x40, 0x47, 0x36, 0x6d, 0xa3, 0x68, 0x84, 0xec, 0x1f, 0x0c, 0x4d, 0xd0, 0x97, 0x92,
	0xc6, 0x59, 0x54, 0x8c, 0xc8, 0x0c, 0x46, 0xb5, 0xf1, 0x7a, 0xdf, 0x6e, 0x6d, 0x49, 0x7f, 0x21,
	0x22, 0x00, 0x4e, 0x3e, 0xae, 0xdf, 0xb5, 0x04, 0xd9, 0x1c, 0xfe, 0x36, 0xa5, 0xf0, 0x7b, 0x63,
	0xab, 0xad, 0x43, 0xfe, 0x1b, 0xf9, 0x04, 0x92, 0xca, 0xd4, 0xaa, 0xa5, 0x43, 0x3c, 0x3f, 0x85,
	0xb4, 0x29, 0x45, 0x8b, 0x42, 0x9a, 0x45, 0x45, 0x1c, 0x8e, 0x75, 0x64, 0x25, 0x2a, 0x45, 0x47,
	0xf8, 0x68, 0x0a, 0xa9, 0x53, 0xf6, 0x09, 0x2d, 0xe8, 0xbf, 0xda, 0x11, 0xb4, 0xfe, 0xf4, 0x3f,
	0x53, 0x2f, 0x5e, 0xd5, 0x4e, 0x9b, 0x9a, 0x8e, 0x03, 0xca, 0x2f, 0xe1, 0x8c, 0xb3, 0x6b, 0x6d,
	0x9d, 0xef, 0x9a, 0xdf, 0xa9, 0x9d, 0xb1, 0x72, 0x65, 0xbc, 0xde, 0x29, 0xf2, 0xbf, 0x9f, 0x78,
	0x29, 0x6f, 0xb5, 0xf3, 0x74, 0x90, 0xc5, 0x45, 0x92, 0x17, 0x40, 0xfa, 0x9d, 0x6e, 0x4c, 0xd9,
	0xbb, 0x1f, 0xd3, 0x04, 0x86, 0x73, 0xc5, 0xf9, 0x39, 0xcc, 0x38, 0xdb, 0x18, 0x29, 0xda, 0x4f,
	0xfd, 0xa4, 0x48, 0x61, 0xce, 0xf8, 0x4a, 0x3d, 0x1f, 0x7d, 0x24, 0xe7, 0x30, 0xe7, 0xec, 0x54,
	0x12, 0xca, 0x38, 0x2f, 0xac, 0xdf, 0xe8, 0x4a, 0x75, 0x67, 0xc2, 0x0a, 0xf2, 0x60, 0x85, 0x0f,
	0xf5, 0x22, 0x24, 0x5f, 0x1b, 0xc4, 0xd8, 0xe0, 0x0a, 0x16, 0xa7, 0x8f, 0xfe, 0xd4, 0xfb, 0x2d,
	0x00, 0x00, 0xff, 0xff, 0x40, 0x88, 0xd7, 0xf2, 0x6f, 0x02, 0x00, 0x00,
}