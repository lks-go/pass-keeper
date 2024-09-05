package client

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

func SetupConfig() *Config {
	cfg := Config{}
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to read app config: %s", err)
	}

	return &cfg
}

type Config struct {
	ServerHost     string `env:"SERVER_HOST" env-default:"localhost:9000"`
	ServerCertPath string `env:"SERVER_CERT_PATH" env-default:"cert/server.crt"`
	EnableTLS      bool   `env:"ENABLE_TLS" env-default:"true"`
}
