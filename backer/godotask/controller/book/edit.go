package book

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)


// EditBook: PUT /api/book/:id
func (ctl *BookController) EditBook(c *gin.Context) {
	id := c.Param("id")
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	if err := ctl.Service.UpdateBook(id, &book); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to edit book",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book edited", "book": book})
}
