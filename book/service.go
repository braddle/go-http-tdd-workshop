package book

type AllBooksProvider interface {
	GetAllBooks() (Books, error)
}
