package client

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog/log"

	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type LoginPassClient interface {
	ListLoginPass(ctx context.Context) ([]entity.DataLoginPass, error)
}

type LoginPass struct {
	client LoginPassClient
}

func (lp *LoginPass) Run(ctx context.Context) error {
	prompt := promptui.Select{
		Label: "Data type",
		Items: []string{OptAdd, OptList, OptGet, OptBack},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	switch result {
	case OptAdd:
		if err := lp.add(ctx); err != nil {
			return fmt.Errorf("failed to add data: %w", err)
		}
	case OptList:
		if err := lp.list(ctx); err != nil {
			return fmt.Errorf("failed to list data: %w", err)
		}
	case OptGet:
		if err := lp.get(ctx); err != nil {
			return fmt.Errorf("failed to get data: %w", err)
		}
	case OptBack:
	}

	return nil
}

func (lp *LoginPass) add(ctx context.Context) error {
	log.Info().Msg("added log and pass")
	return nil
}

func (lp *LoginPass) list(ctx context.Context) error {
	list, err := lp.client.ListLoginPass(ctx)
	if err != nil {
		return fmt.Errorf("failed to get list of logins and passwords: %w", err)
	}

	for _, item := range list {
		fmt.Printf("id: %d;\ttitle: %s\n", item.ID, item.Title)
	}

	return nil
}

func (lp *LoginPass) get(ctx context.Context) error {
	log.Info().Msg("got log and pass")
	return nil
}
