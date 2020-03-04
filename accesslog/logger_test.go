package accesslog_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/braddle/go-http-template/clock"

	"github.com/braddle/go-http-template/accesslog"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type LoggerSuite struct {
	suite.Suite
}

func TestLoggerSuite(t *testing.T) {
	suite.Run(t, new(LoggerSuite))
}

const (
	ResponseContent = "This is the repose body"
	url             = "/people/123456"
	UserAgent       = "Braddle_Client/0.1"
	remoteAddress   = "192.169.10.10"
	requestContent  = "This is a test, oh yes this is a test!"
)

func (s *LoggerSuite) TestAccessLogging() {
	logBuf := bytes.NewBufferString("")
	logrus.SetOutput(logBuf)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	t := time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC)
	al := accesslog.New(clock.Fake(t))

	r := mux.NewRouter()
	r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Second)
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte(ResponseContent))
	})
	r.Use(al.Logger)

	body := bytes.NewBufferString(requestContent)
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Set("User-Agent", UserAgent)
	req.RemoteAddr = remoteAddress
	r.ServeHTTP(httptest.NewRecorder(), req)

	access := make(map[string]interface{})
	sc := bufio.NewScanner(logBuf)
	sc.Scan()

	json.Unmarshal(sc.Bytes(), &access)

	s.Equal("http_access", access["type"])

	// Request
	s.Equal(url, access["request"])
	s.Equal(http.MethodPost, access["method"])
	s.Equal(UserAgent, access["http_user_agent"])
	s.Equal(remoteAddress, access["remote_addr"])
	s.Equal(len(requestContent), int(access["body_bytes_sent"].(float64)))

	// Response
	s.Equal(http.StatusTeapot, int(access["status"].(float64)))
	s.Equal(len(ResponseContent), int(access["size"].(float64)))
	s.Equal(t.Format("02/Jan/2006 03:04:05 -0700"), access["time_local"])
	s.Greater(float64(1), access["request_time"].(float64))
}
