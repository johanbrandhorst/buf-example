package main

import (
	"io/ioutil"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"

	userspb "github.com/johanbrandhorst/buf-example/proto/users/v1"
	"github.com/johanbrandhorst/buf-example/users"
)

func main() {
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to create listener:", err)
	}

	srv := grpc.NewServer()
	userspb.RegisterUserServiceServer(srv, &users.Service{})
	reflection.Register(srv)

	log.Infoln("Serving gRPC on:", listener.Addr().String())
	err = srv.Serve(listener)
	if err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
