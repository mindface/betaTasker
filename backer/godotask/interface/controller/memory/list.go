package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/interface/http/authcontext"
	"github.com/godotask/errors"
	"strconv"
)

// ListMemories: GET /api/memory
func (ctl *MemoryController) ListMemories(c *gin.Context) {
	userID, _ := authcontext.UserID(c)
	memories, err := ctl.Service.ListMemories(userID)
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
		"message": "Memories retrieved",
		"memories": memories,
	})
}

// ListLimitMemories: GET /api/memory/pager
func (ctl *MemoryController) ListMemoriesPager(c *gin.Context) {
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

  memories, total, err := ctl.Service.ListMemoriesPager(userID, limit, offset)
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
    "message":     "Memories retrieved",
    "memories":    memories,
    "meta": gin.H{
      "total":       total,
      "total_pages": totalPages,
      "page":        page,
      "limit":    limit,
    },
  })
}
