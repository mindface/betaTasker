package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookEleteDisplayAction(c *gin.Context) {
	c.HTML(200, "book-add.html", gin.H{})
}

// DeleteBook: DELETE /api/book/:id
func (ctl *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

