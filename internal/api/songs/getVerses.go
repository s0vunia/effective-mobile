package songs

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/s0vunia/effective-mobile/internal/api"
	"github.com/s0vunia/effective-mobile/internal/converter"
	"github.com/s0vunia/effective-mobile/internal/dto"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"github.com/s0vunia/effective-mobile/internal/service"
	"go.uber.org/zap"
)

// GetVerses godoc
// @Summary Получение куплетов песни
// @Description Получение текста песни с пагинацией по куплетам
// @Tags verses
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param limit query int false "Количество куплетов на странице (по умолчанию 10)"
// @Param offset query int false "Смещение (по умолчанию 0)"
// @Success 200 {object} dto.VersesResponse "total и verses"
// @Failure 400 {object} api.Error "Неверный запрос"
// @Failure 404 {object} api.Error "Песня не найдена"
// @Failure 500 {object} api.Error "Внутренняя ошибка сервера"
// @Router /songs/{id}/verses [get]
func (i *Implementation) GetVerses(c echo.Context) error {
	songID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return api.ErrInvalidRequest
	}

	var params dto.VersesParams
	if err := c.Bind(&params); err != nil {
		return api.ErrInvalidRequest
	}

	if err := c.Validate(&params); err != nil {
		return api.ErrInvalidRequest
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	verses, total, err := i.songService.SongVerses(c.Request().Context(), songID, model.Pagination{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	if err != nil {
		logger.Error("Failed to get verses", zap.Error(err))
		switch {
		case errors.Is(err, service.ErrSongNotFound):
			return api.ErrSongNotFound
		}
		return api.ErrInternal
	}

	return c.JSON(http.StatusOK, dto.VersesResponse{
		Total:  total,
		Verses: converter.ToVersesResponse(verses),
	})
}
