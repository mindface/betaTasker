package analyze

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"github.com/godotask/interface/tools"
	"github.com/godotask/errors"
	"fmt"
)

// ListAnalyzeData: GET /api/heuristics/analyze
func (ctl *HeuristicsAnalyzeController) ListAnalyze(c *gin.Context) {
	// 分析データのリストを取得
	analyses, err := ctl.Service.ListAnalyze()
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

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "analysis retrieved",
		"analysis": analyses,
	})
}

// ListAnalyzeData: GET /api/heuristics/analyze/pager
func (ctl *HeuristicsAnalyzeController) ListAnalyzePager(c *gin.Context) {
  pager := tools.ParsePagerQuery(c)
	filter := dtoquery.QueryFilter{
		UserID:  &pager.UserID,
		TaskID:  &pager.TaskID,
		Include: helperquery.ParseIncludeParam(c.Query("include")),
	}
	fmt.Printf("eeeeeeeee")
	fmt.Printf("%d",filter.Include)
  analyses, total, err := ctl.Service.ListAnalysesPager(filter,pager)
  if err != nil {
    appErr := errors.NewAppError(
      errors.SYS_INTERNAL_ERROR,
      errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
      err.Error(),
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
		"analysis": analyses,
		"meta": tools.BuildPageMeta(total, pager.Page, pager.Limit),
	})
}
