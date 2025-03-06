package song

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/repository/song/converter"
	repository "github.com/s0vunia/effective-mobile/internal/repository/song/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"github.com/s0vunia/platform_common/pkg/db"
)

func (r *repo) GetByID(ctx context.Context, id int64) (*model.Song, error) {
	builderSelectOne := sq.Select(idColumn, groupId, titleColumn, releaseDate, linkColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builderSelectOne.ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		Name:     "song_repository.GetByID",
		QueryRaw: query,
	}
	var song repository.Song
	err = r.db.DB().ScanOneContext(ctx, &song, q, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, service.ErrSongNotFound
		}
		return nil, err
	}
	return converter.ToSongFromRepo(&song), nil
}
