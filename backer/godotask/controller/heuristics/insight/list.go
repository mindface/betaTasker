package insight

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/godotask/errors"
)

// ListInsightData: GET /api/heuristics/insight
func (ctl *HeuristicsInsightController) ListInsightData(c *gin.Context) {
	// 分析データのリストを取得
	insights, err := ctl.Service.ListInsight()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list insight data",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "insights retrieved",
		"insights": insights,
	})
}
