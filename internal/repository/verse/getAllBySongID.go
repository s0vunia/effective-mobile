package verse

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/platform_common/pkg/db"
	"go.uber.org/zap"
)

func (r *repo) GetAllBySongID(ctx context.Context, songID int64, limit, offset int) ([]model.Verse, int, error) {
	builder := sq.Select(
		idColumn,
		songIDColumn,
		verseNumColumn,
		textColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		From(tableName).
		Where(sq.Eq{songIDColumn: songID}).
		OrderBy(verseNumColumn + " ASC").
		PlaceholderFormat(sq.Dollar)

	countBuilder := sq.Select("COUNT(*)").
		From(tableName).
		Where(sq.Eq{songIDColumn: songID}).
		PlaceholderFormat(sq.Dollar)

	countQuery, countArgs, err := countBuilder.ToSql()
	if err != nil {
		return nil, 0, err
	}

	var total int
	q := db.Query{
		Name:     "verse_repository.GetAllBySongID.Count",
		QueryRaw: countQuery,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", countArgs))

	err = r.db.DB().ScanOneContext(ctx, &total, q, countArgs...)
	if err != nil {
		return nil, 0, err
	}

	builder = builder.
		Limit(uint64(limit)).
		Offset(uint64(offset))

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, 0, err
	}

	q = db.Query{
		Name:     "verse_repository.GetAllBySongID",
		QueryRaw: query,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", args))

	var verses []model.Verse
	err = r.db.DB().ScanAllContext(ctx, &verses, q, args...)
	if err != nil {
		return nil, 0, err
	}

	return verses, total, nil
}
