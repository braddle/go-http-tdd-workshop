package accesslog

import (
	"net/http"

	"github.com/braddle/go-http-template/clock"

	"github.com/sirupsen/logrus"
)

type AccessLogger struct {
	c clock.CurrentTimeFactory
}

func New(c clock.CurrentTimeFactory) AccessLogger {
	return AccessLogger{c}
}

func (al AccessLogger) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := al.c.GetCurrentTime()
		wl := NewResponseLogger(w)
		next.ServeHTTP(wl, r)
		end := al.c.GetCurrentTime()

		logrus.WithFields(
			logrus.Fields{
				"type":            "http_access",
				"method":          r.Method,
				"request":         r.URL.Path,
				"status":          wl.Status(),
				"size":            wl.Size(),
				"http_user_agent": r.UserAgent(),
				"time_local":      start.Format("02/Jan/2006 03:04:05 -0700"),
				"remote_addr":     r.RemoteAddr,
				"body_bytes_sent": r.ContentLength,
				"request_time":    end.Sub(start).Seconds(),
			}).Info()
	})
}
