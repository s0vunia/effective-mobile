package verse

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"github.com/s0vunia/platform_common/pkg/db"
)

func (r *repo) Update(ctx context.Context, verse *model.Verse) error {
	builder := sq.Update(tableName).
		Set(songIDColumn, verse.SongID).
		Set(verseNumColumn, verse.VerseNumber).
		Set(textColumn, verse.Text).
		Set(updatedAtColumn, verse.UpdatedAt).
		Where(sq.Eq{idColumn: verse.ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "verse_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return service.ErrVerseNotFound
	}
	if err != nil {
		return err
	}

	return nil
}
