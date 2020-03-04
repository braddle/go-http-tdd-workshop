package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type NotFoundSuite struct {
	suite.Suite
}

func TestNotFoundSuite(t *testing.T) {
	suite.Run(t, new(NotFoundSuite))
}

func (s *NotFoundSuite) Test() {
	resp, err := http.Get("http://localhost:8080/never/going/to/exist")

	s.Require().NoError(err)
	s.Equal(http.StatusNotFound, resp.StatusCode)

	bytes, _ := ioutil.ReadAll(resp.Body)
	actBody := string(bytes)

	s.JSONEq(`{"errors":[{"status":"404","title":"Not Found","detail":"The URI /never/going/to/exist did not match any resources"}]}`, actBody)
}
