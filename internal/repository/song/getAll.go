package song

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/repository/song/converter"
	repository "github.com/s0vunia/effective-mobile/internal/repository/song/model"
	"github.com/s0vunia/platform_common/pkg/db"
	"go.uber.org/zap"
)

func (r *repo) GetAll(ctx context.Context, filter model.SongFilter, limit, offset int) ([]model.Song, int, error) {
	builder := sq.Select(
		"songs."+idColumn,
		"songs."+groupId,
		"songs."+titleColumn,
		"songs."+releaseDate,
		"songs."+linkColumn,
		"songs."+createdAtColumn,
		"songs."+updatedAtColumn,
		"groups.name as group_name",
	).From(tableName).PlaceholderFormat(sq.Dollar)

	builder = builder.Join("groups ON songs.group_id = groups.id")

	if filter.Group != "" {
		builder = builder.Where(sq.Like{"groups.name": "%" + filter.Group + "%"})
	}
	if filter.Song != "" {
		builder = builder.Where(sq.Like{"songs." + titleColumn: "%" + filter.Song + "%"})
	}

	countBuilder := sq.Select("COUNT(DISTINCT songs.id)").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Join("groups ON songs.group_id = groups.id")

	if filter.Group != "" {
		countBuilder = countBuilder.Where(sq.Like{"groups.name": "%" + filter.Group + "%"})
	}
	if filter.Song != "" {
		countBuilder = countBuilder.Where(sq.Like{"songs." + titleColumn: "%" + filter.Song + "%"})
	}

	countQuery, countArgs, err := countBuilder.ToSql()
	if err != nil {
		return nil, 0, err
	}

	var total int
	q := db.Query{
		Name:     "song_repository.GetAll.Count",
		QueryRaw: countQuery,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", countArgs))

	err = r.db.DB().ScanOneContext(ctx, &total, q, countArgs...)
	if err != nil {
		return nil, 0, err
	}

	builder = builder.OrderBy("songs." + idColumn).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, 0, err
	}

	q = db.Query{
		Name:     "song_repository.GetAll",
		QueryRaw: query,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", countArgs))

	var songs []repository.Song
	err = r.db.DB().ScanAllContext(ctx, &songs, q, args...)
	if err != nil {
		return nil, 0, err
	}

	result := make([]model.Song, 0, len(songs))
	for _, song := range songs {
		result = append(result, *converter.ToSongFromRepo(&song))
	}

	return result, total, nil
}
