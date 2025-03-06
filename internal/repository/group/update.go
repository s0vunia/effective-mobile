package group

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"github.com/s0vunia/platform_common/pkg/db"
)

func (r *repo) Update(ctx context.Context, group *model.Group) error {
	builder := sq.Update(tableName).
		Set(nameColumn, group.Name).
		Set(updatedAtColumn, group.UpdatedAt).
		Where(sq.Eq{idColumn: group.ID}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "group_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return service.ErrGroupNotFound
	}

	return nil
}
