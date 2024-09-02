package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/lks-go/pass-keeper/internal/app"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msg("Setup config")
	cfg := app.SetupServerAPPConfig()

	log.Info().Msg("Build service")
	app := app.NewServerAPP(cfg)
	if err := app.Build(); err != nil {
		log.Error().Err(err).Msg("filed to build app")
		return
	}

	log.Info().Msg("Running service")
	if err := app.Run(); err != nil {
		log.Error().Err(err).Msg("filed to build app")
		return
	}

	log.Info().Msg("Service successfully stopped")
}
