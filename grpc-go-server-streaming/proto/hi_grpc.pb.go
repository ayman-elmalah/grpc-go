// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/hi.proto

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

// HiServiceClient is the client API for HiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HiServiceClient interface {
	SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (HiService_SayHiClient, error)
}

type hiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHiServiceClient(cc grpc.ClientConnInterface) HiServiceClient {
	return &hiServiceClient{cc}
}

func (c *hiServiceClient) SayHi(ctx context.Context, in *HiRequest, opts ...grpc.CallOption) (HiService_SayHiClient, error) {
	stream, err := c.cc.NewStream(ctx, &HiService_ServiceDesc.Streams[0], "/hi.HiService/SayHi", opts...)
	if err != nil {
		return nil, err
	}
	x := &hiServiceSayHiClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HiService_SayHiClient interface {
	Recv() (*HiResponse, error)
	grpc.ClientStream
}

type hiServiceSayHiClient struct {
	grpc.ClientStream
}

func (x *hiServiceSayHiClient) Recv() (*HiResponse, error) {
	m := new(HiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HiServiceServer is the server API for HiService service.
// All implementations must embed UnimplementedHiServiceServer
// for forward compatibility
type HiServiceServer interface {
	SayHi(*HiRequest, HiService_SayHiServer) error
	mustEmbedUnimplementedHiServiceServer()
}

// UnimplementedHiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHiServiceServer struct {
}

func (UnimplementedHiServiceServer) SayHi(*HiRequest, HiService_SayHiServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (UnimplementedHiServiceServer) mustEmbedUnimplementedHiServiceServer() {}

// UnsafeHiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HiServiceServer will
// result in compilation errors.
type UnsafeHiServiceServer interface {
	mustEmbedUnimplementedHiServiceServer()
}

func RegisterHiServiceServer(s grpc.ServiceRegistrar, srv HiServiceServer) {
	s.RegisterService(&HiService_ServiceDesc, srv)
}

func _HiService_SayHi_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HiRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HiServiceServer).SayHi(m, &hiServiceSayHiServer{stream})
}

type HiService_SayHiServer interface {
	Send(*HiResponse) error
	grpc.ServerStream
}

type hiServiceSayHiServer struct {
	grpc.ServerStream
}

func (x *hiServiceSayHiServer) Send(m *HiResponse) error {
	return x.ServerStream.SendMsg(m)
}

// HiService_ServiceDesc is the grpc.ServiceDesc for HiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hi.HiService",
	HandlerType: (*HiServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHi",
			Handler:       _HiService_SayHi_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/hi.proto",
}
