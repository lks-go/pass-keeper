package auth

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
)

type Storage interface {
	Auth(ctx context.Context, login string, password string) (token string, err error)
	Reg(ctx context.Context, login string, password string) error
}

type Auth struct {
	Storage Storage
}

func (lp *Auth) Auth(ctx context.Context) (string, error) {
	prompt := promptui.Prompt{
		Label: "Input login",
	}

	login, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	prompt = promptui.Prompt{
		Label: "Input password",
	}

	pass, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	token, err := lp.Storage.Auth(ctx, login, pass)
	if err != nil {
		return "", fmt.Errorf("failed to auth: %w", err)
	}

	return token, nil
}

func (lp *Auth) Reg(ctx context.Context) error {
	prompt := promptui.Prompt{
		Label: "Input login",
	}

	login, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	prompt = promptui.Prompt{
		Label: "Input password",
	}

	pass, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	if err := lp.Storage.Reg(ctx, login, pass); err != nil {
		return fmt.Errorf("failed to auth: %w", err)
	}

	return nil
}
