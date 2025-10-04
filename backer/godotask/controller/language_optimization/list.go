package language_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// ListLanguageOptimizations: GET /api/language_optimization
func (ctl *LanguageOptimizationController) ListLanguageOptimizations(c *gin.Context) {
	languageOptimizations, err := ctl.Service.ListLanguageOptimizations()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list language optimizations",
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
		"message": "Language optimizations retrieved",
		"LanguageOptimizations": languageOptimizations,
	})
}