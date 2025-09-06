package multimodal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CompareImages - 画像比較
func (ctrl *MultimodalController) CompareImages(ctx *gin.Context) {
	var req struct {
		ImageA string `json:"imageA" binding:"required"`
		ImageB string `json:"imageB" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "不正なリクエスト",
			"detail": err.Error(),
		})
		return
	}

	comparison, err := ctrl.Service.CompareImages(req.ImageA, req.ImageB)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "画像比較失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"comparison": comparison,
	})
}