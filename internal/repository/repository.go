package repository

import (
	"context"

	"github.com/Danya97i/chat-server/internal/models"
)

// ChatRepository interface
type ChatRepository interface {
	Create(ctx context.Context, title string) (int64, error)
	AddChatUsers(ctx context.Context, id int64, users []string) error
	Delete(ctx context.Context, id int64) error
}

// LogRepository interface
type LogRepository interface {
	Save(ctx context.Context, logInfo models.LogInfo) error
}
