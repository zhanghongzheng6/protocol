// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package msg

import (
	context "context"
	sdkws "github.com/openimsdk/protocol/sdkws"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// 获取最小最大seq（包括用户的，以及指定群组的）
	GetMaxSeq(ctx context.Context, in *sdkws.GetMaxSeqReq, opts ...grpc.CallOption) (*sdkws.GetMaxSeqResp, error)
	// 获取会话列表的最大seq
	GetMaxSeqs(ctx context.Context, in *GetMaxSeqsReq, opts ...grpc.CallOption) (*SeqsInfoResp, error)
	// 获取会话列表已读的最大seq
	GetHasReadSeqs(ctx context.Context, in *GetHasReadSeqsReq, opts ...grpc.CallOption) (*SeqsInfoResp, error)
	// 获取最新消息
	GetMsgByConversationIDs(ctx context.Context, in *GetMsgByConversationIDsReq, opts ...grpc.CallOption) (*GetMsgByConversationIDsResp, error)
	GetConversationMaxSeq(ctx context.Context, in *GetConversationMaxSeqReq, opts ...grpc.CallOption) (*GetConversationMaxSeqResp, error)
	// 拉取历史消息（包括用户的，以及指定群组的）
	PullMessageBySeqs(ctx context.Context, in *sdkws.PullMessageBySeqsReq, opts ...grpc.CallOption) (*sdkws.PullMessageBySeqsResp, error)
	SearchMessage(ctx context.Context, in *SearchMessageReq, opts ...grpc.CallOption) (*SearchMessageResp, error)
	// 发送消息
	SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error)
	// 全量清空指定会话消息 重置min seq 比最大seq大1
	ClearConversationsMsg(ctx context.Context, in *ClearConversationsMsgReq, opts ...grpc.CallOption) (*ClearConversationsMsgResp, error)
	// 删除用户全部消息 重置min seq 比最大seq大1
	UserClearAllMsg(ctx context.Context, in *UserClearAllMsgReq, opts ...grpc.CallOption) (*UserClearAllMsgResp, error)
	// 用户标记删除部分消息by Seq
	DeleteMsgs(ctx context.Context, in *DeleteMsgsReq, opts ...grpc.CallOption) (*DeleteMsgsResp, error)
	// seq物理删除消息
	DeleteMsgPhysicalBySeq(ctx context.Context, in *DeleteMsgPhysicalBySeqReq, opts ...grpc.CallOption) (*DeleteMsgPhysicalBySeqResp, error)
	// 物理删除消息by 时间
	DeleteMsgPhysical(ctx context.Context, in *DeleteMsgPhysicalReq, opts ...grpc.CallOption) (*DeleteMsgPhysicalResp, error)
	// 设置消息是否发送成功-针对api发送的消息
	SetSendMsgStatus(ctx context.Context, in *SetSendMsgStatusReq, opts ...grpc.CallOption) (*SetSendMsgStatusResp, error)
	// 获取消息发送状态
	GetSendMsgStatus(ctx context.Context, in *GetSendMsgStatusReq, opts ...grpc.CallOption) (*GetSendMsgStatusResp, error)
	RevokeMsg(ctx context.Context, in *RevokeMsgReq, opts ...grpc.CallOption) (*RevokeMsgResp, error)
	// mark as read
	MarkMsgsAsRead(ctx context.Context, in *MarkMsgsAsReadReq, opts ...grpc.CallOption) (*MarkMsgsAsReadResp, error)
	MarkConversationAsRead(ctx context.Context, in *MarkConversationAsReadReq, opts ...grpc.CallOption) (*MarkConversationAsReadResp, error)
	SetConversationHasReadSeq(ctx context.Context, in *SetConversationHasReadSeqReq, opts ...grpc.CallOption) (*SetConversationHasReadSeqResp, error)
	GetConversationsHasReadAndMaxSeq(ctx context.Context, in *GetConversationsHasReadAndMaxSeqReq, opts ...grpc.CallOption) (*GetConversationsHasReadAndMaxSeqResp, error)
	GetActiveUser(ctx context.Context, in *GetActiveUserReq, opts ...grpc.CallOption) (*GetActiveUserResp, error)
	GetActiveGroup(ctx context.Context, in *GetActiveGroupReq, opts ...grpc.CallOption) (*GetActiveGroupResp, error)
	GetServerTime(ctx context.Context, in *GetServerTimeReq, opts ...grpc.CallOption) (*GetServerTimeResp, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GetMaxSeq(ctx context.Context, in *sdkws.GetMaxSeqReq, opts ...grpc.CallOption) (*sdkws.GetMaxSeqResp, error) {
	out := new(sdkws.GetMaxSeqResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetMaxSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetMaxSeqs(ctx context.Context, in *GetMaxSeqsReq, opts ...grpc.CallOption) (*SeqsInfoResp, error) {
	out := new(SeqsInfoResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetMaxSeqs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetHasReadSeqs(ctx context.Context, in *GetHasReadSeqsReq, opts ...grpc.CallOption) (*SeqsInfoResp, error) {
	out := new(SeqsInfoResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetHasReadSeqs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetMsgByConversationIDs(ctx context.Context, in *GetMsgByConversationIDsReq, opts ...grpc.CallOption) (*GetMsgByConversationIDsResp, error) {
	out := new(GetMsgByConversationIDsResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetMsgByConversationIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetConversationMaxSeq(ctx context.Context, in *GetConversationMaxSeqReq, opts ...grpc.CallOption) (*GetConversationMaxSeqResp, error) {
	out := new(GetConversationMaxSeqResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetConversationMaxSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PullMessageBySeqs(ctx context.Context, in *sdkws.PullMessageBySeqsReq, opts ...grpc.CallOption) (*sdkws.PullMessageBySeqsResp, error) {
	out := new(sdkws.PullMessageBySeqsResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/PullMessageBySeqs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SearchMessage(ctx context.Context, in *SearchMessageReq, opts ...grpc.CallOption) (*SearchMessageResp, error) {
	out := new(SearchMessageResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/SearchMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendMsg(ctx context.Context, in *SendMsgReq, opts ...grpc.CallOption) (*SendMsgResp, error) {
	out := new(SendMsgResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/SendMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClearConversationsMsg(ctx context.Context, in *ClearConversationsMsgReq, opts ...grpc.CallOption) (*ClearConversationsMsgResp, error) {
	out := new(ClearConversationsMsgResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/ClearConversationsMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UserClearAllMsg(ctx context.Context, in *UserClearAllMsgReq, opts ...grpc.CallOption) (*UserClearAllMsgResp, error) {
	out := new(UserClearAllMsgResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/UserClearAllMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteMsgs(ctx context.Context, in *DeleteMsgsReq, opts ...grpc.CallOption) (*DeleteMsgsResp, error) {
	out := new(DeleteMsgsResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/DeleteMsgs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteMsgPhysicalBySeq(ctx context.Context, in *DeleteMsgPhysicalBySeqReq, opts ...grpc.CallOption) (*DeleteMsgPhysicalBySeqResp, error) {
	out := new(DeleteMsgPhysicalBySeqResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/DeleteMsgPhysicalBySeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteMsgPhysical(ctx context.Context, in *DeleteMsgPhysicalReq, opts ...grpc.CallOption) (*DeleteMsgPhysicalResp, error) {
	out := new(DeleteMsgPhysicalResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/DeleteMsgPhysical", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetSendMsgStatus(ctx context.Context, in *SetSendMsgStatusReq, opts ...grpc.CallOption) (*SetSendMsgStatusResp, error) {
	out := new(SetSendMsgStatusResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/SetSendMsgStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetSendMsgStatus(ctx context.Context, in *GetSendMsgStatusReq, opts ...grpc.CallOption) (*GetSendMsgStatusResp, error) {
	out := new(GetSendMsgStatusResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetSendMsgStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevokeMsg(ctx context.Context, in *RevokeMsgReq, opts ...grpc.CallOption) (*RevokeMsgResp, error) {
	out := new(RevokeMsgResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/RevokeMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MarkMsgsAsRead(ctx context.Context, in *MarkMsgsAsReadReq, opts ...grpc.CallOption) (*MarkMsgsAsReadResp, error) {
	out := new(MarkMsgsAsReadResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/MarkMsgsAsRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MarkConversationAsRead(ctx context.Context, in *MarkConversationAsReadReq, opts ...grpc.CallOption) (*MarkConversationAsReadResp, error) {
	out := new(MarkConversationAsReadResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/MarkConversationAsRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetConversationHasReadSeq(ctx context.Context, in *SetConversationHasReadSeqReq, opts ...grpc.CallOption) (*SetConversationHasReadSeqResp, error) {
	out := new(SetConversationHasReadSeqResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/SetConversationHasReadSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetConversationsHasReadAndMaxSeq(ctx context.Context, in *GetConversationsHasReadAndMaxSeqReq, opts ...grpc.CallOption) (*GetConversationsHasReadAndMaxSeqResp, error) {
	out := new(GetConversationsHasReadAndMaxSeqResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetConversationsHasReadAndMaxSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetActiveUser(ctx context.Context, in *GetActiveUserReq, opts ...grpc.CallOption) (*GetActiveUserResp, error) {
	out := new(GetActiveUserResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetActiveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetActiveGroup(ctx context.Context, in *GetActiveGroupReq, opts ...grpc.CallOption) (*GetActiveGroupResp, error) {
	out := new(GetActiveGroupResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetActiveGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetServerTime(ctx context.Context, in *GetServerTimeReq, opts ...grpc.CallOption) (*GetServerTimeResp, error) {
	out := new(GetServerTimeResp)
	err := c.cc.Invoke(ctx, "/openim.msg.msg/GetServerTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// 获取最小最大seq（包括用户的，以及指定群组的）
	GetMaxSeq(context.Context, *sdkws.GetMaxSeqReq) (*sdkws.GetMaxSeqResp, error)
	// 获取会话列表的最大seq
	GetMaxSeqs(context.Context, *GetMaxSeqsReq) (*SeqsInfoResp, error)
	// 获取会话列表已读的最大seq
	GetHasReadSeqs(context.Context, *GetHasReadSeqsReq) (*SeqsInfoResp, error)
	// 获取最新消息
	GetMsgByConversationIDs(context.Context, *GetMsgByConversationIDsReq) (*GetMsgByConversationIDsResp, error)
	GetConversationMaxSeq(context.Context, *GetConversationMaxSeqReq) (*GetConversationMaxSeqResp, error)
	// 拉取历史消息（包括用户的，以及指定群组的）
	PullMessageBySeqs(context.Context, *sdkws.PullMessageBySeqsReq) (*sdkws.PullMessageBySeqsResp, error)
	SearchMessage(context.Context, *SearchMessageReq) (*SearchMessageResp, error)
	// 发送消息
	SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error)
	// 全量清空指定会话消息 重置min seq 比最大seq大1
	ClearConversationsMsg(context.Context, *ClearConversationsMsgReq) (*ClearConversationsMsgResp, error)
	// 删除用户全部消息 重置min seq 比最大seq大1
	UserClearAllMsg(context.Context, *UserClearAllMsgReq) (*UserClearAllMsgResp, error)
	// 用户标记删除部分消息by Seq
	DeleteMsgs(context.Context, *DeleteMsgsReq) (*DeleteMsgsResp, error)
	// seq物理删除消息
	DeleteMsgPhysicalBySeq(context.Context, *DeleteMsgPhysicalBySeqReq) (*DeleteMsgPhysicalBySeqResp, error)
	// 物理删除消息by 时间
	DeleteMsgPhysical(context.Context, *DeleteMsgPhysicalReq) (*DeleteMsgPhysicalResp, error)
	// 设置消息是否发送成功-针对api发送的消息
	SetSendMsgStatus(context.Context, *SetSendMsgStatusReq) (*SetSendMsgStatusResp, error)
	// 获取消息发送状态
	GetSendMsgStatus(context.Context, *GetSendMsgStatusReq) (*GetSendMsgStatusResp, error)
	RevokeMsg(context.Context, *RevokeMsgReq) (*RevokeMsgResp, error)
	// mark as read
	MarkMsgsAsRead(context.Context, *MarkMsgsAsReadReq) (*MarkMsgsAsReadResp, error)
	MarkConversationAsRead(context.Context, *MarkConversationAsReadReq) (*MarkConversationAsReadResp, error)
	SetConversationHasReadSeq(context.Context, *SetConversationHasReadSeqReq) (*SetConversationHasReadSeqResp, error)
	GetConversationsHasReadAndMaxSeq(context.Context, *GetConversationsHasReadAndMaxSeqReq) (*GetConversationsHasReadAndMaxSeqResp, error)
	GetActiveUser(context.Context, *GetActiveUserReq) (*GetActiveUserResp, error)
	GetActiveGroup(context.Context, *GetActiveGroupReq) (*GetActiveGroupResp, error)
	GetServerTime(context.Context, *GetServerTimeReq) (*GetServerTimeResp, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) GetMaxSeq(context.Context, *sdkws.GetMaxSeqReq) (*sdkws.GetMaxSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaxSeq not implemented")
}
func (UnimplementedMsgServer) GetMaxSeqs(context.Context, *GetMaxSeqsReq) (*SeqsInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaxSeqs not implemented")
}
func (UnimplementedMsgServer) GetHasReadSeqs(context.Context, *GetHasReadSeqsReq) (*SeqsInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHasReadSeqs not implemented")
}
func (UnimplementedMsgServer) GetMsgByConversationIDs(context.Context, *GetMsgByConversationIDsReq) (*GetMsgByConversationIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMsgByConversationIDs not implemented")
}
func (UnimplementedMsgServer) GetConversationMaxSeq(context.Context, *GetConversationMaxSeqReq) (*GetConversationMaxSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationMaxSeq not implemented")
}
func (UnimplementedMsgServer) PullMessageBySeqs(context.Context, *sdkws.PullMessageBySeqsReq) (*sdkws.PullMessageBySeqsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PullMessageBySeqs not implemented")
}
func (UnimplementedMsgServer) SearchMessage(context.Context, *SearchMessageReq) (*SearchMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMessage not implemented")
}
func (UnimplementedMsgServer) SendMsg(context.Context, *SendMsgReq) (*SendMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsg not implemented")
}
func (UnimplementedMsgServer) ClearConversationsMsg(context.Context, *ClearConversationsMsgReq) (*ClearConversationsMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearConversationsMsg not implemented")
}
func (UnimplementedMsgServer) UserClearAllMsg(context.Context, *UserClearAllMsgReq) (*UserClearAllMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserClearAllMsg not implemented")
}
func (UnimplementedMsgServer) DeleteMsgs(context.Context, *DeleteMsgsReq) (*DeleteMsgsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMsgs not implemented")
}
func (UnimplementedMsgServer) DeleteMsgPhysicalBySeq(context.Context, *DeleteMsgPhysicalBySeqReq) (*DeleteMsgPhysicalBySeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMsgPhysicalBySeq not implemented")
}
func (UnimplementedMsgServer) DeleteMsgPhysical(context.Context, *DeleteMsgPhysicalReq) (*DeleteMsgPhysicalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMsgPhysical not implemented")
}
func (UnimplementedMsgServer) SetSendMsgStatus(context.Context, *SetSendMsgStatusReq) (*SetSendMsgStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSendMsgStatus not implemented")
}
func (UnimplementedMsgServer) GetSendMsgStatus(context.Context, *GetSendMsgStatusReq) (*GetSendMsgStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSendMsgStatus not implemented")
}
func (UnimplementedMsgServer) RevokeMsg(context.Context, *RevokeMsgReq) (*RevokeMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeMsg not implemented")
}
func (UnimplementedMsgServer) MarkMsgsAsRead(context.Context, *MarkMsgsAsReadReq) (*MarkMsgsAsReadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkMsgsAsRead not implemented")
}
func (UnimplementedMsgServer) MarkConversationAsRead(context.Context, *MarkConversationAsReadReq) (*MarkConversationAsReadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MarkConversationAsRead not implemented")
}
func (UnimplementedMsgServer) SetConversationHasReadSeq(context.Context, *SetConversationHasReadSeqReq) (*SetConversationHasReadSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConversationHasReadSeq not implemented")
}
func (UnimplementedMsgServer) GetConversationsHasReadAndMaxSeq(context.Context, *GetConversationsHasReadAndMaxSeqReq) (*GetConversationsHasReadAndMaxSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationsHasReadAndMaxSeq not implemented")
}
func (UnimplementedMsgServer) GetActiveUser(context.Context, *GetActiveUserReq) (*GetActiveUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveUser not implemented")
}
func (UnimplementedMsgServer) GetActiveGroup(context.Context, *GetActiveGroupReq) (*GetActiveGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActiveGroup not implemented")
}
func (UnimplementedMsgServer) GetServerTime(context.Context, *GetServerTimeReq) (*GetServerTimeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerTime not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_GetMaxSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sdkws.GetMaxSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetMaxSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetMaxSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetMaxSeq(ctx, req.(*sdkws.GetMaxSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetMaxSeqs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMaxSeqsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetMaxSeqs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetMaxSeqs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetMaxSeqs(ctx, req.(*GetMaxSeqsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetHasReadSeqs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHasReadSeqsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetHasReadSeqs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetHasReadSeqs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetHasReadSeqs(ctx, req.(*GetHasReadSeqsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetMsgByConversationIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMsgByConversationIDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetMsgByConversationIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetMsgByConversationIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetMsgByConversationIDs(ctx, req.(*GetMsgByConversationIDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetConversationMaxSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationMaxSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetConversationMaxSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetConversationMaxSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetConversationMaxSeq(ctx, req.(*GetConversationMaxSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PullMessageBySeqs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sdkws.PullMessageBySeqsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PullMessageBySeqs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/PullMessageBySeqs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PullMessageBySeqs(ctx, req.(*sdkws.PullMessageBySeqsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SearchMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SearchMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/SearchMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SearchMessage(ctx, req.(*SearchMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/SendMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendMsg(ctx, req.(*SendMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClearConversationsMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearConversationsMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClearConversationsMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/ClearConversationsMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClearConversationsMsg(ctx, req.(*ClearConversationsMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UserClearAllMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserClearAllMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UserClearAllMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/UserClearAllMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UserClearAllMsg(ctx, req.(*UserClearAllMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteMsgs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMsgsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteMsgs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/DeleteMsgs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteMsgs(ctx, req.(*DeleteMsgsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteMsgPhysicalBySeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMsgPhysicalBySeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteMsgPhysicalBySeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/DeleteMsgPhysicalBySeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteMsgPhysicalBySeq(ctx, req.(*DeleteMsgPhysicalBySeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteMsgPhysical_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMsgPhysicalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteMsgPhysical(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/DeleteMsgPhysical",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteMsgPhysical(ctx, req.(*DeleteMsgPhysicalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetSendMsgStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSendMsgStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetSendMsgStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/SetSendMsgStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetSendMsgStatus(ctx, req.(*SetSendMsgStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetSendMsgStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSendMsgStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetSendMsgStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetSendMsgStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetSendMsgStatus(ctx, req.(*GetSendMsgStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevokeMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevokeMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/RevokeMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevokeMsg(ctx, req.(*RevokeMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MarkMsgsAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkMsgsAsReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MarkMsgsAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/MarkMsgsAsRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MarkMsgsAsRead(ctx, req.(*MarkMsgsAsReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MarkConversationAsRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MarkConversationAsReadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MarkConversationAsRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/MarkConversationAsRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MarkConversationAsRead(ctx, req.(*MarkConversationAsReadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetConversationHasReadSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConversationHasReadSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetConversationHasReadSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/SetConversationHasReadSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetConversationHasReadSeq(ctx, req.(*SetConversationHasReadSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetConversationsHasReadAndMaxSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsHasReadAndMaxSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetConversationsHasReadAndMaxSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetConversationsHasReadAndMaxSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetConversationsHasReadAndMaxSeq(ctx, req.(*GetConversationsHasReadAndMaxSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetActiveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetActiveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetActiveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetActiveUser(ctx, req.(*GetActiveUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetActiveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActiveGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetActiveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetActiveGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetActiveGroup(ctx, req.(*GetActiveGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetServerTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServerTimeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetServerTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.msg.msg/GetServerTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetServerTime(ctx, req.(*GetServerTimeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openim.msg.msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMaxSeq",
			Handler:    _Msg_GetMaxSeq_Handler,
		},
		{
			MethodName: "GetMaxSeqs",
			Handler:    _Msg_GetMaxSeqs_Handler,
		},
		{
			MethodName: "GetHasReadSeqs",
			Handler:    _Msg_GetHasReadSeqs_Handler,
		},
		{
			MethodName: "GetMsgByConversationIDs",
			Handler:    _Msg_GetMsgByConversationIDs_Handler,
		},
		{
			MethodName: "GetConversationMaxSeq",
			Handler:    _Msg_GetConversationMaxSeq_Handler,
		},
		{
			MethodName: "PullMessageBySeqs",
			Handler:    _Msg_PullMessageBySeqs_Handler,
		},
		{
			MethodName: "SearchMessage",
			Handler:    _Msg_SearchMessage_Handler,
		},
		{
			MethodName: "SendMsg",
			Handler:    _Msg_SendMsg_Handler,
		},
		{
			MethodName: "ClearConversationsMsg",
			Handler:    _Msg_ClearConversationsMsg_Handler,
		},
		{
			MethodName: "UserClearAllMsg",
			Handler:    _Msg_UserClearAllMsg_Handler,
		},
		{
			MethodName: "DeleteMsgs",
			Handler:    _Msg_DeleteMsgs_Handler,
		},
		{
			MethodName: "DeleteMsgPhysicalBySeq",
			Handler:    _Msg_DeleteMsgPhysicalBySeq_Handler,
		},
		{
			MethodName: "DeleteMsgPhysical",
			Handler:    _Msg_DeleteMsgPhysical_Handler,
		},
		{
			MethodName: "SetSendMsgStatus",
			Handler:    _Msg_SetSendMsgStatus_Handler,
		},
		{
			MethodName: "GetSendMsgStatus",
			Handler:    _Msg_GetSendMsgStatus_Handler,
		},
		{
			MethodName: "RevokeMsg",
			Handler:    _Msg_RevokeMsg_Handler,
		},
		{
			MethodName: "MarkMsgsAsRead",
			Handler:    _Msg_MarkMsgsAsRead_Handler,
		},
		{
			MethodName: "MarkConversationAsRead",
			Handler:    _Msg_MarkConversationAsRead_Handler,
		},
		{
			MethodName: "SetConversationHasReadSeq",
			Handler:    _Msg_SetConversationHasReadSeq_Handler,
		},
		{
			MethodName: "GetConversationsHasReadAndMaxSeq",
			Handler:    _Msg_GetConversationsHasReadAndMaxSeq_Handler,
		},
		{
			MethodName: "GetActiveUser",
			Handler:    _Msg_GetActiveUser_Handler,
		},
		{
			MethodName: "GetActiveGroup",
			Handler:    _Msg_GetActiveGroup_Handler,
		},
		{
			MethodName: "GetServerTime",
			Handler:    _Msg_GetServerTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "msg/msg.proto",
}
