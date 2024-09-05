package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/lks-go/pass-keeper/internal/app/backend"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msg("Setup config")
	cfg := backend.SetupConfig()

	log.Info().Msg("Building app")
	app := backend.New(cfg)
	if err := app.Build(); err != nil {
		log.Error().Err(err).Msg("filed to build app")
		return
	}

	log.Info().Msg("Running app")
	if err := app.Run(); err != nil {
		log.Error().Err(err).Msg("filed to build app")
		return
	}

	log.Info().Msg("App successfully stopped")
}
