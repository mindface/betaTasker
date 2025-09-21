package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookListDisplayAction(c *gin.Context) {
	c.HTML(200, "book-list.html", gin.H{
		"list": "[]",
	})
}

// func ApiBookListDisplayAction(c *gin.Context) {
// 	list := model.GetBookList()
// 	if list == "nil" {
// 		list = "[]"
// 	}
// 	c.JSON(200, gin.H{
// 		"list": list,
// 	})
// }

// ListBooks: GET /api/book
func (ctl *BookController) ListBooks(c *gin.Context) {
	books, err := ctl.Service.ListBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list books"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"books": books})
}