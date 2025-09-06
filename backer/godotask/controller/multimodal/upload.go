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

// UploadImage - 画像アップロード
func (ctrl *MultimodalController) UploadImage(ctx *gin.Context) {
	userIDStr := ctx.PostForm("userId")
	taskIDStr := ctx.PostForm("taskId")

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

	// 画像ファイルの処理
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "画像ファイルが必要です",
		})
		return
	}
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

	// 画像を保存
	imageURL := fmt.Sprintf("/uploads/images/%s_%s", uuid.New().String(), header.Filename)
	
	// 実際の保存処理を実行
	result, err := ctrl.Service.SaveImage(file, imageURL, uint(userID), uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "画像保存失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"imageUrl": imageURL,
		"result":   result,
	})
}