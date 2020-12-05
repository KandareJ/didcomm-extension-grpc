// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package messaging

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// DIDCommPlainClient is the client API for DIDCommPlain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DIDCommPlainClient interface {
	Unary(ctx context.Context, in *CoreMessage, opts ...grpc.CallOption) (*CoreMessage, error)
	ServerStreaming(ctx context.Context, in *CoreMessage, opts ...grpc.CallOption) (DIDCommPlain_ServerStreamingClient, error)
	ClientStreaming(ctx context.Context, opts ...grpc.CallOption) (DIDCommPlain_ClientStreamingClient, error)
	BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (DIDCommPlain_BidirectionalStreamingClient, error)
}

type dIDCommPlainClient struct {
	cc grpc.ClientConnInterface
}

func NewDIDCommPlainClient(cc grpc.ClientConnInterface) DIDCommPlainClient {
	return &dIDCommPlainClient{cc}
}

func (c *dIDCommPlainClient) Unary(ctx context.Context, in *CoreMessage, opts ...grpc.CallOption) (*CoreMessage, error) {
	out := new(CoreMessage)
	err := c.cc.Invoke(ctx, "/didcomm.messaging.DIDCommPlain/Unary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dIDCommPlainClient) ServerStreaming(ctx context.Context, in *CoreMessage, opts ...grpc.CallOption) (DIDCommPlain_ServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DIDCommPlain_serviceDesc.Streams[0], "/didcomm.messaging.DIDCommPlain/ServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &dIDCommPlainServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DIDCommPlain_ServerStreamingClient interface {
	Recv() (*CoreMessage, error)
	grpc.ClientStream
}

type dIDCommPlainServerStreamingClient struct {
	grpc.ClientStream
}

func (x *dIDCommPlainServerStreamingClient) Recv() (*CoreMessage, error) {
	m := new(CoreMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dIDCommPlainClient) ClientStreaming(ctx context.Context, opts ...grpc.CallOption) (DIDCommPlain_ClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DIDCommPlain_serviceDesc.Streams[1], "/didcomm.messaging.DIDCommPlain/ClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &dIDCommPlainClientStreamingClient{stream}
	return x, nil
}

type DIDCommPlain_ClientStreamingClient interface {
	Send(*CoreMessage) error
	CloseAndRecv() (*CoreMessage, error)
	grpc.ClientStream
}

type dIDCommPlainClientStreamingClient struct {
	grpc.ClientStream
}

func (x *dIDCommPlainClientStreamingClient) Send(m *CoreMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dIDCommPlainClientStreamingClient) CloseAndRecv() (*CoreMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CoreMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dIDCommPlainClient) BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (DIDCommPlain_BidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DIDCommPlain_serviceDesc.Streams[2], "/didcomm.messaging.DIDCommPlain/BidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &dIDCommPlainBidirectionalStreamingClient{stream}
	return x, nil
}

type DIDCommPlain_BidirectionalStreamingClient interface {
	Send(*CoreMessage) error
	Recv() (*CoreMessage, error)
	grpc.ClientStream
}

type dIDCommPlainBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *dIDCommPlainBidirectionalStreamingClient) Send(m *CoreMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dIDCommPlainBidirectionalStreamingClient) Recv() (*CoreMessage, error) {
	m := new(CoreMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DIDCommPlainServer is the server API for DIDCommPlain service.
// All implementations must embed UnimplementedDIDCommPlainServer
// for forward compatibility
type DIDCommPlainServer interface {
	Unary(context.Context, *CoreMessage) (*CoreMessage, error)
	ServerStreaming(*CoreMessage, DIDCommPlain_ServerStreamingServer) error
	ClientStreaming(DIDCommPlain_ClientStreamingServer) error
	BidirectionalStreaming(DIDCommPlain_BidirectionalStreamingServer) error
	mustEmbedUnimplementedDIDCommPlainServer()
}

// UnimplementedDIDCommPlainServer must be embedded to have forward compatible implementations.
type UnimplementedDIDCommPlainServer struct {
}

func (UnimplementedDIDCommPlainServer) Unary(context.Context, *CoreMessage) (*CoreMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary not implemented")
}
func (UnimplementedDIDCommPlainServer) ServerStreaming(*CoreMessage, DIDCommPlain_ServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreaming not implemented")
}
func (UnimplementedDIDCommPlainServer) ClientStreaming(DIDCommPlain_ClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreaming not implemented")
}
func (UnimplementedDIDCommPlainServer) BidirectionalStreaming(DIDCommPlain_BidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreaming not implemented")
}
func (UnimplementedDIDCommPlainServer) mustEmbedUnimplementedDIDCommPlainServer() {}

// UnsafeDIDCommPlainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DIDCommPlainServer will
// result in compilation errors.
type UnsafeDIDCommPlainServer interface {
	mustEmbedUnimplementedDIDCommPlainServer()
}

func RegisterDIDCommPlainServer(s grpc.ServiceRegistrar, srv DIDCommPlainServer) {
	s.RegisterService(&_DIDCommPlain_serviceDesc, srv)
}

func _DIDCommPlain_Unary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoreMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DIDCommPlainServer).Unary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didcomm.messaging.DIDCommPlain/Unary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DIDCommPlainServer).Unary(ctx, req.(*CoreMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DIDCommPlain_ServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CoreMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DIDCommPlainServer).ServerStreaming(m, &dIDCommPlainServerStreamingServer{stream})
}

type DIDCommPlain_ServerStreamingServer interface {
	Send(*CoreMessage) error
	grpc.ServerStream
}

type dIDCommPlainServerStreamingServer struct {
	grpc.ServerStream
}

func (x *dIDCommPlainServerStreamingServer) Send(m *CoreMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _DIDCommPlain_ClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DIDCommPlainServer).ClientStreaming(&dIDCommPlainClientStreamingServer{stream})
}

type DIDCommPlain_ClientStreamingServer interface {
	SendAndClose(*CoreMessage) error
	Recv() (*CoreMessage, error)
	grpc.ServerStream
}

type dIDCommPlainClientStreamingServer struct {
	grpc.ServerStream
}

func (x *dIDCommPlainClientStreamingServer) SendAndClose(m *CoreMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dIDCommPlainClientStreamingServer) Recv() (*CoreMessage, error) {
	m := new(CoreMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DIDCommPlain_BidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DIDCommPlainServer).BidirectionalStreaming(&dIDCommPlainBidirectionalStreamingServer{stream})
}

type DIDCommPlain_BidirectionalStreamingServer interface {
	Send(*CoreMessage) error
	Recv() (*CoreMessage, error)
	grpc.ServerStream
}

type dIDCommPlainBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *dIDCommPlainBidirectionalStreamingServer) Send(m *CoreMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dIDCommPlainBidirectionalStreamingServer) Recv() (*CoreMessage, error) {
	m := new(CoreMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DIDCommPlain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "didcomm.messaging.DIDCommPlain",
	HandlerType: (*DIDCommPlainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary",
			Handler:    _DIDCommPlain_Unary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreaming",
			Handler:       _DIDCommPlain_ServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreaming",
			Handler:       _DIDCommPlain_ClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreaming",
			Handler:       _DIDCommPlain_BidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "didcomm.proto",
}

// DIDCommEncryptedClient is the client API for DIDCommEncrypted service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DIDCommEncryptedClient interface {
	Unary(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (*EncryptedMessage, error)
	ServerStreaming(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (DIDCommEncrypted_ServerStreamingClient, error)
	ClientStreaming(ctx context.Context, opts ...grpc.CallOption) (DIDCommEncrypted_ClientStreamingClient, error)
	BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (DIDCommEncrypted_BidirectionalStreamingClient, error)
}

type dIDCommEncryptedClient struct {
	cc grpc.ClientConnInterface
}

func NewDIDCommEncryptedClient(cc grpc.ClientConnInterface) DIDCommEncryptedClient {
	return &dIDCommEncryptedClient{cc}
}

func (c *dIDCommEncryptedClient) Unary(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (*EncryptedMessage, error) {
	out := new(EncryptedMessage)
	err := c.cc.Invoke(ctx, "/didcomm.messaging.DIDCommEncrypted/Unary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dIDCommEncryptedClient) ServerStreaming(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (DIDCommEncrypted_ServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DIDCommEncrypted_serviceDesc.Streams[0], "/didcomm.messaging.DIDCommEncrypted/ServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &dIDCommEncryptedServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DIDCommEncrypted_ServerStreamingClient interface {
	Recv() (*EncryptedMessage, error)
	grpc.ClientStream
}

type dIDCommEncryptedServerStreamingClient struct {
	grpc.ClientStream
}

func (x *dIDCommEncryptedServerStreamingClient) Recv() (*EncryptedMessage, error) {
	m := new(EncryptedMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dIDCommEncryptedClient) ClientStreaming(ctx context.Context, opts ...grpc.CallOption) (DIDCommEncrypted_ClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DIDCommEncrypted_serviceDesc.Streams[1], "/didcomm.messaging.DIDCommEncrypted/ClientStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &dIDCommEncryptedClientStreamingClient{stream}
	return x, nil
}

type DIDCommEncrypted_ClientStreamingClient interface {
	Send(*EncryptedMessage) error
	CloseAndRecv() (*EncryptedMessage, error)
	grpc.ClientStream
}

type dIDCommEncryptedClientStreamingClient struct {
	grpc.ClientStream
}

func (x *dIDCommEncryptedClientStreamingClient) Send(m *EncryptedMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dIDCommEncryptedClientStreamingClient) CloseAndRecv() (*EncryptedMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EncryptedMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dIDCommEncryptedClient) BidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (DIDCommEncrypted_BidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DIDCommEncrypted_serviceDesc.Streams[2], "/didcomm.messaging.DIDCommEncrypted/BidirectionalStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &dIDCommEncryptedBidirectionalStreamingClient{stream}
	return x, nil
}

type DIDCommEncrypted_BidirectionalStreamingClient interface {
	Send(*EncryptedMessage) error
	Recv() (*EncryptedMessage, error)
	grpc.ClientStream
}

type dIDCommEncryptedBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *dIDCommEncryptedBidirectionalStreamingClient) Send(m *EncryptedMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dIDCommEncryptedBidirectionalStreamingClient) Recv() (*EncryptedMessage, error) {
	m := new(EncryptedMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DIDCommEncryptedServer is the server API for DIDCommEncrypted service.
// All implementations must embed UnimplementedDIDCommEncryptedServer
// for forward compatibility
type DIDCommEncryptedServer interface {
	Unary(context.Context, *EncryptedMessage) (*EncryptedMessage, error)
	ServerStreaming(*EncryptedMessage, DIDCommEncrypted_ServerStreamingServer) error
	ClientStreaming(DIDCommEncrypted_ClientStreamingServer) error
	BidirectionalStreaming(DIDCommEncrypted_BidirectionalStreamingServer) error
	mustEmbedUnimplementedDIDCommEncryptedServer()
}

// UnimplementedDIDCommEncryptedServer must be embedded to have forward compatible implementations.
type UnimplementedDIDCommEncryptedServer struct {
}

func (UnimplementedDIDCommEncryptedServer) Unary(context.Context, *EncryptedMessage) (*EncryptedMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unary not implemented")
}
func (UnimplementedDIDCommEncryptedServer) ServerStreaming(*EncryptedMessage, DIDCommEncrypted_ServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreaming not implemented")
}
func (UnimplementedDIDCommEncryptedServer) ClientStreaming(DIDCommEncrypted_ClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreaming not implemented")
}
func (UnimplementedDIDCommEncryptedServer) BidirectionalStreaming(DIDCommEncrypted_BidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreaming not implemented")
}
func (UnimplementedDIDCommEncryptedServer) mustEmbedUnimplementedDIDCommEncryptedServer() {}

// UnsafeDIDCommEncryptedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DIDCommEncryptedServer will
// result in compilation errors.
type UnsafeDIDCommEncryptedServer interface {
	mustEmbedUnimplementedDIDCommEncryptedServer()
}

func RegisterDIDCommEncryptedServer(s grpc.ServiceRegistrar, srv DIDCommEncryptedServer) {
	s.RegisterService(&_DIDCommEncrypted_serviceDesc, srv)
}

func _DIDCommEncrypted_Unary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DIDCommEncryptedServer).Unary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didcomm.messaging.DIDCommEncrypted/Unary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DIDCommEncryptedServer).Unary(ctx, req.(*EncryptedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _DIDCommEncrypted_ServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EncryptedMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DIDCommEncryptedServer).ServerStreaming(m, &dIDCommEncryptedServerStreamingServer{stream})
}

type DIDCommEncrypted_ServerStreamingServer interface {
	Send(*EncryptedMessage) error
	grpc.ServerStream
}

type dIDCommEncryptedServerStreamingServer struct {
	grpc.ServerStream
}

func (x *dIDCommEncryptedServerStreamingServer) Send(m *EncryptedMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _DIDCommEncrypted_ClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DIDCommEncryptedServer).ClientStreaming(&dIDCommEncryptedClientStreamingServer{stream})
}

type DIDCommEncrypted_ClientStreamingServer interface {
	SendAndClose(*EncryptedMessage) error
	Recv() (*EncryptedMessage, error)
	grpc.ServerStream
}

type dIDCommEncryptedClientStreamingServer struct {
	grpc.ServerStream
}

func (x *dIDCommEncryptedClientStreamingServer) SendAndClose(m *EncryptedMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dIDCommEncryptedClientStreamingServer) Recv() (*EncryptedMessage, error) {
	m := new(EncryptedMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DIDCommEncrypted_BidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DIDCommEncryptedServer).BidirectionalStreaming(&dIDCommEncryptedBidirectionalStreamingServer{stream})
}

type DIDCommEncrypted_BidirectionalStreamingServer interface {
	Send(*EncryptedMessage) error
	Recv() (*EncryptedMessage, error)
	grpc.ServerStream
}

type dIDCommEncryptedBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *dIDCommEncryptedBidirectionalStreamingServer) Send(m *EncryptedMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dIDCommEncryptedBidirectionalStreamingServer) Recv() (*EncryptedMessage, error) {
	m := new(EncryptedMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DIDCommEncrypted_serviceDesc = grpc.ServiceDesc{
	ServiceName: "didcomm.messaging.DIDCommEncrypted",
	HandlerType: (*DIDCommEncryptedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unary",
			Handler:    _DIDCommEncrypted_Unary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreaming",
			Handler:       _DIDCommEncrypted_ServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreaming",
			Handler:       _DIDCommEncrypted_ClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreaming",
			Handler:       _DIDCommEncrypted_BidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "didcomm.proto",
}

// DIDCommSimplexClient is the client API for DIDCommSimplex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DIDCommSimplexClient interface {
	Simplex(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (*NoOp, error)
}

type dIDCommSimplexClient struct {
	cc grpc.ClientConnInterface
}

func NewDIDCommSimplexClient(cc grpc.ClientConnInterface) DIDCommSimplexClient {
	return &dIDCommSimplexClient{cc}
}

func (c *dIDCommSimplexClient) Simplex(ctx context.Context, in *EncryptedMessage, opts ...grpc.CallOption) (*NoOp, error) {
	out := new(NoOp)
	err := c.cc.Invoke(ctx, "/didcomm.messaging.DIDCommSimplex/Simplex", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DIDCommSimplexServer is the server API for DIDCommSimplex service.
// All implementations must embed UnimplementedDIDCommSimplexServer
// for forward compatibility
type DIDCommSimplexServer interface {
	Simplex(context.Context, *EncryptedMessage) (*NoOp, error)
	mustEmbedUnimplementedDIDCommSimplexServer()
}

// UnimplementedDIDCommSimplexServer must be embedded to have forward compatible implementations.
type UnimplementedDIDCommSimplexServer struct {
}

func (UnimplementedDIDCommSimplexServer) Simplex(context.Context, *EncryptedMessage) (*NoOp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Simplex not implemented")
}
func (UnimplementedDIDCommSimplexServer) mustEmbedUnimplementedDIDCommSimplexServer() {}

// UnsafeDIDCommSimplexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DIDCommSimplexServer will
// result in compilation errors.
type UnsafeDIDCommSimplexServer interface {
	mustEmbedUnimplementedDIDCommSimplexServer()
}

func RegisterDIDCommSimplexServer(s grpc.ServiceRegistrar, srv DIDCommSimplexServer) {
	s.RegisterService(&_DIDCommSimplex_serviceDesc, srv)
}

func _DIDCommSimplex_Simplex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptedMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DIDCommSimplexServer).Simplex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/didcomm.messaging.DIDCommSimplex/Simplex",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DIDCommSimplexServer).Simplex(ctx, req.(*EncryptedMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _DIDCommSimplex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "didcomm.messaging.DIDCommSimplex",
	HandlerType: (*DIDCommSimplexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Simplex",
			Handler:    _DIDCommSimplex_Simplex_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "didcomm.proto",
}