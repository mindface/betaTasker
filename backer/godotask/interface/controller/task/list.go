package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"

	"github.com/godotask/interface/http/authcontext" 
	"github.com/godotask/errors"
	"github.com/godotask/interface/tools"
)

// ListTasks: GET /api/task
func (ctl *TaskController) ListTasks(c *gin.Context) {
	userID, _ := authcontext.UserID(c)
	// userID でフィルタしてタスク取得
	tasks, err := ctl.Service.ListTasks(userID)
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
		"message": "Tasks retrieved",
		"tasks": tasks,
	})
}

// ListTasksPager: GET /api/task/pager
func (ctl *TaskController) ListTasksPager(c *gin.Context) {
  pager := tools.ParsePagerQuery(c)
  filter := dtoquery.QueryFilter{
    UserID:  &pager.UserID,
    TaskID:  &pager.TaskID,
    Include: helperquery.ParseIncludeParam(c.Query("include")),
  }

  // Service 側で total も返す想定
  tasks, total, err := ctl.Service.ListTasksPager(filter,pager)
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

  totalPages := 0
  if total > 0 {
    totalPages = int((total + int64(pager.Limit) - 1) / int64(pager.Limit))
  }

  c.JSON(http.StatusOK, gin.H{
    "success":     true,
    "message":     "Tasks retrieved",
    "tasks":       tasks,
    "meta": gin.H{
      "total":       total,
      "total_pages": totalPages,
      "page":        pager.Page,
      "limit":       pager.Limit,
    },
  })
}
