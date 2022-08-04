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

// MuntahClient is the client API for Muntah service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MuntahClient interface {
	GetMuntah(ctx context.Context, in *GetMuntahRequest, opts ...grpc.CallOption) (*GetMuntahResponse, error)
}

type muntahClient struct {
	cc grpc.ClientConnInterface
}

func NewMuntahClient(cc grpc.ClientConnInterface) MuntahClient {
	return &muntahClient{cc}
}

func (c *muntahClient) GetMuntah(ctx context.Context, in *GetMuntahRequest, opts ...grpc.CallOption) (*GetMuntahResponse, error) {
	out := new(GetMuntahResponse)
	err := c.cc.Invoke(ctx, "/Muntah/GetMuntah", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MuntahServer is the server API for Muntah service.
// All implementations must embed UnimplementedMuntahServer
// for forward compatibility
type MuntahServer interface {
	GetMuntah(context.Context, *GetMuntahRequest) (*GetMuntahResponse, error)
	mustEmbedUnimplementedMuntahServer()
}

// UnimplementedMuntahServer must be embedded to have forward compatible implementations.
type UnimplementedMuntahServer struct {
}

func (UnimplementedMuntahServer) GetMuntah(context.Context, *GetMuntahRequest) (*GetMuntahResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMuntah not implemented")
}
func (UnimplementedMuntahServer) mustEmbedUnimplementedMuntahServer() {}

// UnsafeMuntahServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MuntahServer will
// result in compilation errors.
type UnsafeMuntahServer interface {
	mustEmbedUnimplementedMuntahServer()
}

func RegisterMuntahServer(s grpc.ServiceRegistrar, srv MuntahServer) {
	s.RegisterService(&Muntah_ServiceDesc, srv)
}

func _Muntah_GetMuntah_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMuntahRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MuntahServer).GetMuntah(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Muntah/GetMuntah",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MuntahServer).GetMuntah(ctx, req.(*GetMuntahRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Muntah_ServiceDesc is the grpc.ServiceDesc for Muntah service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Muntah_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Muntah",
	HandlerType: (*MuntahServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMuntah",
			Handler:    _Muntah_GetMuntah_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/muntah.proto",
}