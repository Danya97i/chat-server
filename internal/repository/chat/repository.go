package chat

import (
	"context"

	"github.com/Danya97i/chat-server/internal/client/db"
	"github.com/Danya97i/chat-server/internal/repository"
	"github.com/Masterminds/squirrel"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ChatRepository {
	return &repo{db: db}
}

// Create – метод репозитория для создания чата
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
	r.db.DB().ScanOneContext(ctx, &chatID, query, args...)

	return chatID, nil
}

func (r *repo) AddChatUsers(ctx context.Context, chatID int64, userEmails []string) error {
	insertUsersQueryBuilder := squirrel.Insert("chat_users").
		PlaceholderFormat(squirrel.Dollar).
		Columns("chat_id", "user_id")
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

// Delete – метод репозитория для удаления чата
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
