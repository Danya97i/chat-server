package chat

import (
	"context"

	"github.com/Danya97i/chat-server/internal/models"
)

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
	if err != nil {
		return 0, err
	}
	return id, nil
}
