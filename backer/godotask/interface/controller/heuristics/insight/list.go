package insight

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"github.com/godotask/interface/tools"
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

// ListInsightData: GET /api/heuristics/insight/pager
func (ctl *HeuristicsInsightController) ListInsightPager(c *gin.Context) {
	// 分析データのリストを取得
  pager := tools.ParsePagerQuery(c)
	filter := dtoquery.QueryFilter{
		UserID:  &pager.UserID,
		TaskID:  &pager.TaskID,
		Include: helperquery.ParseIncludeParam(c.Query("include")),
	}
	insights, total, err := ctl.Service.ListInsightPager(filter,pager)
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
		"meta": tools.BuildPageMeta(total, pager.Page, pager.Limit),
	})
}
