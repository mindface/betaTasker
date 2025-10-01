package memory

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// ListMemories: GET /api/memory
func (ctl *MemoryController) ListMemories(c *gin.Context) {
	memories, err := ctl.Service.ListMemories()
	fmt.Printf("ListMemories called, found %d memories\n", len(memories))
	if err != nil {
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
	c.JSON(http.StatusOK, gin.H{"memories": memories})
}
