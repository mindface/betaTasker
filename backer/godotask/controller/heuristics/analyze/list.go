package analyze

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/godotask/errors"
)

// ListAnalyzeData: GET /api/heuristics/analyze
func (ctl *AnalyzeController) ListAnalyzeData(c *gin.Context) {
	// 分析データのリストを取得
	analyses, err := ctl.Service.ListAnalyses()
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list analyze data",
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
		"message": "analysis retrieved",
		"analysis": analyses,
	})
}
