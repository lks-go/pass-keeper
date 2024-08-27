package grpchandler

import (
	"context"
	"fmt"

	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
	grpc_api.UnimplementedPassKeeperServer
}

func (h *Handler) RegisterUser(ctx context.Context, request *grpc_api.RegisterUserRequest) (*grpc_api.RegisterUserResponse, error) {
	fmt.Println(request)

	////TODO implement me
	//panic("implement me")

	return &grpc_api.RegisterUserResponse{}, nil
}

func (h *Handler) AuthUser(ctx context.Context, request *grpc_api.AuthUserRequest) (*grpc_api.AuthUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
