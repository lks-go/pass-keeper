package client

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog/log"
)

type Deps struct {
	LoginPassClient LoginPassClient
}

func New(d Deps) *Client {
	return &Client{
		loginPass: &LoginPass{
			client: d.LoginPassClient,
		},
	}
}

type Client struct {
	loginPass *LoginPass
}

func (c *Client) Run(ctx context.Context) error {

LOOP:
	for {
		prompt := promptui.Select{
			Label: "Select option:",
			Items: []string{OptLoginPass, OptTextData, OptCards, OptBinaryData, OptExit},
		}

		_, result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed %w", err)
		}

		switch result {
		case OptLoginPass:
			if err := c.loginPass.Run(ctx); err != nil {
				log.Err(err).Msg("login pass failed")
			}
		case OptExit:
			fmt.Println("exit")
			break LOOP
		}
	}

	return nil
}
