package client

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"github.com/lks-go/pass-keeper/internal/service/client"
	"github.com/lks-go/pass-keeper/internal/transport/backend_client"
)

type App struct {
	cfg    *Config
	client *client.Client
}

func New(cfg *Config) *App {
	return &App{
		cfg: cfg,
	}
}

func (app *App) Build() error {

	backendClientConfig := backend_client.Config{
		Host:      app.cfg.ServerHost,
		CertPath:  app.cfg.ServerCertPath,
		EnableTLS: app.cfg.EnableTLS,
	}
	backendClient, err := backend_client.New(&backendClientConfig)
	if err != nil {
		return fmt.Errorf("failed to get backend client: %w", err)
	}

	app.client = client.New(backendClient)

	return nil
}

func (app *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := app.client.Run(ctx); err != nil {
			return fmt.Errorf("failed to run client: %w", err)
		}

		stop()

		return nil
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("group error: %w", err)
	}

	return nil

}
