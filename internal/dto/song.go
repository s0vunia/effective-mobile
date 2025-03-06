package dto

import "time"

type CreateSongRequest struct {
	GroupTitle  string       `json:"group_title" validate:"required,min=1,max=255"`
	Title       string       `json:"title" validate:"required,min=1,max=255"`
	ReleaseDate string       `json:"release_date" `
	Link        string       `json:"link" validate:"omitempty,url"`
	Verses      []VerseInput `json:"verses" validate:"dive"`
}

type UpdateSongRequest struct {
	GroupID     *int64  `json:"group_id"`
	GroupTitle  *string `json:"group_title" validate:"omitempty,min=1,max=255"`
	Title       *string `json:"title" validate:"omitempty,min=1,max=255"`
	ReleaseDate *string `json:"release_date"`
	Link        *string `json:"link" validate:"omitempty,url"`
}

type LibraryParams struct {
	Group       string `query:"group"`
	Song        string `query:"song"`
	ReleaseDate string `query:"release_date"`
	Link        string `query:"link"`
	Verse       string `query:"verse"`
	Limit       int    `query:"limit" validate:"omitempty,min=1,max=100"`
	Offset      int    `query:"offset" validate:"omitempty,min=0"`
}

type LibraryResponse struct {
	Total int            `json:"total"`
	Songs []SongResponse `json:"songs"`
}

type SongResponse struct {
	ID          int64     `json:"id"`
	Group       string    `query:"group"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	Link        string    `json:"link"`
}
