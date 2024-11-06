package middleware

import (
	"avitoTest/backend/pkg/errorsx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LoggingMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Запоминаем время начала запроса
		c.Next()            // Выполняем обработчик и остальные middleware

		duration := time.Since(start)
		status := c.Writer.Status()
		path := c.Request.URL.Path

		// Проверяем, были ли ошибки в запросе
		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				if customErr, ok := e.Err.(*errorsx.Errorx); ok {
					logger.Error("Запрос завершился с ошибкой",
						zap.String("code", customErr.Code),
						zap.String("message", customErr.Message),
						zap.String("op", customErr.Op),
						zap.Error(customErr.Err), // Логируем вложенную ошибку
						zap.String("path", path),
						zap.Int("status", status),
						zap.Duration("duration", duration),
					)
				} else {
					// Логирование неизвестных ошибок
					logger.Error("Запрос завершился с неизвестной ошибкой",
						zap.Error(e.Err),
						zap.String("path", path),
						zap.Int("status", status),
						zap.Duration("duration", duration),
					)
				}
			}
		} else {
			// Логирование успешного запроса
			logger.Info("Запрос выполнен успешно",
				zap.String("path", path),
				zap.String("method", c.Request.Method),
				zap.Int("status", status),
				zap.Duration("duration", duration),
			)
		}
	}
}
