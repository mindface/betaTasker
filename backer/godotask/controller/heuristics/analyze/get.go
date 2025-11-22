package analyze

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/godotask/errors"
)

// GetAnalyzeData: GET /api/heuristics/analyze/:id
func (ctl *AnalyzeController) GetAnalyzeData(c *gin.Context) {
	id := c.Param("id")

	analysis, err := ctl.Service.GetAnalysisById(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			err.Error()+" | Failed to get analyze data",
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
		"message": "analysis retrieved",
		"analysis": analysis,
	})
}