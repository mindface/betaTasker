package language_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeleteLanguageOptimization: DELETE /api/language_optimization/:id
func (ctl *LanguageOptimizationController) DeleteLanguageOptimization(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteLanguageOptimization(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Failed to delete language optimization",
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
		"message": "Language optimization deleted",
		"language_optimization_id": id,})
}
