package rest

import (
	"encoding/json"
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
	//books, _ := h.p.GetAllBooks()
	books := book.Books{}
	b, _ := json.Marshal(books)

	resp.Write(b)
}
