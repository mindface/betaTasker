package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"github.com/godotask/controller/user" 
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
	userID, ok := user.GetUserIDFromContext(c)
	if !ok {
		appErr := errors.NewAppError(
			errors.AUTH_UNAUTHORIZED,
			"Unauthorized",
			"User ID not found in context",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	books, err := ctl.Service.ListBooks(userID)
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list books",
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
		"message": "Books retrieved",
		"books": books,
	})
}