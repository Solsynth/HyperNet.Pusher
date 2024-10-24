// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: pusher.proto

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
	PusherService_PushNotification_FullMethodName      = "/proto.PusherService/PushNotification"
	PusherService_PushNotificationBatch_FullMethodName = "/proto.PusherService/PushNotificationBatch"
	PusherService_DeliverEmail_FullMethodName          = "/proto.PusherService/DeliverEmail"
	PusherService_DeliverEmailBatch_FullMethodName     = "/proto.PusherService/DeliverEmailBatch"
)

// PusherServiceClient is the client API for PusherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PusherServiceClient interface {
	PushNotification(ctx context.Context, in *PushNotificationRequest, opts ...grpc.CallOption) (*DeliveryResponse, error)
	PushNotificationBatch(ctx context.Context, in *PushNotificationBatchRequest, opts ...grpc.CallOption) (*DeliveryResponse, error)
	DeliverEmail(ctx context.Context, in *DeliverEmailRequest, opts ...grpc.CallOption) (*DeliveryResponse, error)
	DeliverEmailBatch(ctx context.Context, in *DeliverEmailBatchRequest, opts ...grpc.CallOption) (*DeliveryResponse, error)
}

type pusherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPusherServiceClient(cc grpc.ClientConnInterface) PusherServiceClient {
	return &pusherServiceClient{cc}
}

func (c *pusherServiceClient) PushNotification(ctx context.Context, in *PushNotificationRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, PusherService_PushNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pusherServiceClient) PushNotificationBatch(ctx context.Context, in *PushNotificationBatchRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, PusherService_PushNotificationBatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pusherServiceClient) DeliverEmail(ctx context.Context, in *DeliverEmailRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, PusherService_DeliverEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pusherServiceClient) DeliverEmailBatch(ctx context.Context, in *DeliverEmailBatchRequest, opts ...grpc.CallOption) (*DeliveryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResponse)
	err := c.cc.Invoke(ctx, PusherService_DeliverEmailBatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PusherServiceServer is the server API for PusherService service.
// All implementations must embed UnimplementedPusherServiceServer
// for forward compatibility.
type PusherServiceServer interface {
	PushNotification(context.Context, *PushNotificationRequest) (*DeliveryResponse, error)
	PushNotificationBatch(context.Context, *PushNotificationBatchRequest) (*DeliveryResponse, error)
	DeliverEmail(context.Context, *DeliverEmailRequest) (*DeliveryResponse, error)
	DeliverEmailBatch(context.Context, *DeliverEmailBatchRequest) (*DeliveryResponse, error)
	mustEmbedUnimplementedPusherServiceServer()
}

// UnimplementedPusherServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPusherServiceServer struct{}

func (UnimplementedPusherServiceServer) PushNotification(context.Context, *PushNotificationRequest) (*DeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushNotification not implemented")
}
func (UnimplementedPusherServiceServer) PushNotificationBatch(context.Context, *PushNotificationBatchRequest) (*DeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushNotificationBatch not implemented")
}
func (UnimplementedPusherServiceServer) DeliverEmail(context.Context, *DeliverEmailRequest) (*DeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverEmail not implemented")
}
func (UnimplementedPusherServiceServer) DeliverEmailBatch(context.Context, *DeliverEmailBatchRequest) (*DeliveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverEmailBatch not implemented")
}
func (UnimplementedPusherServiceServer) mustEmbedUnimplementedPusherServiceServer() {}
func (UnimplementedPusherServiceServer) testEmbeddedByValue()                       {}

// UnsafePusherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PusherServiceServer will
// result in compilation errors.
type UnsafePusherServiceServer interface {
	mustEmbedUnimplementedPusherServiceServer()
}

func RegisterPusherServiceServer(s grpc.ServiceRegistrar, srv PusherServiceServer) {
	// If the following call pancis, it indicates UnimplementedPusherServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PusherService_ServiceDesc, srv)
}

func _PusherService_PushNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServiceServer).PushNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PusherService_PushNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServiceServer).PushNotification(ctx, req.(*PushNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PusherService_PushNotificationBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushNotificationBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServiceServer).PushNotificationBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PusherService_PushNotificationBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServiceServer).PushNotificationBatch(ctx, req.(*PushNotificationBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PusherService_DeliverEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServiceServer).DeliverEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PusherService_DeliverEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServiceServer).DeliverEmail(ctx, req.(*DeliverEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PusherService_DeliverEmailBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverEmailBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PusherServiceServer).DeliverEmailBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PusherService_DeliverEmailBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PusherServiceServer).DeliverEmailBatch(ctx, req.(*DeliverEmailBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PusherService_ServiceDesc is the grpc.ServiceDesc for PusherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PusherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PusherService",
	HandlerType: (*PusherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushNotification",
			Handler:    _PusherService_PushNotification_Handler,
		},
		{
			MethodName: "PushNotificationBatch",
			Handler:    _PusherService_PushNotificationBatch_Handler,
		},
		{
			MethodName: "DeliverEmail",
			Handler:    _PusherService_DeliverEmail_Handler,
		},
		{
			MethodName: "DeliverEmailBatch",
			Handler:    _PusherService_DeliverEmailBatch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pusher.proto",
}