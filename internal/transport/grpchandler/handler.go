package grpchandler

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/lks-go/pass-keeper/internal/service/entity"
	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

type Service interface {
	RegisterUser(ctx context.Context, login, password string) (string, error)
	AuthUser(ctx context.Context, login string, password string) (string, error)

	AddDataLoginPass(ctx context.Context, ownerLogin string, data *entity.DataLoginPass) (int32, error)
	DataLoginPassList(ctx context.Context, ownerLogin string) ([]entity.DataLoginPass, error)
	DataLoginPass(ctx context.Context, ownerLogin string, ID int32) (*entity.DataLoginPass, error)

	AddDataText(ctx context.Context, ownerLogin string, data *entity.DataText) (int32, error)
	DataTextList(ctx context.Context, ownerLogin string) ([]entity.DataText, error)
	DataText(ctx context.Context, ownerLogin string, ID int32) (*entity.DataText, error)

	AddDataCard(ctx context.Context, ownerLogin string, data *entity.DataCard) (int32, error)
	DataCardList(ctx context.Context, ownerLogin string) ([]entity.DataCard, error)
	DataCard(ctx context.Context, ownerLogin string, ID int32) (*entity.DataCard, error)

	AddDataBinary(ctx context.Context, ownerLogin string, binary *entity.DataBinary) (int32, error)
	AddDataBinaryTitle(ctx context.Context, ownerLogin string, binary *entity.DataBinary) (int32, error)
	DataBinaryList(ctx context.Context, ownerLogin string) ([]entity.DataBinary, error)
	DataBinary(ctx context.Context, ownerLogin string, binaryID int32) (stream <-chan byte, err <-chan error)
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
		case errors.Is(err, entity.ErrAlreadyExists):
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
		case errors.Is(err, entity.ErrUsersPasswordNotMatch):
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		case errors.Is(err, entity.ErrUserNotFound):
			log.Warn().Str("login", request.Login).Msg("user not found")
			return nil, status.Error(codes.NotFound, (codes.NotFound).String())
		default:
			log.Error().Err(err)
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	return &grpc_api.AuthUserResponse{Token: jwtString}, nil
}

func (h *Handler) AddDataLoginPass(ctx context.Context, request *grpc_api.AddDataLoginPassRequest) (*grpc_api.AddDataResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data := entity.DataLoginPass{
		Title:    request.Title,
		Login:    request.Login,
		Password: request.Pass,
	}
	id, err := h.service.AddDataLoginPass(ctx, ownerLogin, &data)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			log.Warn().Str("login", ownerLogin).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		default:
			log.Error().Err(err).Msg("failed to add data login and pass")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	return &grpc_api.AddDataResponse{Id: id}, nil
}

func (h *Handler) GetDataLoginPassList(ctx context.Context, _ *grpc_api.GetDataListRequest) (*grpc_api.GetDataListResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data, err := h.service.DataLoginPassList(ctx, ownerLogin)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			log.Warn().Str("login", ownerLogin).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		case errors.Is(err, entity.ErrNoData):
			return nil, status.Error(codes.NotFound, (codes.NotFound).String())
		default:
			log.Error().Err(err).Msg("failed to get login and pass list")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	list := make([]*grpc_api.GetDataListResponse_Data, 0, len(data))
	for _, d := range data {
		respData := grpc_api.GetDataListResponse_Data{
			Id:    d.ID,
			Title: d.Title,
		}

		list = append(list, &respData)
	}

	return &grpc_api.GetDataListResponse{List: list}, nil
}

func (h *Handler) GetDataLoginPass(ctx context.Context, request *grpc_api.GetDataRequest) (*grpc_api.GetDataLoginPassResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data, err := h.service.DataLoginPass(ctx, ownerLogin, request.Id)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			log.Warn().Str("login", ownerLogin).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		case errors.Is(err, entity.ErrNoData):
			return nil, status.Error(codes.NotFound, (codes.NotFound).String())
		default:
			log.Error().Err(err).Msg("failed to get login and pass list")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	response := grpc_api.GetDataLoginPassResponse{
		Id:    data.ID,
		Title: data.Title,
		Login: data.Login,
		Pass:  data.Password,
	}

	return &response, nil
}

func (h *Handler) AddDataText(ctx context.Context, request *grpc_api.AddDataTextRequest) (*grpc_api.AddDataResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data := entity.DataText{
		Title: request.Title,
		Text:  request.Text,
	}
	id, err := h.service.AddDataText(ctx, ownerLogin, &data)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			log.Warn().Str("login", ownerLogin).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		default:
			log.Error().Err(err).Msg("failed to add data login and pass")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	return &grpc_api.AddDataResponse{Id: id}, nil
}

func (h *Handler) GetDataTextList(ctx context.Context, _ *grpc_api.GetDataListRequest) (*grpc_api.GetDataListResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data, err := h.service.DataTextList(ctx, ownerLogin)
	if err != nil {
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrUserNotFound):
				log.Warn().Str("login", ownerLogin).Msg("user not found")
				return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
			case errors.Is(err, entity.ErrNoData):
				return nil, status.Error(codes.NotFound, (codes.NotFound).String())
			default:
				log.Error().Err(err).Msg("failed to get text list")
				return nil, status.Error(codes.Internal, (codes.Internal).String())
			}
		}
	}

	list := make([]*grpc_api.GetDataListResponse_Data, 0, len(data))
	for _, d := range data {
		respData := grpc_api.GetDataListResponse_Data{
			Id:    d.ID,
			Title: d.Title,
		}

		list = append(list, &respData)
	}

	return &grpc_api.GetDataListResponse{List: list}, nil
}

func (h *Handler) GetDataText(ctx context.Context, request *grpc_api.GetDataRequest) (*grpc_api.GetDataTextResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data, err := h.service.DataText(ctx, ownerLogin, request.Id)
	if err != nil {
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrUserNotFound):
				log.Warn().Str("login", ownerLogin).Msg("user not found")
				return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
			case errors.Is(err, entity.ErrNoData):
				return nil, status.Error(codes.NotFound, (codes.NotFound).String())
			default:
				log.Error().Err(err).Msg("failed to get text")
				return nil, status.Error(codes.Internal, (codes.Internal).String())
			}
		}
	}

	response := grpc_api.GetDataTextResponse{
		Id:    data.ID,
		Title: data.Title,
		Text:  data.Text,
	}

	return &response, nil
}

func (h *Handler) AddDataCard(ctx context.Context, request *grpc_api.AddDataCardRequest) (*grpc_api.AddDataResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data := entity.DataCard{
		Title:   request.Title,
		Number:  request.Number,
		Owner:   request.Owner,
		ExpDate: request.ExpDate,
		CVCCode: request.CvcCode,
	}
	id, err := h.service.AddDataCard(ctx, ownerLogin, &data)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			log.Warn().Str("login", ownerLogin).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		default:
			log.Error().Err(err).Msg("failed to add card data")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	return &grpc_api.AddDataResponse{Id: id}, nil
}

func (h *Handler) GetDataCardList(ctx context.Context, _ *grpc_api.GetDataListRequest) (*grpc_api.GetDataListResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data, err := h.service.DataCardList(ctx, ownerLogin)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			log.Warn().Str("login", ownerLogin).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		case errors.Is(err, entity.ErrNoData):
			return nil, status.Error(codes.NotFound, (codes.NotFound).String())
		default:
			log.Error().Err(err).Msg("failed to get card list")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	list := make([]*grpc_api.GetDataListResponse_Data, 0, len(data))
	for _, d := range data {
		respData := grpc_api.GetDataListResponse_Data{
			Id:    d.ID,
			Title: d.Title,
		}

		list = append(list, &respData)
	}

	return &grpc_api.GetDataListResponse{List: list}, nil
}

func (h *Handler) GetDataCard(ctx context.Context, request *grpc_api.GetDataRequest) (*grpc_api.GetDataCardResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data, err := h.service.DataCard(ctx, ownerLogin, request.Id)
	if err != nil {
		if err != nil {
			switch {
			case errors.Is(err, entity.ErrUserNotFound):
				log.Warn().Str("login", ownerLogin).Msg("user not found")
				return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
			case errors.Is(err, entity.ErrNoData):
				return nil, status.Error(codes.NotFound, (codes.NotFound).String())
			default:
				log.Error().Err(err).Msg("failed to get card")
				return nil, status.Error(codes.Internal, (codes.Internal).String())
			}
		}
	}

	response := grpc_api.GetDataCardResponse{
		Id:      data.ID,
		Title:   data.Title,
		Number:  data.Number,
		Owner:   data.Owner,
		ExpDate: data.ExpDate,
		CvcCode: data.CVCCode,
	}

	return &response, nil
}

func (h *Handler) AddDataBinary(stream grpc.ClientStreamingServer[grpc_api.AddDataBinaryRequest, grpc_api.AddDataResponse]) error {
	ownerLogin, err := userLogin(stream.Context())
	if err != nil {
		return status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	var resultID int32
	ch := make(chan byte, 0)
	binary := entity.DataBinary{
		Body: ch,
	}

	g, ctx := errgroup.WithContext(stream.Context())

	g.Go(func() error {
		id, err := h.service.AddDataBinary(ctx, ownerLogin, &binary)
		if err != nil {
			return fmt.Errorf("filed to add binary data: %w", err)
		}

		resultID = id
		return nil
	})

	g.Go(func() error {
		defer close(ch)

	LOOP:
		for {
			req, err := stream.Recv()
			if err != nil {
				switch {
				case err == io.EOF:
					break LOOP
				default:
					return fmt.Errorf("failed to get streaming data: %w", err)
				}
			}

			for _, bodyByte := range req.Body {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case ch <- bodyByte:
				}
			}
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		log.Error().Err(err).Msg("group error")
		return status.Error(codes.Internal, (codes.Internal).String())
	}

	return stream.SendAndClose(&grpc_api.AddDataResponse{
		Id: resultID,
	})
}

func (h *Handler) AddDataBinaryTitle(ctx context.Context, request *grpc_api.AddDataBinaryTitleRequest) (*grpc_api.AddDataResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data := entity.DataBinary{
		ID:    request.Id,
		Title: request.Title,
	}

	id, err := h.service.AddDataBinaryTitle(ctx, ownerLogin, &data)
	if err != nil {
		log.Error().Err(err)
		return nil, status.Error(codes.Internal, (codes.Internal).String())
	}

	return &grpc_api.AddDataResponse{Id: id}, nil
}

func (h *Handler) GetDataBinaryList(ctx context.Context, _ *grpc_api.GetDataListRequest) (*grpc_api.GetDataListResponse, error) {
	ownerLogin, err := userLogin(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	data, err := h.service.DataBinaryList(ctx, ownerLogin)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrUserNotFound):
			log.Warn().Str("login", ownerLogin).Msg("user not found")
			return nil, status.Error(codes.PermissionDenied, (codes.PermissionDenied).String())
		case errors.Is(err, entity.ErrNoData):
			return nil, status.Error(codes.NotFound, (codes.NotFound).String())
		default:
			log.Error().Err(err).Msg("failed to get card list")
			return nil, status.Error(codes.Internal, (codes.Internal).String())
		}
	}

	list := make([]*grpc_api.GetDataListResponse_Data, 0, len(data))
	for _, d := range data {
		respData := grpc_api.GetDataListResponse_Data{
			Id:    d.ID,
			Title: d.Title,
		}

		list = append(list, &respData)
	}

	return &grpc_api.GetDataListResponse{List: list}, nil
}

func (h *Handler) GetDataBinary(request *grpc_api.GetDataRequest, stream grpc.ServerStreamingServer[grpc_api.GetDataBinaryResponse]) error {
	ownerLogin, err := userLogin(stream.Context())
	if err != nil {
		return status.Error(codes.InvalidArgument, (codes.InvalidArgument).String())
	}

	g, ctx := errgroup.WithContext(stream.Context())
	dataStream, chErr := h.service.DataBinary(ctx, ownerLogin, request.Id)

	g.Go(func() error {
		for {
			select {
			case err := <-chErr:
				return err
			case <-ctx.Done():
				return ctx.Err()
			case b := <-dataStream:
				if err := stream.Send(&grpc_api.GetDataBinaryResponse{Body: []byte{b}}); err != nil {
					return fmt.Errorf("failed to send stream request: %w", err)
				}
			}
		}

	})

	if err := g.Wait(); err != nil {
		log.Error().Err(err).Msg("failed to get binary")
		return status.Error(codes.Internal, (codes.Internal).String())
	}

	return nil
}

func userLogin(ctx context.Context) (string, error) {
	data, err := outgoingMetaData(ctx, entity.UserLoginHeaderName)
	if err != nil {
		return "", fmt.Errorf("failed to get metadata: %w", err)
	}

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
