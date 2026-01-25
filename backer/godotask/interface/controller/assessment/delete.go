package assessment

import (
	"net/http"
	"github.com/gin-gonic/gin"	
	"github.com/godotask/errors"
)

// DeleteAssessment: DELETE /api/assessment/:id
func (ctl *AssessmentController) DeleteAssessment(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeleteAssessment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete assessment"})
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Failed to delete assessment",
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
		"message": "Assessment deleted",
		"assessment_id": id,
	})
}
