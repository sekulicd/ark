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

// ExplorerServiceClient is the client API for ExplorerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExplorerServiceClient interface {
	GetCommitmentTx(ctx context.Context, in *GetCommitmentTxRequest, opts ...grpc.CallOption) (*GetCommitmentTxResponse, error)
	GetVtxoTree(ctx context.Context, in *GetVtxoTreeRequest, opts ...grpc.CallOption) (*GetVtxoTreeResponse, error)
	GetForfeitTxs(ctx context.Context, in *GetForfeitTxsRequest, opts ...grpc.CallOption) (*GetForfeitTxsResponse, error)
	GetConnectors(ctx context.Context, in *GetConnectorsRequest, opts ...grpc.CallOption) (*GetConnectorsResponse, error)
	GetSpendableVtxos(ctx context.Context, in *GetSpendableVtxosRequest, opts ...grpc.CallOption) (*GetSpendableVtxosResponse, error)
	GetTransactionHistory(ctx context.Context, in *GetTransactionHistoryRequest, opts ...grpc.CallOption) (*GetTransactionHistoryResponse, error)
	GetVtxoChain(ctx context.Context, in *GetVtxoChainRequest, opts ...grpc.CallOption) (*GetVtxoChainResponse, error)
	GetVirtualTxs(ctx context.Context, in *GetVirtualTxsRequest, opts ...grpc.CallOption) (*GetVirtualTxsResponse, error)
	GetSweptCommitmentTx(ctx context.Context, in *GetSweptCommitmentTxRequest, opts ...grpc.CallOption) (*GetSweptCommitmentTxResponse, error)
	SubscribeForAddresses(ctx context.Context, in *SubscribeForAddressesRequest, opts ...grpc.CallOption) (ExplorerService_SubscribeForAddressesClient, error)
}

type explorerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExplorerServiceClient(cc grpc.ClientConnInterface) ExplorerServiceClient {
	return &explorerServiceClient{cc}
}

func (c *explorerServiceClient) GetCommitmentTx(ctx context.Context, in *GetCommitmentTxRequest, opts ...grpc.CallOption) (*GetCommitmentTxResponse, error) {
	out := new(GetCommitmentTxResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetCommitmentTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetVtxoTree(ctx context.Context, in *GetVtxoTreeRequest, opts ...grpc.CallOption) (*GetVtxoTreeResponse, error) {
	out := new(GetVtxoTreeResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetVtxoTree", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetForfeitTxs(ctx context.Context, in *GetForfeitTxsRequest, opts ...grpc.CallOption) (*GetForfeitTxsResponse, error) {
	out := new(GetForfeitTxsResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetForfeitTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetConnectors(ctx context.Context, in *GetConnectorsRequest, opts ...grpc.CallOption) (*GetConnectorsResponse, error) {
	out := new(GetConnectorsResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetConnectors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetSpendableVtxos(ctx context.Context, in *GetSpendableVtxosRequest, opts ...grpc.CallOption) (*GetSpendableVtxosResponse, error) {
	out := new(GetSpendableVtxosResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetSpendableVtxos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetTransactionHistory(ctx context.Context, in *GetTransactionHistoryRequest, opts ...grpc.CallOption) (*GetTransactionHistoryResponse, error) {
	out := new(GetTransactionHistoryResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetTransactionHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetVtxoChain(ctx context.Context, in *GetVtxoChainRequest, opts ...grpc.CallOption) (*GetVtxoChainResponse, error) {
	out := new(GetVtxoChainResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetVtxoChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetVirtualTxs(ctx context.Context, in *GetVirtualTxsRequest, opts ...grpc.CallOption) (*GetVirtualTxsResponse, error) {
	out := new(GetVirtualTxsResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetVirtualTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) GetSweptCommitmentTx(ctx context.Context, in *GetSweptCommitmentTxRequest, opts ...grpc.CallOption) (*GetSweptCommitmentTxResponse, error) {
	out := new(GetSweptCommitmentTxResponse)
	err := c.cc.Invoke(ctx, "/ark.v1.ExplorerService/GetSweptCommitmentTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *explorerServiceClient) SubscribeForAddresses(ctx context.Context, in *SubscribeForAddressesRequest, opts ...grpc.CallOption) (ExplorerService_SubscribeForAddressesClient, error) {
	stream, err := c.cc.NewStream(ctx, &ExplorerService_ServiceDesc.Streams[0], "/ark.v1.ExplorerService/SubscribeForAddresses", opts...)
	if err != nil {
		return nil, err
	}
	x := &explorerServiceSubscribeForAddressesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ExplorerService_SubscribeForAddressesClient interface {
	Recv() (*SubscribeForAddressesResponse, error)
	grpc.ClientStream
}

type explorerServiceSubscribeForAddressesClient struct {
	grpc.ClientStream
}

func (x *explorerServiceSubscribeForAddressesClient) Recv() (*SubscribeForAddressesResponse, error) {
	m := new(SubscribeForAddressesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ExplorerServiceServer is the server API for ExplorerService service.
// All implementations should embed UnimplementedExplorerServiceServer
// for forward compatibility
type ExplorerServiceServer interface {
	GetCommitmentTx(context.Context, *GetCommitmentTxRequest) (*GetCommitmentTxResponse, error)
	GetVtxoTree(context.Context, *GetVtxoTreeRequest) (*GetVtxoTreeResponse, error)
	GetForfeitTxs(context.Context, *GetForfeitTxsRequest) (*GetForfeitTxsResponse, error)
	GetConnectors(context.Context, *GetConnectorsRequest) (*GetConnectorsResponse, error)
	GetSpendableVtxos(context.Context, *GetSpendableVtxosRequest) (*GetSpendableVtxosResponse, error)
	GetTransactionHistory(context.Context, *GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error)
	GetVtxoChain(context.Context, *GetVtxoChainRequest) (*GetVtxoChainResponse, error)
	GetVirtualTxs(context.Context, *GetVirtualTxsRequest) (*GetVirtualTxsResponse, error)
	GetSweptCommitmentTx(context.Context, *GetSweptCommitmentTxRequest) (*GetSweptCommitmentTxResponse, error)
	SubscribeForAddresses(*SubscribeForAddressesRequest, ExplorerService_SubscribeForAddressesServer) error
}

// UnimplementedExplorerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedExplorerServiceServer struct {
}

func (UnimplementedExplorerServiceServer) GetCommitmentTx(context.Context, *GetCommitmentTxRequest) (*GetCommitmentTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommitmentTx not implemented")
}
func (UnimplementedExplorerServiceServer) GetVtxoTree(context.Context, *GetVtxoTreeRequest) (*GetVtxoTreeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVtxoTree not implemented")
}
func (UnimplementedExplorerServiceServer) GetForfeitTxs(context.Context, *GetForfeitTxsRequest) (*GetForfeitTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForfeitTxs not implemented")
}
func (UnimplementedExplorerServiceServer) GetConnectors(context.Context, *GetConnectorsRequest) (*GetConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnectors not implemented")
}
func (UnimplementedExplorerServiceServer) GetSpendableVtxos(context.Context, *GetSpendableVtxosRequest) (*GetSpendableVtxosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSpendableVtxos not implemented")
}
func (UnimplementedExplorerServiceServer) GetTransactionHistory(context.Context, *GetTransactionHistoryRequest) (*GetTransactionHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactionHistory not implemented")
}
func (UnimplementedExplorerServiceServer) GetVtxoChain(context.Context, *GetVtxoChainRequest) (*GetVtxoChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVtxoChain not implemented")
}
func (UnimplementedExplorerServiceServer) GetVirtualTxs(context.Context, *GetVirtualTxsRequest) (*GetVirtualTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVirtualTxs not implemented")
}
func (UnimplementedExplorerServiceServer) GetSweptCommitmentTx(context.Context, *GetSweptCommitmentTxRequest) (*GetSweptCommitmentTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSweptCommitmentTx not implemented")
}
func (UnimplementedExplorerServiceServer) SubscribeForAddresses(*SubscribeForAddressesRequest, ExplorerService_SubscribeForAddressesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeForAddresses not implemented")
}

// UnsafeExplorerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExplorerServiceServer will
// result in compilation errors.
type UnsafeExplorerServiceServer interface {
	mustEmbedUnimplementedExplorerServiceServer()
}

func RegisterExplorerServiceServer(s grpc.ServiceRegistrar, srv ExplorerServiceServer) {
	s.RegisterService(&ExplorerService_ServiceDesc, srv)
}

func _ExplorerService_GetCommitmentTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommitmentTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetCommitmentTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetCommitmentTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetCommitmentTx(ctx, req.(*GetCommitmentTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetVtxoTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVtxoTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetVtxoTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetVtxoTree",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetVtxoTree(ctx, req.(*GetVtxoTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetForfeitTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForfeitTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetForfeitTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetForfeitTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetForfeitTxs(ctx, req.(*GetForfeitTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetConnectors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetConnectors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetConnectors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetConnectors(ctx, req.(*GetConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetSpendableVtxos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSpendableVtxosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetSpendableVtxos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetSpendableVtxos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetSpendableVtxos(ctx, req.(*GetSpendableVtxosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetTransactionHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetTransactionHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetTransactionHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetTransactionHistory(ctx, req.(*GetTransactionHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetVtxoChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVtxoChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetVtxoChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetVtxoChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetVtxoChain(ctx, req.(*GetVtxoChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetVirtualTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVirtualTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetVirtualTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetVirtualTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetVirtualTxs(ctx, req.(*GetVirtualTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_GetSweptCommitmentTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSweptCommitmentTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExplorerServiceServer).GetSweptCommitmentTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ark.v1.ExplorerService/GetSweptCommitmentTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExplorerServiceServer).GetSweptCommitmentTx(ctx, req.(*GetSweptCommitmentTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExplorerService_SubscribeForAddresses_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeForAddressesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ExplorerServiceServer).SubscribeForAddresses(m, &explorerServiceSubscribeForAddressesServer{stream})
}

type ExplorerService_SubscribeForAddressesServer interface {
	Send(*SubscribeForAddressesResponse) error
	grpc.ServerStream
}

type explorerServiceSubscribeForAddressesServer struct {
	grpc.ServerStream
}

func (x *explorerServiceSubscribeForAddressesServer) Send(m *SubscribeForAddressesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ExplorerService_ServiceDesc is the grpc.ServiceDesc for ExplorerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExplorerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ark.v1.ExplorerService",
	HandlerType: (*ExplorerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCommitmentTx",
			Handler:    _ExplorerService_GetCommitmentTx_Handler,
		},
		{
			MethodName: "GetVtxoTree",
			Handler:    _ExplorerService_GetVtxoTree_Handler,
		},
		{
			MethodName: "GetForfeitTxs",
			Handler:    _ExplorerService_GetForfeitTxs_Handler,
		},
		{
			MethodName: "GetConnectors",
			Handler:    _ExplorerService_GetConnectors_Handler,
		},
		{
			MethodName: "GetSpendableVtxos",
			Handler:    _ExplorerService_GetSpendableVtxos_Handler,
		},
		{
			MethodName: "GetTransactionHistory",
			Handler:    _ExplorerService_GetTransactionHistory_Handler,
		},
		{
			MethodName: "GetVtxoChain",
			Handler:    _ExplorerService_GetVtxoChain_Handler,
		},
		{
			MethodName: "GetVirtualTxs",
			Handler:    _ExplorerService_GetVirtualTxs_Handler,
		},
		{
			MethodName: "GetSweptCommitmentTx",
			Handler:    _ExplorerService_GetSweptCommitmentTx_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeForAddresses",
			Handler:       _ExplorerService_SubscribeForAddresses_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ark/v1/explorer.proto",
}
