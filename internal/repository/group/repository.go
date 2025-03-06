package group

import (
	"github.com/s0vunia/effective-mobile/internal/repository"
	"github.com/s0vunia/platform_common/pkg/db"
)

const (
	tableName = "groups"

	idColumn        = "id"
	nameColumn      = "name"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.GroupRepository {
	return &repo{
		db: db,
	}
}
