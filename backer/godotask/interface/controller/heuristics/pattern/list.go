package pattern

import (
  "net/http"

  "github.com/gin-gonic/gin"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"github.com/godotask/interface/http/authcontext"
	"github.com/godotask/interface/tools"
  "github.com/godotask/errors"
)

// ListPatternData: GET /api/pattern/pattern
func (ctl *HeuristicsPatternController) ListPatternData(c *gin.Context) {
	userID, _ := authcontext.UserID(c)
	// 分析データのリストを取得
	patterns, err := ctl.Service.ListPattern(userID)
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list pattern data",
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
		"message": "patterns retrieved",
		"patterns": patterns,
	})
}

// ListPatternPager: GET /api/pattern/pattern/pager
func (ctl *HeuristicsPatternController) ListPatternPager(c *gin.Context) {
  pager := tools.ParsePagerQuery(c)
	filter := dtoquery.QueryFilter{
		UserID:  &pager.UserID,
		TaskID:  &pager.TaskID,
		Include: helperquery.ParseIncludeParam(c.Query("include")),
	}
  patterns, total, err := ctl.Service.ListPatternPager(filter, pager)
  if err != nil {
    appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list pattern data",
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
		"message": "patterns retrieved",
		"patterns": patterns,
		"meta": tools.BuildPageMeta(total, pager.Page, pager.Limit),
	})
}
