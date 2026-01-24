package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/errors"
)

// GET /api/memory/context/:code
func (ctl *MemoryController) GetMemoryContextByCode(c *gin.Context) {
	code := c.Param("code") // 例: MA-Q-02
	var contexts []model.MemoryContext
	if err := ctl.Service.FindMemoryContextsByCode(code, &contexts); err != nil {
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
