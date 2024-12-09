// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: server.proto

package server

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
	AuctionService_Register_FullMethodName     = "/server.AuctionService/Register"
	AuctionService_Login_FullMethodName        = "/server.AuctionService/Login"
	AuctionService_CreateLot_FullMethodName    = "/server.AuctionService/CreateLot"
	AuctionService_StartAuction_FullMethodName = "/server.AuctionService/StartAuction"
	AuctionService_PlaceBid_FullMethodName     = "/server.AuctionService/PlaceBid"
	AuctionService_Pay_FullMethodName          = "/server.AuctionService/Pay"
)

// AuctionServiceClient is the client API for AuctionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для управления аукционами
type AuctionServiceClient interface {
	// Метод для регистрации пользователя
	Register(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*JWT, error)
	// Метод для входа пользователя
	Login(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*JWT, error)
	// Метод для создания лота
	CreateLot(ctx context.Context, in *Lot, opts ...grpc.CallOption) (*ID, error)
	// Метод для начала аукциона
	StartAuction(ctx context.Context, in *Auction, opts ...grpc.CallOption) (*ID, error)
	// Метод для размещения ставки
	PlaceBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*ID, error)
	// Метод для оплаты
	Pay(ctx context.Context, in *Money, opts ...grpc.CallOption) (*Empty, error)
}

type auctionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionServiceClient(cc grpc.ClientConnInterface) AuctionServiceClient {
	return &auctionServiceClient{cc}
}

func (c *auctionServiceClient) Register(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*JWT, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JWT)
	err := c.cc.Invoke(ctx, AuctionService_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) Login(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*JWT, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JWT)
	err := c.cc.Invoke(ctx, AuctionService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) CreateLot(ctx context.Context, in *Lot, opts ...grpc.CallOption) (*ID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ID)
	err := c.cc.Invoke(ctx, AuctionService_CreateLot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) StartAuction(ctx context.Context, in *Auction, opts ...grpc.CallOption) (*ID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ID)
	err := c.cc.Invoke(ctx, AuctionService_StartAuction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) PlaceBid(ctx context.Context, in *Bid, opts ...grpc.CallOption) (*ID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ID)
	err := c.cc.Invoke(ctx, AuctionService_PlaceBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) Pay(ctx context.Context, in *Money, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, AuctionService_Pay_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServiceServer is the server API for AuctionService service.
// All implementations must embed UnimplementedAuctionServiceServer
// for forward compatibility.
//
// Сервис для управления аукционами
type AuctionServiceServer interface {
	// Метод для регистрации пользователя
	Register(context.Context, *UserData) (*JWT, error)
	// Метод для входа пользователя
	Login(context.Context, *UserData) (*JWT, error)
	// Метод для создания лота
	CreateLot(context.Context, *Lot) (*ID, error)
	// Метод для начала аукциона
	StartAuction(context.Context, *Auction) (*ID, error)
	// Метод для размещения ставки
	PlaceBid(context.Context, *Bid) (*ID, error)
	// Метод для оплаты
	Pay(context.Context, *Money) (*Empty, error)
	mustEmbedUnimplementedAuctionServiceServer()
}

// UnimplementedAuctionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuctionServiceServer struct{}

func (UnimplementedAuctionServiceServer) Register(context.Context, *UserData) (*JWT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuctionServiceServer) Login(context.Context, *UserData) (*JWT, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuctionServiceServer) CreateLot(context.Context, *Lot) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLot not implemented")
}
func (UnimplementedAuctionServiceServer) StartAuction(context.Context, *Auction) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartAuction not implemented")
}
func (UnimplementedAuctionServiceServer) PlaceBid(context.Context, *Bid) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBid not implemented")
}
func (UnimplementedAuctionServiceServer) Pay(context.Context, *Money) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}
func (UnimplementedAuctionServiceServer) mustEmbedUnimplementedAuctionServiceServer() {}
func (UnimplementedAuctionServiceServer) testEmbeddedByValue()                        {}

// UnsafeAuctionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServiceServer will
// result in compilation errors.
type UnsafeAuctionServiceServer interface {
	mustEmbedUnimplementedAuctionServiceServer()
}

func RegisterAuctionServiceServer(s grpc.ServiceRegistrar, srv AuctionServiceServer) {
	// If the following call panics, it indicates UnimplementedAuctionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuctionService_ServiceDesc, srv)
}

func _AuctionService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Register(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Login(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_CreateLot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Lot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).CreateLot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_CreateLot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).CreateLot(ctx, req.(*Lot))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_StartAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Auction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).StartAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_StartAuction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).StartAuction(ctx, req.(*Auction))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_PlaceBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).PlaceBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_PlaceBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).PlaceBid(ctx, req.(*Bid))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Money)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_Pay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).Pay(ctx, req.(*Money))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionService_ServiceDesc is the grpc.ServiceDesc for AuctionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.AuctionService",
	HandlerType: (*AuctionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _AuctionService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _AuctionService_Login_Handler,
		},
		{
			MethodName: "CreateLot",
			Handler:    _AuctionService_CreateLot_Handler,
		},
		{
			MethodName: "StartAuction",
			Handler:    _AuctionService_StartAuction_Handler,
		},
		{
			MethodName: "PlaceBid",
			Handler:    _AuctionService_PlaceBid_Handler,
		},
		{
			MethodName: "Pay",
			Handler:    _AuctionService_Pay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}