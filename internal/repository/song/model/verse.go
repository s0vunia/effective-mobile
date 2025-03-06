package model

import "time"

type Verse struct {
	ID          int64     `db:"id"`
	SongID      int64     `db:"song_id"`
	VerseNumber int       `db:"verse_number"`
	Text        string    `db:"text"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
