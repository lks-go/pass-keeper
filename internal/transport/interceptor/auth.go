package interceptor

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/lks-go/pass-keeper/internal/lib/token"
	"github.com/lks-go/pass-keeper/internal/service/entity"
)

func NewAuth(t *token.Token) *Auth {
	return &Auth{token: t}
}

type Auth struct {
	token *token.Token
}

func (a *Auth) CheckAccess(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if strings.Contains(info.FullMethod, "RegisterUser") || strings.Contains(info.FullMethod, "AuthUser") {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}

	var login string
	var claims *token.Claims
	var err error

	authToken, ok := md[entity.AuthTokenHeader]
	if !ok {
		return nil, status.Error(codes.PermissionDenied, fmt.Sprintf("%s: %s", (codes.PermissionDenied).String(), "missing auth token"))
	}

	claims, err = a.token.ParseJWTToken(authToken[0])
	if err != nil {
		switch {
		case errors.Is(err, token.ErrInvalidToken):
			return nil, status.Error(codes.InvalidArgument, token.ErrInvalidToken.Error())
		case errors.Is(err, token.ErrTokenExpired):
			return nil, status.Error(codes.InvalidArgument, token.ErrTokenExpired.Error())
		default:
			log.Error().Err(err).Msg("failed to parse jwt")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	if claims != nil && claims.Login == "" {
		return nil, status.Error(codes.Unauthenticated, (codes.Unauthenticated).String())
	}

	if claims != nil {
		login = claims.Login
	}

	ctx = metadata.AppendToOutgoingContext(ctx, entity.UserLoginHeaderName, login)
	return handler(ctx, req)
}
