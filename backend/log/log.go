package log

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Middleware для логирования запросов и ответов
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(r.Context())
		start := time.Now()

		log.WithFields(logrus.Fields{
			"requestID": reqID,
			"method":    r.Method,
			"path":      r.URL.Path,
		}).Info("Started request")

		// Обработка запроса
		next.ServeHTTP(w, r)

		log.WithFields(logrus.Fields{
			"requestID": reqID,
			"method":    r.Method,
			"path":      r.URL.Path,
			"duration":  time.Since(start),
		}).Info("Completed request")
	})
}
