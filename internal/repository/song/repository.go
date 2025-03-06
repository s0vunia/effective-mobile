package song

import (
	"github.com/s0vunia/effective-mobile/internal/repository"
	"github.com/s0vunia/platform_common/pkg/db"
)

const (
	tableName = "songs"

	idColumn        = "id"
	groupId         = "group_id"
	titleColumn     = "title"
	releaseDate     = "release_date"
	linkColumn      = "link"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// NewRepository creates a new user repository.
func NewRepository(db db.Client) repository.SongRepository {
	return &repo{db: db}
}
