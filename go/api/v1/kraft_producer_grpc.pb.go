// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api_v1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProducerServiceClient is the client API for ProducerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProducerServiceClient interface {
	Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type producerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProducerServiceClient(cc grpc.ClientConnInterface) ProducerServiceClient {
	return &producerServiceClient{cc}
}

func (c *producerServiceClient) Health(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.v1.ProducerService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *producerServiceClient) CreateEvent(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/api.v1.ProducerService/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProducerServiceServer is the server API for ProducerService service.
// All implementations must embed UnimplementedProducerServiceServer
// for forward compatibility
type ProducerServiceServer interface {
	Health(context.Context, *empty.Empty) (*empty.Empty, error)
	CreateEvent(context.Context, *EventRequest) (*EventResponse, error)
	mustEmbedUnimplementedProducerServiceServer()
}

// UnimplementedProducerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProducerServiceServer struct {
}

func (UnimplementedProducerServiceServer) Health(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedProducerServiceServer) CreateEvent(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (UnimplementedProducerServiceServer) mustEmbedUnimplementedProducerServiceServer() {}

// UnsafeProducerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProducerServiceServer will
// result in compilation errors.
type UnsafeProducerServiceServer interface {
	mustEmbedUnimplementedProducerServiceServer()
}

func RegisterProducerServiceServer(s grpc.ServiceRegistrar, srv ProducerServiceServer) {
	s.RegisterService(&ProducerService_ServiceDesc, srv)
}

func _ProducerService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProducerService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).Health(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProducerService_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProducerServiceServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.ProducerService/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProducerServiceServer).CreateEvent(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProducerService_ServiceDesc is the grpc.ServiceDesc for ProducerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProducerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.ProducerService",
	HandlerType: (*ProducerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Health",
			Handler:    _ProducerService_Health_Handler,
		},
		{
			MethodName: "CreateEvent",
			Handler:    _ProducerService_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api/v1/kraft_producer.proto",
}
