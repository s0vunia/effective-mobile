package song

import (
	"context"
	"fmt"

	"github.com/s0vunia/effective-mobile/internal/model"
)

func (s serv) Update(ctx context.Context, songID int64, update model.SongUpdate) error {
	if update.GroupTitle != nil {
		group, err := s.groupRepository.GetByName(ctx, *update.GroupTitle)
		if err != nil {
			return fmt.Errorf("failed to find group by title: %w", err)
		}
		update.GroupID = &group.ID
	}
	return s.songRepository.Update(ctx, songID, &update)
}
