package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Danya97i/chat-server/internal/client/auth"
)

// NewAccessInterceptor returns a grpc.UnaryServerInterceptor that verifies access token
func NewAccessInterceptor(accessClient auth.AccessClient) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, fmt.Errorf("metadata is not provided")
		}

		ctx = metadata.NewOutgoingContext(ctx, md)
		if err := accessClient.Verifiy(ctx, info.FullMethod); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
