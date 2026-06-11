package main

import (
	"book-api/internal/book/adapters"
	"book-api/internal/book/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize core business logic
	bookService := domain.NewBookService()

	// Inject business logic into the HTTP adapter handler
	bookHandler := adapters.NewHTTPHandler(bookService)

	// Define the REST routes
	r.GET("/books/:id", bookHandler.GetBook)

	// Fire up the server
	r.Run(":9090")
}
