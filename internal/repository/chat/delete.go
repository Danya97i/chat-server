package chat

import (
	"context"

	"github.com/Danya97i/platform_common/pkg/db"
	"github.com/Masterminds/squirrel"
)

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
