package chat

import (
	"context"
	"log"

	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
)

// CreateChat - метод для создания нового чата
func (s *Server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	log.Println("create chat request", req)

	chatID, err := s.chatService.Create(ctx, req.Title, req.UserEmails)
	if err != nil {
		return nil, err
	}
	return &pb.CreateChatResponse{Id: chatID}, nil
}
