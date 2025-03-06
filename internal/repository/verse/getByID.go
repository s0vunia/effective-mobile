package verse

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"github.com/s0vunia/platform_common/pkg/db"
	"go.uber.org/zap"
)

func (r *repo) GetByID(ctx context.Context, id int64) (*model.Verse, error) {
	builder := sq.Select(
		idColumn,
		songIDColumn,
		verseNumColumn,
		textColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "verse_repository.GetByID",
		QueryRaw: query,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", args))

	var verse model.Verse
	err = r.db.DB().ScanOneContext(ctx, &verse, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, service.ErrVerseNotFound
		}
		return nil, err
	}

	return &verse, nil
}
