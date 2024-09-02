package grpchandler

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/lks-go/pass-keeper/internal/service"
	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

type Service interface {
	RegisterUser(ctx context.Context, u service.User) (string, error)
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
	u := service.User{
		Login:    request.Login,
		Password: request.Password,
	}

	userId, err := h.service.RegisterUser(ctx, u)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, (codes.AlreadyExists).String())
		default:
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	log.Info().Str("user_id", userId).Msg("User registered")

	return &grpc_api.RegisterUserResponse{}, nil
}

func (h *Handler) AuthUser(ctx context.Context, request *grpc_api.AuthUserRequest) (*grpc_api.AuthUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
