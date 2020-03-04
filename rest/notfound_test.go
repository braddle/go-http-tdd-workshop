package rest_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/braddle/go-http-template/rest"
	"github.com/stretchr/testify/suite"
)

type NotFoundSuite struct {
	suite.Suite
}

func TestNotFoundSuite(t *testing.T) {
	suite.Run(t, new(NotFoundSuite))
}

func (s *NotFoundSuite) TestHealthyService() {
	h := rest.NotFound{}

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/will-not-exist", nil)
	h.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusNotFound, recorder.Code)
	s.JSONEq(
		`{"errors":[{"status":"404","title":"Not Found","detail":"The URI /will-not-exist did not match any resources"}]}`,
		string(body),
	)
}
