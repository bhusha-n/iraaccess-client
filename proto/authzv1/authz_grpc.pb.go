package authzv1

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
	AccessManager_SetRole_FullMethodName              = "/iraaccess.v1.AccessManager/SetRole"
	AccessManager_RemoveRole_FullMethodName           = "/iraaccess.v1.AccessManager/RemoveRole"
	AccessManager_GrantResourceAccess_FullMethodName  = "/iraaccess.v1.AccessManager/GrantResourceAccess"
	AccessManager_RevokeResourceAccess_FullMethodName = "/iraaccess.v1.AccessManager/RevokeResourceAccess"
	AccessManager_CheckAccess_FullMethodName          = "/iraaccess.v1.AccessManager/CheckAccess"
	AccessManager_GetUserAccessTree_FullMethodName    = "/iraaccess.v1.AccessManager/GetUserAccessTree"
	AccessManager_GetResourceTypes_FullMethodName     = "/iraaccess.v1.AccessManager/GetResourceTypes"
)

// AccessManagerClient is the client API for AccessManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccessManagerClient interface {
	SetRole(ctx context.Context, in *SetRoleRequest, opts ...grpc.CallOption) (*SetRoleResponse, error)
	RemoveRole(ctx context.Context, in *RemoveRoleRequest, opts ...grpc.CallOption) (*RemoveRoleResponse, error)
	GrantResourceAccess(ctx context.Context, in *GrantResourceAccessRequest, opts ...grpc.CallOption) (*GrantResourceAccessResponse, error)
	RevokeResourceAccess(ctx context.Context, in *RevokeResourceAccessRequest, opts ...grpc.CallOption) (*RevokeResourceAccessResponse, error)
	CheckAccess(ctx context.Context, in *CheckAccessRequest, opts ...grpc.CallOption) (*CheckAccessResponse, error)
	GetUserAccessTree(ctx context.Context, in *GetUserAccessTreeRequest, opts ...grpc.CallOption) (*GetUserAccessTreeResponse, error)
	GetResourceTypes(ctx context.Context, in *GetResourceTypesRequest, opts ...grpc.CallOption) (*GetResourceTypesResponse, error)
}

type accessManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewAccessManagerClient(cc grpc.ClientConnInterface) AccessManagerClient {
	return &accessManagerClient{cc}
}

func (c *accessManagerClient) SetRole(ctx context.Context, in *SetRoleRequest, opts ...grpc.CallOption) (*SetRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetRoleResponse)
	err := c.cc.Invoke(ctx, AccessManager_SetRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) RemoveRole(ctx context.Context, in *RemoveRoleRequest, opts ...grpc.CallOption) (*RemoveRoleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveRoleResponse)
	err := c.cc.Invoke(ctx, AccessManager_RemoveRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) GrantResourceAccess(ctx context.Context, in *GrantResourceAccessRequest, opts ...grpc.CallOption) (*GrantResourceAccessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GrantResourceAccessResponse)
	err := c.cc.Invoke(ctx, AccessManager_GrantResourceAccess_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) RevokeResourceAccess(ctx context.Context, in *RevokeResourceAccessRequest, opts ...grpc.CallOption) (*RevokeResourceAccessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeResourceAccessResponse)
	err := c.cc.Invoke(ctx, AccessManager_RevokeResourceAccess_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) CheckAccess(ctx context.Context, in *CheckAccessRequest, opts ...grpc.CallOption) (*CheckAccessResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAccessResponse)
	err := c.cc.Invoke(ctx, AccessManager_CheckAccess_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) GetUserAccessTree(ctx context.Context, in *GetUserAccessTreeRequest, opts ...grpc.CallOption) (*GetUserAccessTreeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserAccessTreeResponse)
	err := c.cc.Invoke(ctx, AccessManager_GetUserAccessTree_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accessManagerClient) GetResourceTypes(ctx context.Context, in *GetResourceTypesRequest, opts ...grpc.CallOption) (*GetResourceTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResourceTypesResponse)
	err := c.cc.Invoke(ctx, AccessManager_GetResourceTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccessManagerServer is the server API for AccessManager service.
// All implementations must embed UnimplementedAccessManagerServer
// for forward compatibility.
type AccessManagerServer interface {
	SetRole(context.Context, *SetRoleRequest) (*SetRoleResponse, error)
	RemoveRole(context.Context, *RemoveRoleRequest) (*RemoveRoleResponse, error)
	GrantResourceAccess(context.Context, *GrantResourceAccessRequest) (*GrantResourceAccessResponse, error)
	RevokeResourceAccess(context.Context, *RevokeResourceAccessRequest) (*RevokeResourceAccessResponse, error)
	CheckAccess(context.Context, *CheckAccessRequest) (*CheckAccessResponse, error)
	GetUserAccessTree(context.Context, *GetUserAccessTreeRequest) (*GetUserAccessTreeResponse, error)
	GetResourceTypes(context.Context, *GetResourceTypesRequest) (*GetResourceTypesResponse, error)
	mustEmbedUnimplementedAccessManagerServer()
}

// UnimplementedAccessManagerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAccessManagerServer struct{}

func (UnimplementedAccessManagerServer) SetRole(context.Context, *SetRoleRequest) (*SetRoleResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method SetRole not implemented")
}
func (UnimplementedAccessManagerServer) RemoveRole(context.Context, *RemoveRoleRequest) (*RemoveRoleResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method RemoveRole not implemented")
}
func (UnimplementedAccessManagerServer) GrantResourceAccess(context.Context, *GrantResourceAccessRequest) (*GrantResourceAccessResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GrantResourceAccess not implemented")
}
func (UnimplementedAccessManagerServer) RevokeResourceAccess(context.Context, *RevokeResourceAccessRequest) (*RevokeResourceAccessResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method RevokeResourceAccess not implemented")
}
func (UnimplementedAccessManagerServer) CheckAccess(context.Context, *CheckAccessRequest) (*CheckAccessResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CheckAccess not implemented")
}
func (UnimplementedAccessManagerServer) GetUserAccessTree(context.Context, *GetUserAccessTreeRequest) (*GetUserAccessTreeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetUserAccessTree not implemented")
}
func (UnimplementedAccessManagerServer) GetResourceTypes(context.Context, *GetResourceTypesRequest) (*GetResourceTypesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetResourceTypes not implemented")
}
func (UnimplementedAccessManagerServer) mustEmbedUnimplementedAccessManagerServer() {}
func (UnimplementedAccessManagerServer) testEmbeddedByValue()                       {}

// UnsafeAccessManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccessManagerServer will
// result in compilation errors.
type UnsafeAccessManagerServer interface {
	mustEmbedUnimplementedAccessManagerServer()
}

func RegisterAccessManagerServer(s grpc.ServiceRegistrar, srv AccessManagerServer) {
	// If the following call panics, it indicates UnimplementedAccessManagerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AccessManager_ServiceDesc, srv)
}

func _AccessManager_SetRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).SetRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessManager_SetRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).SetRole(ctx, req.(*SetRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_RemoveRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).RemoveRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessManager_RemoveRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).RemoveRole(ctx, req.(*RemoveRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_GrantResourceAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GrantResourceAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).GrantResourceAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessManager_GrantResourceAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).GrantResourceAccess(ctx, req.(*GrantResourceAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_RevokeResourceAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeResourceAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).RevokeResourceAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessManager_RevokeResourceAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).RevokeResourceAccess(ctx, req.(*RevokeResourceAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_CheckAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).CheckAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessManager_CheckAccess_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).CheckAccess(ctx, req.(*CheckAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_GetUserAccessTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAccessTreeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).GetUserAccessTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessManager_GetUserAccessTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).GetUserAccessTree(ctx, req.(*GetUserAccessTreeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccessManager_GetResourceTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourceTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccessManagerServer).GetResourceTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccessManager_GetResourceTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccessManagerServer).GetResourceTypes(ctx, req.(*GetResourceTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccessManager_ServiceDesc is the grpc.ServiceDesc for AccessManager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccessManager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iraaccess.v1.AccessManager",
	HandlerType: (*AccessManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetRole",
			Handler:    _AccessManager_SetRole_Handler,
		},
		{
			MethodName: "RemoveRole",
			Handler:    _AccessManager_RemoveRole_Handler,
		},
		{
			MethodName: "GrantResourceAccess",
			Handler:    _AccessManager_GrantResourceAccess_Handler,
		},
		{
			MethodName: "RevokeResourceAccess",
			Handler:    _AccessManager_RevokeResourceAccess_Handler,
		},
		{
			MethodName: "CheckAccess",
			Handler:    _AccessManager_CheckAccess_Handler,
		},
		{
			MethodName: "GetUserAccessTree",
			Handler:    _AccessManager_GetUserAccessTree_Handler,
		},
		{
			MethodName: "GetResourceTypes",
			Handler:    _AccessManager_GetResourceTypes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/authzv1/authz.proto",
}
