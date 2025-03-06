package model

import "time"

// Verse представляет куплет песни
type Verse struct {
	ID          int64
	SongID      int64
	VerseNumber int
	Text        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
