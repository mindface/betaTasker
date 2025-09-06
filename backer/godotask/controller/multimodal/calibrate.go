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

// CalibrateUser - ユーザーキャリブレーション
func (ctrl *MultimodalController) CalibrateUser(ctx *gin.Context) {
	userIDStr := ctx.PostForm("userId")
	referenceObject := ctx.PostForm("referenceObject")

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "無効なユーザーID",
		})
		return
	}

	if referenceObject == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "参照オブジェクトが必要です",
		})
		return
	}

	var imageURL string
	
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
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "JPEGまたはPNG画像のみサポートされています",
		})
		return
	}

	// 画像を保存
	imageURL = fmt.Sprintf("/uploads/calibration/%s_%s", uuid.New().String(), header.Filename)
	
	// キャリブレーション処理を実行
	result, err := ctrl.Service.CalibrateUser(uint(userID), referenceObject, imageURL, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "キャリブレーション失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

// GetCalibration - キャリブレーション取得
func (ctrl *MultimodalController) GetCalibration(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ユーザーIDが必要です",
		})
		return
	}

	calibration, err := ctrl.Service.GetUserCalibration(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "キャリブレーション取得失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   calibration,
	})
}