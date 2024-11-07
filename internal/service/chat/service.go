package chat

import (
	"github.com/Danya97i/platform_common/pkg/db"

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
