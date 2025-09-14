package heuristics

import (
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"github.com/godotask/model"
)

// DetectPatterns: GET /api/heuristics/patterns
func (ctl *HeuristicsController) DetectPatterns(c *gin.Context) {
	// クエリパラメータの取得
	userID := c.Query("user_id")
	dataType := c.DefaultQuery("type", "all")
	period := c.DefaultQuery("period", "7d")

	patterns, err := ctl.Service.DetectPatterns(userID, dataType, period)
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
		"message": "パターン検出が完了しました",
		"data": gin.H{
			"patterns": patterns,
			"metadata": gin.H{
				"user_id":   userID,
				"data_type": dataType,
				"period":    period,
			},
		},
	})
}

// TrainModel: POST /api/heuristics/patterns/train
func (ctl *HeuristicsController) TrainModel(c *gin.Context) {
	var request model.HeuristicsTrainRequest
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

	// モデルタイプの検証
	if request.ModelType == "" {
		appErr := errors.NewAppError(
			errors.VAL_MISSING_FIELD,
			errors.GetErrorMessage(errors.VAL_MISSING_FIELD),
			"モデルタイプは必須項目です",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	result, err := ctl.Service.TrainModel(&request)
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
		"message": "モデルトレーニングが開始されました",
		"data": gin.H{
			"training_id": result.ID,
			"status":      result.Status,
			"model_type":  request.ModelType,
		},
	})
}