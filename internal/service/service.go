package service

import "context"

type Config struct {
	UserPassSalt string
}

func New(cfg *Config) *Service {
	return &Service{cfg: cfg}
}

type Service struct {
	cfg *Config
}

type User struct {
	Login    string
	Password string
}

// Auth authenticates and authorizes the user by login and password
// If auth succeed returns user's a new JWT
func (s *Service) Auth(ctx context.Context, u User) (string, error) {
	panic("implement me")
	return "", nil
}
