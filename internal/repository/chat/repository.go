package chat

import (
	"context"

	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"

	"github.com/Danya97i/chat-server/internal/repository"
)

type repo struct {
	db db.Client
}

// NewRepository makes new repository
func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

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

// AddChatUsers – save users to chat in database
func (r *repo) AddChatUsers(ctx context.Context, chatID int64, userEmails []string) error {
	insertUsersQueryBuilder := squirrel.Insert("chat_users").
		PlaceholderFormat(squirrel.Dollar).
		Columns("chat_id", "user_email")
	for _, email := range userEmails {
		insertUsersQueryBuilder = insertUsersQueryBuilder.Values(chatID, email)
	}
	insertUsersQuery, args, err := insertUsersQueryBuilder.ToSql()
	if err != nil {
		return err
	}
	query := db.Query{
		RawQuery: insertUsersQuery,
	}
	_, err = r.db.DB().ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

// Delete – deletes chat from database
func (r *repo) Delete(ctx context.Context, id int64) error {
	deleteChatQueryBuilder := squirrel.Delete("chats").
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": id})
	deleteChatQuery, args, err := deleteChatQueryBuilder.ToSql()
	if err != nil {
		return err
	}
	q := db.Query{
		RawQuery: deleteChatQuery,
	}
	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return nil
}
