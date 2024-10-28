package repository

import "context"

// ChatRepository interface
type ChatRepository interface {
	Create(ctx context.Context, title string) (int64, error)
	AddChatUsers(ctx context.Context, id int64, users []string) error
	Delete(ctx context.Context, id int64) error
}
