package backend_client

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/lks-go/pass-keeper/internal/service/entity"
	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

type Config struct {
	Host      string
	CertPath  string
	EnableTLS bool
}

func New(cfg *Config) (*BackendClient, error) {
	opts := []grpc.DialOption{}

	if cfg.EnableTLS {
		creds, err := credentials.NewClientTLSFromFile(cfg.CertPath, "")
		if err != nil {
			return nil, fmt.Errorf("could not load tls cert: %w", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.NewClient(cfg.Host, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to makee new client: %w", err)
	}

	client := grpc_api.NewPassKeeperClient(conn)

	return &BackendClient{client: client}, nil
}

type BackendClient struct {
	client grpc_api.PassKeeperClient
}

func (c *BackendClient) ListLoginPass(ctx context.Context, token string) ([]entity.DataLoginPass, error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.client.GetDataLoginPassList(ctx, &grpc_api.GetDataListRequest{})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Message() {
			case entity.ErrMissingToken.Error():
				return nil, entity.ErrMissingToken
			default:
				return nil, fmt.Errorf("request failed: %w", err)
			}
		}

		return nil, fmt.Errorf("request failed: %w", err)
	}

	list := make([]entity.DataLoginPass, 0, len(resp.List))
	for _, data := range resp.List {
		list = append(list, entity.DataLoginPass{
			ID:    data.Id,
			Title: data.Title,
		})
	}

	return list, nil
}

func (c *BackendClient) LoginPassData(ctx context.Context, id int32, token string) (*entity.DataLoginPass, error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.client.GetDataLoginPass(ctx, &grpc_api.GetDataRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	data := entity.DataLoginPass{
		ID:       resp.Id,
		Title:    resp.Title,
		Login:    resp.Login,
		Password: resp.Pass,
	}

	return &data, nil
}

func (c *BackendClient) Reg(ctx context.Context, login string, password string) error {
	req := grpc_api.RegisterUserRequest{
		Login:    login,
		Password: password,
	}
	_, err := c.client.RegisterUser(ctx, &req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	return nil
}

func (c *BackendClient) Auth(ctx context.Context, login string, password string) (token string, err error) {
	req := grpc_api.AuthUserRequest{
		Login:    login,
		Password: password,
	}
	resp, err := c.client.AuthUser(ctx, &req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}

	return resp.Token, nil
}
