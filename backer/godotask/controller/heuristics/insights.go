package heuristics

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

// ListInsights: GET /api/heuristics/insights
func (ctl *HeuristicsController) ListInsights(c *gin.Context) {
	// クエリパラメータの取得
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")
	userID := c.Query("user_id")
	
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	insights, total, err := ctl.Service.GetInsights(userID, limit, offset)
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

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"insights": insights,
			"total":    total,
			"limit":    limit,
			"offset":   offset,
		},
	})
}

// GetInsight: GET /api/heuristics/insights/:id
func (ctl *HeuristicsController) GetInsight(c *gin.Context) {
	id := c.Param("id")
	
	insight, err := ctl.Service.GetInsightById(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			"指定されたインサイトが見つかりません",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data": gin.H{
			"insight": insight,
		},
	})
}