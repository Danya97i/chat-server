package env

import (
	"errors"
	"os"

	"github.com/Danya97i/chat-server/internal/config"
)

var _ config.AuthClientConfig = (*authClientConifg)(nil)

type authClientConifg struct {
	address  string
	certFile string
}

// NewAuthClientConfig returns a new AuthClientConfig
func NewAuthClientConfig() (*authClientConifg, error) {
	address := os.Getenv("AUTH_CLIENT_ADDRESS")
	if len(address) == 0 {
		return nil, errors.New("AUTH_CLIENT_ADDRESS is empty")
	}
	certFile := os.Getenv("AUTH_CLIENT_CERT_FILE")
	if len(certFile) == 0 {
		return nil, errors.New("AUTH_CLIENT_CERT_FILE is empty")
	}
	return &authClientConifg{address: address, certFile: certFile}, nil
}

func (c *authClientConifg) Address() string {
	return c.address
}

func (c *authClientConifg) CertFile() string {
	return c.certFile
}
