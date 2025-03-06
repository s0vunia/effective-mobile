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

	versesToCreate := make([]model.Verse, 0, len(song.Verses))
	for _, verse := range song.Verses {
		verseToCreate := model.Verse{
			SongID:      songID,
			VerseNumber: verse.VerseNumber,
			Text:        verse.Text,
		}
		versesToCreate = append(versesToCreate, verseToCreate)
	}

	_, err = s.verseRepository.CreateBatch(ctx, versesToCreate)
	if err != nil {
		return 0, err
	}

	return songID, nil
}
