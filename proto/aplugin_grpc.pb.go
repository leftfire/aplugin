// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ITSClient is the client API for ITS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ITSClient interface {
	UseStremFunc(ctx context.Context, in *FuncReq, opts ...grpc.CallOption) (ITS_UseStremFuncClient, error)
	UseFunc(ctx context.Context, in *FuncReq, opts ...grpc.CallOption) (*FuncResp, error)
}

type iTSClient struct {
	cc grpc.ClientConnInterface
}

func NewITSClient(cc grpc.ClientConnInterface) ITSClient {
	return &iTSClient{cc}
}

func (c *iTSClient) UseStremFunc(ctx context.Context, in *FuncReq, opts ...grpc.CallOption) (ITS_UseStremFuncClient, error) {
	stream, err := c.cc.NewStream(ctx, &ITS_ServiceDesc.Streams[0], "/proto.ITS/UseStremFunc", opts...)
	if err != nil {
		return nil, err
	}
	x := &iTSUseStremFuncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ITS_UseStremFuncClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type iTSUseStremFuncClient struct {
	grpc.ClientStream
}

func (x *iTSUseStremFuncClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *iTSClient) UseFunc(ctx context.Context, in *FuncReq, opts ...grpc.CallOption) (*FuncResp, error) {
	out := new(FuncResp)
	err := c.cc.Invoke(ctx, "/proto.ITS/UseFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ITSServer is the server API for ITS service.
// All implementations must embed UnimplementedITSServer
// for forward compatibility
type ITSServer interface {
	UseStremFunc(*FuncReq, ITS_UseStremFuncServer) error
	UseFunc(context.Context, *FuncReq) (*FuncResp, error)
	mustEmbedUnimplementedITSServer()
}

// UnimplementedITSServer must be embedded to have forward compatible implementations.
type UnimplementedITSServer struct {
}

func (UnimplementedITSServer) UseStremFunc(*FuncReq, ITS_UseStremFuncServer) error {
	return status.Errorf(codes.Unimplemented, "method UseStremFunc not implemented")
}
func (UnimplementedITSServer) UseFunc(context.Context, *FuncReq) (*FuncResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseFunc not implemented")
}
func (UnimplementedITSServer) mustEmbedUnimplementedITSServer() {}

// UnsafeITSServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ITSServer will
// result in compilation errors.
type UnsafeITSServer interface {
	mustEmbedUnimplementedITSServer()
}

func RegisterITSServer(s grpc.ServiceRegistrar, srv ITSServer) {
	s.RegisterService(&ITS_ServiceDesc, srv)
}

func _ITS_UseStremFunc_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FuncReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ITSServer).UseStremFunc(m, &iTSUseStremFuncServer{stream})
}

type ITS_UseStremFuncServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type iTSUseStremFuncServer struct {
	grpc.ServerStream
}

func (x *iTSUseStremFuncServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _ITS_UseFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ITSServer).UseFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ITS/UseFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ITSServer).UseFunc(ctx, req.(*FuncReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ITS_ServiceDesc is the grpc.ServiceDesc for ITS service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ITS_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ITS",
	HandlerType: (*ITSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UseFunc",
			Handler:    _ITS_UseFunc_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UseStremFunc",
			Handler:       _ITS_UseStremFunc_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "aplugin.proto",
}