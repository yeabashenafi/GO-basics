package main

import (
	"book-api/internal/book/adapters"
	"book-api/internal/book/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize core business logic
	bookService := service.NewBookService()

	// Inject business logic into the HTTP adapter handler
	bookHandler := adapters.NewHTTPHandler(bookService)

	// Define the REST routes
	r.GET("/books/:id", bookHandler.GetBook)
	r.GET("/books", bookHandler.GetBooks)
	r.POST("/books", bookHandler.CreateBook)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)
	r.POST("/books/:id/upload-cover", bookHandler.UploadCover)

	// Fire up the server
	r.Run(":9090")
}
