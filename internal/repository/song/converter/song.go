package converter

import (
	"github.com/s0vunia/effective-mobile/internal/model"
	modelRepo "github.com/s0vunia/effective-mobile/internal/repository/song/model"
)

func ToSongFromRepo(song *modelRepo.Song) *model.Song {
	return &model.Song{
		ID:      song.ID,
		GroupID: song.GroupID,
		Group: model.Group{
			ID:   song.GroupID,
			Name: song.GroupName,
		},
		Title:       song.Title,
		ReleaseDate: song.ReleaseDate,
		Link:        song.Link,
		CreatedAt:   song.CreatedAt,
		UpdatedAt:   song.UpdatedAt,
		Verses:      nil,
	}
}

func ToGroupFromRepo(group *modelRepo.Group) *model.Group {
	return &model.Group{
		ID:        group.ID,
		Name:      group.Name,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
		Songs:     nil,
	}
}
