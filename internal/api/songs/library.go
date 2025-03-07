package songs

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/s0vunia/effective-mobile/internal/api"
	"github.com/s0vunia/effective-mobile/internal/converter"
	"github.com/s0vunia/effective-mobile/internal/dto"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"github.com/s0vunia/effective-mobile/internal/model"
	"go.uber.org/zap"
)

// Library godoc
// @Summary Получение списка песен
// @Description Получение списка песен с фильтрацией и пагинацией
// @Tags songs
// @Accept json
// @Produce json
// @Param group query string false "Фильтр по названию группы"
// @Param song query string false "Фильтр по названию песни"
// @Param release_date query string false "Фильтр по дате выпуска"
// @Param link query string false "Фильтр по ссылке"
// @Param verse query string false "Фильтр по куплету"
// @Param limit query int false "Количество записей на странице (по умолчанию 10)"
// @Param offset query int false "Смещение (по умолчанию 0)"
// @Success 200 {object} dto.LibraryResponse
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /songs [get]
func (i *Implementation) Library(c echo.Context) error {
	logger.Debug("Library request received")
	var params dto.LibraryParams
	if err := c.Bind(&params); err != nil {
		return api.ErrInvalidRequest
	}

	if err := c.Validate(&params); err != nil {
		return api.ValidationError(err)
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	var releaseDate time.Time
	var err error

	if params.ReleaseDate != "" {
		releaseDate, err = time.Parse("2006-01-02", params.ReleaseDate)
		if err != nil {
			logger.Error("Invalid release date format", zap.Error(err))
			return api.ValidationError(err)
		}
	}

	songs, total, err := i.songService.Songs(c.Request().Context(), model.SongFilter{
		Group:       params.Group,
		Song:        params.Song,
		ReleaseDate: releaseDate,
		Link:        params.Link,
		Verse:       params.Verse,
	}, model.Pagination{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
	logger.Debug("Got songs", zap.Any("songs", songs), zap.Int("total", total), zap.Int("limit", params.Limit), zap.Int("offset", params.Offset))

	if err != nil {
		logger.Error("Failed to get songs", zap.Error(err))
		return api.ErrInternal
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"total": total,
		"songs": converter.ToSongsResponse(songs),
	})
}
