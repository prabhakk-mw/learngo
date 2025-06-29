// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: proto/capservice.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CapService_Capitalize_FullMethodName = "/capitalize_service.CapService/Capitalize"
)

// CapServiceClient is the client API for CapService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CapServiceClient interface {
	Capitalize(ctx context.Context, in *CapRequest, opts ...grpc.CallOption) (*CapResponse, error)
}

type capServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCapServiceClient(cc grpc.ClientConnInterface) CapServiceClient {
	return &capServiceClient{cc}
}

func (c *capServiceClient) Capitalize(ctx context.Context, in *CapRequest, opts ...grpc.CallOption) (*CapResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CapResponse)
	err := c.cc.Invoke(ctx, CapService_Capitalize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CapServiceServer is the server API for CapService service.
// All implementations must embed UnimplementedCapServiceServer
// for forward compatibility.
type CapServiceServer interface {
	Capitalize(context.Context, *CapRequest) (*CapResponse, error)
	mustEmbedUnimplementedCapServiceServer()
}

// UnimplementedCapServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCapServiceServer struct{}

func (UnimplementedCapServiceServer) Capitalize(context.Context, *CapRequest) (*CapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Capitalize not implemented")
}
func (UnimplementedCapServiceServer) mustEmbedUnimplementedCapServiceServer() {}
func (UnimplementedCapServiceServer) testEmbeddedByValue()                    {}

// UnsafeCapServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CapServiceServer will
// result in compilation errors.
type UnsafeCapServiceServer interface {
	mustEmbedUnimplementedCapServiceServer()
}

func RegisterCapServiceServer(s grpc.ServiceRegistrar, srv CapServiceServer) {
	// If the following call pancis, it indicates UnimplementedCapServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CapService_ServiceDesc, srv)
}

func _CapService_Capitalize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapServiceServer).Capitalize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapService_Capitalize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapServiceServer).Capitalize(ctx, req.(*CapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CapService_ServiceDesc is the grpc.ServiceDesc for CapService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CapService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "capitalize_service.CapService",
	HandlerType: (*CapServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Capitalize",
			Handler:    _CapService_Capitalize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/capservice.proto",
}
