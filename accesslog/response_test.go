package accesslog_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/braddle/go-http-template/accesslog"

	"github.com/stretchr/testify/suite"
)

type ResponseLoggerSuite struct {
	suite.Suite
}

func TestResponseLoggerSuite(t *testing.T) {
	suite.Run(t, new(ResponseLoggerSuite))
}

func (s *ResponseLoggerSuite) TestStatusCode() {
	w := httptest.NewRecorder()
	l := accesslog.NewResponseLogger(w)

	l.WriteHeader(http.StatusTeapot)

	s.Equal(http.StatusTeapot, l.Status())
	s.Equal(http.StatusTeapot, w.Code)
}

func (s *ResponseLoggerSuite) TestContentLength() {
	w := httptest.NewRecorder()
	l := accesslog.NewResponseLogger(w)

	c := "This is the response"

	b, err := l.Write([]byte(c))

	s.NoError(err)
	s.Equal(c, w.Body.String())
	s.Equal(b, l.Size())
}

func (s *ResponseLoggerSuite) TestHeader() {
	w := httptest.NewRecorder()
	l := accesslog.NewResponseLogger(w)

	s.Equal(w.Header(), l.Header())
}
