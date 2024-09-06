package client

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type LoginPassClient interface {
	ListLoginPass(ctx context.Context, token string) ([]entity.DataLoginPass, error)
	LoginPassData(ctx context.Context, token string, id int32) (*entity.DataLoginPass, error)
	LoginPassAdd(ctx context.Context, token string, title, login, pass string) (id int32, err error)
}

type LoginPass struct {
	client LoginPassClient
	token  string
}

func (lp *LoginPass) SetToken(t string) {
	lp.token = t
}

func (lp *LoginPass) Run(ctx context.Context) error {
	prompt := promptui.Select{
		Label: "Choose action",
		Items: []string{OptAdd, OptList, OptBack},
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
	case OptBack:
	}

	return nil
}

func (lp *LoginPass) add(ctx context.Context) error {
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

func (lp *LoginPass) list(ctx context.Context) error {
	list, err := lp.client.ListLoginPass(ctx, lp.token)
	if err != nil {
		return fmt.Errorf("failed to get list of logins and passwords: %w", err)
	}

	itmes := make([]string, 0, len(list))
	for _, item := range list {
		itmes = append(itmes, item.Title)
	}

	itmes = append(itmes, OptBack)

	prompt := promptui.Select{
		Label: "Choose creds",
		Items: itmes,
	}

	n, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	if result == OptBack {
		return nil
	}

	if err := lp.get(ctx, list[n].ID); err != nil {
		return fmt.Errorf("failed to get chosen data: %w", err)
	}

	return nil
}

func (lp *LoginPass) get(ctx context.Context, id int32) error {

	data, err := lp.client.LoginPassData(ctx, lp.token, id)
	if err != nil {
		return fmt.Errorf("failed to get login and pass: %w", err)
	}

	fmt.Printf("%s\nlogin: %s\nPassword: %s\n\n", data.Title, data.Login, data.Password)

	return nil
}
