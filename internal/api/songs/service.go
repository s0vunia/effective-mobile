package songs

import (
	"github.com/labstack/echo/v4"
	"github.com/s0vunia/effective-mobile/internal/service"
)

// @title Music Library API
// @version 1.0
// @description API для работы с музыкальной библиотекой

// @host localhost:8080
// @BasePath /api/v1
// @schemes http

type Implementation struct {
	songService service.SongService
}

func NewImplementation(songService service.SongService) *Implementation {
	return &Implementation{
		songService: songService,
	}
}

func (i *Implementation) RegisterHandlers(e *echo.Echo, middlewares ...echo.MiddlewareFunc) {
	// Группа для API песен
	songs := e.Group("/api/v1/songs", middlewares...)

	// Основные операции с песнями
	songs.GET("", i.Library)
	songs.POST("", i.Create)
	songs.PUT("/:id", i.Update)
	songs.DELETE("/:id", i.Delete)

	// Операции с куплетами
	songs.GET("/:id/verses", i.GetVerses)
}
