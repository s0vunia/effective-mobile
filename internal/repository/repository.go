package repository

import (
	"context"

	"github.com/s0vunia/effective-mobile/internal/model"
)

type SongRepository interface {
	GetByID(ctx context.Context, id int64) (*model.Song, error)
	GetAll(ctx context.Context, filter model.SongFilter, limit, offset int) ([]model.Song, int, error)
	Create(ctx context.Context, song *model.SongCreate) (int64, error)
	Update(ctx context.Context, id int64, song *model.SongUpdate) error
	Delete(ctx context.Context, id int64) error
}

type GroupRepository interface {
	GetByID(ctx context.Context, id uint) (*model.Group, error)
	GetByName(ctx context.Context, name string) (*model.Group, error)
	Create(ctx context.Context, group *model.Group) (int64, error)
	Update(ctx context.Context, group *model.Group) error
	Delete(ctx context.Context, id uint) error
}

type VerseRepository interface {
	GetByID(ctx context.Context, id int64) (*model.Verse, error)
	GetAllBySongID(ctx context.Context, songID int64, limit, offset int) ([]model.Verse, int, error)
	Create(ctx context.Context, verse *model.Verse) (int64, error)
	Update(ctx context.Context, verse *model.Verse) error
	Delete(ctx context.Context, id int64) error
}
