package model

import "time"

// Song представляет песню в библиотеке
type Song struct {
	ID          int64
	GroupID     int64
	Group       Group
	Title       string
	ReleaseDate time.Time
	Link        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Verses      []Verse
}

type SongCreate struct {
	GroupID     int64
	GroupTitle  string
	Title       string
	ReleaseDate time.Time
	Link        string
	Verses      []Verse
}

type SongUpdate struct {
	GroupID     *int64
	GroupTitle  *string
	Title       *string
	ReleaseDate *time.Time
	Link        *string
}
