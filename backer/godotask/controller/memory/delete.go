package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeleteMemory: DELETE /api/memory/:id
func (ctl *MemoryController) DeleteMemory(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteMemory(id); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Memory deleted"})
}
