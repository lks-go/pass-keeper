package grpchandler

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lks-go/pass-keeper/internal/service/server"
	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

type Service interface {
	RegisterUser(ctx context.Context, login, password string) (string, error)
	AuthUser(ctx context.Context, login string, password string) (string, error)
	AddDataLoginPass(ctx context.Context, userLogin string, data server.DataLoginPass) error
}

func New(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

type Handler struct {
	service Service

	grpc_api.UnimplementedPassKeeperServer
}

func (h *Handler) RegisterUser(ctx context.Context, request *grpc_api.RegisterUserRequest) (*grpc_api.RegisterUserResponse, error) {
	userId, err := h.service.RegisterUser(ctx, request.Login, request.Password)
	if err != nil {
		switch {
		case errors.Is(err, server.ErrAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, (codes.AlreadyExists).String())
		default:
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	log.Info().Str("user_id", userId).Msg("User registered")

	return &grpc_api.RegisterUserResponse{}, nil
}

func (h *Handler) AuthUser(ctx context.Context, request *grpc_api.AuthUserRequest) (*grpc_api.AuthUserResponse, error) {
	jwtString, err := h.service.AuthUser(ctx, request.Login, request.Password)
	if err != nil {
		switch {
		case errors.Is(err, server.ErrUsersPasswordNotMatch):
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		case errors.Is(err, server.ErrNotFound):
			return nil, status.Error(codes.NotFound, (codes.NotFound).String())
		default:
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	return &grpc_api.AuthUserResponse{Token: jwtString}, nil
}

func (h *Handler) AddDataLoginPass(ctx context.Context, request *grpc_api.AddDataLoginPassRequest) (*grpc_api.AddDataLoginPassResponse, error) {
	data := server.DataLoginPass{
		Title:    request.Title,
		Login:    request.Login,
		Password: request.Pass,
	}
	err := h.service.AddDataLoginPass(ctx, request.Login, data)
	if err != nil {
		switch {
		case errors.Is(err, server.ErrUserNotFound):
			log.Warn().Str("login", request.Login).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		}
	}

	return &grpc_api.AddDataLoginPassResponse{}, nil
}
