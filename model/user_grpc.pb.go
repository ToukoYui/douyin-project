// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: idl/pb/user.proto

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

// UserSrvClient is the client API for UserSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSrvClient interface {
	Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error)
	Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error)
	GetUserInfo(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error)
}

type userSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSrvClient(cc grpc.ClientConnInterface) UserSrvClient {
	return &userSrvClient{cc}
}

func (c *userSrvClient) Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error) {
	out := new(DouyinUserRegisterResponse)
	err := c.cc.Invoke(ctx, "/user.UserSrv/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvClient) Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error) {
	out := new(DouyinUserLoginResponse)
	err := c.cc.Invoke(ctx, "/user.UserSrv/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSrvClient) GetUserInfo(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error) {
	out := new(DouyinUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserSrv/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSrvServer is the server API for UserSrv service.
// All implementations must embed UnimplementedUserSrvServer
// for forward compatibility
type UserSrvServer interface {
	Register(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error)
	Login(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error)
	GetUserInfo(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error)
	mustEmbedUnimplementedUserSrvServer()
}

// UnimplementedUserSrvServer must be embedded to have forward compatible implementations.
type UnimplementedUserSrvServer struct {
}

func (UnimplementedUserSrvServer) Register(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserSrvServer) Login(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserSrvServer) GetUserInfo(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserSrvServer) mustEmbedUnimplementedUserSrvServer() {}

// UnsafeUserSrvServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSrvServer will
// result in compilation errors.
type UnsafeUserSrvServer interface {
	mustEmbedUnimplementedUserSrvServer()
}

func RegisterUserSrvServer(s grpc.ServiceRegistrar, srv UserSrvServer) {
	s.RegisterService(&UserSrv_ServiceDesc, srv)
}

func _UserSrv_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSrv/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).Register(ctx, req.(*DouyinUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSrv_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSrv/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).Login(ctx, req.(*DouyinUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSrv_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSrvServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserSrv/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSrvServer).GetUserInfo(ctx, req.(*DouyinUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSrv_ServiceDesc is the grpc.ServiceDesc for UserSrv service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSrv_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserSrv",
	HandlerType: (*UserSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserSrv_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserSrv_Login_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserSrv_GetUserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/pb/user.proto",
}