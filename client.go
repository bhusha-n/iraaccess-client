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

type Client struct {
	conn      *grpc.ClientConn
	api       authzv1.IraAccessServiceClient
	appID     string
	appSecret string
}

func NewConnection(jsonPath string) (*Client, error) {
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

	return &Client{
		conn:      conn,
		api:       authzv1.NewIraAccessServiceClient(conn),
		appID:     cfg.AppID,
		appSecret: cfg.AppSecret,
	}, nil
}

func (c *Client) IraCloseConn() error {
	return c.conn.Close()
}

func (c *Client) authed(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "app_id", c.appID, "app_secret", c.appSecret)
}

func (c *Client) IraGrantTenantAccess(ctx context.Context, grantedBy, userID, tenantID string) error {
	return c.grantAccess(ctx, grantedBy, userID, tenantID, false)
}

func (c *Client) IraGrantAdminAccess(ctx context.Context, grantedBy, userID string) error {
	return c.grantAccess(ctx, grantedBy, userID, "", true)
}

func (c *Client) grantAccess(ctx context.Context, grantedBy, userID, tenantID string, asAdmin bool) error {
	resp, err := c.api.GrantAccess(c.authed(ctx), &authzv1.GrantAccessRequest{
		GrantedBy: grantedBy,
		UserId:    userID,
		TenantId:  tenantID,
		AsAdmin:   asAdmin,
	})
	if err != nil {
		return err
	}
	if !resp.GetSuccess() {
		return fmt.Errorf("iam_client: grant access was not successful")
	}
	return nil
}

func (c *Client) IraCheckAccess(ctx context.Context, userID, tenantID string) (bool, error) {
	resp, err := c.api.CheckAccess(c.authed(ctx), &authzv1.CheckAccessRequest{
		UserId:   userID,
		TenantId: tenantID,
	})
	if err != nil {
		return false, err
	}

	return resp.GetAllowed(), nil
}

func (c *Client) IraCheckAdminAccess(ctx context.Context, userId string) (bool, error) {
	resp, err := c.api.CheckAdminAccess(c.authed(ctx), &authzv1.CheckAdminAccessRequest{
		UserId: userId,
	})
	c.api.CheckAdminAccess(c.authed(ctx), &authzv1.CheckAdminAccessRequest{})
	if err != nil {
		return false, err
	}

	return resp.GetAllowed(), nil
}
