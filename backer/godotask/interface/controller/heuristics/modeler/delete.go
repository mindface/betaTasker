package modeler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeleteModelerData: DELETE /api/heuristics/modeler/:id
func (ctl *HeuristicsModelerController) DeleteModelerData(c *gin.Context) {
	id := c.Param("id")

	if err := ctl.Service.DeleteModelerData(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error()+" | Failed to delete Modeler data",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":               true,
		"message":               "Modeler data deleted",
		"modeler_id": id,
	})
}