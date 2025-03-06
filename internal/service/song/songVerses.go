package song

import (
	"context"

	"github.com/s0vunia/effective-mobile/internal/model"
)

func (s serv) SongVerses(ctx context.Context, songID int64, pagination model.Pagination) ([]model.Verse, int, error) {
	_, err := s.songRepository.GetByID(ctx, songID)
	if err != nil {
		return []model.Verse{}, 0, err
	}
	return s.verseRepository.GetAllBySongID(ctx, songID, pagination.Limit, pagination.Offset)
}
