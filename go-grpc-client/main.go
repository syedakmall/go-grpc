package main

import (
	"context"
	"log"

	"github.com/syedakmall/go-grpc-client/todos"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure)
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	t := todos.NewTodoClient(conn)

	response, err := t.ReadTodos(context.Background(), &todos.Void{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)
}
