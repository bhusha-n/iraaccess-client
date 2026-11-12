package iraaccess

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	authzv1 "github.com/bhusha-n/iraaccess-client/proto/authzv1"
)

type (
	Action        = authzv1.Action
	Role          = authzv1.Role
	ResourceScope = authzv1.ResourceScope
)

const (
	ActionView   = authzv1.Action_ACTION_VIEW
	ActionCreate = authzv1.Action_ACTION_CREATE
	ActionUpdate = authzv1.Action_ACTION_UPDATE
	ActionDelete = authzv1.Action_ACTION_DELETE

	RoleTenantMember = authzv1.Role_ROLE_TENANT_MEMBER
	RoleTenantAdmin  = authzv1.Role_ROLE_TENANT_ADMIN
	RoleSuperAdmin   = authzv1.Role_ROLE_SUPER_ADMIN

	ScopeAll      = authzv1.ResourceScope_RESOURCE_SCOPE_ALL
	ScopeSelected = authzv1.ResourceScope_RESOURCE_SCOPE_SELECTED
)

type iraIAMConfig struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	IamURL    string `json:"iam_url"`
}

type Client struct {
	conn      *grpc.ClientConn
	api       authzv1.AccessManagerClient
	appID     string
	appSecret string
}

func InitializeIraAccess(jsonPath string) (*Client, error) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", jsonPath, err)
	}
	var cfg iraIAMConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse %s: %w", jsonPath, err)
	}
	if cfg.AppID == "" || cfg.AppSecret == "" || cfg.IamURL == "" {
		return nil, fmt.Errorf("%s must set app_id, app_secret and iam_url", jsonPath)
	}
	conn, err := grpc.NewClient(cfg.IamURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", cfg.IamURL, err)
	}
	return &Client{
		conn:      conn,
		api:       authzv1.NewAccessManagerClient(conn),
		appID:     cfg.AppID,
		appSecret: cfg.AppSecret,
	}, nil
}

func (c *Client) Close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

func (c *Client) authed(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "app_id", c.appID, "app_secret", c.appSecret)
}

func (c *Client) AssignSuperAdmin(ctx context.Context, targetAppID string, userID string, grantedBy string) error {
	return c.setRole(ctx, targetAppID, "", userID, RoleSuperAdmin, grantedBy)
}

func (c *Client) AssignTenantAdmin(ctx context.Context, targetAppID string, tenantID string, userID string, grantedBy string) error {
	return c.setRole(ctx, targetAppID, tenantID, userID, RoleTenantAdmin, grantedBy)
}

func (c *Client) AssignTenantMember(ctx context.Context, targetAppID string, tenantID string, userID string, resourceType string, scope ResourceScope, resourceIDs []string, actions []Action, grantedBy string) error {
	return c.grantResourceAccess(ctx, targetAppID, tenantID, userID, resourceType, scope, resourceIDs, actions, grantedBy)
}

func (c *Client) RemoveSuperAdmin(ctx context.Context, targetAppID string, userID string, deletedBy string) error {
	return c.removeRole(ctx, targetAppID, "", userID, RoleSuperAdmin, deletedBy)
}

func (c *Client) RemoveTenantAdmin(ctx context.Context, targetAppID string, tenantID string, userID string, deletedBy string) error {
	return c.removeRole(ctx, targetAppID, tenantID, userID, RoleTenantAdmin, deletedBy)
}

func (c *Client) RemoveTenantMember(ctx context.Context, targetAppID string, tenantID string, userID string, deletedBy string) error {
	return c.removeRole(ctx, targetAppID, tenantID, userID, RoleTenantMember, deletedBy)
}

func (c *Client) setRole(ctx context.Context, targetAppID, tenantID, userID string, role Role, grantedBy string) error {
	resp, err := c.api.SetRole(c.authed(ctx), &authzv1.SetRoleRequest{
		TargetAppId: targetAppID, TenantId: tenantID, UserId: userID, Role: role, GrantedBy: grantedBy,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return fmt.Errorf("IAM service did not successfully assign role %s", role)
	}
	return nil
}

func (c *Client) removeRole(ctx context.Context, targetAppID, tenantID, userID string, role Role, deletedBy string) error {
	resp, err := c.api.RemoveRole(c.authed(ctx), &authzv1.RemoveRoleRequest{
		TargetAppId: targetAppID, TenantId: tenantID, UserId: userID, Role: role, DeletedBy: deletedBy,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return fmt.Errorf("IAM service did not successfully remove role %s", role)
	}
	return nil
}

func (c *Client) grantResourceAccess(ctx context.Context, targetAppID, tenantID, userID, resourceType string,
	scope ResourceScope, resourceIDs []string, actions []Action, grantedBy string) error {
	resp, err := c.api.GrantResourceAccess(c.authed(ctx), &authzv1.GrantResourceAccessRequest{
		TargetAppId: targetAppID, TenantId: tenantID, UserId: userID, ResourceType: resourceType,
		Scope: scope, ResourceIds: resourceIDs, Actions: actions, GrantedBy: grantedBy,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return fmt.Errorf("IAM service did not successfully grant resource access")
	}
	return nil
}

func (c *Client) RevokeResourceAccess(ctx context.Context, targetAppID, tenantID, userID, resourceType string,
	scope ResourceScope, resourceIDs []string, actions []Action, deletedBy string) error {
	resp, err := c.api.RevokeResourceAccess(c.authed(ctx), &authzv1.RevokeResourceAccessRequest{
		TargetAppId: targetAppID, TenantId: tenantID, UserId: userID, ResourceType: resourceType,
		Scope: scope, ResourceIds: resourceIDs, Actions: actions, DeletedBy: deletedBy,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return fmt.Errorf("revoke resource access was not successful")
	}
	return nil
}

func (c *Client) CheckAccess(ctx context.Context, userID, tenantID, resourceType, resourceID string, action Action) (bool, error) {
	resp, err := c.api.CheckAccess(c.authed(ctx), &authzv1.CheckAccessRequest{
		UserId: userID, TenantId: tenantID, ResourceType: resourceType, ResourceId: resourceID, Action: action,
	})
	if err != nil {
		return false, err
	}
	return resp.GetAllowed(), nil
}

func (c *Client) GetUserAccessTree(ctx context.Context, userID, targetAppID string) (*authzv1.GetUserAccessTreeResponse, error) {
	return c.api.GetUserAccessTree(c.authed(ctx), &authzv1.GetUserAccessTreeRequest{
		UserId: userID, TargetAppId: targetAppID,
	})
}

// func (c *Client) getResourceTypes(ctx context.Context, targetAppID string) ([]string, error) {
// 	resp, err := c.api.GetResourceTypes(c.authed(ctx), &authzv1.GetResourceTypesRequest{TargetAppId: targetAppID})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.GetResourceTypes(), nil
// }
