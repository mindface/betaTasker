package insight

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/godotask/errors"
)

// GetInsightData: GET /api/heuristics/insight/:id
func (ctl *HeuristicsInsightController) GetInsightData(c *gin.Context) {
	id := c.Param("id")

	insight, err := ctl.Service.GetInsightById(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error()+" | Failed to get insight data",
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
		"message": "insight retrieved",
		"insight": insight,
	})
}