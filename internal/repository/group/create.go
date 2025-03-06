package group

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/platform_common/pkg/db"
)

func (r *repo) Create(ctx context.Context, group *model.Group) (int64, error) {
	builder := sq.Insert(tableName).
		Columns(
			nameColumn,
		).
		Values(
			group.Name,
		).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "group_repository.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil
}
