package book

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

func UpdateBookAction(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "Book ID is required",
			})
			return
	}

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
					"message": "入力値が正しくありません",
					"detail":  err.Error(),
			})
			return
	}

	model.EditBookData(id, input.Title, input.Name, input.Text, input.Disc, input.ImgPath)
	c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Book updated successfully",
	})

	c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Book update successfully",
	})
}
