package dto

type VerseResponse struct {
	ID          int64  `json:"id"`
	SongID      int64  `json:"song_id"`
	VerseNumber int    `json:"verse_number"`
	Text        string `json:"text"`
}

type VerseInput struct {
	VerseNumber int    `json:"verse_number" validate:"required,min=1"`
	Text        string `json:"text" validate:"required"`
}

type VersesParams struct {
	Limit  int `query:"limit" validate:"omitempty,min=1,max=100"`
	Offset int `query:"offset" validate:"omitempty,min=0"`
}

type VersesResponse struct {
	Total  int             `json:"total"`
	Verses []VerseResponse `json:"verses"`
}
