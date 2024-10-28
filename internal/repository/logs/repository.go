package logs

import (
	"context"

	"github.com/Masterminds/squirrel"

	"github.com/Danya97i/chat-server/internal/client/db"
	"github.com/Danya97i/chat-server/internal/models"
	"github.com/Danya97i/chat-server/internal/repository"
)

type repo struct {
	db db.Client
}

// NewRepository creates new repository
func NewRepository(db db.Client) repository.LogRepository {
	return &repo{db: db}
}

// Save saves log info
func (r *repo) Save(ctx context.Context, logInfo models.LogInfo) error {
	insertLogQueryBuilder := squirrel.Insert("chats_logs").
		PlaceholderFormat(squirrel.Dollar).
		Columns("action", "chat_id").
		Values(logInfo.Action, logInfo.ChatID)

	insertLogQuery, args, err := insertLogQueryBuilder.ToSql()
	if err != nil {
		return err
	}
	query := db.Query{RawQuery: insertLogQuery}

	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}
