package book

import (
	"log"
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

func UpdateBookAction(c *gin.Context) {
	id := c.Param("id")
	var form map[string]interface{}
	c.BindJSON(&form)
	log.Print(form)
	name := form["name"].(string)
	title := form["title"].(string)
	text := form["text"].(string)
	disc := form["disc"].(string)
	imgPath := form["imgPath"].(string)
	model.EditBookData(id, title, name, text, disc, imgPath)

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}
