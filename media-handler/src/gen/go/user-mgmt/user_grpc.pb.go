// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user-mgmt/user.proto

package user_mgmt

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

// UserMgmtClient is the client API for UserMgmt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserMgmtClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	InfoUpdate(ctx context.Context, in *InfoUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DummyResponse, error)
}

type userMgmtClient struct {
	cc grpc.ClientConnInterface
}

func NewUserMgmtClient(cc grpc.ClientConnInterface) UserMgmtClient {
	return &userMgmtClient{cc}
}

func (c *userMgmtClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user_mgmt.UserMgmt/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgmtClient) InfoUpdate(ctx context.Context, in *InfoUpdateRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user_mgmt.UserMgmt/InfoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgmtClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user_mgmt.UserMgmt/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userMgmtClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DummyResponse, error) {
	out := new(DummyResponse)
	err := c.cc.Invoke(ctx, "/user_mgmt.UserMgmt/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserMgmtServer is the server API for UserMgmt service.
// All implementations must embed UnimplementedUserMgmtServer
// for forward compatibility
type UserMgmtServer interface {
	AddUser(context.Context, *AddUserRequest) (*UserResponse, error)
	InfoUpdate(context.Context, *InfoUpdateRequest) (*UserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DummyResponse, error)
	mustEmbedUnimplementedUserMgmtServer()
}

// UnimplementedUserMgmtServer must be embedded to have forward compatible implementations.
type UnimplementedUserMgmtServer struct {
}

func (UnimplementedUserMgmtServer) AddUser(context.Context, *AddUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserMgmtServer) InfoUpdate(context.Context, *InfoUpdateRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoUpdate not implemented")
}
func (UnimplementedUserMgmtServer) GetUser(context.Context, *GetUserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserMgmtServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DummyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedUserMgmtServer) mustEmbedUnimplementedUserMgmtServer() {}

// UnsafeUserMgmtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserMgmtServer will
// result in compilation errors.
type UnsafeUserMgmtServer interface {
	mustEmbedUnimplementedUserMgmtServer()
}

func RegisterUserMgmtServer(s grpc.ServiceRegistrar, srv UserMgmtServer) {
	s.RegisterService(&UserMgmt_ServiceDesc, srv)
}

func _UserMgmt_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgmtServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_mgmt.UserMgmt/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgmtServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgmt_InfoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgmtServer).InfoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_mgmt.UserMgmt/InfoUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgmtServer).InfoUpdate(ctx, req.(*InfoUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgmt_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgmtServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_mgmt.UserMgmt/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgmtServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserMgmt_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserMgmtServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_mgmt.UserMgmt/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserMgmtServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserMgmt_ServiceDesc is the grpc.ServiceDesc for UserMgmt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserMgmt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_mgmt.UserMgmt",
	HandlerType: (*UserMgmtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserMgmt_AddUser_Handler,
		},
		{
			MethodName: "InfoUpdate",
			Handler:    _UserMgmt_InfoUpdate_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserMgmt_GetUser_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _UserMgmt_DeleteAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user-mgmt/user.proto",
}
