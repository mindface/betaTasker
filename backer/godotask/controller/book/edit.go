package book

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)


// EditBook: PUT /api/book/:id
func (ctl *BookController) EditBook(c *gin.Context) {
	id := c.Param("id")
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.UpdateBook(id, &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book edited", "book": book})
}
