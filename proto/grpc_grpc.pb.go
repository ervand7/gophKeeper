// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/grpc.proto

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

// GophKeeperClient is the client API for GophKeeper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GophKeeperClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	StreamDB(ctx context.Context, in *StreamDBRequest, opts ...grpc.CallOption) (GophKeeper_StreamDBClient, error)
	CreateCredentials(ctx context.Context, in *CreateCredentialsRequest, opts ...grpc.CallOption) (*CreateCredentialsResponse, error)
	CreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error)
	CreateBinaryData(ctx context.Context, in *CreateBinaryDataRequest, opts ...grpc.CallOption) (*CreateBinaryDataResponse, error)
	CreateBankCard(ctx context.Context, in *CreateBankCardRequest, opts ...grpc.CallOption) (*CreateBankCardResponse, error)
	UpdateCredentials(ctx context.Context, in *UpdateCredentialsRequest, opts ...grpc.CallOption) (*UpdateCredentialsResponse, error)
	UpdateText(ctx context.Context, in *UpdateTextRequest, opts ...grpc.CallOption) (*UpdateTextResponse, error)
	UpdateBinaryData(ctx context.Context, in *UpdateBinaryDataRequest, opts ...grpc.CallOption) (*UpdateBinaryDataResponse, error)
	UpdateBankCard(ctx context.Context, in *UpdateBankCardRequest, opts ...grpc.CallOption) (*UpdateBankCardResponse, error)
	DeleteCredentials(ctx context.Context, in *DeleteCredentialsRequest, opts ...grpc.CallOption) (*DeleteCredentialsResponse, error)
	DeleteText(ctx context.Context, in *DeleteTextRequest, opts ...grpc.CallOption) (*DeleteTextResponse, error)
	DeleteBinaryData(ctx context.Context, in *DeleteBinaryDataRequest, opts ...grpc.CallOption) (*DeleteBinaryDataResponse, error)
	DeleteBankCard(ctx context.Context, in *DeleteBankCardRequest, opts ...grpc.CallOption) (*DeleteBankCardResponse, error)
}

type gophKeeperClient struct {
	cc grpc.ClientConnInterface
}

func NewGophKeeperClient(cc grpc.ClientConnInterface) GophKeeperClient {
	return &gophKeeperClient{cc}
}

func (c *gophKeeperClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) StreamDB(ctx context.Context, in *StreamDBRequest, opts ...grpc.CallOption) (GophKeeper_StreamDBClient, error) {
	stream, err := c.cc.NewStream(ctx, &GophKeeper_ServiceDesc.Streams[0], "/gophkeeper.GophKeeper/StreamDB", opts...)
	if err != nil {
		return nil, err
	}
	x := &gophKeeperStreamDBClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GophKeeper_StreamDBClient interface {
	Recv() (*StreamDBResponse, error)
	grpc.ClientStream
}

type gophKeeperStreamDBClient struct {
	grpc.ClientStream
}

func (x *gophKeeperStreamDBClient) Recv() (*StreamDBResponse, error) {
	m := new(StreamDBResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gophKeeperClient) CreateCredentials(ctx context.Context, in *CreateCredentialsRequest, opts ...grpc.CallOption) (*CreateCredentialsResponse, error) {
	out := new(CreateCredentialsResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/CreateCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) CreateText(ctx context.Context, in *CreateTextRequest, opts ...grpc.CallOption) (*CreateTextResponse, error) {
	out := new(CreateTextResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/CreateText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) CreateBinaryData(ctx context.Context, in *CreateBinaryDataRequest, opts ...grpc.CallOption) (*CreateBinaryDataResponse, error) {
	out := new(CreateBinaryDataResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/CreateBinaryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) CreateBankCard(ctx context.Context, in *CreateBankCardRequest, opts ...grpc.CallOption) (*CreateBankCardResponse, error) {
	out := new(CreateBankCardResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/CreateBankCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UpdateCredentials(ctx context.Context, in *UpdateCredentialsRequest, opts ...grpc.CallOption) (*UpdateCredentialsResponse, error) {
	out := new(UpdateCredentialsResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/UpdateCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UpdateText(ctx context.Context, in *UpdateTextRequest, opts ...grpc.CallOption) (*UpdateTextResponse, error) {
	out := new(UpdateTextResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/UpdateText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UpdateBinaryData(ctx context.Context, in *UpdateBinaryDataRequest, opts ...grpc.CallOption) (*UpdateBinaryDataResponse, error) {
	out := new(UpdateBinaryDataResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/UpdateBinaryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) UpdateBankCard(ctx context.Context, in *UpdateBankCardRequest, opts ...grpc.CallOption) (*UpdateBankCardResponse, error) {
	out := new(UpdateBankCardResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/UpdateBankCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteCredentials(ctx context.Context, in *DeleteCredentialsRequest, opts ...grpc.CallOption) (*DeleteCredentialsResponse, error) {
	out := new(DeleteCredentialsResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/DeleteCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteText(ctx context.Context, in *DeleteTextRequest, opts ...grpc.CallOption) (*DeleteTextResponse, error) {
	out := new(DeleteTextResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/DeleteText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteBinaryData(ctx context.Context, in *DeleteBinaryDataRequest, opts ...grpc.CallOption) (*DeleteBinaryDataResponse, error) {
	out := new(DeleteBinaryDataResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/DeleteBinaryData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gophKeeperClient) DeleteBankCard(ctx context.Context, in *DeleteBankCardRequest, opts ...grpc.CallOption) (*DeleteBankCardResponse, error) {
	out := new(DeleteBankCardResponse)
	err := c.cc.Invoke(ctx, "/gophkeeper.GophKeeper/DeleteBankCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GophKeeperServer is the server API for GophKeeper service.
// All implementations must embed UnimplementedGophKeeperServer
// for forward compatibility
type GophKeeperServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	StreamDB(*StreamDBRequest, GophKeeper_StreamDBServer) error
	CreateCredentials(context.Context, *CreateCredentialsRequest) (*CreateCredentialsResponse, error)
	CreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error)
	CreateBinaryData(context.Context, *CreateBinaryDataRequest) (*CreateBinaryDataResponse, error)
	CreateBankCard(context.Context, *CreateBankCardRequest) (*CreateBankCardResponse, error)
	UpdateCredentials(context.Context, *UpdateCredentialsRequest) (*UpdateCredentialsResponse, error)
	UpdateText(context.Context, *UpdateTextRequest) (*UpdateTextResponse, error)
	UpdateBinaryData(context.Context, *UpdateBinaryDataRequest) (*UpdateBinaryDataResponse, error)
	UpdateBankCard(context.Context, *UpdateBankCardRequest) (*UpdateBankCardResponse, error)
	DeleteCredentials(context.Context, *DeleteCredentialsRequest) (*DeleteCredentialsResponse, error)
	DeleteText(context.Context, *DeleteTextRequest) (*DeleteTextResponse, error)
	DeleteBinaryData(context.Context, *DeleteBinaryDataRequest) (*DeleteBinaryDataResponse, error)
	DeleteBankCard(context.Context, *DeleteBankCardRequest) (*DeleteBankCardResponse, error)
	mustEmbedUnimplementedGophKeeperServer()
}

// UnimplementedGophKeeperServer must be embedded to have forward compatible implementations.
type UnimplementedGophKeeperServer struct {
}

func (UnimplementedGophKeeperServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGophKeeperServer) Auth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedGophKeeperServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedGophKeeperServer) StreamDB(*StreamDBRequest, GophKeeper_StreamDBServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamDB not implemented")
}
func (UnimplementedGophKeeperServer) CreateCredentials(context.Context, *CreateCredentialsRequest) (*CreateCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCredentials not implemented")
}
func (UnimplementedGophKeeperServer) CreateText(context.Context, *CreateTextRequest) (*CreateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateText not implemented")
}
func (UnimplementedGophKeeperServer) CreateBinaryData(context.Context, *CreateBinaryDataRequest) (*CreateBinaryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBinaryData not implemented")
}
func (UnimplementedGophKeeperServer) CreateBankCard(context.Context, *CreateBankCardRequest) (*CreateBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBankCard not implemented")
}
func (UnimplementedGophKeeperServer) UpdateCredentials(context.Context, *UpdateCredentialsRequest) (*UpdateCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCredentials not implemented")
}
func (UnimplementedGophKeeperServer) UpdateText(context.Context, *UpdateTextRequest) (*UpdateTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateText not implemented")
}
func (UnimplementedGophKeeperServer) UpdateBinaryData(context.Context, *UpdateBinaryDataRequest) (*UpdateBinaryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBinaryData not implemented")
}
func (UnimplementedGophKeeperServer) UpdateBankCard(context.Context, *UpdateBankCardRequest) (*UpdateBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBankCard not implemented")
}
func (UnimplementedGophKeeperServer) DeleteCredentials(context.Context, *DeleteCredentialsRequest) (*DeleteCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCredentials not implemented")
}
func (UnimplementedGophKeeperServer) DeleteText(context.Context, *DeleteTextRequest) (*DeleteTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteText not implemented")
}
func (UnimplementedGophKeeperServer) DeleteBinaryData(context.Context, *DeleteBinaryDataRequest) (*DeleteBinaryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBinaryData not implemented")
}
func (UnimplementedGophKeeperServer) DeleteBankCard(context.Context, *DeleteBankCardRequest) (*DeleteBankCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBankCard not implemented")
}
func (UnimplementedGophKeeperServer) mustEmbedUnimplementedGophKeeperServer() {}

// UnsafeGophKeeperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GophKeeperServer will
// result in compilation errors.
type UnsafeGophKeeperServer interface {
	mustEmbedUnimplementedGophKeeperServer()
}

func RegisterGophKeeperServer(s grpc.ServiceRegistrar, srv GophKeeperServer) {
	s.RegisterService(&GophKeeper_ServiceDesc, srv)
}

func _GophKeeper_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_StreamDB_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamDBRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GophKeeperServer).StreamDB(m, &gophKeeperStreamDBServer{stream})
}

type GophKeeper_StreamDBServer interface {
	Send(*StreamDBResponse) error
	grpc.ServerStream
}

type gophKeeperStreamDBServer struct {
	grpc.ServerStream
}

func (x *gophKeeperStreamDBServer) Send(m *StreamDBResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GophKeeper_CreateCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).CreateCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/CreateCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).CreateCredentials(ctx, req.(*CreateCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_CreateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).CreateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/CreateText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).CreateText(ctx, req.(*CreateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_CreateBinaryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBinaryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).CreateBinaryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/CreateBinaryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).CreateBinaryData(ctx, req.(*CreateBinaryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_CreateBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).CreateBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/CreateBankCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).CreateBankCard(ctx, req.(*CreateBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UpdateCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).UpdateCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/UpdateCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).UpdateCredentials(ctx, req.(*UpdateCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UpdateText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).UpdateText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/UpdateText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).UpdateText(ctx, req.(*UpdateTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UpdateBinaryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBinaryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).UpdateBinaryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/UpdateBinaryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).UpdateBinaryData(ctx, req.(*UpdateBinaryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_UpdateBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).UpdateBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/UpdateBankCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).UpdateBankCard(ctx, req.(*UpdateBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/DeleteCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteCredentials(ctx, req.(*DeleteCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/DeleteText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteText(ctx, req.(*DeleteTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteBinaryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBinaryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteBinaryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/DeleteBinaryData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteBinaryData(ctx, req.(*DeleteBinaryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GophKeeper_DeleteBankCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBankCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GophKeeperServer).DeleteBankCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gophkeeper.GophKeeper/DeleteBankCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GophKeeperServer).DeleteBankCard(ctx, req.(*DeleteBankCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GophKeeper_ServiceDesc is the grpc.ServiceDesc for GophKeeper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GophKeeper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gophkeeper.GophKeeper",
	HandlerType: (*GophKeeperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _GophKeeper_Register_Handler,
		},
		{
			MethodName: "Auth",
			Handler:    _GophKeeper_Auth_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _GophKeeper_DeleteUser_Handler,
		},
		{
			MethodName: "CreateCredentials",
			Handler:    _GophKeeper_CreateCredentials_Handler,
		},
		{
			MethodName: "CreateText",
			Handler:    _GophKeeper_CreateText_Handler,
		},
		{
			MethodName: "CreateBinaryData",
			Handler:    _GophKeeper_CreateBinaryData_Handler,
		},
		{
			MethodName: "CreateBankCard",
			Handler:    _GophKeeper_CreateBankCard_Handler,
		},
		{
			MethodName: "UpdateCredentials",
			Handler:    _GophKeeper_UpdateCredentials_Handler,
		},
		{
			MethodName: "UpdateText",
			Handler:    _GophKeeper_UpdateText_Handler,
		},
		{
			MethodName: "UpdateBinaryData",
			Handler:    _GophKeeper_UpdateBinaryData_Handler,
		},
		{
			MethodName: "UpdateBankCard",
			Handler:    _GophKeeper_UpdateBankCard_Handler,
		},
		{
			MethodName: "DeleteCredentials",
			Handler:    _GophKeeper_DeleteCredentials_Handler,
		},
		{
			MethodName: "DeleteText",
			Handler:    _GophKeeper_DeleteText_Handler,
		},
		{
			MethodName: "DeleteBinaryData",
			Handler:    _GophKeeper_DeleteBinaryData_Handler,
		},
		{
			MethodName: "DeleteBankCard",
			Handler:    _GophKeeper_DeleteBankCard_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamDB",
			Handler:       _GophKeeper_StreamDB_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/grpc.proto",
}