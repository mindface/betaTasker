package language_optimization

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// GetLanguageOptimization: GET /api/language_optimization/:id
func (ctl *LanguageOptimizationController) GetLanguageOptimization(c *gin.Context) {
	id := c.Param("id")
	languageOptimization, err := ctl.Service.GetLanguageOptimizationByID(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Language optimization not found",
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
		"message": "Language optimization retrieved",
		"language_optimization": languageOptimization,
	})
}
