package multimodal

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ProcessMultimodal - マルチモーダル処理
func (ctrl *MultimodalController) ProcessMultimodal(ctx *gin.Context) {
	// フォームデータの解析
	text := ctx.PostForm("text")
	userIDStr := ctx.PostForm("userId")
	taskIDStr := ctx.PostForm("taskId")

	if text == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "テキストが必要です",
		})
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "無効なユーザーID",
		})
		return
	}

	taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "無効なタスクID",
		})
		return
	}

	var imageURL string
	
	// 画像ファイルの処理
	if file, header, err := ctx.Request.FormFile("image"); err == nil {
		defer file.Close()
		
		// ファイル拡張子チェック
		ext := strings.ToLower(filepath.Ext(header.Filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "サポートされていない画像形式です",
			})
			return
		}

		// ファイルサイズチェック (10MB制限)
		if header.Size > 10*1024*1024 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "ファイルサイズが大きすぎます (最大10MB)",
			})
			return
		}

		// 画像を保存（実際の実装では外部ストレージを使用）
		imageURL = fmt.Sprintf("/uploads/images/%s_%s", uuid.New().String(), header.Filename)
		
		// ここで実際の画像保存処理を実装
		// 例: AWS S3, Google Cloud Storage, ローカルファイルシステムなど
	}

	// マルチモーダル処理を実行
	result, err := ctrl.Service.ProcessTextAndImage(text, imageURL, uint(userID), uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "マルチモーダル処理失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}