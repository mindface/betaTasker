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

func AddBookAction(c *gin.Context) {
	var input struct {
			Title   string `json:"title" binding:"required"`
			Name    string `json:"name" binding:"required"`
			Text    string `json:"text" binding:"required"`
			Disc    string `json:"disc"`
			ImgPath string `json:"imgPath"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Invalid input data",
					"detail":  err.Error(),
			})
			return
	}

	model.AddBookData(0, input.Title, input.Name, input.Text, input.Disc, input.ImgPath)
	c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Book added successfully",
	})

	c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Book added successfully",
	})

	c.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
}
