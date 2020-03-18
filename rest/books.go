package rest

import (
	"net/http"

	"github.com/braddle/go-http-template/book"
)

type Books struct {
	p book.AllBooksProvider
}

func NewBooksHandler(p book.AllBooksProvider) Books {
	return Books{p}
}

func (h Books) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
}
