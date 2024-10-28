package converter

import (
	"github.com/Danya97i/chat-server/internal/models"
	repoModels "github.com/Danya97i/chat-server/internal/repository/chat/models"
)

func ToChatFromRepo(chat *repoModels.Chat) *models.Chat {
	if chat == nil {
		return nil
	}
	return &models.Chat{
		Id:        chat.Id,
		Title:     chat.Title,
		CreatedAt: chat.CreatedAt,
	}
}
