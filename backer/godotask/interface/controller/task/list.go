package task

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/interface/http/authcontext" 
	"github.com/godotask/errors"

	"strconv"
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
    // クエリパラメータ
    page := 1
    limit := 20
    const maxPerPage = 100
	  userID, _ := authcontext.UserID(c)

    if p := c.Query("page"); p != "" {
      if v, err := strconv.Atoi(p); err == nil && v > 0 {
        page = v
      }
    }
    if pp := c.Query("limit"); pp != "" {
      if v, err := strconv.Atoi(pp); err == nil && v > 0 {
        limit = v
      }
    }
    if limit > maxPerPage {
      limit = maxPerPage
    }

    offset := (page - 1) * limit

    // Service 側で total も返す想定
    tasks, total, err := ctl.Service.ListTasksPager(userID, limit, offset)
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
      totalPages = int((total + int64(limit) - 1) / int64(limit))
    }

    c.JSON(http.StatusOK, gin.H{
      "success":     true,
      "message":     "Tasks retrieved",
      "tasks":       tasks,
      "meta": gin.H{
        "total":       total,
        "total_pages": totalPages,
        "page":        page,
        "limit":       limit,
      },
    })
}
