package analyze

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/godotask/interface/http/authcontext"
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

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "analysis retrieved",
		"analysis": analyses,
	})
}

// ListAnalyzeData: GET /api/heuristics/analyze/pager
func (ctl *AnalyzeController) ListAnalyzePager(c *gin.Context) {
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
	// 分析データのリストを取得
	analyses, total, err := ctl.Service.ListAnalysesPager(userID, offset, limit)
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
    "total": total,
	})
}
