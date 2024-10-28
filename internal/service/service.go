package service

import "context"

// ChatService interface
type ChatService interface {
	Create(ctx context.Context, title string, userEmails []string) (int64, error)
	Delete(ctx context.Context, id int64) error
}
