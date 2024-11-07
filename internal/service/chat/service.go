package chat

import (
	"context"

	"github.com/Danya97i/platform_common/pkg/db"

	"github.com/Danya97i/chat-server/internal/models"
	"github.com/Danya97i/chat-server/internal/repository"
	serv "github.com/Danya97i/chat-server/internal/service"
)

type service struct {
	chatRepo  repository.ChatRepository
	logRepo   repository.LogRepository
	txManager db.TxManager
}

// NewService creates new chat service
func NewService(chatRepo repository.ChatRepository, logRepo repository.LogRepository, txManager db.TxManager) serv.ChatService {
	return &service{
		chatRepo:  chatRepo,
		logRepo:   logRepo,
		txManager: txManager,
	}
}

// Create creates new chat
func (s *service) Create(ctx context.Context, title string, userEmails []string) (int64, error) {
	var id int64
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var txErr error
		id, txErr = s.chatRepo.Create(ctx, title)
		if txErr != nil {
			return txErr
		}
		txErr = s.chatRepo.AddChatUsers(ctx, id, userEmails)
		if txErr != nil {
			return txErr
		}

		// add log
		txErr = s.logRepo.Save(ctx, models.LogInfo{
			ChatID: id,
			Action: models.ActionCreate,
		})
		return txErr
	})
	return id, err
}

// Delete – delete chat
func (s *service) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommited(ctx, func(ctx context.Context) error {
		var txErr error
		txErr = s.chatRepo.Delete(ctx, id)
		if txErr != nil {
			return txErr
		}

		// add log
		txErr = s.logRepo.Save(ctx, models.LogInfo{
			ChatID: id,
			Action: models.ActionDelete,
		})
		return txErr
	})
	return err
}
