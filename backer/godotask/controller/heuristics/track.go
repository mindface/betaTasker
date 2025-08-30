package heuristics

import (
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
	"github.com/godotask/model"
)

// TrackBehavior: POST /api/heuristics/track
func (ctl *HeuristicsController) TrackBehavior(c *gin.Context) {
	var trackData model.HeuristicsTrackingData
	if err := c.ShouldBindJSON(&trackData); err != nil {
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

	// ユーザーIDの検証
	if trackData.UserID == 0 {
		appErr := errors.NewAppError(
			errors.VAL_MISSING_FIELD,
			errors.GetErrorMessage(errors.VAL_MISSING_FIELD),
			"ユーザーIDは必須項目です",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	if err := ctl.Service.TrackUserBehavior(&trackData); err != nil {
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
		"message": "行動データが正常に記録されました",
		"data": gin.H{
			"tracking_id": trackData.ID,
		},
	})
}

// GetTrackingData: GET /api/heuristics/track/:user_id
func (ctl *HeuristicsController) GetTrackingData(c *gin.Context) {
	userID := c.Param("user_id")
	
	trackingData, err := ctl.Service.GetTrackingDataByUserID(userID)
	if err != nil {
		appErr := errors.NewAppError(
			errors.RES_NOT_FOUND,
			errors.GetErrorMessage(errors.RES_NOT_FOUND),
			"指定されたユーザーのトラッキングデータが見つかりません",
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
			"tracking_data": trackingData,
		},
	})
}