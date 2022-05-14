package todos

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) CreateTodo(ctx context.Context, in *TodoText) (*TodoItem, error) {
	log.Printf("Receive message body from client: %s", in.GetText())
	return &TodoItem{Id: 1, Text: in.GetText()}, nil
}

func (s *Server) ReadTodos(ctx context.Context, in *Void) (*TodoItems, error) {
	return &TodoItems{Items: []*TodoItem{
		{Id: 1, Text: "Max"},
		{Id: 2, Text: "Lewis"},
	}}, nil
}

func (s *Server) mustEmbedUnimplementedTodoServer() {}
