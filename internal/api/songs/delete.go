package songs

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/s0vunia/effective-mobile/internal/api"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/service"
	"go.uber.org/zap"
)

// Delete godoc
// @Summary Удаление песни
// @Description Удаление песни по ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 204
// @Failure 400 {object} api.Error "Неверный запрос"
// @Failure 404 {object} api.Error "Песня не найдена"
// @Failure 500 {object} api.Error "Внутренняя ошибка сервера"
// @Router /songs/{id} [delete]
func (i *Implementation) Delete(c echo.Context) error {
	songID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return api.ErrInvalidRequest
	}

	err = i.songService.Delete(c.Request().Context(), songID)
	if err != nil {
		logger.Error("Failed to delete song", zap.Error(err))
		switch {
		case errors.Is(err, service.ErrSongNotFound):
			return api.ErrSongNotFound
		}
		return api.ErrInternal
	}

	return c.NoContent(http.StatusOK)
}
