package quantification_label

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/godotask/model"
)

// CreateLabel - ラベル作成
func (ctrl *QuantificationLabelController) CreateLabel(ctx *gin.Context) {
	var req model.CreateLabelRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "不正なリクエスト",
			"detail": err.Error(),
		})
		return
	}

	// ユーザーIDを取得（認証情報から）
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "認証が必要です",
		})
		return
	}

	label := &model.QuantificationLabel{
		ID:              uuid.New().String(),
		OriginalText:    req.Text,
		NormalizedText:  strings.ToLower(strings.TrimSpace(req.Text)),
		Category:        req.Category,
		Context:         "",
		Domain:          req.Domain,
		ImageURL:        req.ImageURL,
		ImageDescription: req.Description,
		Value:           req.Value,
		Unit:            req.Unit,
		MinRange:        req.Value * 0.9, // デフォルト範囲
		MaxRange:        req.Value * 1.1,
		TypicalValue:    req.Value,
		Precision:       2,
		Confidence:      0.7, // デフォルト信頼度
		AbstractLevel:   "concrete",
		RelatedConcepts: model.JSON(map[string]interface{}{
			"concepts": req.Concepts,
		}),
		SemanticTags: model.JSON(map[string]interface{}{
			"tags": req.Tags,
		}),
		Accuracy:       0.0,
		Consistency:    0.0,
		Reproducibility: 0.0,
		Usability:      0.0,
		Source:         "manual",
		Validated:      false,
		Version:        1,
		CreatedBy:      fmt.Sprintf("%v", userID),
		UpdatedBy:      fmt.Sprintf("%v", userID),
	}

	createdLabel, err := ctrl.Service.CreateLabel(label)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "ラベル作成失敗",
			"detail": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdLabel)
}