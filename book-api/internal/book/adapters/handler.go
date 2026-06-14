package adapters

import (
	"book-api/internal/book/domain"
	"book-api/internal/book/ports"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	// We depend on the port interface, not the concrete struct!
	service ports.BookService
}

func NewHTTPHandler(s ports.BookService) *HTTPHandler {
	return &HTTPHandler{service: s}
}

// Our actual rest handler logic
func (h *HTTPHandler) GetBook(c *gin.Context) {
	id := c.Param("id")

	// call the inner core business logic
	book, err := h.service.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// POST /books
func (h *HTTPHandler) CreateBook(c *gin.Context) {
	var req domain.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json body payload"})
		return
	}

	newBook, err := h.service.CreateBook(req)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newBook)
}

// PUT /books/:id
func (h *HTTPHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var req domain.Book
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalud json body payload"})
		return
	}

	updatedBook, err := h.service.UpdateBook(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, updatedBook)
}

// DELETE /books/:id
func (h *HTTPHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteBook(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// GET /books?page=1&limit=5
func (h *HTTPHandler) GetBooks(c *gin.Context) {
	// Parse query string configurations, provide structural fallbacks
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	items, total, err := h.service.GetBooks(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Standard high-quality pagination envelope payload
	c.JSON(http.StatusOK, gin.H{
		"data":  items,
		"page":  page,
		"limit": limit,
		"total": total,
	})
}
