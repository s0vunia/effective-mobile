package songs

import (
	"errors"
	"net/http"

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
	logger.Info("Create song request")

	var req dto.CreateSongRequest
	if err := c.Bind(&req); err != nil {
		return api.ErrInvalidRequest
	}

	if err := c.Validate(&req); err != nil {
		return api.ErrInvalidRequest
	}

	verses := make([]model.Verse, len(req.Verses))
	for i, v := range req.Verses {
		verses[i] = model.Verse{
			VerseNumber: v.VerseNumber,
			Text:        v.Text,
		}
	}

	songID, err := i.songService.Add(c.Request().Context(), model.SongCreate{
		GroupTitle:  req.GroupTitle,
		Title:       req.Title,
		ReleaseDate: req.ReleaseDate,
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
