// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: v1/save-service.proto

package v1

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

const (
	SaveService_CreateSave_FullMethodName = "/save.v1.SaveService/CreateSave"
	SaveService_GetSave_FullMethodName    = "/save.v1.SaveService/GetSave"
	SaveService_DeleteSave_FullMethodName = "/save.v1.SaveService/DeleteSave"
)

// SaveServiceClient is the client API for SaveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SaveServiceClient interface {
	CreateSave(ctx context.Context, in *CreateSaveRequest, opts ...grpc.CallOption) (*CreateSaveResponse, error)
	GetSave(ctx context.Context, in *GetSaveRequest, opts ...grpc.CallOption) (*GetSaveResponse, error)
	DeleteSave(ctx context.Context, in *DeleteSaveRequest, opts ...grpc.CallOption) (*DeleteSaveResponse, error)
}

type saveServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSaveServiceClient(cc grpc.ClientConnInterface) SaveServiceClient {
	return &saveServiceClient{cc}
}

func (c *saveServiceClient) CreateSave(ctx context.Context, in *CreateSaveRequest, opts ...grpc.CallOption) (*CreateSaveResponse, error) {
	out := new(CreateSaveResponse)
	err := c.cc.Invoke(ctx, SaveService_CreateSave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saveServiceClient) GetSave(ctx context.Context, in *GetSaveRequest, opts ...grpc.CallOption) (*GetSaveResponse, error) {
	out := new(GetSaveResponse)
	err := c.cc.Invoke(ctx, SaveService_GetSave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *saveServiceClient) DeleteSave(ctx context.Context, in *DeleteSaveRequest, opts ...grpc.CallOption) (*DeleteSaveResponse, error) {
	out := new(DeleteSaveResponse)
	err := c.cc.Invoke(ctx, SaveService_DeleteSave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SaveServiceServer is the server API for SaveService service.
// All implementations must embed UnimplementedSaveServiceServer
// for forward compatibility
type SaveServiceServer interface {
	CreateSave(context.Context, *CreateSaveRequest) (*CreateSaveResponse, error)
	GetSave(context.Context, *GetSaveRequest) (*GetSaveResponse, error)
	DeleteSave(context.Context, *DeleteSaveRequest) (*DeleteSaveResponse, error)
	mustEmbedUnimplementedSaveServiceServer()
}

// UnimplementedSaveServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSaveServiceServer struct {
}

func (UnimplementedSaveServiceServer) CreateSave(context.Context, *CreateSaveRequest) (*CreateSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSave not implemented")
}
func (UnimplementedSaveServiceServer) GetSave(context.Context, *GetSaveRequest) (*GetSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSave not implemented")
}
func (UnimplementedSaveServiceServer) DeleteSave(context.Context, *DeleteSaveRequest) (*DeleteSaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSave not implemented")
}
func (UnimplementedSaveServiceServer) mustEmbedUnimplementedSaveServiceServer() {}

// UnsafeSaveServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SaveServiceServer will
// result in compilation errors.
type UnsafeSaveServiceServer interface {
	mustEmbedUnimplementedSaveServiceServer()
}

func RegisterSaveServiceServer(s grpc.ServiceRegistrar, srv SaveServiceServer) {
	s.RegisterService(&SaveService_ServiceDesc, srv)
}

func _SaveService_CreateSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaveServiceServer).CreateSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaveService_CreateSave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaveServiceServer).CreateSave(ctx, req.(*CreateSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaveService_GetSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaveServiceServer).GetSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaveService_GetSave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaveServiceServer).GetSave(ctx, req.(*GetSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SaveService_DeleteSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SaveServiceServer).DeleteSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SaveService_DeleteSave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SaveServiceServer).DeleteSave(ctx, req.(*DeleteSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SaveService_ServiceDesc is the grpc.ServiceDesc for SaveService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SaveService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "save.v1.SaveService",
	HandlerType: (*SaveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSave",
			Handler:    _SaveService_CreateSave_Handler,
		},
		{
			MethodName: "GetSave",
			Handler:    _SaveService_GetSave_Handler,
		},
		{
			MethodName: "DeleteSave",
			Handler:    _SaveService_DeleteSave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/save-service.proto",
}
