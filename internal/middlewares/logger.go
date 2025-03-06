package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/s0vunia/effective-mobile/internal/logger"
	"go.uber.org/zap"
)

// LogMiddleware logs HTTP requests and responses.
func LogMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			err := next(c)

			if err != nil {
				logger.Error(err.Error(),
					zap.String("method", req.Method),
					zap.String("path", req.URL.Path),
					zap.Int("status", res.Status),
				)
			}

			logger.Info("request",
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.Int("status", res.Status),
				zap.Int64("size", res.Size),
				zap.String("ip", c.RealIP()),
				zap.Duration("duration", time.Since(start)),
			)

			return err
		}
	}
}
