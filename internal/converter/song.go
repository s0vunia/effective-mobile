package converter

import (
	"github.com/s0vunia/effective-mobile/internal/dto"
	"github.com/s0vunia/effective-mobile/internal/model"
)

func ToSongsResponse(songs []model.Song) []dto.SongResponse {
	var result []dto.SongResponse
	for _, song := range songs {
		result = append(result, ToSongResponse(song))
	}
	return result
}

func ToSongResponse(song model.Song) dto.SongResponse {
	return dto.SongResponse{
		ID:          song.ID,
		Group:       song.Group.Name,
		Title:       song.Title,
		ReleaseDate: song.ReleaseDate,
		Link:        song.Link,
	}
}
