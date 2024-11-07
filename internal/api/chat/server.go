package chat

import (
	"github.com/Danya97i/chat-server/internal/service"
	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
)

// Server - имплементация gRPC-сервера
type Server struct {
	pb.UnimplementedChatV1Server
	chatService service.ChatService
}

// NewServer - создает новый gRPC-сервер
func NewServer(chatService service.ChatService) *Server {
	return &Server{chatService: chatService}
}
