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

type iraIAMConfig struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	IamURL    string `json:"iam_url"`
}

type client struct {
	conn      *grpc.ClientConn
	api       authzv1.IraAccessServiceClient
	appID     string
	appSecret string
}

func InitializeIraAccess(jsonPath string) (*client, error) {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("iam_client: failed to read %s: %w", jsonPath, err)
	}

	var cfg iraIAMConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("iam_client: failed to parse %s: %w", jsonPath, err)
	}
	if cfg.AppID == "" || cfg.AppSecret == "" || cfg.IamURL == "" {
		return nil, fmt.Errorf("iam_client: %s must set app_id, app_secret and iam_url", jsonPath)
	}

	conn, err := grpc.NewClient(cfg.IamURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("iam_client: failed to connect to %s: %w", cfg.IamURL, err)
	}

	return &client{
		conn:      conn,
		api:       authzv1.NewIraAccessServiceClient(conn),
		appID:     cfg.AppID,
		appSecret: cfg.AppSecret,
	}, nil
}

func (c *client) CloseConn() error {
	return c.conn.Close()
}

func (c *client) authed(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "app_id", c.appID, "app_secret", c.appSecret)
}

func (c *client) GrantTenantAccessAs(ctx context.Context, targetAppID, grantedBy, userID, tenantID string) error {
	return c.grantAccess(ctx, targetAppID, grantedBy, userID, tenantID, false)
}

func (c *client) GrantAdminAccessAs(ctx context.Context, targetAppID, grantedBy, userID string) error {
	return c.grantAccess(ctx, targetAppID, grantedBy, userID, "", true)
}

func (c *client) grantAccess(ctx context.Context, targetAppID, grantedBy, userID, tenantID string, asAdmin bool) error {
	resp, err := c.api.GrantAccess(c.authed(ctx), &authzv1.GrantAccessRequest{
		GrantedBy:   grantedBy,
		UserId:      userID,
		TenantId:    tenantID,
		AsAdmin:     asAdmin,
		TargetAppId: targetAppID,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return fmt.Errorf("ERROR : grant access was not successful")
	}
	return nil
}

func (c *client) CheckAccess(ctx context.Context, userID, tenantID string) (bool, error) {
	resp, err := c.api.CheckAccess(c.authed(ctx), &authzv1.CheckAccessRequest{
		UserId:   userID,
		TenantId: tenantID,
	})
	if err != nil {
		return false, err
	}

	return resp.GetAllowed(), nil
}

func (c *client) DeleteAccess(ctx context.Context, targetAppID, userID, TenantID, DeletedBy string) error {
	resp, err := c.api.DeleteAccess(c.authed(ctx), &authzv1.DeleteAccessRequest{
		UserId:      userID,
		TenantId:    TenantID,
		DeletedBy:   DeletedBy,
		TargetAppId: targetAppID,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return fmt.Errorf("ERROR : delete access was not successful")
	}
	return nil
}

func (c *client) DeleteAdminAccess(ctx context.Context, targetAppID, userID, DeletedBy string) error {
	resp, err := c.api.DeleteAdminAccess(c.authed(ctx), &authzv1.DeleteAdminAccessRequest{
		UserId:      userID,
		DeletedBy:   DeletedBy,
		TargetAppId: targetAppID,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return fmt.Errorf("ERROR : delete access was not successfull")
	}
	return nil
}

func (c *client) CheckAdminAccess(ctx context.Context, userId string) (bool, error) {
	resp, err := c.api.CheckAdminAccess(c.authed(ctx), &authzv1.CheckAdminAccessRequest{
		UserId: userId,
	})

	if err != nil {
		return false, err
	}

	return resp.GetAllowed(), nil
}
