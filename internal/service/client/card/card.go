package card

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type Storage interface {
	ListCard(ctx context.Context, token string) ([]entity.DataCard, error)
	CardData(ctx context.Context, token string, id int32) (*entity.DataCard, error)
	CardAdd(ctx context.Context, token string, card *entity.DataCard) (id int32, err error)
}

type Card struct {
	Storage Storage
	token   string
}

func (c *Card) SetToken(t string) {
	c.token = t
}

func (c *Card) Run(ctx context.Context) error {
	prompt := promptui.Select{
		Label: "Choose action",
		Items: []string{entity.OptAdd, entity.OptList, entity.OptBack},
	}

	_, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	switch result {
	case entity.OptAdd:
		if err := c.add(ctx); err != nil {
			return fmt.Errorf("failed to add data: %w", err)
		}
	case entity.OptList:
		if err := c.list(ctx); err != nil {
			return fmt.Errorf("failed to list data: %w", err)
		}
	case entity.OptBack:
	}

	return nil
}

func (c *Card) add(ctx context.Context) error {
	prompt := promptui.Prompt{
		Label: "Input title",
	}

	title, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	prompt = promptui.Prompt{
		Label: "Input card number",
	}

	number, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	prompt = promptui.Prompt{
		Label: "Input card owner",
	}

	owner, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	prompt = promptui.Prompt{
		Label: "Input expiration date",
	}

	expDate, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	prompt = promptui.Prompt{
		Label: "Input cvc",
	}

	code, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	data := entity.DataCard{
		Title:   title,
		Number:  number,
		Owner:   owner,
		ExpDate: expDate,
		CVCCode: code,
	}

	id, err := c.Storage.CardAdd(ctx, c.token, &data)
	if err != nil {
		return fmt.Errorf("failed to add login and pass: %w", err)
	}

	fmt.Printf("Record id: %d", id)

	return nil
}

func (c *Card) list(ctx context.Context) error {
	list, err := c.Storage.ListCard(ctx, c.token)
	if err != nil {
		return fmt.Errorf("failed to get list of logins and passwords: %w", err)
	}

	itmes := make([]string, 0, len(list))
	for _, item := range list {
		itmes = append(itmes, item.Title)
	}

	itmes = append(itmes, entity.OptBack)

	prompt := promptui.Select{
		Label: "Choose card",
		Items: itmes,
	}

	n, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	if result == entity.OptBack {
		return nil
	}

	if err := c.get(ctx, list[n].ID); err != nil {
		return fmt.Errorf("failed to get chosen data: %w", err)
	}

	return nil
}

func (c *Card) get(ctx context.Context, id int32) error {

	data, err := c.Storage.CardData(ctx, c.token, id)
	if err != nil {
		return fmt.Errorf("failed to get login and pass: %w", err)
	}

	fmt.Printf("%s\nCard number:\t\t%s\nCard owner:\t\t%s\nExpiration date:\t%s\nCVC:\t\t\t%s\n\n", data.Title, data.Number, data.Owner, data.ExpDate, data.CVCCode)

	return nil
}
