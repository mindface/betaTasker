package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"fmt"
)

// GetMemory: GET /api/memory/:id
func (ctl *MemoryController) GetMemory(c *gin.Context) {
	id := c.Param("id")
	memory, err := ctl.Service.GetMemoryByID(id)
	fmt.Printf("GetMemory: Retrieved memory: %+v\n", memory)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			"Memory not found",
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
		"message": "Memory retrieved",
		"memory": memory,
	})
}
