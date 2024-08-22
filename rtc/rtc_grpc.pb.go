// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rtc

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

// RtcServiceClient is the client API for RtcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RtcServiceClient interface {
	SignalMessageAssemble(ctx context.Context, in *SignalMessageAssembleReq, opts ...grpc.CallOption) (*SignalMessageAssembleResp, error)
	SignalGetRoomByGroupID(ctx context.Context, in *SignalGetRoomByGroupIDReq, opts ...grpc.CallOption) (*SignalGetRoomByGroupIDResp, error)
	SignalGetTokenByRoomID(ctx context.Context, in *SignalGetTokenByRoomIDReq, opts ...grpc.CallOption) (*SignalGetTokenByRoomIDResp, error)
	SignalGetRooms(ctx context.Context, in *SignalGetRoomsReq, opts ...grpc.CallOption) (*SignalGetRoomsResp, error)
	GetSignalInvitationInfo(ctx context.Context, in *GetSignalInvitationInfoReq, opts ...grpc.CallOption) (*GetSignalInvitationInfoResp, error)
	GetSignalInvitationInfoStartApp(ctx context.Context, in *GetSignalInvitationInfoStartAppReq, opts ...grpc.CallOption) (*GetSignalInvitationInfoStartAppResp, error)
	// custom signal
	SignalSendCustomSignal(ctx context.Context, in *SignalSendCustomSignalReq, opts ...grpc.CallOption) (*SignalSendCustomSignalResp, error)
	// rtc cms
	GetSignalInvitationRecords(ctx context.Context, in *GetSignalInvitationRecordsReq, opts ...grpc.CallOption) (*GetSignalInvitationRecordsResp, error)
	DeleteSignalRecords(ctx context.Context, in *DeleteSignalRecordsReq, opts ...grpc.CallOption) (*DeleteSignalRecordsResp, error)
}

type rtcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRtcServiceClient(cc grpc.ClientConnInterface) RtcServiceClient {
	return &rtcServiceClient{cc}
}

func (c *rtcServiceClient) SignalMessageAssemble(ctx context.Context, in *SignalMessageAssembleReq, opts ...grpc.CallOption) (*SignalMessageAssembleResp, error) {
	out := new(SignalMessageAssembleResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/SignalMessageAssemble", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalGetRoomByGroupID(ctx context.Context, in *SignalGetRoomByGroupIDReq, opts ...grpc.CallOption) (*SignalGetRoomByGroupIDResp, error) {
	out := new(SignalGetRoomByGroupIDResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/SignalGetRoomByGroupID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalGetTokenByRoomID(ctx context.Context, in *SignalGetTokenByRoomIDReq, opts ...grpc.CallOption) (*SignalGetTokenByRoomIDResp, error) {
	out := new(SignalGetTokenByRoomIDResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/SignalGetTokenByRoomID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalGetRooms(ctx context.Context, in *SignalGetRoomsReq, opts ...grpc.CallOption) (*SignalGetRoomsResp, error) {
	out := new(SignalGetRoomsResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/SignalGetRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) GetSignalInvitationInfo(ctx context.Context, in *GetSignalInvitationInfoReq, opts ...grpc.CallOption) (*GetSignalInvitationInfoResp, error) {
	out := new(GetSignalInvitationInfoResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/GetSignalInvitationInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) GetSignalInvitationInfoStartApp(ctx context.Context, in *GetSignalInvitationInfoStartAppReq, opts ...grpc.CallOption) (*GetSignalInvitationInfoStartAppResp, error) {
	out := new(GetSignalInvitationInfoStartAppResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/GetSignalInvitationInfoStartApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) SignalSendCustomSignal(ctx context.Context, in *SignalSendCustomSignalReq, opts ...grpc.CallOption) (*SignalSendCustomSignalResp, error) {
	out := new(SignalSendCustomSignalResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/SignalSendCustomSignal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) GetSignalInvitationRecords(ctx context.Context, in *GetSignalInvitationRecordsReq, opts ...grpc.CallOption) (*GetSignalInvitationRecordsResp, error) {
	out := new(GetSignalInvitationRecordsResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/GetSignalInvitationRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rtcServiceClient) DeleteSignalRecords(ctx context.Context, in *DeleteSignalRecordsReq, opts ...grpc.CallOption) (*DeleteSignalRecordsResp, error) {
	out := new(DeleteSignalRecordsResp)
	err := c.cc.Invoke(ctx, "/openim.rtc.RtcService/DeleteSignalRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RtcServiceServer is the server API for RtcService service.
// All implementations must embed UnimplementedRtcServiceServer
// for forward compatibility
type RtcServiceServer interface {
	SignalMessageAssemble(context.Context, *SignalMessageAssembleReq) (*SignalMessageAssembleResp, error)
	SignalGetRoomByGroupID(context.Context, *SignalGetRoomByGroupIDReq) (*SignalGetRoomByGroupIDResp, error)
	SignalGetTokenByRoomID(context.Context, *SignalGetTokenByRoomIDReq) (*SignalGetTokenByRoomIDResp, error)
	SignalGetRooms(context.Context, *SignalGetRoomsReq) (*SignalGetRoomsResp, error)
	GetSignalInvitationInfo(context.Context, *GetSignalInvitationInfoReq) (*GetSignalInvitationInfoResp, error)
	GetSignalInvitationInfoStartApp(context.Context, *GetSignalInvitationInfoStartAppReq) (*GetSignalInvitationInfoStartAppResp, error)
	// custom signal
	SignalSendCustomSignal(context.Context, *SignalSendCustomSignalReq) (*SignalSendCustomSignalResp, error)
	// rtc cms
	GetSignalInvitationRecords(context.Context, *GetSignalInvitationRecordsReq) (*GetSignalInvitationRecordsResp, error)
	DeleteSignalRecords(context.Context, *DeleteSignalRecordsReq) (*DeleteSignalRecordsResp, error)
	mustEmbedUnimplementedRtcServiceServer()
}

// UnimplementedRtcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRtcServiceServer struct {
}

func (UnimplementedRtcServiceServer) SignalMessageAssemble(context.Context, *SignalMessageAssembleReq) (*SignalMessageAssembleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalMessageAssemble not implemented")
}
func (UnimplementedRtcServiceServer) SignalGetRoomByGroupID(context.Context, *SignalGetRoomByGroupIDReq) (*SignalGetRoomByGroupIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalGetRoomByGroupID not implemented")
}
func (UnimplementedRtcServiceServer) SignalGetTokenByRoomID(context.Context, *SignalGetTokenByRoomIDReq) (*SignalGetTokenByRoomIDResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalGetTokenByRoomID not implemented")
}
func (UnimplementedRtcServiceServer) SignalGetRooms(context.Context, *SignalGetRoomsReq) (*SignalGetRoomsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalGetRooms not implemented")
}
func (UnimplementedRtcServiceServer) GetSignalInvitationInfo(context.Context, *GetSignalInvitationInfoReq) (*GetSignalInvitationInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalInvitationInfo not implemented")
}
func (UnimplementedRtcServiceServer) GetSignalInvitationInfoStartApp(context.Context, *GetSignalInvitationInfoStartAppReq) (*GetSignalInvitationInfoStartAppResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalInvitationInfoStartApp not implemented")
}
func (UnimplementedRtcServiceServer) SignalSendCustomSignal(context.Context, *SignalSendCustomSignalReq) (*SignalSendCustomSignalResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignalSendCustomSignal not implemented")
}
func (UnimplementedRtcServiceServer) GetSignalInvitationRecords(context.Context, *GetSignalInvitationRecordsReq) (*GetSignalInvitationRecordsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSignalInvitationRecords not implemented")
}
func (UnimplementedRtcServiceServer) DeleteSignalRecords(context.Context, *DeleteSignalRecordsReq) (*DeleteSignalRecordsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSignalRecords not implemented")
}
func (UnimplementedRtcServiceServer) mustEmbedUnimplementedRtcServiceServer() {}

// UnsafeRtcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RtcServiceServer will
// result in compilation errors.
type UnsafeRtcServiceServer interface {
	mustEmbedUnimplementedRtcServiceServer()
}

func RegisterRtcServiceServer(s grpc.ServiceRegistrar, srv RtcServiceServer) {
	s.RegisterService(&RtcService_ServiceDesc, srv)
}

func _RtcService_SignalMessageAssemble_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalMessageAssembleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalMessageAssemble(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/SignalMessageAssemble",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalMessageAssemble(ctx, req.(*SignalMessageAssembleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalGetRoomByGroupID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalGetRoomByGroupIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalGetRoomByGroupID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/SignalGetRoomByGroupID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalGetRoomByGroupID(ctx, req.(*SignalGetRoomByGroupIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalGetTokenByRoomID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalGetTokenByRoomIDReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalGetTokenByRoomID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/SignalGetTokenByRoomID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalGetTokenByRoomID(ctx, req.(*SignalGetTokenByRoomIDReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalGetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalGetRoomsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalGetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/SignalGetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalGetRooms(ctx, req.(*SignalGetRoomsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_GetSignalInvitationInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalInvitationInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).GetSignalInvitationInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/GetSignalInvitationInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).GetSignalInvitationInfo(ctx, req.(*GetSignalInvitationInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_GetSignalInvitationInfoStartApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalInvitationInfoStartAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).GetSignalInvitationInfoStartApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/GetSignalInvitationInfoStartApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).GetSignalInvitationInfoStartApp(ctx, req.(*GetSignalInvitationInfoStartAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_SignalSendCustomSignal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalSendCustomSignalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).SignalSendCustomSignal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/SignalSendCustomSignal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).SignalSendCustomSignal(ctx, req.(*SignalSendCustomSignalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_GetSignalInvitationRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSignalInvitationRecordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).GetSignalInvitationRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/GetSignalInvitationRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).GetSignalInvitationRecords(ctx, req.(*GetSignalInvitationRecordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _RtcService_DeleteSignalRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSignalRecordsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RtcServiceServer).DeleteSignalRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/openim.rtc.RtcService/DeleteSignalRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RtcServiceServer).DeleteSignalRecords(ctx, req.(*DeleteSignalRecordsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RtcService_ServiceDesc is the grpc.ServiceDesc for RtcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RtcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "openim.rtc.RtcService",
	HandlerType: (*RtcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignalMessageAssemble",
			Handler:    _RtcService_SignalMessageAssemble_Handler,
		},
		{
			MethodName: "SignalGetRoomByGroupID",
			Handler:    _RtcService_SignalGetRoomByGroupID_Handler,
		},
		{
			MethodName: "SignalGetTokenByRoomID",
			Handler:    _RtcService_SignalGetTokenByRoomID_Handler,
		},
		{
			MethodName: "SignalGetRooms",
			Handler:    _RtcService_SignalGetRooms_Handler,
		},
		{
			MethodName: "GetSignalInvitationInfo",
			Handler:    _RtcService_GetSignalInvitationInfo_Handler,
		},
		{
			MethodName: "GetSignalInvitationInfoStartApp",
			Handler:    _RtcService_GetSignalInvitationInfoStartApp_Handler,
		},
		{
			MethodName: "SignalSendCustomSignal",
			Handler:    _RtcService_SignalSendCustomSignal_Handler,
		},
		{
			MethodName: "GetSignalInvitationRecords",
			Handler:    _RtcService_GetSignalInvitationRecords_Handler,
		},
		{
			MethodName: "DeleteSignalRecords",
			Handler:    _RtcService_DeleteSignalRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rtc/rtc.proto",
}
