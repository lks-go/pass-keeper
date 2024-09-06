package client

import (
	"context"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"

	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type TextClient interface {
	ListLoginPass(ctx context.Context, token string) ([]entity.DataLoginPass, error)
	LoginPassData(ctx context.Context, token string, id int32) (*entity.DataLoginPass, error)
	LoginPassAdd(ctx context.Context, token string, title, login, pass string) (id int32, err error)
}

type Text struct {
	client TextClient
	token  string
}

func (lp *Text) SetToken(t string) {
	lp.token = t
}

func (lp *Text) Run(ctx context.Context) error {
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

func (lp *Text) add(ctx context.Context) error {
	prompt := promptui.Prompt{
		Label: "Input title",
	}

	title, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	prompt = promptui.Prompt{
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

	id, err := lp.client.LoginPassAdd(ctx, lp.token, title, login, pass)
	if err != nil {
		return fmt.Errorf("failed to add login and pass: %w", err)
	}

	fmt.Printf("Record id: %d", id)

	return nil
}

func (lp *Text) list(ctx context.Context) error {
	list, err := lp.client.ListLoginPass(ctx, lp.token)
	if err != nil {
		return fmt.Errorf("failed to get list of logins and passwords: %w", err)
	}

	for _, item := range list {
		fmt.Printf("id: %d;\ttitle: %s\n", item.ID, item.Title)
	}

	return nil
}

func (lp *Text) get(ctx context.Context) error {
	prompt := promptui.Prompt{
		Label: "Chose and input id",
	}

	input, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	id, err := strconv.Atoi(input)
	if err != nil {
		return fmt.Errorf("invalid number: %w", err)
	}

	data, err := lp.client.LoginPassData(ctx, lp.token, int32(id))
	if err != nil {
		return fmt.Errorf("failed to get login and pass: %w", err)
	}

	fmt.Printf("id: %d;\title: %s\nlogin: %s\nPassword: %s\n\n", data.ID, data.Title, data.Login, data.Password)

	return nil
}
