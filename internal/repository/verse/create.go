package verse

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/platform_common/pkg/db"
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

	var id int64
	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}
	return id, nil
}
