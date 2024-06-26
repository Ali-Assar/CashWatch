// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: types/expense.proto

package types

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

// ExpenseServiceClient is the client API for ExpenseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpenseServiceClient interface {
	InsertCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error)
	GetCategoryById(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Category, error)
	UpdateCategoryById(ctx context.Context, in *Category, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteCategoryById(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	InsertExpense(ctx context.Context, in *Expense, opts ...grpc.CallOption) (*Expense, error)
	GetExpenseById(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*Expense, error)
	UpdateExpenseById(ctx context.Context, in *Expense, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteExpenseById(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type expenseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpenseServiceClient(cc grpc.ClientConnInterface) ExpenseServiceClient {
	return &expenseServiceClient{cc}
}

func (c *expenseServiceClient) InsertCategory(ctx context.Context, in *Category, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/users.ExpenseService/InsertCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) GetCategoryById(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*Category, error) {
	out := new(Category)
	err := c.cc.Invoke(ctx, "/users.ExpenseService/GetCategoryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) UpdateCategoryById(ctx context.Context, in *Category, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/users.ExpenseService/UpdateCategoryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) DeleteCategoryById(ctx context.Context, in *CategoryRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/users.ExpenseService/DeleteCategoryById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) InsertExpense(ctx context.Context, in *Expense, opts ...grpc.CallOption) (*Expense, error) {
	out := new(Expense)
	err := c.cc.Invoke(ctx, "/users.ExpenseService/InsertExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) GetExpenseById(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*Expense, error) {
	out := new(Expense)
	err := c.cc.Invoke(ctx, "/users.ExpenseService/GetExpenseById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) UpdateExpenseById(ctx context.Context, in *Expense, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/users.ExpenseService/UpdateExpenseById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseServiceClient) DeleteExpenseById(ctx context.Context, in *ExpenseRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/users.ExpenseService/DeleteExpenseById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpenseServiceServer is the server API for ExpenseService service.
// All implementations must embed UnimplementedExpenseServiceServer
// for forward compatibility
type ExpenseServiceServer interface {
	InsertCategory(context.Context, *Category) (*Category, error)
	GetCategoryById(context.Context, *CategoryRequest) (*Category, error)
	UpdateCategoryById(context.Context, *Category) (*empty.Empty, error)
	DeleteCategoryById(context.Context, *CategoryRequest) (*empty.Empty, error)
	InsertExpense(context.Context, *Expense) (*Expense, error)
	GetExpenseById(context.Context, *ExpenseRequest) (*Expense, error)
	UpdateExpenseById(context.Context, *Expense) (*empty.Empty, error)
	DeleteExpenseById(context.Context, *ExpenseRequest) (*empty.Empty, error)
	mustEmbedUnimplementedExpenseServiceServer()
}

// UnimplementedExpenseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExpenseServiceServer struct {
}

func (UnimplementedExpenseServiceServer) InsertCategory(context.Context, *Category) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertCategory not implemented")
}
func (UnimplementedExpenseServiceServer) GetCategoryById(context.Context, *CategoryRequest) (*Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategoryById not implemented")
}
func (UnimplementedExpenseServiceServer) UpdateCategoryById(context.Context, *Category) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCategoryById not implemented")
}
func (UnimplementedExpenseServiceServer) DeleteCategoryById(context.Context, *CategoryRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategoryById not implemented")
}
func (UnimplementedExpenseServiceServer) InsertExpense(context.Context, *Expense) (*Expense, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertExpense not implemented")
}
func (UnimplementedExpenseServiceServer) GetExpenseById(context.Context, *ExpenseRequest) (*Expense, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpenseById not implemented")
}
func (UnimplementedExpenseServiceServer) UpdateExpenseById(context.Context, *Expense) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExpenseById not implemented")
}
func (UnimplementedExpenseServiceServer) DeleteExpenseById(context.Context, *ExpenseRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExpenseById not implemented")
}
func (UnimplementedExpenseServiceServer) mustEmbedUnimplementedExpenseServiceServer() {}

// UnsafeExpenseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpenseServiceServer will
// result in compilation errors.
type UnsafeExpenseServiceServer interface {
	mustEmbedUnimplementedExpenseServiceServer()
}

func RegisterExpenseServiceServer(s grpc.ServiceRegistrar, srv ExpenseServiceServer) {
	s.RegisterService(&ExpenseService_ServiceDesc, srv)
}

func _ExpenseService_InsertCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).InsertCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.ExpenseService/InsertCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).InsertCategory(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_GetCategoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).GetCategoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.ExpenseService/GetCategoryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).GetCategoryById(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_UpdateCategoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Category)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).UpdateCategoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.ExpenseService/UpdateCategoryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).UpdateCategoryById(ctx, req.(*Category))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_DeleteCategoryById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).DeleteCategoryById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.ExpenseService/DeleteCategoryById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).DeleteCategoryById(ctx, req.(*CategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_InsertExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Expense)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).InsertExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.ExpenseService/InsertExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).InsertExpense(ctx, req.(*Expense))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_GetExpenseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).GetExpenseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.ExpenseService/GetExpenseById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).GetExpenseById(ctx, req.(*ExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_UpdateExpenseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Expense)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).UpdateExpenseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.ExpenseService/UpdateExpenseById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).UpdateExpenseById(ctx, req.(*Expense))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseService_DeleteExpenseById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseServiceServer).DeleteExpenseById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.ExpenseService/DeleteExpenseById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseServiceServer).DeleteExpenseById(ctx, req.(*ExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpenseService_ServiceDesc is the grpc.ServiceDesc for ExpenseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpenseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.ExpenseService",
	HandlerType: (*ExpenseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertCategory",
			Handler:    _ExpenseService_InsertCategory_Handler,
		},
		{
			MethodName: "GetCategoryById",
			Handler:    _ExpenseService_GetCategoryById_Handler,
		},
		{
			MethodName: "UpdateCategoryById",
			Handler:    _ExpenseService_UpdateCategoryById_Handler,
		},
		{
			MethodName: "DeleteCategoryById",
			Handler:    _ExpenseService_DeleteCategoryById_Handler,
		},
		{
			MethodName: "InsertExpense",
			Handler:    _ExpenseService_InsertExpense_Handler,
		},
		{
			MethodName: "GetExpenseById",
			Handler:    _ExpenseService_GetExpenseById_Handler,
		},
		{
			MethodName: "UpdateExpenseById",
			Handler:    _ExpenseService_UpdateExpenseById_Handler,
		},
		{
			MethodName: "DeleteExpenseById",
			Handler:    _ExpenseService_DeleteExpenseById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types/expense.proto",
}
