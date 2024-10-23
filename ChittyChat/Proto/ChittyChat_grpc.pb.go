// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: Proto/ChittyChat.proto

package chittychatproto

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
	ChittyChatService_EnterChat_FullMethodName = "/chittychatproto.ChittyChatService/EnterChat"
)

// ChittyChatServiceClient is the client API for ChittyChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatServiceClient interface {
	EnterChat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[MessageRequest, MessageReply], error)
}

type chittyChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatServiceClient(cc grpc.ClientConnInterface) ChittyChatServiceClient {
	return &chittyChatServiceClient{cc}
}

func (c *chittyChatServiceClient) EnterChat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[MessageRequest, MessageReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[0], ChittyChatService_EnterChat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MessageRequest, MessageReply]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChittyChatService_EnterChatClient = grpc.BidiStreamingClient[MessageRequest, MessageReply]

// ChittyChatServiceServer is the server API for ChittyChatService service.
// All implementations must embed UnimplementedChittyChatServiceServer
// for forward compatibility.
type ChittyChatServiceServer interface {
	EnterChat(grpc.BidiStreamingServer[MessageRequest, MessageReply]) error
	mustEmbedUnimplementedChittyChatServiceServer()
}

// UnimplementedChittyChatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChittyChatServiceServer struct{}

func (UnimplementedChittyChatServiceServer) EnterChat(grpc.BidiStreamingServer[MessageRequest, MessageReply]) error {
	return status.Errorf(codes.Unimplemented, "method EnterChat not implemented")
}
func (UnimplementedChittyChatServiceServer) mustEmbedUnimplementedChittyChatServiceServer() {}
func (UnimplementedChittyChatServiceServer) testEmbeddedByValue()                           {}

// UnsafeChittyChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServiceServer will
// result in compilation errors.
type UnsafeChittyChatServiceServer interface {
	mustEmbedUnimplementedChittyChatServiceServer()
}

func RegisterChittyChatServiceServer(s grpc.ServiceRegistrar, srv ChittyChatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChittyChatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChittyChatService_ServiceDesc, srv)
}

func _ChittyChatService_EnterChat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChittyChatServiceServer).EnterChat(&grpc.GenericServerStream[MessageRequest, MessageReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChittyChatService_EnterChatServer = grpc.BidiStreamingServer[MessageRequest, MessageReply]

// ChittyChatService_ServiceDesc is the grpc.ServiceDesc for ChittyChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chittychatproto.ChittyChatService",
	HandlerType: (*ChittyChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EnterChat",
			Handler:       _ChittyChatService_EnterChat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "Proto/ChittyChat.proto",
}