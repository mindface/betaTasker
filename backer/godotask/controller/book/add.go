package book

import (
	"log"
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

func AddBookAction(c *gin.Context) {
	var form map[string]interface{}
	c.BindJSON(&form)
	log.Print(form)
	name := form["name"].(string)
	title := form["title"].(string)
	text := form["text"].(string)
	disc := form["disc"].(string)
	imgPath := form["imgPath"].(string)

	model.AddBookData(0, title, name, text, disc, imgPath)

	c.JSON(http.StatusOK, gin.H{"str": "OK"})
}
