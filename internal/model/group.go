package model

import "time"

// Group представляет музыкальную группу
type Group struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Songs     []Song
}
