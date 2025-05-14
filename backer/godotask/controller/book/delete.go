package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

func BookEleteDisplayAction(c *gin.Context) {
	c.HTML(200, "book-add.html", gin.H{})
}

func DeleteBookAction(c *gin.Context) {
	id := c.Param("id")
	model.DeleteBookData(id)

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
