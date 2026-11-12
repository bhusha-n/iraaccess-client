package authzv1

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	Action_ACTION_UNSPECIFIED Action = 0
	Action_ACTION_VIEW        Action = 1
	Action_ACTION_CREATE      Action = 2
	Action_ACTION_UPDATE      Action = 3
	Action_ACTION_DELETE      Action = 4
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_VIEW",
		2: "ACTION_CREATE",
		3: "ACTION_UPDATE",
		4: "ACTION_DELETE",
	}
	Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_VIEW":        1,
		"ACTION_CREATE":      2,
		"ACTION_UPDATE":      3,
		"ACTION_DELETE":      4,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_authzv1_authz_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_proto_authzv1_authz_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{0}
}

type Role int32

const (
	Role_ROLE_UNSPECIFIED   Role = 0
	Role_ROLE_TENANT_MEMBER Role = 1
	Role_ROLE_TENANT_ADMIN  Role = 2
	Role_ROLE_SUPER_ADMIN   Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ROLE_TENANT_MEMBER",
		2: "ROLE_TENANT_ADMIN",
		3: "ROLE_SUPER_ADMIN",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED":   0,
		"ROLE_TENANT_MEMBER": 1,
		"ROLE_TENANT_ADMIN":  2,
		"ROLE_SUPER_ADMIN":   3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_authzv1_authz_proto_enumTypes[1].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_proto_authzv1_authz_proto_enumTypes[1]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{1}
}

type ResourceScope int32

const (
	ResourceScope_RESOURCE_SCOPE_UNSPECIFIED ResourceScope = 0
	ResourceScope_RESOURCE_SCOPE_ALL         ResourceScope = 1
	ResourceScope_RESOURCE_SCOPE_SELECTED    ResourceScope = 2
)

// Enum value maps for ResourceScope.
var (
	ResourceScope_name = map[int32]string{
		0: "RESOURCE_SCOPE_UNSPECIFIED",
		1: "RESOURCE_SCOPE_ALL",
		2: "RESOURCE_SCOPE_SELECTED",
	}
	ResourceScope_value = map[string]int32{
		"RESOURCE_SCOPE_UNSPECIFIED": 0,
		"RESOURCE_SCOPE_ALL":         1,
		"RESOURCE_SCOPE_SELECTED":    2,
	}
)

func (x ResourceScope) Enum() *ResourceScope {
	p := new(ResourceScope)
	*p = x
	return p
}

func (x ResourceScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceScope) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_authzv1_authz_proto_enumTypes[2].Descriptor()
}

func (ResourceScope) Type() protoreflect.EnumType {
	return &file_proto_authzv1_authz_proto_enumTypes[2]
}

func (x ResourceScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceScope.Descriptor instead.
func (ResourceScope) EnumDescriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{2}
}

// SetRole assigns TENANT_ADMIN (tenant_id required) or SUPER_ADMIN (tenant_id must be empty).
// granted_by is for audit logging only, never for authorization.
type SetRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TargetAppId   string                 `protobuf:"bytes,1,opt,name=target_app_id,json=targetAppId,proto3" json:"target_app_id,omitempty"`
	TenantId      string                 `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role          Role                   `protobuf:"varint,4,opt,name=role,proto3,enum=iraaccess.v1.Role" json:"role,omitempty"`
	GrantedBy     string                 `protobuf:"bytes,5,opt,name=granted_by,json=grantedBy,proto3" json:"granted_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetRoleRequest) Reset() {
	*x = SetRoleRequest{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoleRequest) ProtoMessage() {}

func (x *SetRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoleRequest.ProtoReflect.Descriptor instead.
func (*SetRoleRequest) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{0}
}

func (x *SetRoleRequest) GetTargetAppId() string {
	if x != nil {
		return x.TargetAppId
	}
	return ""
}

func (x *SetRoleRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *SetRoleRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetRoleRequest) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *SetRoleRequest) GetGrantedBy() string {
	if x != nil {
		return x.GrantedBy
	}
	return ""
}

type SetRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ZedToken      string                 `protobuf:"bytes,2,opt,name=zed_token,json=zedToken,proto3" json:"zed_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetRoleResponse) Reset() {
	*x = SetRoleResponse{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRoleResponse) ProtoMessage() {}

func (x *SetRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRoleResponse.ProtoReflect.Descriptor instead.
func (*SetRoleResponse) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{1}
}

func (x *SetRoleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SetRoleResponse) GetZedToken() string {
	if x != nil {
		return x.ZedToken
	}
	return ""
}

type RemoveRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TargetAppId   string                 `protobuf:"bytes,1,opt,name=target_app_id,json=targetAppId,proto3" json:"target_app_id,omitempty"`
	TenantId      string                 `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Role          Role                   `protobuf:"varint,4,opt,name=role,proto3,enum=iraaccess.v1.Role" json:"role,omitempty"`
	DeletedBy     string                 `protobuf:"bytes,5,opt,name=deleted_by,json=deletedBy,proto3" json:"deleted_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRoleRequest) Reset() {
	*x = RemoveRoleRequest{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoleRequest) ProtoMessage() {}

func (x *RemoveRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoleRequest.ProtoReflect.Descriptor instead.
func (*RemoveRoleRequest) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveRoleRequest) GetTargetAppId() string {
	if x != nil {
		return x.TargetAppId
	}
	return ""
}

func (x *RemoveRoleRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *RemoveRoleRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RemoveRoleRequest) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *RemoveRoleRequest) GetDeletedBy() string {
	if x != nil {
		return x.DeletedBy
	}
	return ""
}

type RemoveRoleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ZedToken      string                 `protobuf:"bytes,2,opt,name=zed_token,json=zedToken,proto3" json:"zed_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveRoleResponse) Reset() {
	*x = RemoveRoleResponse{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveRoleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRoleResponse) ProtoMessage() {}

func (x *RemoveRoleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRoleResponse.ProtoReflect.Descriptor instead.
func (*RemoveRoleResponse) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveRoleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RemoveRoleResponse) GetZedToken() string {
	if x != nil {
		return x.ZedToken
	}
	return ""
}

// scope ALL: resource_ids must be empty. scope SELECTED: resource_ids required, CREATE not allowed.
type GrantResourceAccessRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TargetAppId   string                 `protobuf:"bytes,1,opt,name=target_app_id,json=targetAppId,proto3" json:"target_app_id,omitempty"`
	TenantId      string                 `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResourceType  string                 `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Scope         ResourceScope          `protobuf:"varint,5,opt,name=scope,proto3,enum=iraaccess.v1.ResourceScope" json:"scope,omitempty"`
	ResourceIds   []string               `protobuf:"bytes,6,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
	Actions       []Action               `protobuf:"varint,7,rep,packed,name=actions,proto3,enum=iraaccess.v1.Action" json:"actions,omitempty"`
	GrantedBy     string                 `protobuf:"bytes,8,opt,name=granted_by,json=grantedBy,proto3" json:"granted_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantResourceAccessRequest) Reset() {
	*x = GrantResourceAccessRequest{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantResourceAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantResourceAccessRequest) ProtoMessage() {}

func (x *GrantResourceAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantResourceAccessRequest.ProtoReflect.Descriptor instead.
func (*GrantResourceAccessRequest) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{4}
}

func (x *GrantResourceAccessRequest) GetTargetAppId() string {
	if x != nil {
		return x.TargetAppId
	}
	return ""
}

func (x *GrantResourceAccessRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *GrantResourceAccessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GrantResourceAccessRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *GrantResourceAccessRequest) GetScope() ResourceScope {
	if x != nil {
		return x.Scope
	}
	return ResourceScope_RESOURCE_SCOPE_UNSPECIFIED
}

func (x *GrantResourceAccessRequest) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *GrantResourceAccessRequest) GetActions() []Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *GrantResourceAccessRequest) GetGrantedBy() string {
	if x != nil {
		return x.GrantedBy
	}
	return ""
}

type GrantResourceAccessResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ZedToken      string                 `protobuf:"bytes,2,opt,name=zed_token,json=zedToken,proto3" json:"zed_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GrantResourceAccessResponse) Reset() {
	*x = GrantResourceAccessResponse{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GrantResourceAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantResourceAccessResponse) ProtoMessage() {}

func (x *GrantResourceAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantResourceAccessResponse.ProtoReflect.Descriptor instead.
func (*GrantResourceAccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{5}
}

func (x *GrantResourceAccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GrantResourceAccessResponse) GetZedToken() string {
	if x != nil {
		return x.ZedToken
	}
	return ""
}

type RevokeResourceAccessRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TargetAppId   string                 `protobuf:"bytes,1,opt,name=target_app_id,json=targetAppId,proto3" json:"target_app_id,omitempty"`
	TenantId      string                 `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	UserId        string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResourceType  string                 `protobuf:"bytes,4,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Scope         ResourceScope          `protobuf:"varint,5,opt,name=scope,proto3,enum=iraaccess.v1.ResourceScope" json:"scope,omitempty"`
	ResourceIds   []string               `protobuf:"bytes,6,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
	Actions       []Action               `protobuf:"varint,7,rep,packed,name=actions,proto3,enum=iraaccess.v1.Action" json:"actions,omitempty"`
	DeletedBy     string                 `protobuf:"bytes,8,opt,name=deleted_by,json=deletedBy,proto3" json:"deleted_by,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeResourceAccessRequest) Reset() {
	*x = RevokeResourceAccessRequest{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeResourceAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeResourceAccessRequest) ProtoMessage() {}

func (x *RevokeResourceAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeResourceAccessRequest.ProtoReflect.Descriptor instead.
func (*RevokeResourceAccessRequest) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{6}
}

func (x *RevokeResourceAccessRequest) GetTargetAppId() string {
	if x != nil {
		return x.TargetAppId
	}
	return ""
}

func (x *RevokeResourceAccessRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *RevokeResourceAccessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RevokeResourceAccessRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *RevokeResourceAccessRequest) GetScope() ResourceScope {
	if x != nil {
		return x.Scope
	}
	return ResourceScope_RESOURCE_SCOPE_UNSPECIFIED
}

func (x *RevokeResourceAccessRequest) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *RevokeResourceAccessRequest) GetActions() []Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *RevokeResourceAccessRequest) GetDeletedBy() string {
	if x != nil {
		return x.DeletedBy
	}
	return ""
}

type RevokeResourceAccessResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ZedToken      string                 `protobuf:"bytes,2,opt,name=zed_token,json=zedToken,proto3" json:"zed_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RevokeResourceAccessResponse) Reset() {
	*x = RevokeResourceAccessResponse{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RevokeResourceAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeResourceAccessResponse) ProtoMessage() {}

func (x *RevokeResourceAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeResourceAccessResponse.ProtoReflect.Descriptor instead.
func (*RevokeResourceAccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{7}
}

func (x *RevokeResourceAccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *RevokeResourceAccessResponse) GetZedToken() string {
	if x != nil {
		return x.ZedToken
	}
	return ""
}

// resource_id is ignored for ACTION_CREATE.
type CheckAccessRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Action        Action                 `protobuf:"varint,2,opt,name=action,proto3,enum=iraaccess.v1.Action" json:"action,omitempty"`
	ResourceType  string                 `protobuf:"bytes,3,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	ResourceId    string                 `protobuf:"bytes,4,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	TenantId      string                 `protobuf:"bytes,5,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckAccessRequest) Reset() {
	*x = CheckAccessRequest{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessRequest) ProtoMessage() {}

func (x *CheckAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessRequest.ProtoReflect.Descriptor instead.
func (*CheckAccessRequest) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{8}
}

func (x *CheckAccessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CheckAccessRequest) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

func (x *CheckAccessRequest) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *CheckAccessRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CheckAccessRequest) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

type CheckAccessResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Allowed       bool                   `protobuf:"varint,1,opt,name=allowed,proto3" json:"allowed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckAccessResponse) Reset() {
	*x = CheckAccessResponse{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAccessResponse) ProtoMessage() {}

func (x *CheckAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckAccessResponse.ProtoReflect.Descriptor instead.
func (*CheckAccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{9}
}

func (x *CheckAccessResponse) GetAllowed() bool {
	if x != nil {
		return x.Allowed
	}
	return false
}

// target_app_id empty = caller's own app.
type GetUserAccessTreeRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TargetAppId   string                 `protobuf:"bytes,2,opt,name=target_app_id,json=targetAppId,proto3" json:"target_app_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserAccessTreeRequest) Reset() {
	*x = GetUserAccessTreeRequest{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserAccessTreeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAccessTreeRequest) ProtoMessage() {}

func (x *GetUserAccessTreeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAccessTreeRequest.ProtoReflect.Descriptor instead.
func (*GetUserAccessTreeRequest) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{10}
}

func (x *GetUserAccessTreeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserAccessTreeRequest) GetTargetAppId() string {
	if x != nil {
		return x.TargetAppId
	}
	return ""
}

type GetUserAccessTreeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Application   *ApplicationAccess     `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserAccessTreeResponse) Reset() {
	*x = GetUserAccessTreeResponse{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserAccessTreeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAccessTreeResponse) ProtoMessage() {}

func (x *GetUserAccessTreeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAccessTreeResponse.ProtoReflect.Descriptor instead.
func (*GetUserAccessTreeResponse) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserAccessTreeResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserAccessTreeResponse) GetApplication() *ApplicationAccess {
	if x != nil {
		return x.Application
	}
	return nil
}

type ApplicationAccess struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AppId         string                 `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Role          Role                   `protobuf:"varint,2,opt,name=role,proto3,enum=iraaccess.v1.Role" json:"role,omitempty"`
	Tenants       []*TenantAccess        `protobuf:"bytes,3,rep,name=tenants,proto3" json:"tenants,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApplicationAccess) Reset() {
	*x = ApplicationAccess{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApplicationAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplicationAccess) ProtoMessage() {}

func (x *ApplicationAccess) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplicationAccess.ProtoReflect.Descriptor instead.
func (*ApplicationAccess) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{12}
}

func (x *ApplicationAccess) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *ApplicationAccess) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *ApplicationAccess) GetTenants() []*TenantAccess {
	if x != nil {
		return x.Tenants
	}
	return nil
}

type TenantAccess struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TenantId      string                 `protobuf:"bytes,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Role          Role                   `protobuf:"varint,2,opt,name=role,proto3,enum=iraaccess.v1.Role" json:"role,omitempty"`
	ResourceTypes []*ResourceTypeAccess  `protobuf:"bytes,3,rep,name=resource_types,json=resourceTypes,proto3" json:"resource_types,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TenantAccess) Reset() {
	*x = TenantAccess{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TenantAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TenantAccess) ProtoMessage() {}

func (x *TenantAccess) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TenantAccess.ProtoReflect.Descriptor instead.
func (*TenantAccess) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{13}
}

func (x *TenantAccess) GetTenantId() string {
	if x != nil {
		return x.TenantId
	}
	return ""
}

func (x *TenantAccess) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *TenantAccess) GetResourceTypes() []*ResourceTypeAccess {
	if x != nil {
		return x.ResourceTypes
	}
	return nil
}

type ResourceTypeAccess struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceType  string                 `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	Grants        []*ResourceGrant       `protobuf:"bytes,2,rep,name=grants,proto3" json:"grants,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceTypeAccess) Reset() {
	*x = ResourceTypeAccess{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceTypeAccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceTypeAccess) ProtoMessage() {}

func (x *ResourceTypeAccess) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceTypeAccess.ProtoReflect.Descriptor instead.
func (*ResourceTypeAccess) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{14}
}

func (x *ResourceTypeAccess) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ResourceTypeAccess) GetGrants() []*ResourceGrant {
	if x != nil {
		return x.Grants
	}
	return nil
}

type ResourceGrant struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Scope         ResourceScope          `protobuf:"varint,1,opt,name=scope,proto3,enum=iraaccess.v1.ResourceScope" json:"scope,omitempty"`
	ResourceIds   []string               `protobuf:"bytes,2,rep,name=resource_ids,json=resourceIds,proto3" json:"resource_ids,omitempty"`
	Actions       []Action               `protobuf:"varint,3,rep,packed,name=actions,proto3,enum=iraaccess.v1.Action" json:"actions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceGrant) Reset() {
	*x = ResourceGrant{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceGrant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceGrant) ProtoMessage() {}

func (x *ResourceGrant) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceGrant.ProtoReflect.Descriptor instead.
func (*ResourceGrant) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{15}
}

func (x *ResourceGrant) GetScope() ResourceScope {
	if x != nil {
		return x.Scope
	}
	return ResourceScope_RESOURCE_SCOPE_UNSPECIFIED
}

func (x *ResourceGrant) GetResourceIds() []string {
	if x != nil {
		return x.ResourceIds
	}
	return nil
}

func (x *ResourceGrant) GetActions() []Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

// target_app_id empty = caller's own app.
type GetResourceTypesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TargetAppId   string                 `protobuf:"bytes,1,opt,name=target_app_id,json=targetAppId,proto3" json:"target_app_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResourceTypesRequest) Reset() {
	*x = GetResourceTypesRequest{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResourceTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceTypesRequest) ProtoMessage() {}

func (x *GetResourceTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceTypesRequest.ProtoReflect.Descriptor instead.
func (*GetResourceTypesRequest) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{16}
}

func (x *GetResourceTypesRequest) GetTargetAppId() string {
	if x != nil {
		return x.TargetAppId
	}
	return ""
}

type GetResourceTypesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ResourceTypes []string               `protobuf:"bytes,1,rep,name=resource_types,json=resourceTypes,proto3" json:"resource_types,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResourceTypesResponse) Reset() {
	*x = GetResourceTypesResponse{}
	mi := &file_proto_authzv1_authz_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResourceTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourceTypesResponse) ProtoMessage() {}

func (x *GetResourceTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_authzv1_authz_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourceTypesResponse.ProtoReflect.Descriptor instead.
func (*GetResourceTypesResponse) Descriptor() ([]byte, []int) {
	return file_proto_authzv1_authz_proto_rawDescGZIP(), []int{17}
}

func (x *GetResourceTypesResponse) GetResourceTypes() []string {
	if x != nil {
		return x.ResourceTypes
	}
	return nil
}

var File_proto_authzv1_authz_proto protoreflect.FileDescriptor

const file_proto_authzv1_authz_proto_rawDesc = "" +
	"\n" +
	"\x19proto/authzv1/authz.proto\x12\firaaccess.v1\"\xb1\x01\n" +
	"\x0eSetRoleRequest\x12\"\n" +
	"\rtarget_app_id\x18\x01 \x01(\tR\vtargetAppId\x12\x1b\n" +
	"\ttenant_id\x18\x02 \x01(\tR\btenantId\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userId\x12&\n" +
	"\x04role\x18\x04 \x01(\x0e2\x12.iraaccess.v1.RoleR\x04role\x12\x1d\n" +
	"\n" +
	"granted_by\x18\x05 \x01(\tR\tgrantedBy\"H\n" +
	"\x0fSetRoleResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x1b\n" +
	"\tzed_token\x18\x02 \x01(\tR\bzedToken\"\xb4\x01\n" +
	"\x11RemoveRoleRequest\x12\"\n" +
	"\rtarget_app_id\x18\x01 \x01(\tR\vtargetAppId\x12\x1b\n" +
	"\ttenant_id\x18\x02 \x01(\tR\btenantId\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userId\x12&\n" +
	"\x04role\x18\x04 \x01(\x0e2\x12.iraaccess.v1.RoleR\x04role\x12\x1d\n" +
	"\n" +
	"deleted_by\x18\x05 \x01(\tR\tdeletedBy\"K\n" +
	"\x12RemoveRoleResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x1b\n" +
	"\tzed_token\x18\x02 \x01(\tR\bzedToken\"\xc0\x02\n" +
	"\x1aGrantResourceAccessRequest\x12\"\n" +
	"\rtarget_app_id\x18\x01 \x01(\tR\vtargetAppId\x12\x1b\n" +
	"\ttenant_id\x18\x02 \x01(\tR\btenantId\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userId\x12#\n" +
	"\rresource_type\x18\x04 \x01(\tR\fresourceType\x121\n" +
	"\x05scope\x18\x05 \x01(\x0e2\x1b.iraaccess.v1.ResourceScopeR\x05scope\x12!\n" +
	"\fresource_ids\x18\x06 \x03(\tR\vresourceIds\x12.\n" +
	"\aactions\x18\a \x03(\x0e2\x14.iraaccess.v1.ActionR\aactions\x12\x1d\n" +
	"\n" +
	"granted_by\x18\b \x01(\tR\tgrantedBy\"T\n" +
	"\x1bGrantResourceAccessResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x1b\n" +
	"\tzed_token\x18\x02 \x01(\tR\bzedToken\"\xc1\x02\n" +
	"\x1bRevokeResourceAccessRequest\x12\"\n" +
	"\rtarget_app_id\x18\x01 \x01(\tR\vtargetAppId\x12\x1b\n" +
	"\ttenant_id\x18\x02 \x01(\tR\btenantId\x12\x17\n" +
	"\auser_id\x18\x03 \x01(\tR\x06userId\x12#\n" +
	"\rresource_type\x18\x04 \x01(\tR\fresourceType\x121\n" +
	"\x05scope\x18\x05 \x01(\x0e2\x1b.iraaccess.v1.ResourceScopeR\x05scope\x12!\n" +
	"\fresource_ids\x18\x06 \x03(\tR\vresourceIds\x12.\n" +
	"\aactions\x18\a \x03(\x0e2\x14.iraaccess.v1.ActionR\aactions\x12\x1d\n" +
	"\n" +
	"deleted_by\x18\b \x01(\tR\tdeletedBy\"U\n" +
	"\x1cRevokeResourceAccessResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x1b\n" +
	"\tzed_token\x18\x02 \x01(\tR\bzedToken\"\xbe\x01\n" +
	"\x12CheckAccessRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12,\n" +
	"\x06action\x18\x02 \x01(\x0e2\x14.iraaccess.v1.ActionR\x06action\x12#\n" +
	"\rresource_type\x18\x03 \x01(\tR\fresourceType\x12\x1f\n" +
	"\vresource_id\x18\x04 \x01(\tR\n" +
	"resourceId\x12\x1b\n" +
	"\ttenant_id\x18\x05 \x01(\tR\btenantId\"/\n" +
	"\x13CheckAccessResponse\x12\x18\n" +
	"\aallowed\x18\x01 \x01(\bR\aallowed\"W\n" +
	"\x18GetUserAccessTreeRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\"\n" +
	"\rtarget_app_id\x18\x02 \x01(\tR\vtargetAppId\"w\n" +
	"\x19GetUserAccessTreeResponse\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12A\n" +
	"\vapplication\x18\x02 \x01(\v2\x1f.iraaccess.v1.ApplicationAccessR\vapplication\"\x88\x01\n" +
	"\x11ApplicationAccess\x12\x15\n" +
	"\x06app_id\x18\x01 \x01(\tR\x05appId\x12&\n" +
	"\x04role\x18\x02 \x01(\x0e2\x12.iraaccess.v1.RoleR\x04role\x124\n" +
	"\atenants\x18\x03 \x03(\v2\x1a.iraaccess.v1.TenantAccessR\atenants\"\x9c\x01\n" +
	"\fTenantAccess\x12\x1b\n" +
	"\ttenant_id\x18\x01 \x01(\tR\btenantId\x12&\n" +
	"\x04role\x18\x02 \x01(\x0e2\x12.iraaccess.v1.RoleR\x04role\x12G\n" +
	"\x0eresource_types\x18\x03 \x03(\v2 .iraaccess.v1.ResourceTypeAccessR\rresourceTypes\"n\n" +
	"\x12ResourceTypeAccess\x12#\n" +
	"\rresource_type\x18\x01 \x01(\tR\fresourceType\x123\n" +
	"\x06grants\x18\x02 \x03(\v2\x1b.iraaccess.v1.ResourceGrantR\x06grants\"\x95\x01\n" +
	"\rResourceGrant\x121\n" +
	"\x05scope\x18\x01 \x01(\x0e2\x1b.iraaccess.v1.ResourceScopeR\x05scope\x12!\n" +
	"\fresource_ids\x18\x02 \x03(\tR\vresourceIds\x12.\n" +
	"\aactions\x18\x03 \x03(\x0e2\x14.iraaccess.v1.ActionR\aactions\"=\n" +
	"\x17GetResourceTypesRequest\x12\"\n" +
	"\rtarget_app_id\x18\x01 \x01(\tR\vtargetAppId\"A\n" +
	"\x18GetResourceTypesResponse\x12%\n" +
	"\x0eresource_types\x18\x01 \x03(\tR\rresourceTypes*j\n" +
	"\x06Action\x12\x16\n" +
	"\x12ACTION_UNSPECIFIED\x10\x00\x12\x0f\n" +
	"\vACTION_VIEW\x10\x01\x12\x11\n" +
	"\rACTION_CREATE\x10\x02\x12\x11\n" +
	"\rACTION_UPDATE\x10\x03\x12\x11\n" +
	"\rACTION_DELETE\x10\x04*a\n" +
	"\x04Role\x12\x14\n" +
	"\x10ROLE_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12ROLE_TENANT_MEMBER\x10\x01\x12\x15\n" +
	"\x11ROLE_TENANT_ADMIN\x10\x02\x12\x14\n" +
	"\x10ROLE_SUPER_ADMIN\x10\x03*d\n" +
	"\rResourceScope\x12\x1e\n" +
	"\x1aRESOURCE_SCOPE_UNSPECIFIED\x10\x00\x12\x16\n" +
	"\x12RESOURCE_SCOPE_ALL\x10\x01\x12\x1b\n" +
	"\x17RESOURCE_SCOPE_SELECTED\x10\x022\xa0\x05\n" +
	"\rAccessManager\x12F\n" +
	"\aSetRole\x12\x1c.iraaccess.v1.SetRoleRequest\x1a\x1d.iraaccess.v1.SetRoleResponse\x12O\n" +
	"\n" +
	"RemoveRole\x12\x1f.iraaccess.v1.RemoveRoleRequest\x1a .iraaccess.v1.RemoveRoleResponse\x12j\n" +
	"\x13GrantResourceAccess\x12(.iraaccess.v1.GrantResourceAccessRequest\x1a).iraaccess.v1.GrantResourceAccessResponse\x12m\n" +
	"\x14RevokeResourceAccess\x12).iraaccess.v1.RevokeResourceAccessRequest\x1a*.iraaccess.v1.RevokeResourceAccessResponse\x12R\n" +
	"\vCheckAccess\x12 .iraaccess.v1.CheckAccessRequest\x1a!.iraaccess.v1.CheckAccessResponse\x12d\n" +
	"\x11GetUserAccessTree\x12&.iraaccess.v1.GetUserAccessTreeRequest\x1a'.iraaccess.v1.GetUserAccessTreeResponse\x12a\n" +
	"\x10GetResourceTypes\x12%.iraaccess.v1.GetResourceTypesRequest\x1a&.iraaccess.v1.GetResourceTypesResponseB<Z:github.com/bhusha-n/IraAcesssManager/proto/authzv1;authzv1b\x06proto3"

var (
	file_proto_authzv1_authz_proto_rawDescOnce sync.Once
	file_proto_authzv1_authz_proto_rawDescData []byte
)

func file_proto_authzv1_authz_proto_rawDescGZIP() []byte {
	file_proto_authzv1_authz_proto_rawDescOnce.Do(func() {
		file_proto_authzv1_authz_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_authzv1_authz_proto_rawDesc), len(file_proto_authzv1_authz_proto_rawDesc)))
	})
	return file_proto_authzv1_authz_proto_rawDescData
}

var file_proto_authzv1_authz_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_proto_authzv1_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file_proto_authzv1_authz_proto_goTypes = []any{
	(Action)(0),                          // 0: iraaccess.v1.Action
	(Role)(0),                            // 1: iraaccess.v1.Role
	(ResourceScope)(0),                   // 2: iraaccess.v1.ResourceScope
	(*SetRoleRequest)(nil),               // 3: iraaccess.v1.SetRoleRequest
	(*SetRoleResponse)(nil),              // 4: iraaccess.v1.SetRoleResponse
	(*RemoveRoleRequest)(nil),            // 5: iraaccess.v1.RemoveRoleRequest
	(*RemoveRoleResponse)(nil),           // 6: iraaccess.v1.RemoveRoleResponse
	(*GrantResourceAccessRequest)(nil),   // 7: iraaccess.v1.GrantResourceAccessRequest
	(*GrantResourceAccessResponse)(nil),  // 8: iraaccess.v1.GrantResourceAccessResponse
	(*RevokeResourceAccessRequest)(nil),  // 9: iraaccess.v1.RevokeResourceAccessRequest
	(*RevokeResourceAccessResponse)(nil), // 10: iraaccess.v1.RevokeResourceAccessResponse
	(*CheckAccessRequest)(nil),           // 11: iraaccess.v1.CheckAccessRequest
	(*CheckAccessResponse)(nil),          // 12: iraaccess.v1.CheckAccessResponse
	(*GetUserAccessTreeRequest)(nil),     // 13: iraaccess.v1.GetUserAccessTreeRequest
	(*GetUserAccessTreeResponse)(nil),    // 14: iraaccess.v1.GetUserAccessTreeResponse
	(*ApplicationAccess)(nil),            // 15: iraaccess.v1.ApplicationAccess
	(*TenantAccess)(nil),                 // 16: iraaccess.v1.TenantAccess
	(*ResourceTypeAccess)(nil),           // 17: iraaccess.v1.ResourceTypeAccess
	(*ResourceGrant)(nil),                // 18: iraaccess.v1.ResourceGrant
	(*GetResourceTypesRequest)(nil),      // 19: iraaccess.v1.GetResourceTypesRequest
	(*GetResourceTypesResponse)(nil),     // 20: iraaccess.v1.GetResourceTypesResponse
}
var file_proto_authzv1_authz_proto_depIdxs = []int32{
	1,  // 0: iraaccess.v1.SetRoleRequest.role:type_name -> iraaccess.v1.Role
	1,  // 1: iraaccess.v1.RemoveRoleRequest.role:type_name -> iraaccess.v1.Role
	2,  // 2: iraaccess.v1.GrantResourceAccessRequest.scope:type_name -> iraaccess.v1.ResourceScope
	0,  // 3: iraaccess.v1.GrantResourceAccessRequest.actions:type_name -> iraaccess.v1.Action
	2,  // 4: iraaccess.v1.RevokeResourceAccessRequest.scope:type_name -> iraaccess.v1.ResourceScope
	0,  // 5: iraaccess.v1.RevokeResourceAccessRequest.actions:type_name -> iraaccess.v1.Action
	0,  // 6: iraaccess.v1.CheckAccessRequest.action:type_name -> iraaccess.v1.Action
	15, // 7: iraaccess.v1.GetUserAccessTreeResponse.application:type_name -> iraaccess.v1.ApplicationAccess
	1,  // 8: iraaccess.v1.ApplicationAccess.role:type_name -> iraaccess.v1.Role
	16, // 9: iraaccess.v1.ApplicationAccess.tenants:type_name -> iraaccess.v1.TenantAccess
	1,  // 10: iraaccess.v1.TenantAccess.role:type_name -> iraaccess.v1.Role
	17, // 11: iraaccess.v1.TenantAccess.resource_types:type_name -> iraaccess.v1.ResourceTypeAccess
	18, // 12: iraaccess.v1.ResourceTypeAccess.grants:type_name -> iraaccess.v1.ResourceGrant
	2,  // 13: iraaccess.v1.ResourceGrant.scope:type_name -> iraaccess.v1.ResourceScope
	0,  // 14: iraaccess.v1.ResourceGrant.actions:type_name -> iraaccess.v1.Action
	3,  // 15: iraaccess.v1.AccessManager.SetRole:input_type -> iraaccess.v1.SetRoleRequest
	5,  // 16: iraaccess.v1.AccessManager.RemoveRole:input_type -> iraaccess.v1.RemoveRoleRequest
	7,  // 17: iraaccess.v1.AccessManager.GrantResourceAccess:input_type -> iraaccess.v1.GrantResourceAccessRequest
	9,  // 18: iraaccess.v1.AccessManager.RevokeResourceAccess:input_type -> iraaccess.v1.RevokeResourceAccessRequest
	11, // 19: iraaccess.v1.AccessManager.CheckAccess:input_type -> iraaccess.v1.CheckAccessRequest
	13, // 20: iraaccess.v1.AccessManager.GetUserAccessTree:input_type -> iraaccess.v1.GetUserAccessTreeRequest
	19, // 21: iraaccess.v1.AccessManager.GetResourceTypes:input_type -> iraaccess.v1.GetResourceTypesRequest
	4,  // 22: iraaccess.v1.AccessManager.SetRole:output_type -> iraaccess.v1.SetRoleResponse
	6,  // 23: iraaccess.v1.AccessManager.RemoveRole:output_type -> iraaccess.v1.RemoveRoleResponse
	8,  // 24: iraaccess.v1.AccessManager.GrantResourceAccess:output_type -> iraaccess.v1.GrantResourceAccessResponse
	10, // 25: iraaccess.v1.AccessManager.RevokeResourceAccess:output_type -> iraaccess.v1.RevokeResourceAccessResponse
	12, // 26: iraaccess.v1.AccessManager.CheckAccess:output_type -> iraaccess.v1.CheckAccessResponse
	14, // 27: iraaccess.v1.AccessManager.GetUserAccessTree:output_type -> iraaccess.v1.GetUserAccessTreeResponse
	20, // 28: iraaccess.v1.AccessManager.GetResourceTypes:output_type -> iraaccess.v1.GetResourceTypesResponse
	22, // [22:29] is the sub-list for method output_type
	15, // [15:22] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_proto_authzv1_authz_proto_init() }
func file_proto_authzv1_authz_proto_init() {
	if File_proto_authzv1_authz_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_authzv1_authz_proto_rawDesc), len(file_proto_authzv1_authz_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_authzv1_authz_proto_goTypes,
		DependencyIndexes: file_proto_authzv1_authz_proto_depIdxs,
		EnumInfos:         file_proto_authzv1_authz_proto_enumTypes,
		MessageInfos:      file_proto_authzv1_authz_proto_msgTypes,
	}.Build()
	File_proto_authzv1_authz_proto = out.File
	file_proto_authzv1_authz_proto_goTypes = nil
	file_proto_authzv1_authz_proto_depIdxs = nil
}
