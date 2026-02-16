package modeler

import (
	"net/http"

  "github.com/gin-gonic/gin"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"github.com/godotask/interface/http/authcontext"
	"github.com/godotask/interface/tools"
  "github.com/godotask/errors"
)

// ListModelerData: GET /api/heuristics/modeler
func (ctl *HeuristicsModelerController) ListModelerData(c *gin.Context) {
	userID, _ := authcontext.UserID(c)
	// モデルデータのリストを取得
	modeler, err := ctl.Service.ListModeler(userID)
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list modeler data",
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
		"message": "modeler data retrieved",
		"modeler": modeler,
	})
}


// ListModelerPager: GET /api/heuristics/modeler/pager
func (ctl *HeuristicsModelerController) ListModelerPager(c *gin.Context) {
  pager := tools.ParsePagerQuery(c)
	filter := dtoquery.QueryFilter{
		UserID:  &pager.UserID,
		TaskID:  &pager.TaskID,
		Include: helperquery.ParseIncludeParam(c.Query("include")),
	}
  modelers, total, err := ctl.Service.ListModelerPager(filter, pager)
  if err != nil {
    appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to list modeler data",
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
		"message": "modelers retrieved",
		"modelers": modelers,
		"meta": tools.BuildPageMeta(total, pager.Page, pager.Limit),
	})
}

