package chat

import (
	"context"

	"github.com/Danya97i/chat-server/internal/models"
)

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
