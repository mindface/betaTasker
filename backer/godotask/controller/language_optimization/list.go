package language_optimization

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/interface/http/authcontext"
	"github.com/godotask/errors"
)

// ListLanguageOptimizations: GET /api/language_optimization
func (ctl *LanguageOptimizationController) ListLanguageOptimizations(c *gin.Context) {
  userID, _ := authcontext.UserID(c)
	languageOptimizations, err := ctl.Service.ListLanguageOptimizations(userID)
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
		"language_optimizations": languageOptimizations,
	})
}