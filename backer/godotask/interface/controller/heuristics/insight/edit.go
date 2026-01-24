package insight

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/errors"
)

// EditIinsightData: PUT /api/heuristics/insight/:id
func (ctl *HeuristicsInsightController) EditInsightsData(c *gin.Context) {
	id := c.Param("id") // URLパラメータからIDを取得

	var insight model.HeuristicsInsight
	// リクエストボディをバインド
	if err := c.ShouldBindJSON(&insight); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	// 分析データを更新
	if err := ctl.Service.UpdateInsightData(id, &insight); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error()+" | Failed to update insight data",
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
		"message": "Knowledge pattern edited",
		"insight": insight,
	})
}