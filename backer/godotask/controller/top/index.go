package top

import (
	"github.com/gin-gonic/gin"
)

func IndexDisplayAction(c *gin.Context) {
	c.HTML(200, "book-add.html", gin.H{})
}
