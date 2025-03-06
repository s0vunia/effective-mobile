package song

import (
	"context"
)

func (s serv) Delete(ctx context.Context, songID int64) error {
	return s.songRepository.Delete(ctx, songID)
}
