package ports

import (
	"book-api/internal/book/domain"
)

// What the outside world(gin) can call inside the core
type BookService interface {
	GetBookByID(id string) (domain.Book, error)
	GetBooks(page int, limit int) ([]domain.Book, int, error)
	CreateBook(book domain.Book) (domain.Book, error)
	UpdateBook(id string, book domain.Book) (domain.Book, error)
	DeleteBook(id string) error
	// interface for fiie upload
	UploadBookCover(bookID string, file domain.FileUploadReq) (string, error)
}
