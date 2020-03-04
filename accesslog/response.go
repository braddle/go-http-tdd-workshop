package accesslog

import "net/http"

type ResponseLogger struct {
	w      http.ResponseWriter
	status int
	size   int
}

func NewResponseLogger(w http.ResponseWriter) *ResponseLogger {
	return &ResponseLogger{w, 0, 0}
}

func (l *ResponseLogger) WriteHeader(statusCode int) {
	l.w.WriteHeader(statusCode)
	l.status = statusCode
}

func (l *ResponseLogger) Status() int {
	return l.status
}

func (l *ResponseLogger) Size() int {
	return l.size
}

func (l *ResponseLogger) Write(bytes []byte) (int, error) {
	s, err := l.w.Write(bytes)
	l.size += s

	return s, err
}

func (l *ResponseLogger) Header() http.Header {
	return l.w.Header()
}
