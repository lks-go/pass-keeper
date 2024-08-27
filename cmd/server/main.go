package main

import (
	"log"

	"github.com/lks-go/pass-keeper/internal/app"
)

func main() {
	app := app.NewServerAPP(app.SetupServerAPPConfig())
	if err := app.Build(); err != nil {
		log.Fatalf("filed to build app: %s", err)
	}

	log.Println("running service")
	if err := app.Run(); err != nil {
		log.Fatalf("filed to build app: %s", err)
	}

	log.Println("service successfully stopped")
}
