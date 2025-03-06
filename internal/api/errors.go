package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Error struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"invalid request"`
}

var (
	ErrInternal       = echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	ErrInvalidRequest = echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	ErrGroupNotFound  = echo.NewHTTPError(http.StatusNotFound, "group not found")
	ErrSongNotFound   = echo.NewHTTPError(http.StatusNotFound, "song not found")
)
