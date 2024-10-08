package backend

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

	"github.com/lks-go/pass-keeper/internal/lib/crypt"
	"github.com/lks-go/pass-keeper/internal/lib/password"
	"github.com/lks-go/pass-keeper/internal/lib/token"
	"github.com/lks-go/pass-keeper/internal/service"
	"github.com/lks-go/pass-keeper/internal/transport/grpchandler"
	"github.com/lks-go/pass-keeper/internal/transport/interceptor"
	"github.com/lks-go/pass-keeper/internal/transport/storage"
	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

type App struct {
	grpcHandler     grpc_api.PassKeeperServer
	pool            *sql.DB
	authInterceptor *interceptor.Auth
	config          *Config
}

func New(cfg *Config) *App {
	return &App{
		config: cfg,
	}
}

func (app *App) Build() error {
	pool, err := setupDB(app.config.DatabaseDSN)
	if err != nil {
		return fmt.Errorf("failed to setup DB: %w", err)
	}

	if err := RunMigrations(app.config.DatabaseDSN, "./migrations"); err != nil {
		return fmt.Errorf("failed to run migraions: %w", err)
	}

	storage := storage.New(pool)
	passwordHasher := password.New(app.config.UserPassSalt)

	token, err := token.New(app.config.TokenSecretKey, app.config.TokenExpirationTime)
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	crypt, err := crypt.New(app.config.CryptSecretKey)
	if err != nil {
		return fmt.Errorf("failed to get crypt: %w", err)
	}

	servCfg := service.ServerConfig{BinaryChunkSize: app.config.BinaryChunkSize}
	servDeps := service.ServerDeps{
		Storage:      storage,
		PasswordHash: passwordHasher,
		Token:        token,
		Crypt:        crypt,
	}
	service := service.NewBackend(servCfg, servDeps)

	grpcHandler := grpchandler.New(service)

	app.pool = pool
	app.grpcHandler = grpcHandler
	app.authInterceptor = interceptor.NewAuth(token)

	return nil
}

func (app *App) Run() error {
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

func (app *App) Exit() error {
	if err := app.pool.Close(); err != nil {
		return fmt.Errorf("failed to close pool: %w", err)
	}

	return nil
}

func (app *App) startGRPCServer(ctx context.Context) error {
	listen, err := net.Listen("tcp", app.config.GRPCNetAddress)
	if err != nil {
		return fmt.Errorf("filed to start listen address %s: %w", app.config.GRPCNetAddress, err)
	}

	serverOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(app.authInterceptor.CheckAccess),
		grpc.StreamInterceptor(app.authInterceptor.CheckAccessStream),
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
