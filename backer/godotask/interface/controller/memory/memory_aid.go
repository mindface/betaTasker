package memory

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/errors"
)

// GET /api/memory/aid/:code
func (ctl *MemoryController) GetMemoryAidByCode(c *gin.Context) {
	code := c.Param("code")
	fmt.Printf("GetMemoryAidByCode code: %s", code)
	var contexts []model.MemoryContext
	err := ctl.Service.FindMemoryAidsByCode(code, &contexts)
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
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Memory contexts retrieved",
		"contexts": contexts,
	})
}
