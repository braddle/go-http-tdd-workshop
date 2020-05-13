package rest_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/braddle/go-http-template/rest"
	"github.com/stretchr/testify/suite"
)

type HealthCheckSuite struct {
	suite.Suite
}

func TestHealthCheckSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckSuite))
}

func (s *HealthCheckSuite) TestHealthyService() {
	h := rest.HealthCheck{}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/healthcheck", nil)
	h.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(`{"status": "OK", "errors": ["string"]}`, string(body))
}
