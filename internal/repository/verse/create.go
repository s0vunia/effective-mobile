package verse

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/platform_common/pkg/db"
	"go.uber.org/zap"
)

func (r *repo) Create(ctx context.Context, verse *model.Verse) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(
			songIDColumn,
			verseNumColumn,
			textColumn,
		).
		Values(
			verse.SongID,
			verse.VerseNumber,
			verse.Text,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "verse_repository.Create",
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

func (r *repo) CreateBatch(ctx context.Context, verses []model.Verse) ([]int64, error) {
	if len(verses) == 0 {
		return nil, nil
	}

	builder := sq.Insert(tableName).
		Columns(songIDColumn, verseNumColumn, textColumn).
		PlaceholderFormat(sq.Dollar)

	for _, verse := range verses {
		builder = builder.Values(verse.SongID, verse.VerseNumber, verse.Text)
	}

	query, args, err := builder.Suffix("RETURNING id").ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "verse_repository.Create",
		QueryRaw: query,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", args))

	var ids []int64
	err = r.db.DB().ScanAllContext(ctx, &ids, q, args...)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
