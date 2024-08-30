PROJECT_PATH=$(shell pwd)
PROJECT_BIN_PATH=$(PROJECT_PATH)/bin
PROJECT_API_PATH=$(PROJECT_PATH)/api

export GOBIN := $(PROJECT_BIN_PATH)

-include local.env

PB_REL="https://github.com/protocolbuffers/protobuf/releases"

tools: install_protoc
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/vektra/mockery/v2@v2.43.2
	go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.1


install_protoc:
	curl -LO $(PB_REL)/download/v27.3/protoc-27.3-osx-x86_64.zip && \
	unzip protoc-27.3-osx-x86_64.zip -d ./bin && \
	rm protoc-27.3-osx-x86_64.zip

generate:
	go generate ./...
	go mod tidy

cert_gen:
	cd ./cert && sh gen.sh

# alternative variant to generate grpc client/server
generate_grpc:
	mkdir -p ./pkg/grpc && \
	$(PROJECT_BIN_PATH)/bin/protoc \
	--proto_path=$(PROJECT_API_PATH) \
	--go_out=./pkg/grpc --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/grpc --go-grpc_opt=paths=source_relative pass-keeper.proto

run_server:
	go run cmd/server/main.go

new_migration:
	$(GOPATH)/bin/migrate create -ext sql -dir ./migrations -seq $(filter-out $@, $(MAKECMDGOALS))