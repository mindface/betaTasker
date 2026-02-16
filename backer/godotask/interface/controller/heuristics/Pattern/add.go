package pattern

import (
	"net/http"

	"github.com/godotask/infrastructure/db/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

func (ctl *HeuristicsPatternController) AddPatternData(c *gin.Context) {
	var pattern model.HeuristicsPattern

	// リクエストボディをバインド 
	if err := c.ShouldBindJSON(&pattern); err != nil {
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

	// 分析データを追加
	putPattern, err := ctl.Service.CreatePatternData(&pattern)
	if err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error() + " | Failed to add pattern",
		)
		c.JSON(appErr.HTTPStatus, gin.H{
			"code":    appErr.Code,
			"message": appErr.Message,
			"detail":  appErr.Detail,
		})
		return
	}

	// 成功レスポンス
	c.JSON(http.StatusCreated, putPattern)
}
