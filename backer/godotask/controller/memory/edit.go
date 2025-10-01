package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
	"github.com/godotask/errors"
)

// EditMemory: PUT /api/memory/:id
func (ctl *MemoryController) EditMemory(c *gin.Context) {
	id := c.Param("id")
	var memory model.Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
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
	if err := ctl.Service.UpdateMemory(id, &memory); err != nil {
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
	if err := ctl.Service.UpdateMemory(id, &memory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to edit memory"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Memory edited", "memory": memory})
}
