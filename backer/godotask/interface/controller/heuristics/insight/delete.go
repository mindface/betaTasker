package insight

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeleteInsightData: DELETE /api/heuristics/insight/:id
func (ctl *HeuristicsInsightController) DeleteInsightData(c *gin.Context) {
	id := c.Param("id")

	if err := ctl.Service.DeleteInsightData(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error()+" | Failed to delete insight data",
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
		"message":               "Insight data deleted",
		"insight_id": id,
	})
}