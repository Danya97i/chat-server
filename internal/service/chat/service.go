package chat

import (
	"context"

	"github.com/Danya97i/chat-server/internal/client/db"
	"github.com/Danya97i/chat-server/internal/repository"
	serv "github.com/Danya97i/chat-server/internal/service"
)

type service struct {
	chatRepo  repository.ChatRepository
	txManager db.TxManager
}

func NewService(chatRepo repository.ChatRepository, txManager db.TxManager) serv.ChatService {
	return &service{
		chatRepo:  chatRepo,
		txManager: txManager,
	}
}

func (s *service) Create(ctx context.Context, title string, userEmails []string) (int64, error) {
	var id int64
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.chatRepo.Create(ctx, title)
		if errTx != nil {
			return errTx
		}
		errTx = s.chatRepo.AddChatUsers(ctx, id, userEmails)
		return nil
	})
	return id, err
}

func (s *service) Delete(ctx context.Context, id int64) error {
	return s.chatRepo.Delete(ctx, id)
}
