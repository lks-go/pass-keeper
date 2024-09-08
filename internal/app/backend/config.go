package backend

import (
	"log"
	"time"

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
	GRPCNetAddress      string        `env:"GRPC_NET_ADDRESS" env-default:":9000"`
	DatabaseDSN         string        `env:"DATABASE_DSN" env-required:"true"`
	UserPassSalt        string        `env:"USER_PASS_SALT" env-required:"true"`
	EnableTLS           bool          `env:"ENABLE_TLS" env-default:"true"`
	CertServerCRTPath   string        `env:"SERVER_CRT_PATH" env-default:"cert/server.crt"`
	CertServerKeyPath   string        `env:"SERVER_KEY_PATH" env-default:"cert/server.key"`
	TokenSecretKey      string        `env:"TOKEN_SECRET_KEY" env-required:"true"`
	TokenExpirationTime time.Duration `env:"TOKEN_EXPIRATION_TIME" env-default:"10m"`
	CryptSecretKey      string        `env:"CRYPT_SECRET_KEY" env-required:"true"`
	BinaryChunkSize     int           `env:"BINARY_CHUNK_SIZE" env-default:"1024"`
}
