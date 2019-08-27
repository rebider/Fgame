// Code generated by protoc-gen-go. DO NOT EDIT.
// source: player_manage.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PlayerServerInfo struct {
	UserId     int64  `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	ServerId   int32  `protobuf:"varint,2,opt,name=serverId" json:"serverId,omitempty"`
	PlayerId   int64  `protobuf:"varint,3,opt,name=playerId" json:"playerId,omitempty"`
	PlayerName string `protobuf:"bytes,4,opt,name=playerName" json:"playerName,omitempty"`
	Role       int32  `protobuf:"varint,5,opt,name=role" json:"role,omitempty"`
	Sex        int32  `protobuf:"varint,6,opt,name=sex" json:"sex,omitempty"`
	Level      int32  `protobuf:"varint,7,opt,name=level" json:"level,omitempty"`
	ZhuanShu   int32  `protobuf:"varint,8,opt,name=zhuanShu" json:"zhuanShu,omitempty"`
}

func (m *PlayerServerInfo) Reset()                    { *m = PlayerServerInfo{} }
func (m *PlayerServerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerServerInfo) ProtoMessage()               {}
func (*PlayerServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *PlayerServerInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PlayerServerInfo) GetServerId() int32 {
	if m != nil {
		return m.ServerId
	}
	return 0
}

func (m *PlayerServerInfo) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *PlayerServerInfo) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *PlayerServerInfo) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *PlayerServerInfo) GetSex() int32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

func (m *PlayerServerInfo) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *PlayerServerInfo) GetZhuanShu() int32 {
	if m != nil {
		return m.ZhuanShu
	}
	return 0
}

// 获取用户的角色列表
type PlayerListRequest struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *PlayerListRequest) Reset()                    { *m = PlayerListRequest{} }
func (m *PlayerListRequest) String() string            { return proto.CompactTextString(m) }
func (*PlayerListRequest) ProtoMessage()               {}
func (*PlayerListRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *PlayerListRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

// 获取用户的角色列表回复
type PlayerListResponse struct {
	PlayerList []*PlayerServerInfo `protobuf:"bytes,1,rep,name=playerList" json:"playerList,omitempty"`
}

func (m *PlayerListResponse) Reset()                    { *m = PlayerListResponse{} }
func (m *PlayerListResponse) String() string            { return proto.CompactTextString(m) }
func (*PlayerListResponse) ProtoMessage()               {}
func (*PlayerListResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *PlayerListResponse) GetPlayerList() []*PlayerServerInfo {
	if m != nil {
		return m.PlayerList
	}
	return nil
}

// 同步玩家的角色信息
type PlayerInfoSyncRequest struct {
	PlayerInfo *PlayerServerInfo `protobuf:"bytes,1,opt,name=playerInfo" json:"playerInfo,omitempty"`
}

func (m *PlayerInfoSyncRequest) Reset()                    { *m = PlayerInfoSyncRequest{} }
func (m *PlayerInfoSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfoSyncRequest) ProtoMessage()               {}
func (*PlayerInfoSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *PlayerInfoSyncRequest) GetPlayerInfo() *PlayerServerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

// 同步玩家的角色信息回复
type PlayerInfoSyncResponse struct {
}

func (m *PlayerInfoSyncResponse) Reset()                    { *m = PlayerInfoSyncResponse{} }
func (m *PlayerInfoSyncResponse) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfoSyncResponse) ProtoMessage()               {}
func (*PlayerInfoSyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func init() {
	proto.RegisterType((*PlayerServerInfo)(nil), "pb.PlayerServerInfo")
	proto.RegisterType((*PlayerListRequest)(nil), "pb.PlayerListRequest")
	proto.RegisterType((*PlayerListResponse)(nil), "pb.PlayerListResponse")
	proto.RegisterType((*PlayerInfoSyncRequest)(nil), "pb.PlayerInfoSyncRequest")
	proto.RegisterType((*PlayerInfoSyncResponse)(nil), "pb.PlayerInfoSyncResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PlayerManage service

type PlayerManageClient interface {
	// 获取用户的角色列表
	GetPlayerServerList(ctx context.Context, in *PlayerListRequest, opts ...grpc.CallOption) (*PlayerListResponse, error)
	// 同步用户的玩家角色
	SyncPlayerServerInfo(ctx context.Context, in *PlayerInfoSyncRequest, opts ...grpc.CallOption) (*PlayerInfoSyncResponse, error)
}

type playerManageClient struct {
	cc *grpc.ClientConn
}

func NewPlayerManageClient(cc *grpc.ClientConn) PlayerManageClient {
	return &playerManageClient{cc}
}

func (c *playerManageClient) GetPlayerServerList(ctx context.Context, in *PlayerListRequest, opts ...grpc.CallOption) (*PlayerListResponse, error) {
	out := new(PlayerListResponse)
	err := grpc.Invoke(ctx, "/pb.PlayerManage/GetPlayerServerList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playerManageClient) SyncPlayerServerInfo(ctx context.Context, in *PlayerInfoSyncRequest, opts ...grpc.CallOption) (*PlayerInfoSyncResponse, error) {
	out := new(PlayerInfoSyncResponse)
	err := grpc.Invoke(ctx, "/pb.PlayerManage/SyncPlayerServerInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PlayerManage service

type PlayerManageServer interface {
	// 获取用户的角色列表
	GetPlayerServerList(context.Context, *PlayerListRequest) (*PlayerListResponse, error)
	// 同步用户的玩家角色
	SyncPlayerServerInfo(context.Context, *PlayerInfoSyncRequest) (*PlayerInfoSyncResponse, error)
}

func RegisterPlayerManageServer(s *grpc.Server, srv PlayerManageServer) {
	s.RegisterService(&_PlayerManage_serviceDesc, srv)
}

func _PlayerManage_GetPlayerServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerManageServer).GetPlayerServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlayerManage/GetPlayerServerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerManageServer).GetPlayerServerList(ctx, req.(*PlayerListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayerManage_SyncPlayerServerInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayerInfoSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerManageServer).SyncPlayerServerInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.PlayerManage/SyncPlayerServerInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerManageServer).SyncPlayerServerInfo(ctx, req.(*PlayerInfoSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlayerManage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PlayerManage",
	HandlerType: (*PlayerManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayerServerList",
			Handler:    _PlayerManage_GetPlayerServerList_Handler,
		},
		{
			MethodName: "SyncPlayerServerInfo",
			Handler:    _PlayerManage_SyncPlayerServerInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "player_manage.proto",
}

func init() { proto.RegisterFile("player_manage.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4b, 0x4e, 0xc3, 0x30,
	0x10, 0xc5, 0x4d, 0x5b, 0xca, 0xc0, 0xa2, 0x4c, 0x3f, 0x32, 0x59, 0xa0, 0xc8, 0xab, 0x48, 0x48,
	0x5d, 0x14, 0xce, 0x00, 0x2a, 0xa2, 0x80, 0xd2, 0x03, 0xa0, 0x94, 0x0c, 0x14, 0x29, 0x4d, 0x82,
	0x9d, 0x54, 0x94, 0x0b, 0x71, 0x27, 0x4e, 0x83, 0x6c, 0x93, 0x34, 0x0a, 0x20, 0x76, 0x33, 0xef,
	0x8d, 0xdf, 0x7b, 0x63, 0x1b, 0x06, 0x59, 0x1c, 0x6e, 0x49, 0x3e, 0xac, 0xc3, 0x24, 0x7c, 0xa6,
	0x49, 0x26, 0xd3, 0x3c, 0xc5, 0x56, 0xb6, 0x14, 0x9f, 0x0c, 0xfa, 0xf7, 0x86, 0x5b, 0x90, 0xdc,
	0x90, 0x9c, 0x25, 0x4f, 0x29, 0x8e, 0xa1, 0x5b, 0x28, 0x92, 0xb3, 0x88, 0x33, 0x8f, 0xf9, 0x4e,
	0xf0, 0xdd, 0xa1, 0x0b, 0x3d, 0x65, 0xa7, 0x22, 0xde, 0xf2, 0x98, 0xdf, 0x09, 0xaa, 0x5e, 0x73,
	0xd6, 0x63, 0x16, 0x71, 0xc7, 0x9c, 0xaa, 0x7a, 0x3c, 0x05, 0xb0, 0xf5, 0x6d, 0xb8, 0x26, 0xde,
	0xf6, 0x98, 0x7f, 0x10, 0xd4, 0x10, 0x44, 0x68, 0xcb, 0x34, 0x26, 0xde, 0x31, 0x9a, 0xa6, 0xc6,
	0x3e, 0x38, 0x8a, 0xde, 0x78, 0xd7, 0x40, 0xba, 0xc4, 0x21, 0x74, 0x62, 0xda, 0x50, 0xcc, 0xf7,
	0x0d, 0x66, 0x1b, 0xed, 0xfb, 0xbe, 0x2a, 0xc2, 0x64, 0xb1, 0x2a, 0x78, 0xcf, 0x66, 0x2a, 0x7b,
	0x71, 0x06, 0xc7, 0x76, 0xb7, 0x9b, 0x17, 0x95, 0x07, 0xf4, 0x5a, 0x90, 0xca, 0xff, 0x5a, 0x4e,
	0x5c, 0x03, 0xd6, 0x87, 0x55, 0x96, 0x26, 0x8a, 0xf0, 0xa2, 0x8c, 0xae, 0x51, 0xce, 0x3c, 0xc7,
	0x3f, 0x9c, 0x0e, 0x27, 0xd9, 0x72, 0xd2, 0xbc, 0xb4, 0xa0, 0x36, 0x27, 0xe6, 0x30, 0xb2, 0xbc,
	0x66, 0x16, 0xdb, 0xe4, 0xb1, 0x34, 0xaf, 0xe4, 0x34, 0x61, 0x02, 0xfc, 0x23, 0xa7, 0x6b, 0xc1,
	0x61, 0xdc, 0x94, 0xb3, 0xf1, 0xa6, 0x1f, 0x0c, 0x8e, 0x2c, 0x35, 0x37, 0x2f, 0x8b, 0x97, 0x30,
	0xb8, 0xa2, 0xbc, 0xae, 0xa6, 0x03, 0xe1, 0x68, 0xe7, 0x51, 0xbb, 0x0b, 0x77, 0xdc, 0x84, 0xad,
	0xac, 0xd8, 0xc3, 0x3b, 0x18, 0x6a, 0xa3, 0x1f, 0x5f, 0xe3, 0x64, 0x77, 0xa2, 0xb1, 0x9b, 0xeb,
	0xfe, 0x46, 0x95, 0x82, 0xcb, 0xae, 0xf9, 0x73, 0xe7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x44,
	0xf9, 0x4e, 0x62, 0x8a, 0x02, 0x00, 0x00,
}
