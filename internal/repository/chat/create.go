package chat

import (
	"context"

	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"
)

// Create - creates new chat in database
func (r *repo) Create(ctx context.Context, title string) (int64, error) {
	insertChatQueryBuilder := squirrel.Insert("chats").
		PlaceholderFormat(squirrel.Dollar).
		Columns("title").
		Values(title).
		Suffix("returning id")
	insertChatQuery, args, err := insertChatQueryBuilder.ToSql()
	if err != nil {
		return 0, err
	}
	query := db.Query{
		RawQuery: insertChatQuery,
	}
	var chatID int64
	if err := r.db.DB().ScanOneContext(ctx, &chatID, query, args...); err != nil {
		return 0, err
	}

	return chatID, nil
}
