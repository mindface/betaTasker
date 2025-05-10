package book

import (
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
)

func BookListDisplayAction(c *gin.Context) {
	c.HTML(200, "book-list.html", gin.H{
		"list": model.GetBookList(),
	})
}

func ApiBookListDisplayAction(c *gin.Context) {
	list := model.GetBookList()
	if list == "nil" {
		list = "[]"
	}
	c.JSON(200, gin.H{
		"list": list,
	})
}
