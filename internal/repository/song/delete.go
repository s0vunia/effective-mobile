package song

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/service"
	"github.com/s0vunia/platform_common/pkg/db"
	"go.uber.org/zap"
)

func (r *repo) Delete(ctx context.Context, id int64) error {
	builder := sq.Delete(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "song_repository.Delete",
		QueryRaw: query,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", args))

	result, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return service.ErrSongNotFound
	}

	return nil
}
