package service

import (
	"book-api/internal/book/domain"
	"errors"
	"strconv"
	"sync"
)

type BookService struct {
	mu    sync.Mutex // Protects the slice from concurrent read/write data races
	books []domain.Book
}

func NewBookService() *BookService {
	return &BookService{
		books: []domain.Book{
			{ID: "42", Title: "Anna's Diary", Author: "Agatha Christie"},
		},
	}
}

// Implements the ports.BookService Interface
func (s *BookService) GetBookByID(id string) (domain.Book, error) {
	s.mu.Lock() // Allow multiple concurrent readers
	defer s.mu.Unlock()

	for _, b := range s.books {
		if b.ID == id {
			return b, nil
		}
	}

	return domain.Book{}, errors.New("book not found")
}

// CREATE
func (s *BookService) CreateBook(book domain.Book) (domain.Book, error) {
	s.mu.Lock() // Allow multiple concurrent readers
	defer s.mu.Unlock()

	// Enforce basic validation rules
	if book.Title == "" || book.Author == "" {
		return domain.Book{}, errors.New("title and author are required")
	}

	// Auto-generate an incremental ID based on length for mock purposes
	book.ID = strconv.Itoa(len(s.books) + 42)

	s.books = append(s.books, book)

	return book, nil
}

// UPDATE
func (s *BookService) UpdateBook(id string, updatedBook domain.Book) (domain.Book, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if updatedBook.Title == "" || updatedBook.Author == "" {
		return domain.Book{}, errors.New("title and author are required")
	}

	for i, b := range s.books {
		if b.ID == id {
			s.books[i].Title = updatedBook.Title
			s.books[i].Author = updatedBook.Author
			return s.books[i], nil
		}
	}

	return domain.Book{}, errors.New("book not found to update")
}

// DELETE
func (s *BookService) DeleteBook(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, b := range s.books {
		if b.ID == id {
			// Remove from slice cleaninly by splicing around the index
			s.books = append(s.books[:i], s.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found to delete")
}

// List with pagination
func (s *BookService) GetBooks(page int, limit int) ([]domain.Book, int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	totalItems := len(s.books)
	if totalItems == 0 {
		return []domain.Book{}, 0, nil
	}

	// fallback defaults
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	// calculate math slice indexes
	startIndex := (page - 1) * limit
	endIndex := startIndex + limit

	// Bounds checking to prevent "slice bounds out of range" panic
	if startIndex >= totalItems {
		return []domain.Book{}, totalItems, nil // empty page requested
	}

	if endIndex > totalItems {
		endIndex = totalItems
	}

	paginatedBooks := s.books[startIndex:endIndex]

	return paginatedBooks, totalItems, nil

}
