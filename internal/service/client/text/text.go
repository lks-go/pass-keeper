package text

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type Storage interface {
	TextAdd(ctx context.Context, token string, title, text string) (id int32, err error)
	ListText(ctx context.Context, token string) ([]entity.DataText, error)
	TextData(ctx context.Context, token string, id int32) (*entity.DataText, error)
}

type Text struct {
	Storage Storage
	token   string
}

func (lp *Text) SetToken(t string) {
	lp.token = t
}

func (lp *Text) Run(ctx context.Context) error {
	prompt := promptui.Select{
		Label: "Chose action",
		Items: []string{entity.OptAdd, entity.OptList, entity.OptBack},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	switch result {
	case entity.OptAdd:
		if err := lp.add(ctx); err != nil {
			return fmt.Errorf("failed to add data: %w", err)
		}
	case entity.OptList:
		if err := lp.list(ctx); err != nil {
			return fmt.Errorf("failed to list data: %w", err)
		}
	case entity.OptBack:
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
		Label: "Input text",
	}

	text, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	id, err := lp.Storage.TextAdd(ctx, lp.token, title, text)
	if err != nil {
		return fmt.Errorf("failed to add login and pass: %w", err)
	}

	fmt.Printf("Record id: %d", id)

	return nil
}

func (lp *Text) list(ctx context.Context) error {
	list, err := lp.Storage.ListText(ctx, lp.token)
	if err != nil {
		return fmt.Errorf("failed to get list of logins and passwords: %w", err)
	}

	itmes := make([]string, 0, len(list))
	for _, item := range list {
		itmes = append(itmes, item.Title)
	}

	itmes = append(itmes, entity.OptBack)

	prompt := promptui.Select{
		Label: "Choose creds",
		Items: itmes,
	}

	n, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	if result == entity.OptBack {
		return nil
	}

	if err := lp.get(ctx, list[n].ID); err != nil {
		return fmt.Errorf("failed to get chosen data: %w", err)
	}

	return nil
}

func (lp *Text) get(ctx context.Context, id int32) error {

	data, err := lp.Storage.TextData(ctx, lp.token, id)
	if err != nil {
		return fmt.Errorf("failed to get login and pass: %w", err)
	}

	fmt.Printf("%s\ntext: %s\n\n", data.Title, data.Text)

	return nil
}
