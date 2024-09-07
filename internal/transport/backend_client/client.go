package backend_client

import (
	"context"
	"fmt"
	"io"

	"github.com/rs/zerolog/log"
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

func (c *BackendClient) LoginPassData(ctx context.Context, token string, id int32) (*entity.DataLoginPass, error) {
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

func (c *BackendClient) LoginPassAdd(ctx context.Context, token string, title, login, pass string) (int32, error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.client.AddDataLoginPass(ctx, &grpc_api.AddDataLoginPassRequest{Title: title, Login: login, Pass: pass})
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}

	return resp.Id, nil
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

func (c *BackendClient) TextAdd(ctx context.Context, token string, title, text string) (id int32, err error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.client.AddDataText(ctx, &grpc_api.AddDataTextRequest{Title: title, Text: text})
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}

	return resp.Id, nil
}

func (c *BackendClient) ListText(ctx context.Context, token string) ([]entity.DataText, error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.client.GetDataTextList(ctx, &grpc_api.GetDataListRequest{})
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

	list := make([]entity.DataText, 0, len(resp.List))
	for _, data := range resp.List {
		list = append(list, entity.DataText{
			ID:    data.Id,
			Title: data.Title,
		})
	}

	return list, nil
}

func (c *BackendClient) TextData(ctx context.Context, token string, id int32) (*entity.DataText, error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.client.GetDataText(ctx, &grpc_api.GetDataRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	data := entity.DataText{
		ID:    resp.Id,
		Title: resp.Title,
		Text:  resp.Text,
	}

	return &data, nil
}

func (c *BackendClient) ListCard(ctx context.Context, token string) ([]entity.DataCard, error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.client.GetDataCardList(ctx, &grpc_api.GetDataListRequest{})
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

	list := make([]entity.DataCard, 0, len(resp.List))
	for _, data := range resp.List {
		list = append(list, entity.DataCard{
			ID:    data.Id,
			Title: data.Title,
		})
	}

	return list, nil
}

func (c *BackendClient) CardData(ctx context.Context, token string, id int32) (*entity.DataCard, error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	resp, err := c.client.GetDataCard(ctx, &grpc_api.GetDataRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	data := entity.DataCard{
		ID:      resp.Id,
		Title:   resp.Title,
		Number:  resp.Number,
		Owner:   resp.Owner,
		ExpDate: resp.ExpDate,
		CVCCode: resp.CvcCode,
	}

	return &data, nil
}

func (c *BackendClient) CardAdd(ctx context.Context, token string, card *entity.DataCard) (id int32, err error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	req := grpc_api.AddDataCardRequest{
		Title:   card.Title,
		Number:  card.Number,
		Owner:   card.Owner,
		ExpDate: card.ExpDate,
		CvcCode: card.CVCCode,
	}
	resp, err := c.client.AddDataCard(ctx, &req)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}

	return resp.Id, nil
}

func (c *BackendClient) BinaryAdd(ctx context.Context, token string, binary *entity.DataBinary) (id int32, err error) {
	md := metadata.Pairs("auth_token", token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	stream, err := c.client.AddDataBinary(ctx)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}

	for b := range binary.Body {
		err := stream.Send(&grpc_api.AddDataBinaryRequest{Body: []byte{b}})
		if err != nil {
			if err == io.EOF {
				log.Err(err).Msg("error stream request")
				continue
			}
			return 0, fmt.Errorf("failed to send stream request: %w", err)
		}
	}

	streamResp, err := stream.CloseAndRecv()
	if err != nil {
		return 0, fmt.Errorf("failed to close and receive: %w", err)
	}

	resp, err := c.client.AddDataBinaryTitle(ctx, &grpc_api.AddDataBinaryTitleRequest{Id: streamResp.Id, Title: binary.Title})
	if err != nil {
		return 0, fmt.Errorf("failed to send binary title: %w", err)
	}

	return resp.Id, nil
}

func (c *BackendClient) ListBinary(ctx context.Context, token string) ([]entity.DataBinary, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BackendClient) BinaryData(ctx context.Context, token string, id int32) (*entity.DataBinary, error) {
	//TODO implement me
	panic("implement me")
}
