install:
	go get \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		github.com/bufbuild/buf/cmd/buf \
		github.com/ktr0731/evans

generate:
	buf protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./proto/users/v1/users.proto

lint:
	buf check lint

breaking:
	buf check breaking --against-input .git#branch=initial-service

run:
	go run main.go
