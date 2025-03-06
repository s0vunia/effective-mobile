package service

import (
	"context"

	"github.com/s0vunia/effective-mobile/internal/model"
)

type SongService interface {
	Songs(ctx context.Context, filter model.SongFilter, pagination model.Pagination) ([]model.Song, int, error)
	SongVerses(ctx context.Context, songID int64, pagination model.Pagination) ([]model.Verse, int, error)
	Delete(ctx context.Context, songID int64) error
	Update(ctx context.Context, songID int64, update model.SongUpdate) error
	Add(ctx context.Context, newSong model.SongCreate) (int64, error)
}
