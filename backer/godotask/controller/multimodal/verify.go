package multimodal

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// VerifyQuantification - マルチモーダル結果検証
func (ctrl *MultimodalController) VerifyQuantification(ctx *gin.Context) {
	var req struct {
		DataID   string `json:"dataId" binding:"required"`
		Feedback string `json:"feedback" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "不正なリクエスト",
			"detail": err.Error(),
		})
		return
	}

	// フィードバック値の検証
	validFeedback := []string{"correct", "too_high", "too_low", "incorrect"}
	isValid := false
	for _, valid := range validFeedback {
		if req.Feedback == valid {
			isValid = true
			break
		}
	}

	if !isValid {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "無効なフィードバック値",
			"valid": validFeedback,
		})
		return
	}

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "認証が必要です",
		})
		return
	}

	err := ctrl.Service.VerifyResult(req.DataID, req.Feedback, fmt.Sprintf("%v", userID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "検証処理失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "フィードバックが記録されました",
	})
}

// SubmitFeedback - フィードバック送信
func (ctrl *MultimodalController) SubmitFeedback(ctx *gin.Context) {
	var req struct {
		ProcessID    string  `json:"processId" binding:"required"`
		Accuracy     float64 `json:"accuracy"`
		Comments     string  `json:"comments"`
		Suggestions  string  `json:"suggestions"`
		UserRating   int     `json:"userRating"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "不正なリクエスト",
			"detail": err.Error(),
		})
		return
	}

	// ユーザーIDの取得
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "認証が必要です",
		})
		return
	}

	err := ctrl.Service.SubmitFeedback(req.ProcessID, fmt.Sprintf("%v", userID), req.Accuracy, req.Comments, req.Suggestions, req.UserRating)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "フィードバック送信失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "フィードバックが送信されました",
	})
}