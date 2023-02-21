// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: idl/pb/video.proto

package model

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

// FeedSrvClient is the client API for FeedSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FeedSrvClient interface {
	GetUserFeed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error)
	// rpc GetVideoById (video_id_request) returns (Video) {}
	PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error)
	PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error)
}

type feedSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedSrvClient(cc grpc.ClientConnInterface) FeedSrvClient {
	return &feedSrvClient{cc}
}

func (c *feedSrvClient) GetUserFeed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error) {
	out := new(DouyinFeedResponse)
	err := c.cc.Invoke(ctx, "/video.FeedSrv/GetUserFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedSrvClient) PublishAction(ctx context.Context, in *DouyinPublishActionRequest, opts ...grpc.CallOption) (*DouyinPublishActionResponse, error) {
	out := new(DouyinPublishActionResponse)
	err := c.cc.Invoke(ctx, "/video.FeedSrv/PublishAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedSrvClient) PublishList(ctx context.Context, in *DouyinPublishListRequest, opts ...grpc.CallOption) (*DouyinPublishListResponse, error) {
	out := new(DouyinPublishListResponse)
	err := c.cc.Invoke(ctx, "/video.FeedSrv/PublishList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedSrvServer is the server API for FeedSrv service.
// All implementations must embed UnimplementedFeedSrvServer
// for forward compatibility
type FeedSrvServer interface {
	GetUserFeed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error)
	// rpc GetVideoById (video_id_request) returns (Video) {}
	PublishAction(context.Context, *DouyinPublishActionRequest) (*DouyinPublishActionResponse, error)
	PublishList(context.Context, *DouyinPublishListRequest) (*DouyinPublishListResponse, error)
	mustEmbedUnimplementedFeedSrvServer()
}

// UnimplementedFeedSrvServer must be embedded to have forward compatible implementations.
type UnimplementedFeedSrvServer struct {
}

func (UnimplementedFeedSrvServer) GetUserFeed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFeed not implemented")
}
func (UnimplementedFeedSrvServer) PublishAction(context.Context, *DouyinPublishActionRequest) (*DouyinPublishActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishAction not implemented")
}
func (UnimplementedFeedSrvServer) PublishList(context.Context, *DouyinPublishListRequest) (*DouyinPublishListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishList not implemented")
}
func (UnimplementedFeedSrvServer) mustEmbedUnimplementedFeedSrvServer() {}

// UnsafeFeedSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FeedSrvServer will
// result in compilation errors.
type UnsafeFeedSrvServer interface {
	mustEmbedUnimplementedFeedSrvServer()
}

func RegisterFeedSrvServer(s grpc.ServiceRegistrar, srv FeedSrvServer) {
	s.RegisterService(&FeedSrv_ServiceDesc, srv)
}

func _FeedSrv_GetUserFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedSrvServer).GetUserFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.FeedSrv/GetUserFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedSrvServer).GetUserFeed(ctx, req.(*DouyinFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedSrv_PublishAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinPublishActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedSrvServer).PublishAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.FeedSrv/PublishAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedSrvServer).PublishAction(ctx, req.(*DouyinPublishActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedSrv_PublishList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinPublishListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedSrvServer).PublishList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/video.FeedSrv/PublishList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedSrvServer).PublishList(ctx, req.(*DouyinPublishListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FeedSrv_ServiceDesc is the grpc.ServiceDesc for FeedSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FeedSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.FeedSrv",
	HandlerType: (*FeedSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserFeed",
			Handler:    _FeedSrv_GetUserFeed_Handler,
		},
		{
			MethodName: "PublishAction",
			Handler:    _FeedSrv_PublishAction_Handler,
		},
		{
			MethodName: "PublishList",
			Handler:    _FeedSrv_PublishList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/pb/video.proto",
}