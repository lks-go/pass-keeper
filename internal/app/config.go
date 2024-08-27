package app

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

func SetupServerAPPConfig() *ServerAPPConfig {
	cfg := ServerAPPConfig{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to read app config: %s", err)
	}

	return &cfg
}
