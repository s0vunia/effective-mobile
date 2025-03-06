package verse

import (
	"github.com/s0vunia/effective-mobile/internal/repository"
	"github.com/s0vunia/platform_common/pkg/db"
)

const (
	tableName = "verses"

	idColumn        = "id"
	songIDColumn    = "song_id"
	verseNumColumn  = "verse_number"
	textColumn      = "text"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.VerseRepository {
	return &repo{
		db: db,
	}
}
