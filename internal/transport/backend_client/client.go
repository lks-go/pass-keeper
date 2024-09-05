package backend_client

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func (c *BackendClient) ListLoginPass(ctx context.Context) ([]entity.DataLoginPass, error) {
	log.Info().Msg("Getting request to backend")

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
