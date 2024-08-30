package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/lks-go/pass-keeper/pkg/grpc_api"
)

var cert = "cert/server.crt"

func main() {
	ctx := context.Background()

	creds, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.NewClient("localhost:9000", opts...)
	if err != nil {
		log.Fatalf("failed to makee new client: %s", err)
	}

	client := grpc_api.NewPassKeeperClient(conn)

	req := grpc_api.RegisterUserRequest{
		Login:    "test",
		Password: "pass",
	}
	res, err := client.RegisterUser(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
