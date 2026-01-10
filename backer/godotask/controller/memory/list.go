package memory

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"strconv"
)

// ListMemories: GET /api/memory
func (ctl *MemoryController) ListMemories(c *gin.Context) {
	memories, err := ctl.Service.ListMemories()
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
func (ctl *MemoryController) ListLimitMemories(c *gin.Context) {
    // クエリパラメータ
    page := 1
    perPage := 20
    const maxPerPage = 100

    if p := c.Query("page"); p != "" {
        if v, err := strconv.Atoi(p); err == nil && v > 0 {
            page = v
        }
    }
    if pp := c.Query("per_page"); pp != "" {
        if v, err := strconv.Atoi(pp); err == nil && v > 0 {
            perPage = v
        }
    }
    if perPage > maxPerPage {
        perPage = maxPerPage
    }

    offset := (page - 1) * perPage

    // Service 側で total も返す想定
    memories, total, err := ctl.Service.ListMemoriesTOPager(page, perPage, offset)
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
        totalPages = int((total + int64(perPage) - 1) / int64(perPage))
    }

    c.JSON(http.StatusOK, gin.H{
        "success":     true,
        "message":     "Memories retrieved",
        "memories":    memories,
        "meta": gin.H{
            "total":       total,
            "total_pages": totalPages,
            "page":        page,
            "per_page":    perPage,
        },
    })
}
