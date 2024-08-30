package app

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/lks-go/pass-keeper/internal/interceptor"
	"github.com/lks-go/pass-keeper/internal/transport/grpchandler"
	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

type ServerAPPConfig struct {
	GRPCNetAddress string `env:"GRPC_NET_ADDRESS" env-default:":9000"`
	//DatabaseDSN    string `env:"DATABASE_DSN" env-required:"true"`
	//UserPassSalt   string `env:"USER_PASS_SALT" env-required:"true"`
	DatabaseDSN       string
	UserPassSalt      string
	EnableTLS         bool   `env:"ENABLE_TLS" env-default:"true"`
	CertServerCRTPath string `env:"SERVER_CRT_PATH" env-default:"cert/server.crt"`
	CertServerKeyPath string `env:"SERVER_CRT_PATH" env-default:"cert/server.key"`
}

type ServerAPP struct {
	grpcHandler grpc_api.PassKeeperServer
	pool        *sql.DB

	config *ServerAPPConfig
}

func NewServerAPP(cfg *ServerAPPConfig) *ServerAPP {
	return &ServerAPP{
		config: cfg,
	}
}

func (app *ServerAPP) Build() error {
	//pool, err := setup.DB(app.config.DatabaseDSN)
	//if err != nil {
	//	return fmt.Errorf("failed to setup DB: %w", err)
	//}

	//storage := storage.New(pool)

	//cfg := service.Config{
	//	UserPassSalt: app.config.UserPassSalt,
	//}
	//UserRegistrar := service.NewUserRegistrarWithPassHash(&cfg, service.NewUserRegistrar(&cfg, storage))

	//app.pool = pool
	app.grpcHandler = grpchandler.New()

	return nil
}

func (app *ServerAPP) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := app.startGRPCServer(ctx); err != nil {
			return fmt.Errorf("failed to start grpc server: %w", err)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("group error: %w", err)
	}

	return nil
}

func (app *ServerAPP) Exit() error {
	if err := app.pool.Close(); err != nil {
		return fmt.Errorf("failed to close pool: %w", err)
	}

	return nil
}

func (app *ServerAPP) startGRPCServer(ctx context.Context) error {
	listen, err := net.Listen("tcp", app.config.GRPCNetAddress)
	if err != nil {
		return fmt.Errorf("filed to start listen address %s: %w", app.config.GRPCNetAddress, err)
	}

	serverOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Auth),
	}

	if app.config.EnableTLS {
		creds, err := credentials.NewServerTLSFromFile(app.config.CertServerCRTPath, app.config.CertServerKeyPath)
		if err != nil {
			return fmt.Errorf("could not load TLS keys: %w", err)
		}
		serverOpts = append(serverOpts, grpc.Creds(creds))
	}

	s := grpc.NewServer(serverOpts...)
	grpc_api.RegisterPassKeeperServer(s, app.grpcHandler)

	go func() {
		<-ctx.Done()

		s.Stop()
		listen.Close()
	}()

	if err := s.Serve(listen); err != nil {
		return fmt.Errorf("filed to start serving: %w", err)
	}

	return nil
}
