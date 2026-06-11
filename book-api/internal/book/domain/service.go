package domain

import "errors"

type BookService struct{}

func NewBookService() *BookService {
	return &BookService{}
}

// Implements the ports.BookService Interface
func (s *BookService) GetBookByID(id string) (Book, error) {
	// A Mock business logic for fetching
	if id == "42" {
		return Book{ID: "42", Title: "Anna's Diary", Author: "Agatha Christie"}, nil
	}

	return Book{}, errors.New("book not found")
}
