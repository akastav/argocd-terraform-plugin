// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/ydb/v1/database_service.proto

package ydb

import (
	context "context"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	// Returns the specified database.
	Get(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*Database, error)
	// Retrieves a list of databases.
	List(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error)
	// Creates a new database.
	Create(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Modifies the specified database.
	Update(ctx context.Context, in *UpdateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Starts the specified database.
	Start(ctx context.Context, in *StartDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Stops the specified database.
	Stop(ctx context.Context, in *StopDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified database.
	Delete(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Restores the specified backup
	Restore(ctx context.Context, in *RestoreBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	Backup(ctx context.Context, in *BackupDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) Get(ctx context.Context, in *GetDatabaseRequest, opts ...grpc.CallOption) (*Database, error) {
	out := new(Database)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) List(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error) {
	out := new(ListDatabasesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Create(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Update(ctx context.Context, in *UpdateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Start(ctx context.Context, in *StartDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Stop(ctx context.Context, in *StopDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/ListAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/SetAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/UpdateAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Delete(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Restore(ctx context.Context, in *RestoreBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/Restore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Backup(ctx context.Context, in *BackupDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ydb.v1.DatabaseService/Backup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
// All implementations should embed UnimplementedDatabaseServiceServer
// for forward compatibility
type DatabaseServiceServer interface {
	// Returns the specified database.
	Get(context.Context, *GetDatabaseRequest) (*Database, error)
	// Retrieves a list of databases.
	List(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error)
	// Creates a new database.
	Create(context.Context, *CreateDatabaseRequest) (*operation.Operation, error)
	// Modifies the specified database.
	Update(context.Context, *UpdateDatabaseRequest) (*operation.Operation, error)
	// Starts the specified database.
	Start(context.Context, *StartDatabaseRequest) (*operation.Operation, error)
	// Stops the specified database.
	Stop(context.Context, *StopDatabaseRequest) (*operation.Operation, error)
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
	// Deletes the specified database.
	Delete(context.Context, *DeleteDatabaseRequest) (*operation.Operation, error)
	// Restores the specified backup
	Restore(context.Context, *RestoreBackupRequest) (*operation.Operation, error)
	Backup(context.Context, *BackupDatabaseRequest) (*operation.Operation, error)
}

// UnimplementedDatabaseServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDatabaseServiceServer struct {
}

func (UnimplementedDatabaseServiceServer) Get(context.Context, *GetDatabaseRequest) (*Database, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDatabaseServiceServer) List(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDatabaseServiceServer) Create(context.Context, *CreateDatabaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDatabaseServiceServer) Update(context.Context, *UpdateDatabaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDatabaseServiceServer) Start(context.Context, *StartDatabaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedDatabaseServiceServer) Stop(context.Context, *StopDatabaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedDatabaseServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedDatabaseServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedDatabaseServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedDatabaseServiceServer) Delete(context.Context, *DeleteDatabaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDatabaseServiceServer) Restore(context.Context, *RestoreBackupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}
func (UnimplementedDatabaseServiceServer) Backup(context.Context, *BackupDatabaseRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Backup not implemented")
}

// UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServiceServer will
// result in compilation errors.
type UnsafeDatabaseServiceServer interface {
	mustEmbedUnimplementedDatabaseServiceServer()
}

func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer) {
	s.RegisterService(&DatabaseService_ServiceDesc, srv)
}

func _DatabaseService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Get(ctx, req.(*GetDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).List(ctx, req.(*ListDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Create(ctx, req.(*CreateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Update(ctx, req.(*UpdateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Start(ctx, req.(*StartDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Stop(ctx, req.(*StopDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/ListAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/SetAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/UpdateAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Delete(ctx, req.(*DeleteDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/Restore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Restore(ctx, req.(*RestoreBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ydb.v1.DatabaseService/Backup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Backup(ctx, req.(*BackupDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseService_ServiceDesc is the grpc.ServiceDesc for DatabaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ydb.v1.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DatabaseService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DatabaseService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DatabaseService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DatabaseService_Update_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _DatabaseService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _DatabaseService_Stop_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _DatabaseService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _DatabaseService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _DatabaseService_UpdateAccessBindings_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DatabaseService_Delete_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _DatabaseService_Restore_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _DatabaseService_Backup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ydb/v1/database_service.proto",
}
