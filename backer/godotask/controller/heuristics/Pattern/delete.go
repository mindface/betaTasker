package insight

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// DeletePatternData: DELETE /api/heuristics/pattern/:id
func (ctl *HeuristicsPatternController) DeletePatternData(c *gin.Context) {
	id := c.Param("id")

	if err := ctl.Service.DeletePatternData(id); err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error()+" | Failed to delete Pattern data",
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
		"message":               "Pattern data deleted",
		"pattern_id": id,
	})
}