package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func Auth(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// TODO make auth with JWT
	return handler(ctx, req)
}
