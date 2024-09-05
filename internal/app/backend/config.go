package backend

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

func SetupServerAPPConfig() *Config {
	cfg := Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to read app config: %s", err)
	}

	return &cfg
}
