package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClientAPPConfig struct {
}

type ClientAPP struct {
	cfg *ClientAPPConfig
}

func NewClientAPP(cfg *ClientAPPConfig) *ClientAPP {
	return &ClientAPP{
		cfg: cfg,
	}
}

func (app *ClientAPP) Build() error {

	return nil
}

func (app *ClientAPP) Run() error {

	r := bufio.NewReader(os.Stdin)
	var s string

	for {
		fmt.Fprintf(os.Stderr, "input command")
		s, _ = r.ReadString('\n')
		s = strings.TrimSpace(s)

		fmt.Printf("command is: %s\n", s)

	}

	return nil
}
