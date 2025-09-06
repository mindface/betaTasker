package service

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/godotask/model"
)

type MultimodalService struct {
	db *gorm.DB
}

func NewMultimodalService(db *gorm.DB) *MultimodalService {
	return &MultimodalService{
		db: db,
	}
}

// ProcessTextAndImage - テキストと画像のマルチモーダル処理
func (s *MultimodalService) ProcessTextAndImage(text, imageURL string, userID, taskID uint) (*model.MultimodalData, error) {
	// テキスト特徴抽出
	textFeatures, err := s.extractTextFeatures(text)
	if err != nil {
		return nil, fmt.Errorf("テキスト特徴抽出失敗: %w", err)
	}

	// 画像特徴抽出（画像URLが提供されている場合）
	var imageFeatures *ImageFeatures
	var imageConfidence float64 = 0.0
	
	if imageURL != "" {
		imageFeatures, err = s.extractImageFeatures(imageURL)
		if err != nil {
			// 画像処理に失敗してもテキストのみで処理を継続
			fmt.Printf("画像特徴抽出失敗: %v\n", err)
		} else {
			imageConfidence = imageFeatures.Confidence
		}
	}

	// マルチモーダル融合と定量化
	quantification, mapping, err := s.performQuantification(text, textFeatures, imageFeatures)
	if err != nil {
		return nil, fmt.Errorf("定量化処理失敗: %w", err)
	}

	// データベースに保存
	multimodalData := &model.MultimodalData{
		ID:     uuid.New().String(),
		UserID: userID,
		TaskID: taskID,
		
		// 言語特徴
		Text:           text,
		Tokens:         textFeatures.Tokens,
		SemanticVector: textFeatures.SemanticVector,
		AmbiguityScore: textFeatures.AmbiguityScore,
		
		// 画像特徴
		ImageURL:        imageURL,
		Objects:         convertImageObjects(imageFeatures),
		Measurements:    convertImageMeasurements(imageFeatures),
		ImageConfidence: imageConfidence,
		
		// 関連付け
		MappingType:        mapping.Type,
		CorrelationScore:   mapping.CorrelationScore,
		ContextRelevance:   mapping.ContextRelevance,
		HistoricalAccuracy: mapping.HistoricalAccuracy,
		
		// 定量化結果
		Value:      quantification.Value,
		Unit:       quantification.Unit,
		MinRange:   quantification.MinRange,
		MaxRange:   quantification.MaxRange,
		Confidence: quantification.Confidence,
		
		// メタデータ
		Verified: false,
	}

	if err := s.db.Create(multimodalData).Error; err != nil {
		return nil, fmt.Errorf("データベース保存失敗: %w", err)
	}

	return multimodalData, nil
}

// CalibrateUser - ユーザーキャリブレーション
func (s *MultimodalService) CalibrateUser(userID uint, referenceObject, imageURL string, imageFile multipart.File) (map[string]interface{}, error) {
	// 画像からサイズを測定（実際の実装では画像解析APIを使用）
	measurements, confidence := s.measureFromImage(imageFile, referenceObject)
	
	// キャリブレーションデータを保存
	calibration := &model.UserCalibration{
		ID:              uuid.New().String(),
		UserID:          userID,
		ReferenceObject: referenceObject,
		Measurements:    measurements,
		ImageURL:        imageURL,
		Confidence:      confidence,
	}

	if err := s.db.Create(calibration).Error; err != nil {
		return nil, fmt.Errorf("キャリブレーション保存失敗: %w", err)
	}

	return map[string]interface{}{
		"userId":          userID,
		"referenceObject": referenceObject,
		"measurements":    measurements,
		"imageUrl":        imageURL,
		"confidence":      confidence,
	}, nil
}

// VerifyResult - 結果検証
func (s *MultimodalService) VerifyResult(dataID, feedback, verifierID string) error {
	var data model.MultimodalData
	if err := s.db.Where("id = ?", dataID).First(&data).Error; err != nil {
		return fmt.Errorf("データが見つかりません: %w", err)
	}

	// フィードバックを記録
	data.UserFeedback = feedback
	data.Verified = true

	// フィードバックに基づいて信頼度を調整
	switch feedback {
	case "correct":
		data.Confidence = math.Min(data.Confidence*1.1, 1.0)
	case "too_high", "too_low":
		data.Confidence = data.Confidence * 0.9
	case "incorrect":
		data.Confidence = data.Confidence * 0.7
	}

	return s.db.Save(&data).Error
}

// GetProcessingHistory - 処理履歴取得
func (s *MultimodalService) GetProcessingHistory(userID *uint, limit, offset int) ([]model.MultimodalData, int64, error) {
	var history []model.MultimodalData
	var total int64

	query := s.db.Model(&model.MultimodalData{})
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}

	// 総数取得
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// データ取得
	if err := query.Order("created_at DESC").
		Limit(limit).Offset(offset).
		Find(&history).Error; err != nil {
		return nil, 0, err
	}

	return history, total, nil
}

// GetVisualMetaphors - 視覚的メタファー一覧取得
func (s *MultimodalService) GetVisualMetaphors() ([]model.VisualMetaphor, error) {
	var metaphors []model.VisualMetaphor
	err := s.db.Order("metaphor").Find(&metaphors).Error
	return metaphors, err
}

// CreateVisualMetaphor - 視覚的メタファー作成
func (s *MultimodalService) CreateVisualMetaphor(metaphor *model.VisualMetaphor) (*model.VisualMetaphor, error) {
	if err := s.db.Create(metaphor).Error; err != nil {
		return nil, err
	}
	return metaphor, nil
}

// GetUserCalibrations - ユーザーキャリブレーション取得
func (s *MultimodalService) GetUserCalibrations(userID uint) ([]model.UserCalibration, error) {
	var calibrations []model.UserCalibration
	err := s.db.Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&calibrations).Error
	return calibrations, err
}

// CompareImages - 画像比較
func (s *MultimodalService) CompareImages(imageA, imageB string) (float64, error) {
	// 実際の実装では画像比較アルゴリズムを使用
	// ここではダミーの類似度を返す
	
	// 簡単な類似度計算（実際の実装では画像解析APIを使用）
	similarity := 0.5 + rand.Float64()*0.4 // 0.5-0.9の範囲
	
	return similarity, nil
}

// AnalyzeImage - 画像解析
func (s *MultimodalService) AnalyzeImage(imageData []byte, filename string) (map[string]interface{}, error) {
	// 実際の実装では画像認識APIを使用
	// ここではダミーの解析結果を返す
	
	analysis := map[string]interface{}{
		"filename": filename,
		"size":     len(imageData),
		"objects": []map[string]interface{}{
			{
				"label":      "container",
				"confidence": 0.85,
				"boundingBox": map[string]float64{
					"x": 10, "y": 10, "width": 100, "height": 80,
				},
			},
		},
		"measurements": []map[string]interface{}{
			{
				"type":  "volume",
				"value": 200,
				"unit":  "ml",
			},
		},
		"colors": []string{"blue", "white"},
		"quality": "good",
	}
	
	return analysis, nil
}

// GenerateDescription - 画像から説明生成
func (s *MultimodalService) GenerateDescription(imageURL string, focusRegion interface{}) (string, error) {
	// 実際の実装では画像キャプション生成APIを使用
	// ここではダミーの説明を生成
	
	descriptions := []string{
		"透明なガラス容器に液体が入っています。",
		"白い陶器のカップに飲み物が注がれています。",
		"金属製の計量スプーンに粉末が盛られています。",
		"プラスチック製の容器に食材が保存されています。",
	}
	
	// ランダムに説明を選択
	description := descriptions[rand.Intn(len(descriptions))]
	
	if focusRegion != nil {
		description += " 指定された領域に焦点を当てた解析結果です。"
	}
	
	return description, nil
}

// GetQuantificationSuggestions - 定量化提案取得（マルチモーダル対応）
func (s *MultimodalService) GetQuantificationSuggestions(text, imageURL, domain, context string) ([]map[string]interface{}, error) {
	suggestions := []map[string]interface{}{}
	
	// 1. テキストベースの提案
	textSuggestions := s.getTextBasedSuggestions(text, domain)
	suggestions = append(suggestions, textSuggestions...)
	
	// 2. 画像ベースの提案
	if imageURL != "" {
		imageSuggestions, err := s.getImageBasedSuggestions(imageURL, context)
		if err == nil {
			suggestions = append(suggestions, imageSuggestions...)
		}
	}
	
	// 3. データベース内の類似データからの提案
	dbSuggestions, err := s.getDatabaseSuggestions(text, imageURL, domain)
	if err == nil {
		suggestions = append(suggestions, dbSuggestions...)
	}
	
	// 重複除去と上位5件に制限
	uniqueSuggestions := s.removeDuplicates(suggestions)
	if len(uniqueSuggestions) > 5 {
		uniqueSuggestions = uniqueSuggestions[:5]
	}
	
	return uniqueSuggestions, nil
}

// ExportData - データエクスポート
func (s *MultimodalService) ExportData(userID *uint, fromDate, toDate, format string) (string, error) {
	var data []model.MultimodalData
	
	query := s.db.Model(&model.MultimodalData{})
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	
	if fromDate != "" {
		query = query.Where("created_at >= ?", fromDate)
	}
	
	if toDate != "" {
		query = query.Where("created_at <= ?", toDate)
	}
	
	if err := query.Order("created_at DESC").Find(&data).Error; err != nil {
		return "", err
	}
	
	switch format {
	case "csv":
		return s.exportToCSV(data), nil
	case "xml":
		return s.exportToXML(data), nil
	default: // json
		return s.exportToJSON(data)
	}
}

// プライベートメソッド

// TextFeatures - テキスト特徴
type TextFeatures struct {
	Tokens         model.JSON `json:"tokens"`
	SemanticVector model.JSON `json:"semantic_vector"`
	AmbiguityScore float64    `json:"ambiguity_score"`
}

// ImageFeatures - 画像特徴
type ImageFeatures struct {
	Objects     []ImageObject     `json:"objects"`
	Measurements []ImageMeasurement `json:"measurements"`
	Confidence  float64           `json:"confidence"`
}

type ImageObject struct {
	Label       string             `json:"label"`
	Confidence  float64            `json:"confidence"`
	BoundingBox map[string]float64 `json:"bounding_box"`
}

type ImageMeasurement struct {
	Type  string  `json:"type"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// Quantification - 定量化結果
type Quantification struct {
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
	MinRange   float64 `json:"min_range"`
	MaxRange   float64 `json:"max_range"`
	Confidence float64 `json:"confidence"`
}

// MappingResult - マッピング結果
type MappingResult struct {
	Type               string  `json:"type"`
	CorrelationScore   float64 `json:"correlation_score"`
	ContextRelevance   float64 `json:"context_relevance"`
	HistoricalAccuracy float64 `json:"historical_accuracy"`
}

func (s *MultimodalService) extractTextFeatures(text string) (*TextFeatures, error) {
	// 実際の実装では自然言語処理APIを使用
	
	// トークン化（簡易版）
	tokens := strings.Fields(strings.ToLower(text))
	
	// 意味ベクトル生成（ダミー）
	semanticVector := make([]float64, 10)
	for i := range semanticVector {
		semanticVector[i] = rand.Float64()
	}
	
	// 曖昧さスコア計算（ダミー）
	ambiguityScore := 0.3
	if strings.Contains(text, "ぐらい") || strings.Contains(text, "程度") {
		ambiguityScore = 0.7
	}
	
	return &TextFeatures{
		Tokens:         model.JSON{"tokens": tokens},
		SemanticVector: model.JSON{"vector": semanticVector},
		AmbiguityScore: ambiguityScore,
	}, nil
}

func (s *MultimodalService) extractImageFeatures(imageURL string) (*ImageFeatures, error) {
	// 実際の実装では画像認識APIを使用
	
	return &ImageFeatures{
		Objects: []ImageObject{
			{
				Label:      "container",
				Confidence: 0.85,
				BoundingBox: map[string]float64{
					"x": 10, "y": 10, "width": 100, "height": 80,
				},
			},
		},
		Measurements: []ImageMeasurement{
			{
				Type:  "volume",
				Value: 200,
				Unit:  "ml",
			},
		},
		Confidence: 0.8,
	}, nil
}

func (s *MultimodalService) performQuantification(text string, textFeatures *TextFeatures, imageFeatures *ImageFeatures) (*Quantification, *MappingResult, error) {
	// デフォルト値
	quantification := &Quantification{
		Value:      0,
		Unit:       "",
		MinRange:   0,
		MaxRange:   0,
		Confidence: 0.5,
	}
	
	mapping := &MappingResult{
		Type:               "inferred",
		CorrelationScore:   0.6,
		ContextRelevance:   0.7,
		HistoricalAccuracy: 0.5,
	}
	
	// パターンマッチングによる定量化
	lowerText := strings.ToLower(text)
	
	patterns := map[string]*Quantification{
		"小さじ": {Value: 5.0, Unit: "ml", MinRange: 4.5, MaxRange: 5.5, Confidence: 0.9},
		"大さじ": {Value: 15.0, Unit: "ml", MinRange: 14.0, MaxRange: 16.0, Confidence: 0.9},
		"カップ": {Value: 200.0, Unit: "ml", MinRange: 180.0, MaxRange: 220.0, Confidence: 0.8},
	}
	
	for pattern, result := range patterns {
		if strings.Contains(lowerText, pattern) {
			quantification = result
			mapping.Type = "direct"
			mapping.CorrelationScore = 0.9
			break
		}
	}
	
	// 画像情報による調整
	if imageFeatures != nil {
		for _, measurement := range imageFeatures.Measurements {
			if measurement.Type == "volume" && quantification.Unit == "ml" {
				// 画像から得られた測定値で調整
				confidence := (quantification.Confidence + imageFeatures.Confidence) / 2
				quantification.Confidence = confidence
				mapping.CorrelationScore = confidence
			}
		}
	}
	
	return quantification, mapping, nil
}

func (s *MultimodalService) measureFromImage(imageFile multipart.File, referenceObject string) (model.JSON, float64) {
	// 実際の実装では画像解析による測定
	
	measurements := map[string]float64{}
	confidence := 0.7
	
	switch referenceObject {
	case "hand":
		measurements["width"] = 10.0
		measurements["height"] = 18.0
		confidence = 0.8
	case "finger":
		measurements["width"] = 2.0
		measurements["height"] = 9.0
		confidence = 0.6
	default:
		measurements["width"] = 5.0
		measurements["height"] = 5.0
		confidence = 0.5
	}
	
	return model.JSON(measurements), confidence
}

func convertImageObjects(imageFeatures *ImageFeatures) model.JSON {
	if imageFeatures == nil {
		return nil
	}
	
	objects := make([]map[string]interface{}, len(imageFeatures.Objects))
	for i, obj := range imageFeatures.Objects {
		objects[i] = map[string]interface{}{
			"label":       obj.Label,
			"confidence":  obj.Confidence,
			"bounding_box": obj.BoundingBox,
		}
	}
	
	return model.JSON{"objects": objects}
}

func convertImageMeasurements(imageFeatures *ImageFeatures) model.JSON {
	if imageFeatures == nil {
		return nil
	}
	
	measurements := make([]map[string]interface{}, len(imageFeatures.Measurements))
	for i, measurement := range imageFeatures.Measurements {
		measurements[i] = map[string]interface{}{
			"type":  measurement.Type,
			"value": measurement.Value,
			"unit":  measurement.Unit,
		}
	}
	
	return model.JSON{"measurements": measurements}
}

func (s *MultimodalService) getTextBasedSuggestions(text, domain string) []map[string]interface{} {
	suggestions := []map[string]interface{}{}
	
	// パターンマッチング
	patterns := map[string]map[string]interface{}{
		"小さじ": {"value": 5.0, "unit": "ml", "confidence": 0.9},
		"大さじ": {"value": 15.0, "unit": "ml", "confidence": 0.9},
		"ひとつまみ": {"value": 0.5, "unit": "g", "confidence": 0.6},
	}
	
	lowerText := strings.ToLower(text)
	for pattern, data := range patterns {
		if strings.Contains(lowerText, pattern) {
			suggestion := map[string]interface{}{
				"value":      data["value"],
				"unit":       data["unit"],
				"confidence": data["confidence"],
				"source":     "text_pattern",
			}
			suggestions = append(suggestions, suggestion)
		}
	}
	
	return suggestions
}

func (s *MultimodalService) getImageBasedSuggestions(imageURL, context string) ([]map[string]interface{}, error) {
	// 実際の実装では画像解析APIを使用
	
	suggestions := []map[string]interface{}{
		{
			"value":      150.0,
			"unit":       "ml",
			"confidence": 0.7,
			"source":     "image_analysis",
		},
	}
	
	return suggestions, nil
}

func (s *MultimodalService) getDatabaseSuggestions(text, imageURL, domain string) ([]map[string]interface{}, error) {
	var data []model.MultimodalData
	
	query := s.db.Where("text ILIKE ?", "%"+text+"%")
	if domain != "" {
		// ドメイン情報があれば追加でフィルタ（実際の実装では関連テーブルから取得）
	}
	
	if err := query.Order("confidence DESC").Limit(3).Find(&data).Error; err != nil {
		return nil, err
	}
	
	suggestions := []map[string]interface{}{}
	for _, d := range data {
		suggestion := map[string]interface{}{
			"value":      d.Value,
			"unit":       d.Unit,
			"confidence": d.Confidence * 0.8, // 過去データは信頼度を少し下げる
			"source":     "database_match",
		}
		suggestions = append(suggestions, suggestion)
	}
	
	return suggestions, nil
}

func (s *MultimodalService) removeDuplicates(suggestions []map[string]interface{}) []map[string]interface{} {
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

func (s *MultimodalService) exportToJSON(data []model.MultimodalData) (string, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func (s *MultimodalService) exportToCSV(data []model.MultimodalData) string {
	csv := "ID,UserID,TaskID,Text,Value,Unit,Confidence,CreatedAt\n"
	
	for _, d := range data {
		csv += fmt.Sprintf("%s,%d,%d,%s,%.2f,%s,%.2f,%s\n",
			d.ID, d.UserID, d.TaskID, d.Text, d.Value, d.Unit, d.Confidence, d.CreatedAt.Format("2006-01-02"))
	}
	
	return csv
}

func (s *MultimodalService) exportToXML(data []model.MultimodalData) string {
	xml := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<multimodal_data>\n"
	
	for _, d := range data {
		xml += fmt.Sprintf("  <data id=\"%s\" user_id=\"%d\" task_id=\"%d\">\n", d.ID, d.UserID, d.TaskID)
		xml += fmt.Sprintf("    <text>%s</text>\n", d.Text)
		xml += fmt.Sprintf("    <value>%.2f</value>\n", d.Value)
		xml += fmt.Sprintf("    <unit>%s</unit>\n", d.Unit)
		xml += fmt.Sprintf("    <confidence>%.2f</confidence>\n", d.Confidence)
		xml += fmt.Sprintf("    <created_at>%s</created_at>\n", d.CreatedAt.Format("2006-01-02"))
		xml += "  </data>\n"
	}
	
	xml += "</multimodal_data>"
	return xml
}