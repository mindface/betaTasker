package book

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

type Book struct {
	id      int
	name    string
	text    string
	disc    string
	imgPath string
}

func BookAddDisplayAction(c *gin.Context) {
	c.HTML(200, "book-add.html", gin.H{})
}

// AddBook: POST /api/book
func (ctl *BookController) AddBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctl.Service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book added", "book": book})
}
