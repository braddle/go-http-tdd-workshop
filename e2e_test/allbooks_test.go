package e2e

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AllBooksSuite struct {
	suite.Suite
}

func TestAllBooksSuite(t *testing.T) {
	suite.Run(t, new(AllBooksSuite))
}

func (s *AllBooksSuite) TestGetAllBook() {
	// TODO create books
	s.T().Skip()
	resp, err := http.Get("http://localhost:8080/books")

	s.Require().NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)

	bytes, _ := ioutil.ReadAll(resp.Body)
	actBody := string(bytes)

	s.JSONEq(getAllBooks(), actBody)
}

func getAllBooks() string {
	f, _ := os.Open("../test_data/books.json")
	b, _ := ioutil.ReadAll(f)

	return string(b)
}
