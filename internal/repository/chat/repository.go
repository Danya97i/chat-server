package chat

import (
	"github.com/Danya97i/platform_common/pkg/db"

	"github.com/Danya97i/chat-server/internal/repository"
)

type repo struct {
	db db.Client
}

// NewRepository makes new repository
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}
