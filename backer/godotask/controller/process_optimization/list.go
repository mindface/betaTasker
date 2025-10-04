package process_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// ListProcessOptimizations: GET /api/process_optimization
func (ctl *ProcessOptimizationController) ListProcessOptimizations(c *gin.Context) {
	processOptimizations, err := ctl.Service.ListProcessOptimizations()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list process optimizations",
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
		"message": "Process optimizations retrieved",
		"ProcessOptimizations": processOptimizations,
	})
}