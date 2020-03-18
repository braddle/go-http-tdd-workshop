package rest_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/braddle/go-http-template/book"

	"github.com/braddle/go-http-template/rest"

	"github.com/stretchr/testify/suite"
)

type BooksSuite struct {
	suite.Suite
}

func TestBooksSuite(t *testing.T) {
	suite.Run(t, new(BooksSuite))
}

func (s *BooksSuite) TestAvailableBooksReturned() {
	f := FakeAllBooksProvider{
		b: book.Books{B: []book.Book{
			{
				ISBN10: "1234567890",
				ISBN13: "321-1234567890",
				Title:  "Testing All The Things",
				Author: "Mark Bradley",
				Pages:  418,
			},
		}},
		e: nil,
	}
	h := rest.NewBooksHandler(f)

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	h.ServeHTTP(recorder, req)

	//body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	//s.JSONEq(getAllBooks(), string(body))
}

func getAllBooks() string {
	f, _ := os.Open("../test_data/books.json")
	b, _ := ioutil.ReadAll(f)

	return string(b)
}

type FakeAllBooksProvider struct {
	b book.Books
	e error
}

func (f FakeAllBooksProvider) GetAllBooks() (book.Books, error) {
	return f.b, f.e
}
