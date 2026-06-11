package ports

import "book-api/internal/book/domain"

// What the outside world(gin) can call inside the core
type BookService interface {
	GetBookByID(id string) (domain.Book, error)
}
