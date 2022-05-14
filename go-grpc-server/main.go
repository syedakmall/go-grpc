package main

import (
	"log"
	"net"

	"github.com/syedakmall/go-grpc-cc/todos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	t := todos.Server{}
	grpcServer := grpc.NewServer()

	todos.RegisterTodoServer(grpcServer, &t)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
