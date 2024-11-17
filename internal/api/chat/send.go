package chat

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
)

// SendMesage - метод для отправки сообщения в чат
func (s *Server) SendMesage(_ context.Context, req *pb.SendMessageRequest) (*emptypb.Empty, error) {
	log.Println("send message request", req)
	return nil, nil
}
