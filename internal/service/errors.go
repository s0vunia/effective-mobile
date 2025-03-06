package service

import "errors"

var (
	ErrSongNotFound  = errors.New("song not found")
	ErrGroupNotFound = errors.New("group not found")
	ErrVerseNotFound = errors.New("verse not found")
)
