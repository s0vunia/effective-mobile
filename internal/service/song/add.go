package song

import (
	"context"
	"fmt"

	"github.com/s0vunia/effective-mobile/internal/model"
)

func (s serv) Add(ctx context.Context, song model.SongCreate) (int64, error) {
	group, err := s.groupRepository.GetByName(ctx, song.GroupTitle)
	if err != nil {
		return 0, fmt.Errorf("failed to find group by title: %w", err)
	}

	songToCreate := &model.SongCreate{
		GroupID:     group.ID,
		Title:       song.Title,
		ReleaseDate: song.ReleaseDate,
		Link:        song.Link,
	}

	songID, err := s.songRepository.Create(ctx, songToCreate)
	if err != nil {
		return 0, err
	}

	for _, verse := range song.Verses {
		verseToCreate := &model.Verse{
			SongID:      songID,
			VerseNumber: verse.VerseNumber,
			Text:        verse.Text,
		}
		_, err = s.verseRepository.Create(ctx, verseToCreate)
		if err != nil {
			return 0, fmt.Errorf("failed to create verse: %w", err)
		}
	}

	return songID, nil
}
