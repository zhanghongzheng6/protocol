// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package push

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

// PushMsgServiceClient is the client API for PushMsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushMsgServiceClient interface {
	PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgResp, error)
	DelUserPushToken(ctx context.Context, in *DelUserPushTokenReq, opts ...grpc.CallOption) (*DelUserPushTokenResp, error)
}

type pushMsgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPushMsgServiceClient(cc grpc.ClientConnInterface) PushMsgServiceClient {
	return &pushMsgServiceClient{cc}
}

func (c *pushMsgServiceClient) PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgResp, error) {
	out := new(PushMsgResp)
	err := c.cc.Invoke(ctx, "/openim.push.PushMsgService/PushMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushMsgServiceClient) DelUserPushToken(ctx context.Context, in *DelUserPushTokenReq, opts ...grpc.CallOption) (*DelUserPushTokenResp, error) {
	out := new(DelUserPushTokenResp)
	err := c.cc.Invoke(ctx, "/openim.push.PushMsgService/DelUserPushToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushMsgServiceServer is the server API for PushMsgService service.
// All implementations must embed UnimplementedPushMsgServiceServer
// for forward compatibility
type PushMsgServiceServer interface {
	PushMsg(context.Context, *PushMsgReq) (*PushMsgResp, error)
	DelUserPushToken(context.Context, *DelUserPushTokenReq) (*DelUserPushTokenResp, error)
	mustEmbedUnimplementedPushMsgServiceServer()
}

// UnimplementedPushMsgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPushMsgServiceServer struct {
}

func (UnimplementedPushMsgServiceServer) PushMsg(context.Context, *PushMsgReq) (*PushMsgResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMsg not implemented")
}
func (UnimplementedPushMsgServiceServer) DelUserPushToken(context.Context, *DelUserPushTokenReq) (*DelUserPushTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUserPushToken not implemented")
}
func (UnimplementedPushMsgServiceServer) mustEmbedUnimplementedPushMsgServiceServer() {}

// UnsafePushMsgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushMsgServiceServer will
// result in compilation errors.
type UnsafePushMsgServiceServer interface {
	mustEmbedUnimplementedPushMsgServiceServer()
}

func RegisterPushMsgServiceServer(s grpc.ServiceRegistrar, srv PushMsgServiceServer) {
	s.RegisterService(&PushMsgService_ServiceDesc, srv)
}

func _PushMsgService_PushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushMsgServiceServer).PushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.push.PushMsgService/PushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushMsgServiceServer).PushMsg(ctx, req.(*PushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PushMsgService_DelUserPushToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserPushTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushMsgServiceServer).DelUserPushToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.push.PushMsgService/DelUserPushToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushMsgServiceServer).DelUserPushToken(ctx, req.(*DelUserPushTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PushMsgService_ServiceDesc is the grpc.ServiceDesc for PushMsgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushMsgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openim.push.PushMsgService",
	HandlerType: (*PushMsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMsg",
			Handler:    _PushMsgService_PushMsg_Handler,
		},
		{
			MethodName: "DelUserPushToken",
			Handler:    _PushMsgService_DelUserPushToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push/push.proto",
}
