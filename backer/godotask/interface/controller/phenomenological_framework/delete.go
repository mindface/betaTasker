package phenomenological_framework

import (
	"net/http"

	"github.com/godotask/errors"
	"github.com/gin-gonic/gin"
)

// DeletePhenomenologicalFramework: DELETE /api/phenomenological_framework/:id
func (ctl *PhenomenologicalFrameworkController) DeletePhenomenologicalFramework(c *gin.Context) {
	id := c.Param("id")
	if err := ctl.Service.DeletePhenomenologicalFramework(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete phenomenological framework"})
		return
	}
	if err := ctl.Service.DeletePhenomenologicalFramework(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error() + " | Failed to delete phenomenological framework",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Phenomenological framework deleted"})
}
