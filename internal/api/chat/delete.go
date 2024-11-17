package chat

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Danya97i/chat-server/pkg/chat_v1"
)

// DeleteChat - метод для удаления чата
func (s *Server) DeleteChat(ctx context.Context, req *pb.DeleteChatRequest) (*emptypb.Empty, error) {
	log.Println("create chat request", req)

	if err := s.chatService.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	return nil, nil
}
