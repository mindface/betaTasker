package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

func BookEleteDisplayAction(c *gin.Context) {
	c.HTML(200, "book-add.html", gin.H{})
}

// DeleteBook: DELETE /api/book/:id
func (ctl *BookController) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteBook(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Failed to delete book",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Book deleted",
		"book_id": id,
	})
}

