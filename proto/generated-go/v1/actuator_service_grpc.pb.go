// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/actuator_service.proto

package v1

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

const (
	ActuatorService_GetActuatorInfo_FullMethodName = "/bytebase.v1.ActuatorService/GetActuatorInfo"
)

// ActuatorServiceClient is the client API for ActuatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActuatorServiceClient interface {
	GetActuatorInfo(ctx context.Context, in *GetActuatorInfoRequest, opts ...grpc.CallOption) (*ActuatorInfo, error)
}

type actuatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActuatorServiceClient(cc grpc.ClientConnInterface) ActuatorServiceClient {
	return &actuatorServiceClient{cc}
}

func (c *actuatorServiceClient) GetActuatorInfo(ctx context.Context, in *GetActuatorInfoRequest, opts ...grpc.CallOption) (*ActuatorInfo, error) {
	out := new(ActuatorInfo)
	err := c.cc.Invoke(ctx, ActuatorService_GetActuatorInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActuatorServiceServer is the server API for ActuatorService service.
// All implementations must embed UnimplementedActuatorServiceServer
// for forward compatibility
type ActuatorServiceServer interface {
	GetActuatorInfo(context.Context, *GetActuatorInfoRequest) (*ActuatorInfo, error)
	mustEmbedUnimplementedActuatorServiceServer()
}

// UnimplementedActuatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActuatorServiceServer struct {
}

func (UnimplementedActuatorServiceServer) GetActuatorInfo(context.Context, *GetActuatorInfoRequest) (*ActuatorInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActuatorInfo not implemented")
}
func (UnimplementedActuatorServiceServer) mustEmbedUnimplementedActuatorServiceServer() {}

// UnsafeActuatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActuatorServiceServer will
// result in compilation errors.
type UnsafeActuatorServiceServer interface {
	mustEmbedUnimplementedActuatorServiceServer()
}

func RegisterActuatorServiceServer(s grpc.ServiceRegistrar, srv ActuatorServiceServer) {
	s.RegisterService(&ActuatorService_ServiceDesc, srv)
}

func _ActuatorService_GetActuatorInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActuatorInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActuatorServiceServer).GetActuatorInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActuatorService_GetActuatorInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActuatorServiceServer).GetActuatorInfo(ctx, req.(*GetActuatorInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActuatorService_ServiceDesc is the grpc.ServiceDesc for ActuatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActuatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bytebase.v1.ActuatorService",
	HandlerType: (*ActuatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActuatorInfo",
			Handler:    _ActuatorService_GetActuatorInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/actuator_service.proto",
}
