package auth

import (
	"context"

	access "github.com/Danya97i/auth/pkg/access_v1"
)

// AccessClient is an interface for access client
type AccessClient interface {
	Verifiy(ctx context.Context, endpoint string) error
}

// NewAccessClient creates a new access client
func NewAccessClient(accessClient access.AccessV1Client) AccessClient {
	return &client{
		accessClient: accessClient,
	}
}

type client struct {
	accessClient access.AccessV1Client
}

// Verifiy checks if the endpoint is accessible
func (c *client) Verifiy(ctx context.Context, endpoint string) error {
	req := &access.CheckRequest{
		EndpointAddress: endpoint,
	}

	_, err := c.accessClient.Check(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
