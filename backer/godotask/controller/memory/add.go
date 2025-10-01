package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/model"
	"github.com/godotask/errors"
	"fmt"
)

// AddMemory: POST /api/memory
func (ctl *MemoryController) AddMemory(c *gin.Context) {
	var memory model.Memory
	if err := c.ShouldBindJSON(&memory); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		fmt.Printf("AddMemory BindJSON error: %v\n", err)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	if err := ctl.Service.CreateMemory(&memory); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error(),
		)
		fmt.Printf("AddMemory CreateMemory error: %v\n", err)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	fmt.Printf("AddMemory success: %+v\n", memory)
	c.JSON(http.StatusOK, gin.H{"message": "Memory added", "memory": memory})
}
