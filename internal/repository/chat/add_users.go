package chat

import (
	"context"

	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"
)

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
