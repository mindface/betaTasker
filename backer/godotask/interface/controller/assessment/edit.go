package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/errors"
)

// EditAssessment: PUT /api/assessment/:id
func (ctl *AssessmentController) EditAssessment(c *gin.Context) {
	id := c.Param("id")
	var assessment model.Assessment
	if err := c.ShouldBindJSON(&assessment); err != nil {
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
	if err := ctl.Service.UpdateAssessment(id, &assessment); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to edit assessment",
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
		"message": "Assessment edited",
		"assessment": assessment,
	})
}
