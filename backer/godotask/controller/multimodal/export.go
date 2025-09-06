package multimodal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ExportData - データエクスポート
func (ctrl *MultimodalController) ExportData(ctx *gin.Context) {
	userIDStr := ctx.Query("userId")
	format := ctx.DefaultQuery("format", "json")
	startDate := ctx.Query("startDate")
	endDate := ctx.Query("endDate")

	var userID *uint
	if userIDStr != "" {
		if id, err := strconv.ParseUint(userIDStr, 10, 32); err == nil {
			uid := uint(id)
			userID = &uid
		}
	}

	// サポートされている形式の確認
	validFormats := []string{"json", "csv", "xml"}
	isValidFormat := false
	for _, validFormat := range validFormats {
		if format == validFormat {
			isValidFormat = true
			break
		}
	}

	if !isValidFormat {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "サポートされていない形式です",
			"supported": validFormats,
		})
		return
	}

	data, err := ctrl.Service.ExportData(userID, format, startDate, endDate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "データエクスポート失敗",
			"detail": err.Error(),
		})
		return
	}

	// Content-Typeを形式に応じて設定
	switch format {
	case "csv":
		ctx.Header("Content-Type", "text/csv")
		ctx.Header("Content-Disposition", "attachment; filename=multimodal_data.csv")
	case "xml":
		ctx.Header("Content-Type", "application/xml")
		ctx.Header("Content-Disposition", "attachment; filename=multimodal_data.xml")
	default:
		ctx.Header("Content-Type", "application/json")
	}

	ctx.String(http.StatusOK, data)
}