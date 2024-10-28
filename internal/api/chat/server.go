package chat

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Danya97i/chat-server/internal/service"
	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
)

type Server struct {
	pb.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewServer(chatService service.ChatService) *Server {
	return &Server{chatService: chatService}
}

// CreateChat - метод для создания нового чата
func (s *Server) CreateChat(ctx context.Context, req *pb.CreateChatRequest) (*pb.CreateChatResponse, error) {
	log.Println("create chat request", req)

	chatID, err := s.chatService.Create(ctx, req.Title, req.UserEmails)
	if err != nil {
		return nil, err
	}
	return &pb.CreateChatResponse{Id: chatID}, nil
}

// DeleteChat - метод для удаления чата
func (s *Server) DeleteChat(ctx context.Context, req *pb.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Println("create chat request", req)

	if err := s.chatService.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}

// Метод для отправки сообщения в чат
func (s *Server) SendMesage(_ context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println("send message request", req)
	return nil, nil
}
