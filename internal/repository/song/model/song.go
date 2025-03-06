package model

import "time"

type Song struct {
	ID          int64     `db:"id"`
	GroupID     int64     `db:"group_id"`
	GroupName   string    `db:"group_name"`
	Title       string    `db:"title"`
	ReleaseDate time.Time `db:"release_date"`
	Link        string    `db:"link"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Verses      []Verse   `db:"verses"`
}
