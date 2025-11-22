package insight

import (
	"net/http"

	"github.com/godotask/model"
	"github.com/gin-gonic/gin"
	"github.com/godotask/errors"
)

func (ctl *HeuristicsPatternController) AddInsightData(c *gin.Context) {
    var insight model.HeuristicsPattern

    // リクエストボディをバインド 
    if err := c.ShouldBindJSON(&insight); err != nil {
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
    putInsight, err := ctl.Service.CreateInsightData(&insight)
    if err != nil {
			appErr := errors.NewAppError(
				errors.SYS_INTERNAL_ERROR,
				errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
				err.Error() + " | Failed to add analyze",
			)
			c.JSON(appErr.HTTPStatus, gin.H{
				"code":    appErr.Code,
				"message": appErr.Message,
				"detail":  appErr.Detail,
			})
      return
    }

    // 成功レスポンス
    c.JSON(http.StatusCreated, putInsight)
}
