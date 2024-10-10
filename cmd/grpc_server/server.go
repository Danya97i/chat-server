package main

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
)

type server struct {
	pb.UnimplementedChatV1Server
}

// CreateChat - метод для создания нового чата
func (s *server) CreateChat(_ context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	log.Println("create chat request", req)
	return nil, nil
}

// DeleteChat - метод для удаления чата
func (s *server) Delete(_ context.Context, req *pb.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Println("create chat request", req)
	return nil, nil
}

// Метод для отправки сообщения в чат
func (s *server) SendMesage(_ context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println("send message request", req)
	return nil, nil
}
