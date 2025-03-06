package song

import (
	"context"

	"github.com/s0vunia/effective-mobile/internal/model"
)

func (s serv) Songs(ctx context.Context, filter model.SongFilter, pagination model.Pagination) ([]model.Song, int, error) {
	return s.songRepository.GetAll(ctx, filter, pagination.Limit, pagination.Offset)
}
