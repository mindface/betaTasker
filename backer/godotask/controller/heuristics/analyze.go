package heuristics

import (
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"github.com/godotask/model"
)

// Analyze: POST /api/heuristics/analyze
func (ctl *HeuristicsController) Analyze(c *gin.Context) {
	var request model.HeuristicsAnalysisRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		appErr := errors.NewAppError(
			errors.VAL_INVALID_INPUT,
			errors.GetErrorMessage(errors.VAL_INVALID_INPUT),
			err.Error(),
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	analysis, err := ctl.Service.AnalyzeData(&request)
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
		"message": "分析が正常に完了しました",
		"data": gin.H{
			"analysis": analysis,
		},
	})
}

// GetAnalysis: GET /api/heuristics/analyze/:id
func (ctl *HeuristicsController) GetAnalysis(c *gin.Context) {
	id := c.Param("id")
	
	analysis, err := ctl.Service.GetAnalysisById(id)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			"指定された分析が見つかりません",
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
			"analysis": analysis,
		},
	})
}