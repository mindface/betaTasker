package phenomenological_framework

import (
	"net/http"

	"github.com/godotask/errors"
	"github.com/gin-gonic/gin"
)

// ListPhenomenologicalFrameworks: GET /api/phenomenological_framework
func (ctl *PhenomenologicalFrameworkController) ListPhenomenologicalFrameworks(c *gin.Context) {
	phenomenologicalFrameworks, err := ctl.Service.ListPhenomenologicalFrameworks()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list phenomenological frameworks",
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
		"message": "Phenomenological frameworks retrieved",
		"PhenomenologicalFrameworks": phenomenologicalFrameworks,
	})
}