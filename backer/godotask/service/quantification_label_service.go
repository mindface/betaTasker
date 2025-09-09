package service

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/godotask/model"
)

type QuantificationLabelService struct {
	db *gorm.DB
}

func NewQuantificationLabelService(db *gorm.DB) *QuantificationLabelService {
	return &QuantificationLabelService{
		db: db,
	}
}

// GetAllLabels - 全ラベル取得
func (s *QuantificationLabelService) GetAllLabels() ([]model.QuantificationLabel, error) {
	var labels []model.QuantificationLabel
	
	err := s.db.Preload("Annotations").Preload("Revisions").
		Order("created_at DESC").Find(&labels).Error
	
	return labels, err
}

// GetLabelByID - ID指定でラベル取得
func (s *QuantificationLabelService) GetLabelByID(id string) (*model.QuantificationLabel, error) {
	var label model.QuantificationLabel
	
	err := s.db.Preload("Annotations").Preload("Revisions").
		Where("id = ?", id).First(&label).Error
	
	if err != nil {
		return nil, err
	}
	
	return &label, nil
}

// CreateLabel - ラベル作成
func (s *QuantificationLabelService) CreateLabel(label *model.QuantificationLabel) (*model.QuantificationLabel, error) {
	// 重複チェック
	var existing model.QuantificationLabel
	err := s.db.Where("original_text = ? AND domain = ?", 
		label.OriginalText, label.Domain).First(&existing).Error
	
	if err == nil {
		return nil, fmt.Errorf("同じテキストとドメインのラベルが既に存在します")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// データベースに保存
	if err := s.db.Create(label).Error; err != nil {
		return nil, err
	}

	// 作成履歴を追加
	revision := &model.LabelRevision{
		ID:        uuid.New().String(),
		LabelID:   label.ID,
		Version:   1,
		Changes:   model.JSON{"action": "created"},
		Comment:   "初回作成",
		UserID:    label.CreatedBy,
		Timestamp: time.Now(),
	}
	
	if err := s.db.Create(revision).Error; err != nil {
		// 履歴保存失敗はログに記録するが、処理は継続
		fmt.Printf("Failed to create revision: %v\n", err)
	}

	return label, nil
}

// UpdateLabel - ラベル更新
func (s *QuantificationLabelService) UpdateLabel(id string, updates map[string]interface{}, reason, userID string) (*model.QuantificationLabel, error) {
	// 既存ラベルを取得
	label, err := s.GetLabelByID(id)
	if err != nil {
		return nil, err
	}

	// 変更を記録
	changes := make([]map[string]interface{}, 0)
	
	// 更新を適用
	for field, newValue := range updates {
		var oldValue interface{}
		
		switch field {
		case "normalizedText":
			oldValue = label.NormalizedText
			label.NormalizedText = newValue.(string)
		case "imageDescription":
			oldValue = label.ImageDescription
			label.ImageDescription = newValue.(string)
		case "value":
			oldValue = label.Value
			label.Value = newValue.(float64)
		case "unit":
			oldValue = label.Unit
			label.Unit = newValue.(string)
		case "confidence":
			oldValue = label.Confidence
			label.Confidence = newValue.(float64)
		case "notes":
			oldValue = label.Notes
			label.Notes = newValue.(string)
		default:
			continue // 未対応のフィールドはスキップ
		}
		
		changes = append(changes, map[string]interface{}{
			"field":    field,
			"oldValue": oldValue,
			"newValue": newValue,
			"reason":   reason,
		})
	}

	// バージョンを更新
	label.Version++
	label.UpdatedBy = userID

	// データベースに保存
	if err := s.db.Save(label).Error; err != nil {
		return nil, err
	}

	// 改訂履歴を追加
	if len(changes) > 0 {
		changeMap := map[string]interface{}{
			"changes": changes,
			"version": label.Version,
		}
		if err := s.createRevision(id, changeMap, reason, userID); err != nil {
			// 履歴保存失敗はログに記録するが、処理は継続
			fmt.Printf("Failed to create revision: %v\n", err)
		}
	}

	return label, nil
}

// DeleteLabel - ラベル削除
func (s *QuantificationLabelService) DeleteLabel(id string) error {
	// 関連データも一緒に削除
	tx := s.db.Begin()
	
	// アノテーション削除
	if err := tx.Where("label_id = ?", id).Delete(&model.ImageAnnotation{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// 改訂履歴削除
	if err := tx.Where("label_id = ?", id).Delete(&model.LabelRevision{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	// ラベル削除
	if err := tx.Delete(&model.QuantificationLabel{}, "id = ?", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	return tx.Commit().Error
}

// SearchLabels - ラベル検索
func (s *QuantificationLabelService) SearchLabels(query model.LabelSearchQuery) ([]model.QuantificationLabel, int64, error) {
	var labels []model.QuantificationLabel
	var total int64

	// ベースクエリ
	db := s.db.Model(&model.QuantificationLabel{})

	// フィルタ条件を追加
	if query.Text != "" {
		db = db.Where("original_text ILIKE ? OR normalized_text ILIKE ?", 
			"%"+query.Text+"%", "%"+query.Text+"%")
	}

	if query.Domain != "" {
		db = db.Where("domain = ?", query.Domain)
	}

	if query.Category != "" {
		db = db.Where("category = ?", query.Category)
	}

	if query.Unit != "" {
		db = db.Where("unit = ?", query.Unit)
	}

	if query.MinValue > 0 || query.MaxValue > 0 {
		if query.MinValue > 0 && query.MaxValue > 0 {
			db = db.Where("value BETWEEN ? AND ?", query.MinValue, query.MaxValue)
		} else if query.MinValue > 0 {
			db = db.Where("value >= ?", query.MinValue)
		} else if query.MaxValue > 0 {
			db = db.Where("value <= ?", query.MaxValue)
		}
	}

	if query.MinConfidence > 0 {
		db = db.Where("confidence >= ?", query.MinConfidence)
	}

	if query.Verified != nil {
		db = db.Where("validated = ?", *query.Verified)
	}

	// 日付範囲フィルタ
	if query.From != "" {
		db = db.Where("created_at >= ?", query.From)
	}
	if query.To != "" {
		db = db.Where("created_at <= ?", query.To)
	}

	// 総数を取得
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// ソート
	orderBy := "created_at DESC" // デフォルト
	if query.SortBy != "" {
		direction := "ASC"
		if query.SortOrder == "desc" {
			direction = "DESC"
		}
		
		switch query.SortBy {
		case "confidence", "value", "created_at":
			orderBy = fmt.Sprintf("%s %s", query.SortBy, direction)
		case "relevance":
			if query.Text != "" {
				orderBy = fmt.Sprintf("CASE WHEN original_text ILIKE '%s%%' THEN 1 ELSE 2 END, confidence %s", 
					query.Text, direction)
			}
		}
	}

	// ページネーション
	db = db.Order(orderBy).Limit(query.Limit).Offset(query.Offset)

	// データ取得
	if err := db.Preload("Annotations").Find(&labels).Error; err != nil {
		return nil, 0, err
	}

	return labels, total, nil
}

// VerifyLabel - ラベル検証
func (s *QuantificationLabelService) VerifyLabel(id string, verification map[string]interface{}, verifierID string) (map[string]float64, error) {
	label, err := s.GetLabelByID(id)
	if err != nil {
		return nil, err
	}

	// 検証結果を評価に反映
	accurate, _ := verification["accurate"].(bool)
	consistent, _ := verification["consistency"].(bool)
	reproducible, _ := verification["reproducible"].(bool)
	usable, _ := verification["usable"].(bool)

	if accurate {
		label.Accuracy = updateRunningAverage(label.Accuracy, 1.0, label.VerificationCount+1)
	} else {
		label.Accuracy = updateRunningAverage(label.Accuracy, 0.0, label.VerificationCount+1)
	}

	if consistent {
		label.Consistency = updateRunningAverage(label.Consistency, 1.0, label.VerificationCount+1)
	} else {
		label.Consistency = updateRunningAverage(label.Consistency, 0.0, label.VerificationCount+1)
	}

	if reproducible {
		label.Reproducibility = updateRunningAverage(label.Reproducibility, 1.0, label.VerificationCount+1)
	} else {
		label.Reproducibility = updateRunningAverage(label.Reproducibility, 0.0, label.VerificationCount+1)
	}

	if usable {
		label.Usability = updateRunningAverage(label.Usability, 1.0, label.VerificationCount+1)
	} else {
		label.Usability = updateRunningAverage(label.Usability, 0.0, label.VerificationCount+1)
	}

	label.VerificationCount++
	now := time.Now()
	label.LastVerified = &now

	// 検証結果が良好な場合は検証済みとマーク
	if accurate && consistent && reproducible && usable {
		label.Validated = true
	}

	// データベースに保存
	if err := s.db.Save(label).Error; err != nil {
		return nil, err
	}

	return map[string]float64{
		"accuracy":       label.Accuracy,
		"consistency":    label.Consistency,
		"reproducibility": label.Reproducibility,
		"usability":      label.Usability,
	}, nil
}

// GetStatistics - 統計情報取得
func (s *QuantificationLabelService) GetStatistics() (*model.LabelStatistics, error) {
	var stats model.LabelStatistics
	var total int64

	// 総ラベル数
	s.db.Model(&model.QuantificationLabel{}).Count(&total)
	stats.TotalLabels = int(total)

	// ドメイン別分布
	var domainResults []struct {
		Domain string
		Count  int64
	}
	s.db.Model(&model.QuantificationLabel{}).
		Select("domain, count(*) as count").
		Group("domain").
		Find(&domainResults)

	stats.LabelsByDomain = make(map[string]int)
	for _, result := range domainResults {
		stats.LabelsByDomain[result.Domain] = int(result.Count)
	}

	// カテゴリ別分布
	var categoryResults []struct {
		Category string
		Count    int
	}
	s.db.Model(&model.QuantificationLabel{}).
		Select("category, count(*) as count").
		Group("category").
		Find(&categoryResults)

	stats.LabelsByCategory = make(map[string]int)
	for _, result := range categoryResults {
		stats.LabelsByCategory[result.Category] = int(result.Count)
	}

	// 平均メトリクス
	var avgResults struct {
		AvgConfidence      float64
		AvgAccuracy        float64
		AvgConsistency     float64
		AvgReproducibility float64
	}
	
	s.db.Model(&model.QuantificationLabel{}).
		Select("AVG(confidence) as avg_confidence, AVG(accuracy) as avg_accuracy, AVG(consistency) as avg_consistency, AVG(reproducibility) as avg_reproducibility").
		Scan(&avgResults)

	stats.AverageMetrics = map[string]float64{
		"confidence":      avgResults.AvgConfidence,
		"accuracy":        avgResults.AvgAccuracy,
		"consistency":     avgResults.AvgConsistency,
		"reproducibility": avgResults.AvgReproducibility,
	}

	// 品質分布
	var qualityResults struct {
		HighConfidence int64
		MediumConfidence int64
		LowConfidence int64
		Verified int64
		Unverified int64
	}
	
	s.db.Model(&model.QuantificationLabel{}).
		Select("SUM(CASE WHEN confidence >= 0.8 THEN 1 ELSE 0 END) as high_confidence, " +
			"SUM(CASE WHEN confidence >= 0.5 AND confidence < 0.8 THEN 1 ELSE 0 END) as medium_confidence, " +
			"SUM(CASE WHEN confidence < 0.5 THEN 1 ELSE 0 END) as low_confidence, " +
			"SUM(CASE WHEN validated = true THEN 1 ELSE 0 END) as verified, " +
			"SUM(CASE WHEN validated = false THEN 1 ELSE 0 END) as unverified").
		Scan(&qualityResults)

	stats.Quality = map[string]int{
		"high_confidence": int(qualityResults.HighConfidence),
		"medium_confidence": int(qualityResults.MediumConfidence),
		"low_confidence": int(qualityResults.LowConfidence),
		"verified": int(qualityResults.Verified),
		"unverified": int(qualityResults.Unverified),
	}

	return &stats, nil
}

// SuggestQuantification - 定量化提案
func (s *QuantificationLabelService) SuggestQuantification(text, imageURL, domain string) ([]map[string]interface{}, error) {
	var suggestions []map[string]interface{}

	// 1. 直接マッチング
	var directMatches []model.QuantificationLabel
	s.db.Where("original_text ILIKE ? OR normalized_text ILIKE ?", 
		"%"+text+"%", "%"+text+"%").
		Order("confidence DESC").
		Limit(3).
		Find(&directMatches)

	for _, match := range directMatches {
		suggestions = append(suggestions, map[string]interface{}{
			"value":      match.Value,
			"unit":       match.Unit,
			"confidence": match.Confidence * 0.9, // 直接マッチは少し信頼度を下げる
			"source":     "direct_match",
		})
	}

	// 2. ドメイン内類似検索
	if domain != "" {
		var domainMatches []model.QuantificationLabel
		s.db.Where("domain = ? AND (original_text ILIKE ? OR normalized_text ILIKE ?)", 
			domain, "%"+text+"%", "%"+text+"%").
			Order("confidence DESC").
			Limit(2).
			Find(&domainMatches)

		for _, match := range domainMatches {
			suggestions = append(suggestions, map[string]interface{}{
				"value":      match.Value,
				"unit":       match.Unit,
				"confidence": match.Confidence * 0.8,
				"source":     "domain_match",
			})
		}
	}

	// 3. デフォルト提案（基本的なパターンマッチング）
	defaultSuggestions := s.getDefaultSuggestions(text)
	suggestions = append(suggestions, defaultSuggestions...)

	// 重複を除去し、上位5件に制限
	uniqueSuggestions := removeDuplicateSuggestions(suggestions)
	if len(uniqueSuggestions) > 5 {
		uniqueSuggestions = uniqueSuggestions[:5]
	}

	return uniqueSuggestions, nil
}

// GetLabelHistory - ラベル履歴取得
func (s *QuantificationLabelService) GetLabelHistory(id string) ([]model.LabelRevision, error) {
	var revisions []model.LabelRevision
	
	err := s.db.Where("label_id = ?", id).
		Order("timestamp DESC").
		Find(&revisions).Error
	
	return revisions, err
}

// FindSimilarLabels - 類似ラベル検索
func (s *QuantificationLabelService) FindSimilarLabels(text, imageURL string, value float64, unit string, limit int) ([]model.QuantificationLabel, error) {
	var similar []model.QuantificationLabel
	
	db := s.db.Model(&model.QuantificationLabel{})
	
	if text != "" {
		db = db.Where("original_text ILIKE ? OR normalized_text ILIKE ?", 
			"%"+text+"%", "%"+text+"%")
	}
	
	if value > 0 && unit != "" {
		// 値の範囲で類似度を判定（±20%）
		minVal := value * 0.8
		maxVal := value * 1.2
		db = db.Where("unit = ? AND value BETWEEN ? AND ?", unit, minVal, maxVal)
	}
	
	err := db.Order("confidence DESC").Limit(limit).Find(&similar).Error
	return similar, err
}

// BulkOperation - バルク操作
func (s *QuantificationLabelService) BulkOperation(operation string, labelIDs []string, options map[string]interface{}, userID string) (map[string]interface{}, error) {
	result := map[string]interface{}{
		"processed": 0,
		"failed":    0,
		"details":   []string{},
	}

	switch operation {
	case "verify":
		for _, id := range labelIDs {
			verification := map[string]interface{}{
				"accurate":     true,
				"consistency":  true,
				"reproducible": true,
				"usable":       true,
			}
			
			_, err := s.VerifyLabel(id, verification, userID)
			if err != nil {
				result["failed"] = result["failed"].(int) + 1
				result["details"] = append(result["details"].([]string), 
					fmt.Sprintf("ID %s の検証に失敗: %v", id, err))
			} else {
				result["processed"] = result["processed"].(int) + 1
			}
		}

	case "delete":
		for _, id := range labelIDs {
			err := s.DeleteLabel(id)
			if err != nil {
				result["failed"] = result["failed"].(int) + 1
				result["details"] = append(result["details"].([]string), 
					fmt.Sprintf("ID %s の削除に失敗: %v", id, err))
			} else {
				result["processed"] = result["processed"].(int) + 1
			}
		}

	default:
		return nil, fmt.Errorf("未対応の操作: %s", operation)
	}

	return result, nil
}

// CreateRevision - リビジョン作成（内部用）
func (s *QuantificationLabelService) createRevision(labelID string, changes map[string]interface{}, reason, userID string) error {
	revision := model.LabelRevision{
		ID:        uuid.New().String(),
		LabelID:   labelID,
		Changes:   model.JSON(changes),
		Comment:   reason,
		UserID:    userID,
		Timestamp: time.Now(),
	}
	
	return s.db.Create(&revision).Error
}

// ExportLabels - ラベル一括エクスポート
func (s *QuantificationLabelService) ExportLabels(format string, query model.LabelSearchQuery) ([]byte, string, error) {
	labels, _, err := s.SearchLabels(query)
	if err != nil {
		return nil, "", err
	}
	
	switch format {
	case "json":
		data, err := json.MarshalIndent(labels, "", "  ")
		return data, "application/json", err
		
	case "csv":
		var buf bytes.Buffer
		writer := csv.NewWriter(&buf)
		
		// ヘッダー
		headers := []string{
			"ID", "OriginalText", "NormalizedText", "Category", "Domain",
			"Value", "Unit", "MinRange", "MaxRange", "Confidence", 
			"AbstractLevel", "Source", "Validated", "CreatedAt",
		}
		writer.Write(headers)

		// データ
		for _, label := range labels {
			record := []string{
				label.ID,
				label.OriginalText,
				label.NormalizedText,
				label.Category,
				label.Domain,
				fmt.Sprintf("%.2f", label.Value),
				label.Unit,
				fmt.Sprintf("%.2f", label.MinRange),
				fmt.Sprintf("%.2f", label.MaxRange),
				fmt.Sprintf("%.2f", label.Confidence),
				label.AbstractLevel,
				label.Source,
				fmt.Sprintf("%t", label.Validated),
				label.CreatedAt.Format("2006-01-02 15:04:05"),
			}
			writer.Write(record)
		}
		
		writer.Flush()
		return buf.Bytes(), "text/csv", writer.Error()
		
	default:
		return nil, "", fmt.Errorf("未対応の形式: %s", format)
	}
}

// GetUserStats - ユーザー統計取得
func (s *QuantificationLabelService) GetUserStats(userID string) (map[string]interface{}, error) {
	stats := map[string]interface{}{}
	
	// ユーザーのラベル作成数
	var totalLabels int64
	err := s.db.Model(&model.QuantificationLabel{}).
		Where("created_by = ?", userID).
		Count(&totalLabels).Error
	if err != nil {
		return nil, err
	}
	stats["total_labels"] = totalLabels
	
	// ユーザーの検証済みラベル数
	var verifiedLabels int64
	err = s.db.Model(&model.QuantificationLabel{}).
		Where("created_by = ? AND validated = ?", userID, true).
		Count(&verifiedLabels).Error
	if err != nil {
		return nil, err
	}
	stats["verified_labels"] = verifiedLabels
	
	// ユーザーの平均信頼度
	var avgConfidence float64
	err = s.db.Model(&model.QuantificationLabel{}).
		Where("created_by = ?", userID).
		Select("AVG(confidence)").
		Scan(&avgConfidence).Error
	if err != nil {
		return nil, err
	}
	stats["avg_confidence"] = avgConfidence
	
	// ドメイン別統計
	var domainStats []struct {
		Domain string
		Count  int64
	}
	err = s.db.Model(&model.QuantificationLabel{}).
		Where("created_by = ?", userID).
		Group("domain").
		Select("domain, count(*) as count").
		Scan(&domainStats).Error
	if err != nil {
		return nil, err
	}
	stats["domain_stats"] = domainStats
	
	// 最近のアクティビティ
	var recentLabels []model.QuantificationLabel
	err = s.db.Where("created_by = ?", userID).
		Order("created_at DESC").
		Limit(5).
		Find(&recentLabels).Error
	if err != nil {
		return nil, err
	}
	stats["recent_labels"] = recentLabels
	
	return stats, nil
}

// ヘルパー関数

func updateRunningAverage(currentAvg float64, newValue float64, count int) float64 {
	if count <= 1 {
		return newValue
	}
	return ((currentAvg * float64(count-1)) + newValue) / float64(count)
}

func (s *QuantificationLabelService) getDefaultSuggestions(text string) []map[string]interface{} {
	suggestions := []map[string]interface{}{}
	
	lowerText := strings.ToLower(text)
	
	// 基本的なパターンマッチング
	patterns := map[string]map[string]interface{}{
		"小さじ": {"value": 5.0, "unit": "ml", "confidence": 0.8},
		"大さじ": {"value": 15.0, "unit": "ml", "confidence": 0.8},
		"カップ": {"value": 200.0, "unit": "ml", "confidence": 0.7},
		"ひとつまみ": {"value": 0.5, "unit": "g", "confidence": 0.6},
		"少々": {"value": 0.2, "unit": "g", "confidence": 0.5},
	}
	
	for pattern, data := range patterns {
		if strings.Contains(lowerText, pattern) {
			suggestion := map[string]interface{}{
				"value":      data["value"],
				"unit":       data["unit"],
				"confidence": data["confidence"],
				"source":     "pattern_match",
			}
			suggestions = append(suggestions, suggestion)
		}
	}
	
	return suggestions
}

func removeDuplicateSuggestions(suggestions []map[string]interface{}) []map[string]interface{} {
	seen := make(map[string]bool)
	unique := []map[string]interface{}{}
	
	for _, s := range suggestions {
		key := fmt.Sprintf("%v_%v", s["value"], s["unit"])
		if !seen[key] {
			seen[key] = true
			unique = append(unique, s)
		}
	}

	return unique
}