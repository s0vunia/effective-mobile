package group

import (
	"context"
	"database/sql"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"github.com/s0vunia/platform_common/pkg/db"
	"go.uber.org/zap"
)

func (r *repo) GetByID(ctx context.Context, id uint) (*model.Group, error) {
	builder := sq.Select(idColumn, nameColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "group_repository.GetByID",
		QueryRaw: query,
	}
	logger.Debug("sql query", zap.String("query name", q.Name), zap.String("query raw", q.QueryRaw), zap.Any("args", args))

	var group model.Group
	err = r.db.DB().ScanOneContext(ctx, &group, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, service.ErrGroupNotFound
		}
		return nil, err
	}

	return &group, nil
}
