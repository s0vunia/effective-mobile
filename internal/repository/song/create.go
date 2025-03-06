package song

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/platform_common/pkg/db"
	"go.uber.org/zap"
)

func (r *repo) Create(ctx context.Context, song *model.SongCreate) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(
			groupId,
			titleColumn,
			releaseDate,
			linkColumn,
		).
		Values(
			song.GroupID,
			song.Title,
			song.ReleaseDate,
			song.Link,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "song_repository.Create",
		QueryRaw: query,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", args))

	var id int64
	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}
