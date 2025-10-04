package phenomenological_framework

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// PhenomenologicalFramework: GET /api/phenomenological_framework/:id
func (ctl *PhenomenologicalFrameworkController) GetPhenomenologicalFramework(c *gin.Context) {
	id := c.Param("id")
	phenomenologicalFramework, err := ctl.Service.GetPhenomenologicalFrameworkByID(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Phenomenological framework not found",
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
		"message": "Phenomenological framework retrieved",
		"phenomenological_framework": phenomenologicalFramework,
	})
}
