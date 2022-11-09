// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: chat.proto

package chat

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	ExchangeMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_ExchangeMessageClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ExchangeMessage(ctx context.Context, opts ...grpc.CallOption) (ChatService_ExchangeMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/ChatService/ExchangeMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceExchangeMessageClient{stream}
	return x, nil
}

type ChatService_ExchangeMessageClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServiceExchangeMessageClient struct {
	grpc.ClientStream
}

func (x *chatServiceExchangeMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceExchangeMessageClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	ExchangeMessage(ChatService_ExchangeMessageServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) ExchangeMessage(ChatService_ExchangeMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method ExchangeMessage not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_ExchangeMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).ExchangeMessage(&chatServiceExchangeMessageServer{stream})
}

type ChatService_ExchangeMessageServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServiceExchangeMessageServer struct {
	grpc.ServerStream
}

func (x *chatServiceExchangeMessageServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceExchangeMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExchangeMessage",
			Handler:       _ChatService_ExchangeMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
