package book

type Book struct {
	ISBN10 string `json:"isbn_10"`
	ISBN13 string `json:"isbn_13"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

type Books struct {
	B []Book `json:"books"`
}
