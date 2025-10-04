package phenomenological_framework

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// AddPhenomenologicalFramework: POST /api/phenomenological_framework
func (ctl *PhenomenologicalFrameworkController) AddPhenomenologicalFramework(c *gin.Context) {
	var phenomenologicalFramework model.PhenomenologicalFramework
	if err := c.ShouldBindJSON(&phenomenologicalFramework); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&phenomenologicalFramework); err != nil {
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

	if err := ctl.Service.CreatePhenomenologicalFramework(&phenomenologicalFramework); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to add phenomenological framework",
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
		"message": "Phenomenological framework added",
		"phenomenological_framework": phenomenologicalFramework,
	})
}
