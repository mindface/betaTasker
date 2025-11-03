package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// GetAssessment: GET /api/assessment/:id
func (ctl *AssessmentController) GetAssessment(c *gin.Context) {
	id := c.Param("id")
	assessment, err := ctl.Service.GetAssessmentByID(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Assessments not found",
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
		"message": "Assessment retrieved",
		"assessment": assessment,
	})
}
