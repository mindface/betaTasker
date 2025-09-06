package quantification_label

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLabels - ラベル一覧取得
func (ctrl *QuantificationLabelController) GetLabels(ctx *gin.Context) {
	labels, err := ctrl.Service.GetAllLabels()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "ラベル取得失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"labels": labels,
	})
}