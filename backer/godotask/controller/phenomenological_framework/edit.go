package phenomenological_framework

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
)

// EditPhenomenologicalFramework: PUT /api/phenomenological_framework/:id
func (ctl *PhenomenologicalFrameworkController) EditPhenomenologicalFramework(c *gin.Context) {
	id := c.Param("id")
	var phenomenologicalFramework model.PhenomenologicalFramework
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
	if err := ctl.Service.UpdatePhenomenologicalFramework(id, &phenomenologicalFramework); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to edit phenomenological framework",
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
		"message": "Phenomenological framework edited",
		"phenomenological_framework": phenomenologicalFramework,
	})
}
