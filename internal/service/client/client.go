package client

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/rs/zerolog/log"

	"github.com/lks-go/pass-keeper/internal/service/client/auth"
	"github.com/lks-go/pass-keeper/internal/service/client/binary"
	"github.com/lks-go/pass-keeper/internal/service/client/card"
	"github.com/lks-go/pass-keeper/internal/service/client/loginpass"
	"github.com/lks-go/pass-keeper/internal/service/client/text"
	"github.com/lks-go/pass-keeper/internal/service/entity"
)

type storage interface {
	loginpass.Storage
	auth.Storage
	text.Storage
	card.Storage
	binary.Storage
}

func New(s storage) *Client {
	return &Client{
		auth:      &auth.Auth{Storage: s},
		loginPass: &loginpass.LoginPass{Storage: s},
		text:      &text.Text{Storage: s},
		card:      &card.Card{Storage: s},
		binary:    &binary.Binary{Storage: s},
	}
}

type Client struct {
	auth      *auth.Auth
	loginPass *loginpass.LoginPass
	text      *text.Text
	card      *card.Card
	binary    *binary.Binary

	authenticated bool
}

func (c *Client) SetToken(t string) {
	c.loginPass.SetToken(t)
	c.text.SetToken(t)
	c.card.SetToken(t)
	c.binary.SetToken(t)
}

func (c *Client) Run(ctx context.Context) error {

LOOP:
	for {
		if !c.authenticated {
			authPrompt := promptui.Select{
				Label: "Please log in or register",
				Items: []string{entity.OptLogIn, entity.OptRegister, entity.OptExit},
			}

			_, authResult, err := authPrompt.Run()
			if err != nil {
				return fmt.Errorf("prompt failed %w", err)
			}

			switch authResult {
			case entity.OptLogIn:
				token, err := c.auth.Auth(ctx)
				if err != nil {
					log.Err(err).Msg("login pass failed")
					continue
				}

				c.SetToken(token)
				c.authenticated = true
			case entity.OptRegister:
				err := c.auth.Reg(ctx)
				if err != nil {
					log.Err(err).Msg("registration failed")
					continue
				}

				log.Info().Msg("You successfully registered, use your login and password to authenticate yourself")
			case entity.OptExit:
				fmt.Println("exit")
				break LOOP
			}

			continue
		}

		prompt := promptui.Select{
			Label: "Select category:",
			Items: []string{entity.OptLoginPass, entity.OptTextData, entity.OptCards, entity.OptBinaryData, entity.OptExit},
		}

		_, result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed %w", err)
		}

		switch result {
		case entity.OptLoginPass:
			if err := c.loginPass.Run(ctx); err != nil {
				log.Err(err).Msg("login pass failed")
			}
		case entity.OptTextData:
			if err := c.text.Run(ctx); err != nil {
				log.Err(err).Msg("text failed")
			}
		case entity.OptCards:
			if err := c.card.Run(ctx); err != nil {
				log.Err(err).Msg("text failed")
			}
		case entity.OptBinaryData:
			if err := c.binary.Run(ctx); err != nil {
				log.Err(err).Msg("binary failed")
			}
		case entity.OptExit:
			fmt.Println("exit")
			break LOOP
		}
	}

	return nil
}
