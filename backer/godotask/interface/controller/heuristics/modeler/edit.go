package modeler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/errors"
)

// EditModelerData: PUT /api/heuristics/modeler/:id
func (ctl *HeuristicsModelerController) EditModelerData(c *gin.Context) {
	id := c.Param("id") // URLパラメータからIDを取得

	var modeler model.HeuristicsModeler
	// リクエストボディをバインド
	if err := c.ShouldBindJSON(&modeler); err != nil {
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

	// モデルデータを更新
	if err := ctl.Service.UpdateModelerData(id, &modeler); err != nil {
		appErr := errors.NewAppError(
			errors.SYS_INTERNAL_ERROR,
			errors.GetErrorMessage(errors.SYS_INTERNAL_ERROR),
			err.Error()+" | Failed to update modeler data",
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
		"message": "Modeler data edited",
		"modeler": modeler,
	})
}