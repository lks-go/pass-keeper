package gobone

//go:generate mkdir -p ./pkg/grpc_api
//go:generate ./bin/bin/protoc --proto_path=$PWD/api --go_out=./pkg/grpc_api --go_opt=paths=source_relative --go-grpc_out=./pkg/grpc_api --go-grpc_opt=paths=source_relative pass_keeper.proto
