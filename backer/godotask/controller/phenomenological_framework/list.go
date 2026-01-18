package phenomenological_framework

import (
	"net/http"

	"github.com/godotask/errors"
	"github.com/gin-gonic/gin"
	"github.com/godotask/controller/user"
)

// ListPhenomenologicalFrameworks: GET /api/phenomenological_framework
func (ctl *PhenomenologicalFrameworkController) ListPhenomenologicalFrameworks(c *gin.Context) {
  userID, _ := user.GetUserIDFromContext(c)

	phenomenologicalFrameworks, err := ctl.Service.ListPhenomenologicalFrameworks(userID)
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