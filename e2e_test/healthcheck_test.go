package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HealthCheckSuite struct {
	suite.Suite
}

func TestHealthcheckSuite(t *testing.T) {
	suite.Run(t, new(HealthCheckSuite))
}

func (s *HealthCheckSuite) TestHealthyService() {
	resp, err := http.Get("http://localhost:8080/healthcheck")

	s.Require().NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	bytes, _ := ioutil.ReadAll(resp.Body)
	actBody := string(bytes)

	s.JSONEq(`{"status": "OK", "errors": []}`, actBody)
}
