package grpchandler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/lks-go/pass-keeper/internal/service/entity"
	"github.com/lks-go/pass-keeper/internal/service/server"
	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

type Service interface {
	RegisterUser(ctx context.Context, login, password string) (string, error)
	AuthUser(ctx context.Context, login string, password string) (string, error)
	AddDataLoginPass(ctx context.Context, ownerLogin string, data server.Data) error
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
			log.Error().Err(err)
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
			log.Error().Err(err)
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	return &grpc_api.AuthUserResponse{Token: jwtString}, nil
}

func (h *Handler) AddDataLoginPass(ctx context.Context, request *grpc_api.AddDataLoginPassRequest) (*grpc_api.AddDataLoginPassResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	payload := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{
		Login:    request.Login,
		Password: request.Pass,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal payload")
		return nil, status.Error(codes.Internal, (codes.Internal).String())
	}

	data := server.Data{
		Title:   request.Title,
		Payload: string(payloadBytes),
	}
	err = h.service.AddDataLoginPass(ctx, ownerLogin, data)
	if err != nil {
		switch {
		case errors.Is(err, server.ErrUserNotFound):
			log.Warn().Str("login", request.Login).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		default:
			log.Error().Err(err).Msg("failed to add data login and pass")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	return &grpc_api.AddDataLoginPassResponse{}, nil
}

func userLogin(ctx context.Context) (string, error) {
	data, err := outgoingMetaData(ctx, entity.UserLoginHeaderName)
	if err != nil {
		return "", fmt.Errorf("failed to get metadata: %w", err)
	}

	log.Debug().Str("data", data[0]).Msg("outgoing data")

	return data[0], nil
}

func outgoingMetaData(ctx context.Context, key string) ([]string, error) {
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	value := md.Get(key)
	if len(value) == 0 {
		return nil, fmt.Errorf("%s not supplied", key)
	}

	return value, nil
}
