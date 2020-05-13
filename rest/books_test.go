package rest_test

import (
	"errors"
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
		b: book.Books{
			B: []book.Book{
				{
					ISBN10: "0099576856",
					ISBN13: "978-0099576853",
					Title:  "Casino Royale",
					Author: "Ian Fleming",
					Pages:  256,
				},
				{
					ISBN10: "9781408855652",
					ISBN13: "978-1408855652",
					Title:  "Harry Potter and the Philosopher's Stone",
					Author: "J.K. Rowling",
					Pages:  352,
				},
				{
					ISBN10: "0060987464",
					ISBN13: "978-0060987466",
					Title:  "Long Hard Road Out Of Hell",
					Author: "Marilyn Manson",
					Pages:  288,
				},
			},
		},
		e: nil,
	}
	h := rest.NewBooksHandler(f)

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	h.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(getAllBooks(), string(body))
}

func (s *BooksSuite) TestWhenNoBooksAvailable() {
	f := FakeAllBooksProvider{
		b: book.Books{B: []book.Book{}},
		e: nil,
	}
	h := rest.NewBooksHandler(f)

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	h.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(getEmptyBooks(), string(body))
}

func (s *BooksSuite) TestErrorRetrievingBooks() {
	f := FakeAllBooksProvider{
		b: book.Books{B: []book.Book{}},
		e: errors.New("ITS BROKED"),
	}
	h := rest.NewBooksHandler(f)

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	h.ServeHTTP(recorder, req)

	body, _ := ioutil.ReadAll(recorder.Body)

	s.Equal(http.StatusOK, recorder.Code)
	s.JSONEq(getEmptyBooks(), string(body))
}

func getAllBooks() string {
	f, _ := os.Open("../test_data/books.json")
	b, _ := ioutil.ReadAll(f)

	return string(b)
}

func getEmptyBooks() string {
	return `{"books": []}`
}

type FakeAllBooksProvider struct {
	b book.Books
	e error
}

func (f FakeAllBooksProvider) GetAllBooks() (book.Books, error) {
	return f.b, f.e
}
