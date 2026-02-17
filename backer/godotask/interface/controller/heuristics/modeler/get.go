package modeler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/godotask/errors"
)

// GetModelerData: GET /api/heuristics/modeler/:id
func (ctl *HeuristicsModelerController) GetModelerData(c *gin.Context) {
	id := c.Param("id")

	modeler, err := ctl.Service.GetModelerById(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error()+" | Failed to get modeler data",
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
		"message": "modeler data retrieved",
		"modeler": modeler,
	})
}