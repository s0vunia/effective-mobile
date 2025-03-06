package converter

import (
	"github.com/s0vunia/effective-mobile/internal/dto"
	"github.com/s0vunia/effective-mobile/internal/model"
)

func ToVerseResponse(verse model.Verse) dto.VerseResponse {
	return dto.VerseResponse{
		ID:          verse.ID,
		SongID:      verse.SongID,
		VerseNumber: verse.VerseNumber,
		Text:        verse.Text,
	}
}

func ToVersesResponse(verses []model.Verse) []dto.VerseResponse {
	result := make([]dto.VerseResponse, len(verses))
	for i, verse := range verses {
		result[i] = ToVerseResponse(verse)
	}
	return result
}
