package process_optimization

import (
	"net/http"

	"github.com/godotask/infrastructure/db/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// AddProcessOptimization: POST /api/process_optimization
func (ctl *ProcessOptimizationController) AddProcessOptimization(c *gin.Context) {
	var processOptimization model.ProcessOptimization
	if err := c.ShouldBindJSON(&processOptimization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&processOptimization); err != nil {
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
	if err := ctl.Service.CreateProcessOptimization(&processOptimization); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to add process optimization",
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
		"message": "Process optimization added",
		"process_optimization": processOptimization,
	})
}
