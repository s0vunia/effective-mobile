package songs

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/s0vunia/effective-mobile/internal/api"
	"github.com/s0vunia/effective-mobile/internal/dto"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"go.uber.org/zap"
)

// Create godoc
// @Summary Создание новой песни
// @Description Добавление новой песни с куплетами
// @Tags songs
// @Accept json
// @Produce json
// @Param song body dto.CreateSongRequest true "Данные новой песни"
// @Success 201 {integer} int64 "ID созданной песни"
// @Failure 400 {object} api.Error "Неверный запрос"
// @Failure 404 {object} api.Error "Группа не найдена"
// @Failure 500 {object} api.Error "Внутренняя ошибка сервера"
// @Router /songs [post]
func (i *Implementation) Create(c echo.Context) error {
	logger.Debug("Create song request received")

	var req dto.CreateSongRequest
	if err := c.Bind(&req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return api.ErrInvalidRequest
	}

	logger.Debug("Request parameters", zap.Any("params", req))

	if err := c.Validate(&req); err != nil {
		logger.Error("Validation failed", zap.Error(err))
		return api.ErrInvalidRequest
	}

	verses := make([]model.Verse, len(req.Verses))
	for i, v := range req.Verses {
		verses[i] = model.Verse{
			VerseNumber: v.VerseNumber,
			Text:        v.Text,
		}
	}
	releaseDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		logger.Error("Invalid release date format", zap.Error(err))
		return api.ErrInvalidRequest
	}

	songID, err := i.songService.Add(c.Request().Context(), model.SongCreate{
		GroupTitle:  req.GroupTitle,
		Title:       req.Title,
		ReleaseDate: releaseDate,
		Link:        req.Link,
		Verses:      verses,
	})
	if err != nil {
		logger.Error("Failed to create song", zap.Error(err))
		if errors.Is(err, service.ErrGroupNotFound) {
			return api.ErrGroupNotFound
		}
		return api.ErrInternal
	}

	return c.JSON(http.StatusCreated, songID)
}
