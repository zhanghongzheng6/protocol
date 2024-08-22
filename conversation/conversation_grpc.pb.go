// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package conversation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ConversationClient is the client API for Conversation service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConversationClient interface {
	GetConversation(ctx context.Context, in *GetConversationReq, opts ...grpc.CallOption) (*GetConversationResp, error)
	GetSortedConversationList(ctx context.Context, in *GetSortedConversationListReq, opts ...grpc.CallOption) (*GetSortedConversationListResp, error)
	GetAllConversations(ctx context.Context, in *GetAllConversationsReq, opts ...grpc.CallOption) (*GetAllConversationsResp, error)
	GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error)
	SetConversation(ctx context.Context, in *SetConversationReq, opts ...grpc.CallOption) (*SetConversationResp, error)
	GetRecvMsgNotNotifyUserIDs(ctx context.Context, in *GetRecvMsgNotNotifyUserIDsReq, opts ...grpc.CallOption) (*GetRecvMsgNotNotifyUserIDsResp, error)
	CreateSingleChatConversations(ctx context.Context, in *CreateSingleChatConversationsReq, opts ...grpc.CallOption) (*CreateSingleChatConversationsResp, error)
	CreateGroupChatConversations(ctx context.Context, in *CreateGroupChatConversationsReq, opts ...grpc.CallOption) (*CreateGroupChatConversationsResp, error)
	SetConversationMaxSeq(ctx context.Context, in *SetConversationMaxSeqReq, opts ...grpc.CallOption) (*SetConversationMaxSeqResp, error)
	GetConversationIDs(ctx context.Context, in *GetConversationIDsReq, opts ...grpc.CallOption) (*GetConversationIDsResp, error)
	SetConversations(ctx context.Context, in *SetConversationsReq, opts ...grpc.CallOption) (*SetConversationsResp, error)
	GetUserConversationIDsHash(ctx context.Context, in *GetUserConversationIDsHashReq, opts ...grpc.CallOption) (*GetUserConversationIDsHashResp, error)
	GetConversationsByConversationID(ctx context.Context, in *GetConversationsByConversationIDReq, opts ...grpc.CallOption) (*GetConversationsByConversationIDResp, error)
	GetConversationOfflinePushUserIDs(ctx context.Context, in *GetConversationOfflinePushUserIDsReq, opts ...grpc.CallOption) (*GetConversationOfflinePushUserIDsResp, error)
	GetConversationNotReceiveMessageUserIDs(ctx context.Context, in *GetConversationNotReceiveMessageUserIDsReq, opts ...grpc.CallOption) (*GetConversationNotReceiveMessageUserIDsResp, error)
}

type conversationClient struct {
	cc grpc.ClientConnInterface
}

func NewConversationClient(cc grpc.ClientConnInterface) ConversationClient {
	return &conversationClient{cc}
}

func (c *conversationClient) GetConversation(ctx context.Context, in *GetConversationReq, opts ...grpc.CallOption) (*GetConversationResp, error) {
	out := new(GetConversationResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetSortedConversationList(ctx context.Context, in *GetSortedConversationListReq, opts ...grpc.CallOption) (*GetSortedConversationListResp, error) {
	out := new(GetSortedConversationListResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetSortedConversationList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetAllConversations(ctx context.Context, in *GetAllConversationsReq, opts ...grpc.CallOption) (*GetAllConversationsResp, error) {
	out := new(GetAllConversationsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetAllConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetConversations(ctx context.Context, in *GetConversationsReq, opts ...grpc.CallOption) (*GetConversationsResp, error) {
	out := new(GetConversationsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) SetConversation(ctx context.Context, in *SetConversationReq, opts ...grpc.CallOption) (*SetConversationResp, error) {
	out := new(SetConversationResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/SetConversation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetRecvMsgNotNotifyUserIDs(ctx context.Context, in *GetRecvMsgNotNotifyUserIDsReq, opts ...grpc.CallOption) (*GetRecvMsgNotNotifyUserIDsResp, error) {
	out := new(GetRecvMsgNotNotifyUserIDsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetRecvMsgNotNotifyUserIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) CreateSingleChatConversations(ctx context.Context, in *CreateSingleChatConversationsReq, opts ...grpc.CallOption) (*CreateSingleChatConversationsResp, error) {
	out := new(CreateSingleChatConversationsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/CreateSingleChatConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) CreateGroupChatConversations(ctx context.Context, in *CreateGroupChatConversationsReq, opts ...grpc.CallOption) (*CreateGroupChatConversationsResp, error) {
	out := new(CreateGroupChatConversationsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/CreateGroupChatConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) SetConversationMaxSeq(ctx context.Context, in *SetConversationMaxSeqReq, opts ...grpc.CallOption) (*SetConversationMaxSeqResp, error) {
	out := new(SetConversationMaxSeqResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/SetConversationMaxSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetConversationIDs(ctx context.Context, in *GetConversationIDsReq, opts ...grpc.CallOption) (*GetConversationIDsResp, error) {
	out := new(GetConversationIDsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetConversationIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) SetConversations(ctx context.Context, in *SetConversationsReq, opts ...grpc.CallOption) (*SetConversationsResp, error) {
	out := new(SetConversationsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/SetConversations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetUserConversationIDsHash(ctx context.Context, in *GetUserConversationIDsHashReq, opts ...grpc.CallOption) (*GetUserConversationIDsHashResp, error) {
	out := new(GetUserConversationIDsHashResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetUserConversationIDsHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetConversationsByConversationID(ctx context.Context, in *GetConversationsByConversationIDReq, opts ...grpc.CallOption) (*GetConversationsByConversationIDResp, error) {
	out := new(GetConversationsByConversationIDResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetConversationsByConversationID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetConversationOfflinePushUserIDs(ctx context.Context, in *GetConversationOfflinePushUserIDsReq, opts ...grpc.CallOption) (*GetConversationOfflinePushUserIDsResp, error) {
	out := new(GetConversationOfflinePushUserIDsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetConversationOfflinePushUserIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *conversationClient) GetConversationNotReceiveMessageUserIDs(ctx context.Context, in *GetConversationNotReceiveMessageUserIDsReq, opts ...grpc.CallOption) (*GetConversationNotReceiveMessageUserIDsResp, error) {
	out := new(GetConversationNotReceiveMessageUserIDsResp)
	err := c.cc.Invoke(ctx, "/openim.conversation.conversation/GetConversationNotReceiveMessageUserIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConversationServer is the server API for Conversation service.
// All implementations must embed UnimplementedConversationServer
// for forward compatibility
type ConversationServer interface {
	GetConversation(context.Context, *GetConversationReq) (*GetConversationResp, error)
	GetSortedConversationList(context.Context, *GetSortedConversationListReq) (*GetSortedConversationListResp, error)
	GetAllConversations(context.Context, *GetAllConversationsReq) (*GetAllConversationsResp, error)
	GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error)
	SetConversation(context.Context, *SetConversationReq) (*SetConversationResp, error)
	GetRecvMsgNotNotifyUserIDs(context.Context, *GetRecvMsgNotNotifyUserIDsReq) (*GetRecvMsgNotNotifyUserIDsResp, error)
	CreateSingleChatConversations(context.Context, *CreateSingleChatConversationsReq) (*CreateSingleChatConversationsResp, error)
	CreateGroupChatConversations(context.Context, *CreateGroupChatConversationsReq) (*CreateGroupChatConversationsResp, error)
	SetConversationMaxSeq(context.Context, *SetConversationMaxSeqReq) (*SetConversationMaxSeqResp, error)
	GetConversationIDs(context.Context, *GetConversationIDsReq) (*GetConversationIDsResp, error)
	SetConversations(context.Context, *SetConversationsReq) (*SetConversationsResp, error)
	GetUserConversationIDsHash(context.Context, *GetUserConversationIDsHashReq) (*GetUserConversationIDsHashResp, error)
	GetConversationsByConversationID(context.Context, *GetConversationsByConversationIDReq) (*GetConversationsByConversationIDResp, error)
	GetConversationOfflinePushUserIDs(context.Context, *GetConversationOfflinePushUserIDsReq) (*GetConversationOfflinePushUserIDsResp, error)
	GetConversationNotReceiveMessageUserIDs(context.Context, *GetConversationNotReceiveMessageUserIDsReq) (*GetConversationNotReceiveMessageUserIDsResp, error)
}

// UnimplementedConversationServer must be embedded to have forward compatible implementations.
type UnimplementedConversationServer struct {
}

func (UnimplementedConversationServer) GetConversation(context.Context, *GetConversationReq) (*GetConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversation not implemented")
}
func (UnimplementedConversationServer) GetSortedConversationList(context.Context, *GetSortedConversationListReq) (*GetSortedConversationListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSortedConversationList not implemented")
}
func (UnimplementedConversationServer) GetAllConversations(context.Context, *GetAllConversationsReq) (*GetAllConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConversations not implemented")
}
func (UnimplementedConversationServer) GetConversations(context.Context, *GetConversationsReq) (*GetConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversations not implemented")
}
func (UnimplementedConversationServer) SetConversation(context.Context, *SetConversationReq) (*SetConversationResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConversation not implemented")
}
func (UnimplementedConversationServer) GetRecvMsgNotNotifyUserIDs(context.Context, *GetRecvMsgNotNotifyUserIDsReq) (*GetRecvMsgNotNotifyUserIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecvMsgNotNotifyUserIDs not implemented")
}
func (UnimplementedConversationServer) CreateSingleChatConversations(context.Context, *CreateSingleChatConversationsReq) (*CreateSingleChatConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSingleChatConversations not implemented")
}
func (UnimplementedConversationServer) CreateGroupChatConversations(context.Context, *CreateGroupChatConversationsReq) (*CreateGroupChatConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupChatConversations not implemented")
}
func (UnimplementedConversationServer) SetConversationMaxSeq(context.Context, *SetConversationMaxSeqReq) (*SetConversationMaxSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConversationMaxSeq not implemented")
}
func (UnimplementedConversationServer) GetConversationIDs(context.Context, *GetConversationIDsReq) (*GetConversationIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationIDs not implemented")
}
func (UnimplementedConversationServer) SetConversations(context.Context, *SetConversationsReq) (*SetConversationsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConversations not implemented")
}
func (UnimplementedConversationServer) GetUserConversationIDsHash(context.Context, *GetUserConversationIDsHashReq) (*GetUserConversationIDsHashResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserConversationIDsHash not implemented")
}
func (UnimplementedConversationServer) GetConversationsByConversationID(context.Context, *GetConversationsByConversationIDReq) (*GetConversationsByConversationIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationsByConversationID not implemented")
}
func (UnimplementedConversationServer) GetConversationOfflinePushUserIDs(context.Context, *GetConversationOfflinePushUserIDsReq) (*GetConversationOfflinePushUserIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationOfflinePushUserIDs not implemented")
}
func (UnimplementedConversationServer) GetConversationNotReceiveMessageUserIDs(context.Context, *GetConversationNotReceiveMessageUserIDsReq) (*GetConversationNotReceiveMessageUserIDsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConversationNotReceiveMessageUserIDs not implemented")
}

func RegisterConversationServer(s grpc.ServiceRegistrar, srv ConversationServer) {
	s.RegisterService(&Conversation_ServiceDesc, srv)
}

func _Conversation_GetConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversation(ctx, req.(*GetConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetSortedConversationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSortedConversationListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetSortedConversationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetSortedConversationList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetSortedConversationList(ctx, req.(*GetSortedConversationListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetAllConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetAllConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetAllConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetAllConversations(ctx, req.(*GetAllConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversations(ctx, req.(*GetConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_SetConversation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConversationReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).SetConversation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/SetConversation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).SetConversation(ctx, req.(*SetConversationReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetRecvMsgNotNotifyUserIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecvMsgNotNotifyUserIDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetRecvMsgNotNotifyUserIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetRecvMsgNotNotifyUserIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetRecvMsgNotNotifyUserIDs(ctx, req.(*GetRecvMsgNotNotifyUserIDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_CreateSingleChatConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSingleChatConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).CreateSingleChatConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/CreateSingleChatConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).CreateSingleChatConversations(ctx, req.(*CreateSingleChatConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_CreateGroupChatConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupChatConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).CreateGroupChatConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/CreateGroupChatConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).CreateGroupChatConversations(ctx, req.(*CreateGroupChatConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_SetConversationMaxSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConversationMaxSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).SetConversationMaxSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/SetConversationMaxSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).SetConversationMaxSeq(ctx, req.(*SetConversationMaxSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetConversationIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationIDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversationIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetConversationIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversationIDs(ctx, req.(*GetConversationIDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_SetConversations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConversationsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).SetConversations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/SetConversations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).SetConversations(ctx, req.(*SetConversationsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetUserConversationIDsHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserConversationIDsHashReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetUserConversationIDsHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetUserConversationIDsHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetUserConversationIDsHash(ctx, req.(*GetUserConversationIDsHashReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetConversationsByConversationID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationsByConversationIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversationsByConversationID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetConversationsByConversationID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversationsByConversationID(ctx, req.(*GetConversationsByConversationIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetConversationOfflinePushUserIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationOfflinePushUserIDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversationOfflinePushUserIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetConversationOfflinePushUserIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversationOfflinePushUserIDs(ctx, req.(*GetConversationOfflinePushUserIDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conversation_GetConversationNotReceiveMessageUserIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConversationNotReceiveMessageUserIDsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConversationServer).GetConversationNotReceiveMessageUserIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.conversation.conversation/GetConversationNotReceiveMessageUserIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConversationServer).GetConversationNotReceiveMessageUserIDs(ctx, req.(*GetConversationNotReceiveMessageUserIDsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Conversation_ServiceDesc is the grpc.ServiceDesc for Conversation service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conversation_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openim.conversation.conversation",
	HandlerType: (*ConversationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConversation",
			Handler:    _Conversation_GetConversation_Handler,
		},
		{
			MethodName: "GetSortedConversationList",
			Handler:    _Conversation_GetSortedConversationList_Handler,
		},
		{
			MethodName: "GetAllConversations",
			Handler:    _Conversation_GetAllConversations_Handler,
		},
		{
			MethodName: "GetConversations",
			Handler:    _Conversation_GetConversations_Handler,
		},
		{
			MethodName: "SetConversation",
			Handler:    _Conversation_SetConversation_Handler,
		},
		{
			MethodName: "GetRecvMsgNotNotifyUserIDs",
			Handler:    _Conversation_GetRecvMsgNotNotifyUserIDs_Handler,
		},
		{
			MethodName: "CreateSingleChatConversations",
			Handler:    _Conversation_CreateSingleChatConversations_Handler,
		},
		{
			MethodName: "CreateGroupChatConversations",
			Handler:    _Conversation_CreateGroupChatConversations_Handler,
		},
		{
			MethodName: "SetConversationMaxSeq",
			Handler:    _Conversation_SetConversationMaxSeq_Handler,
		},
		{
			MethodName: "GetConversationIDs",
			Handler:    _Conversation_GetConversationIDs_Handler,
		},
		{
			MethodName: "SetConversations",
			Handler:    _Conversation_SetConversations_Handler,
		},
		{
			MethodName: "GetUserConversationIDsHash",
			Handler:    _Conversation_GetUserConversationIDsHash_Handler,
		},
		{
			MethodName: "GetConversationsByConversationID",
			Handler:    _Conversation_GetConversationsByConversationID_Handler,
		},
		{
			MethodName: "GetConversationOfflinePushUserIDs",
			Handler:    _Conversation_GetConversationOfflinePushUserIDs_Handler,
		},
		{
			MethodName: "GetConversationNotReceiveMessageUserIDs",
			Handler:    _Conversation_GetConversationNotReceiveMessageUserIDs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conversation/conversation.proto",
}
