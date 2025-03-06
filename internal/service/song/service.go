package song

import (
	"github.com/s0vunia/effective-mobile/internal/repository"
	"github.com/s0vunia/effective-mobile/internal/service"
)

type serv struct {
	songRepository  repository.SongRepository
	groupRepository repository.GroupRepository
	verseRepository repository.VerseRepository
}

func NewService(
	songRepository repository.SongRepository,
	groupRepository repository.GroupRepository,
	verseRepository repository.VerseRepository,
) service.SongService {
	return &serv{
		songRepository:  songRepository,
		groupRepository: groupRepository,
		verseRepository: verseRepository,
	}
}
