package process_optimization

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// GetProcessOptimization: GET /api/process_optimization/:id
func (ctl *ProcessOptimizationController) GetProcessOptimization(c *gin.Context) {
	id := c.Param("id")
	processOptimization, err := ctl.Service.GetProcessOptimizationByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Process optimization not found"})
		return
	}
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Process optimization not found",
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
		"message": "Process optimization retrieved",
		"process_optimization": processOptimization,
	})
}
