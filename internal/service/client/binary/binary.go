package binary

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/manifoldco/promptui"
	"golang.org/x/sync/errgroup"

	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type Storage interface {
	BinaryAdd(ctx context.Context, token string, binary *entity.DataBinary) (id int32, err error)
	ListBinary(ctx context.Context, token string) ([]entity.DataBinary, error)
	BinaryData(ctx context.Context, token string, id int32) (*entity.DataBinary, error)
}

type Binary struct {
	Storage Storage
	token   string
}

func (b *Binary) SetToken(t string) {
	b.token = t
}

func (b *Binary) Run(ctx context.Context) error {
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
		if err := b.add(ctx); err != nil {
			return fmt.Errorf("failed to add data: %w", err)
		}
	case entity.OptList:
		if err := b.list(ctx); err != nil {
			return fmt.Errorf("failed to list data: %w", err)
		}
	case entity.OptBack:
	}

	return nil
}

func (b *Binary) add(ctx context.Context) error {
	prompt := promptui.Prompt{
		Label: "Input title",
	}

	title, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	prompt = promptui.Prompt{
		Label: "Input file name",
	}

	file, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("filed to open file: %w", err)
	}
	defer f.Close()

	br := bufio.NewReader(f)
	ch := make(chan byte, 0)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		defer func() {
			close(ch)
		}()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			b, err := br.ReadByte()
			if err != nil {
				if err == io.EOF {
					break LOOP
				}
				return fmt.Errorf("failed to read byte: %w", err)
			}

			ch <- b
		}

		return nil
	})

	var recordId int32
	g.Go(func() error {
		data := entity.DataBinary{
			Title:    title,
			FileName: file,
			Body:     ch,
		}
		id, err := b.Storage.BinaryAdd(ctx, b.token, &data)
		if err != nil {
			return fmt.Errorf("failed to add file: %w", err)
		}

		recordId = id
		return nil
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error group failed: %w", err)
	}

	fmt.Printf("Added record id: %d", recordId)

	return nil
}

func (b *Binary) list(ctx context.Context) error {
	list, err := b.Storage.ListBinary(ctx, b.token)
	if err != nil {
		return fmt.Errorf("failed to get list of logins and passwords: %w", err)
	}

	itmes := make([]string, 0, len(list))
	for _, item := range list {
		itmes = append(itmes, item.Title)
	}

	itmes = append(itmes, entity.OptBack)

	prompt := promptui.Select{
		Label: "Choose file",
		Items: itmes,
	}

	n, result, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("prompt failed: %w", err)
	}

	if result == entity.OptBack {
		return nil
	}

	if err := b.get(ctx, list[n].ID); err != nil {
		return fmt.Errorf("failed to get chosen data: %w", err)
	}

	return nil
}

func (b *Binary) get(ctx context.Context, id int32) error {

	data, err := b.Storage.BinaryData(ctx, b.token, id)
	if err != nil {
		return fmt.Errorf("failed to get login and pass: %w", err)
	}

	fmt.Printf("%s\nfile name: %s\n\n", data.Title, data.FileName)

	return nil
}
