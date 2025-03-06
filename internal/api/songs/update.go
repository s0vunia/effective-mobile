package songs

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/s0vunia/effective-mobile/internal/api"
	"github.com/s0vunia/effective-mobile/internal/dto"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"go.uber.org/zap"
)

// Update godoc
// @Summary Обновление песни
// @Description Изменение данных песни
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body dto.UpdateSongRequest true "Новые данные песни"
// @Success 200
// @Failure 400 {object} api.Error "Неверный запрос"
// @Failure 404 {object} api.Error "Песня не найдена"
// @Failure 500 {object} api.Error "Внутренняя ошибка сервера"
// @Router /songs/{id} [put]
func (i *Implementation) Update(c echo.Context) error {
	logger.Debug("Update song request received")

	songID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("Invalid song ID", zap.Error(err))
		return api.ErrInvalidRequest
	}

	logger.Debug("Updating song with ID", zap.Int64("id", songID))

	var req dto.UpdateSongRequest
	if err := c.Bind(&req); err != nil {
		logger.Error("Failed to bind request", zap.Error(err))
		return api.ErrInvalidRequest
	}

	if err := c.Validate(&req); err != nil {
		return api.ErrInvalidRequest
	}

	updateInput := model.SongUpdate{}

	if req.GroupID != nil {
		updateInput.GroupID = req.GroupID
	}
	if req.GroupTitle != nil {
		updateInput.GroupTitle = req.GroupTitle
	}
	if req.Title != nil {
		updateInput.Title = req.Title
	}
	if req.ReleaseDate != nil {
		updateInput.ReleaseDate = req.ReleaseDate
	}
	if req.Link != nil {
		updateInput.Link = req.Link
	}

	err = i.songService.Update(c.Request().Context(), songID, updateInput)
	if err != nil {
		logger.Error("Failed to update song", zap.Error(err))
		if errors.Is(err, service.ErrGroupNotFound) {
			return api.ErrGroupNotFound
		}
		return api.ErrInternal
	}

	return c.NoContent(http.StatusOK)
}
