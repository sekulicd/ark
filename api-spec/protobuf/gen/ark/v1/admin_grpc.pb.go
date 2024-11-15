// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package arkv1

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

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	GetScheduledSweep(ctx context.Context, in *GetScheduledSweepRequest, opts ...grpc.CallOption) (*GetScheduledSweepResponse, error)
	GetRoundDetails(ctx context.Context, in *GetRoundDetailsRequest, opts ...grpc.CallOption) (*GetRoundDetailsResponse, error)
	GetRounds(ctx context.Context, in *GetRoundsRequest, opts ...grpc.CallOption) (*GetRoundsResponse, error)
	UpdateMarketHour(ctx context.Context, in *UpdateMarketHourRequest, opts ...grpc.CallOption) (*UpdateMarketHourResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) GetScheduledSweep(ctx context.Context, in *GetScheduledSweepRequest, opts ...grpc.CallOption) (*GetScheduledSweepResponse, error) {
	out := new(GetScheduledSweepResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.AdminService/GetScheduledSweep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetRoundDetails(ctx context.Context, in *GetRoundDetailsRequest, opts ...grpc.CallOption) (*GetRoundDetailsResponse, error) {
	out := new(GetRoundDetailsResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.AdminService/GetRoundDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetRounds(ctx context.Context, in *GetRoundsRequest, opts ...grpc.CallOption) (*GetRoundsResponse, error) {
	out := new(GetRoundsResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.AdminService/GetRounds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateMarketHour(ctx context.Context, in *UpdateMarketHourRequest, opts ...grpc.CallOption) (*UpdateMarketHourResponse, error) {
	out := new(UpdateMarketHourResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.AdminService/UpdateMarketHour", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations should embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	GetScheduledSweep(context.Context, *GetScheduledSweepRequest) (*GetScheduledSweepResponse, error)
	GetRoundDetails(context.Context, *GetRoundDetailsRequest) (*GetRoundDetailsResponse, error)
	GetRounds(context.Context, *GetRoundsRequest) (*GetRoundsResponse, error)
	UpdateMarketHour(context.Context, *UpdateMarketHourRequest) (*UpdateMarketHourResponse, error)
}

// UnimplementedAdminServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) GetScheduledSweep(context.Context, *GetScheduledSweepRequest) (*GetScheduledSweepResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScheduledSweep not implemented")
}
func (UnimplementedAdminServiceServer) GetRoundDetails(context.Context, *GetRoundDetailsRequest) (*GetRoundDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoundDetails not implemented")
}
func (UnimplementedAdminServiceServer) GetRounds(context.Context, *GetRoundsRequest) (*GetRoundsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRounds not implemented")
}
func (UnimplementedAdminServiceServer) UpdateMarketHour(context.Context, *UpdateMarketHourRequest) (*UpdateMarketHourResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMarketHour not implemented")
}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_GetScheduledSweep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScheduledSweepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetScheduledSweep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.AdminService/GetScheduledSweep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetScheduledSweep(ctx, req.(*GetScheduledSweepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetRoundDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoundDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetRoundDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.AdminService/GetRoundDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetRoundDetails(ctx, req.(*GetRoundDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetRounds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoundsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetRounds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.AdminService/GetRounds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetRounds(ctx, req.(*GetRoundsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateMarketHour_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMarketHourRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateMarketHour(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.AdminService/UpdateMarketHour",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateMarketHour(ctx, req.(*UpdateMarketHourRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ark.v1.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScheduledSweep",
			Handler:    _AdminService_GetScheduledSweep_Handler,
		},
		{
			MethodName: "GetRoundDetails",
			Handler:    _AdminService_GetRoundDetails_Handler,
		},
		{
			MethodName: "GetRounds",
			Handler:    _AdminService_GetRounds_Handler,
		},
		{
			MethodName: "UpdateMarketHour",
			Handler:    _AdminService_UpdateMarketHour_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ark/v1/admin.proto",
}
