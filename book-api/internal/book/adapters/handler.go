package adapters

import (
	"book-api/internal/book/ports"
	"net/http"

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
