package model

import "time"

type SongFilter struct {
	Group       string
	Song        string
	ReleaseDate time.Time
	Link        string
	Verse       string
}
