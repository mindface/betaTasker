package process_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeleteProcessOptimization: DELETE /api/process_optimization/:id
func (ctl *ProcessOptimizationController) DeleteProcessOptimization(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteProcessOptimization(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Failed to delete process optimization",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Success": true,
		"message": "Process optimization deleted"})
}
