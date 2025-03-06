package song

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"github.com/s0vunia/platform_common/pkg/db"
	"go.uber.org/zap"
)

func (r *repo) Update(ctx context.Context, id int64, song *model.SongUpdate) error {
	builderUpdate := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar)

	if song.Title != nil {
		builderUpdate = builderUpdate.Set(titleColumn, *song.Title)
	}
	if song.ReleaseDate != nil {
		builderUpdate = builderUpdate.Set(releaseDate, *song.ReleaseDate)
	}
	if song.Link != nil {
		builderUpdate = builderUpdate.Set(linkColumn, *song.Link)
	}
	if song.GroupID != nil {
		builderUpdate = builderUpdate.Set(groupId, *song.GroupID)
	}
	builderUpdate = builderUpdate.
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "song_repository.Update",
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
