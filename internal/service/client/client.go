package client

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog/log"
)

type BackendClient interface {
	LoginPassClient
	AuthClient
}

func New(back BackendClient) *Client {
	return &Client{
		loginPass: &LoginPass{client: back},
		auth:      &Auth{client: back},
	}
}

type Client struct {
	loginPass *LoginPass
	auth      *Auth

	authenticated bool
}

func (c *Client) SetToken(t string) {
	c.loginPass.SetToenk(t)
}

func (c *Client) Run(ctx context.Context) error {

LOOP:
	for {
		if !c.authenticated {
			authPrompt := promptui.Select{
				Label: "Please log in or register",
				Items: []string{OptLogIn, OptRegister, OptExit},
			}

			_, authResult, err := authPrompt.Run()
			if err != nil {
				return fmt.Errorf("prompt failed %w", err)
			}

			switch authResult {
			case OptLogIn:
				token, err := c.auth.Auth(ctx)
				if err != nil {
					log.Err(err).Msg("login pass failed")
					continue
				}

				c.SetToken(token)
				c.authenticated = true
			case OptRegister:
				err := c.auth.Reg(ctx)
				if err != nil {
					log.Err(err).Msg("registration failed")
					continue
				}

				log.Info().Msg("You successfully registered, use your login and password to authenticate yourself")
			case OptExit:
				fmt.Println("exit")
				break LOOP
			}

			continue
		}

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
