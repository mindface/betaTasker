package heuristics

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// UserAction ユーザーアクション
type UserAction struct {
	Timestamp  int64                  `json:"timestamp"`
	ActionType string                 `json:"actionType"`
	ElementID  string                 `json:"elementId"`
	Context    map[string]interface{} `json:"context"`
	Duration   int64                  `json:"duration,omitempty"`
	Sequence   int                    `json:"sequence,omitempty"`
}

// Pattern 検出されたパターン
type Pattern struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Frequency  int     `json:"frequency"`
	Confidence float64 `json:"confidence"`
	Actions    []UserAction `json:"actions"`
	Heuristic  string  `json:"heuristic,omitempty"`
}

// HeuristicModel ヒューリスティクスモデル
type HeuristicModel struct {
	Patterns    []Pattern `json:"patterns"`
	Accuracy    float64   `json:"accuracy"`
	LastUpdated time.Time `json:"lastUpdated"`
}

// MLPipeline 機械学習パイプライン
type MLPipeline struct {
	mu              sync.RWMutex
	actionBuffer    []UserAction
	patterns        map[string]*Pattern
	modelVersion    int
	cycleInterval   time.Duration
	lastCycleTime   time.Time
	patternHistory  []HeuristicModel
}

// NewMLPipeline パイプラインの初期化
func NewMLPipeline() *MLPipeline {
	pipeline := &MLPipeline{
		actionBuffer:   make([]UserAction, 0, 1000),
		patterns:       make(map[string]*Pattern),
		cycleInterval:  5 * time.Second,
		patternHistory: make([]HeuristicModel, 0, 100),
	}
	
	// バックグラウンドでサイクル実行
	go pipeline.startCycle()
	
	return pipeline
}

// RecordAction アクションの記録
func (ml *MLPipeline) RecordAction(action UserAction) {
	ml.mu.Lock()
	defer ml.mu.Unlock()
	
	// バッファに追加（最大1000件）
	ml.actionBuffer = append(ml.actionBuffer, action)
	if len(ml.actionBuffer) > 1000 {
		ml.actionBuffer = ml.actionBuffer[1:]
	}
}

// startCycle サイクルの開始
func (ml *MLPipeline) startCycle() {
	ticker := time.NewTicker(ml.cycleInterval)
	defer ticker.Stop()
	
	for range ticker.C {
		ml.runCycle()
	}
}

// runCycle サイクル実行
func (ml *MLPipeline) runCycle() {
	ml.mu.Lock()
	defer ml.mu.Unlock()
	
	if len(ml.actionBuffer) < 10 {
		return
	}
	
	// パターン抽出
	patterns := ml.extractPatterns(ml.actionBuffer)
	
	// ヒューリスティクス推論
	ml.inferHeuristics(patterns)
	
	// モデル更新
	model := ml.updateModel(patterns)
	
	// 履歴に追加
	ml.patternHistory = append(ml.patternHistory, model)
	if len(ml.patternHistory) > 100 {
		ml.patternHistory = ml.patternHistory[1:]
	}

	ml.lastCycleTime = time.Now()
}

// extractPatterns パターン抽出
func (ml *MLPipeline) extractPatterns(actions []UserAction) []Pattern {
	patterns := []Pattern{}
	
	// N-gramパターン抽出（2-5gram）
	for n := 2; n <= 5 && n <= len(actions); n++ {
		ngrams := ml.extractNgrams(actions, n)
		patterns = append(patterns, ngrams...)
	}
	
	// 時間的パターン抽出
	temporalPatterns := ml.extractTemporalPatterns(actions)
	patterns = append(patterns, temporalPatterns...)
	
	// 頻度パターン抽出
	frequencyPatterns := ml.extractFrequencyPatterns(actions)
	patterns = append(patterns, frequencyPatterns...)
	
	return patterns
}

// extractNgrams N-gramパターン抽出
func (ml *MLPipeline) extractNgrams(actions []UserAction, n int) []Pattern {
	patterns := []Pattern{}
	sequences := make(map[string]int)
	
	for i := 0; i <= len(actions)-n; i++ {
		sequence := ""
		for j := 0; j < n; j++ {
			if j > 0 {
				sequence += "->"
			}
			sequence += actions[i+j].ActionType
		}
		sequences[sequence]++
	}
	
	// 閾値（3回以上）を超えるパターンのみ
	for seq, count := range sequences {
		if count >= 3 {
			patterns = append(patterns, Pattern{
				ID:         generatePatternID(),
				Name:       seq,
				Frequency:  count,
				Confidence: float64(count) / float64(len(actions)),
			})
		}
	}
	
	return patterns
}

// extractTemporalPatterns 時間的パターン抽出
func (ml *MLPipeline) extractTemporalPatterns(actions []UserAction) []Pattern {
	patterns := []Pattern{}
	
	// 時間的クラスタリング（2秒以内のアクション）
	clusters := [][]UserAction{}
	currentCluster := []UserAction{}
	
	for i, action := range actions {
		if i == 0 || action.Timestamp-actions[i-1].Timestamp <= 2000 {
			currentCluster = append(currentCluster, action)
		} else {
			if len(currentCluster) >= 2 {
				clusters = append(clusters, currentCluster)
			}
			currentCluster = []UserAction{action}
		}
	}
	
	if len(currentCluster) >= 2 {
		clusters = append(clusters, currentCluster)
	}
	
	// クラスターをパターンに変換
	for _, cluster := range clusters {
		pattern := Pattern{
			ID:         generatePatternID(),
			Name:       "temporal_cluster",
			Frequency:  len(cluster),
			Confidence: 0.7,
			Actions:    cluster,
			Heuristic:  "temporal_proximity",
		}
		patterns = append(patterns, pattern)
	}
	
	return patterns
}

// extractFrequencyPatterns 頻度パターン抽出
func (ml *MLPipeline) extractFrequencyPatterns(actions []UserAction) []Pattern {
	patterns := []Pattern{}
	actionCounts := make(map[string]int)
	
	// アクションタイプごとの頻度計算
	for _, action := range actions {
		actionCounts[action.ActionType]++
	}
	
	// 高頻度アクションをパターンとして記録
	for actionType, count := range actionCounts {
		if count >= len(actions)/4 { // 25%以上の頻度
			patterns = append(patterns, Pattern{
				ID:         generatePatternID(),
				Name:       "high_frequency_" + actionType,
				Frequency:  count,
				Confidence: float64(count) / float64(len(actions)),
				Heuristic:  "availability_heuristic",
			})
		}
	}
	
	return patterns
}

// inferHeuristics ヒューリスティクス推論
func (ml *MLPipeline) inferHeuristics(patterns []Pattern) {
	for i := range patterns {
		if patterns[i].Heuristic == "" {
			patterns[i].Heuristic = ml.inferHeuristicType(&patterns[i])
		}
	}
}

// inferHeuristicType ヒューリスティクスタイプ推論
func (ml *MLPipeline) inferHeuristicType(pattern *Pattern) string {
	// パターン名から推論
	if contains(pattern.Name, "->") {
		parts := splitSequence(pattern.Name)
		
		// 繰り返しパターン
		if allSame(parts) {
			return "confirmation_bias"
		}
		
		// 後退パターン
		if contains(pattern.Name, "back") || contains(pattern.Name, "cancel") {
			return "loss_aversion"
		}
		
		// 進行パターン
		if contains(pattern.Name, "next") || contains(pattern.Name, "continue") {
			return "progressive_disclosure"
		}
	}
	
	// 高頻度パターン
	if pattern.Frequency > 10 {
		return "habit_formation"
	}
	
	// 信頼度が高い
	if pattern.Confidence > 0.8 {
		return "anchoring_bias"
	}
	
	return "general_pattern"
}

// updateModel モデル更新
func (ml *MLPipeline) updateModel(patterns []Pattern) HeuristicModel {
	ml.modelVersion++
	
	// 既存パターンとマージ
	for _, p := range patterns {
		if existing, ok := ml.patterns[p.ID]; ok {
			// 信頼度を更新（移動平均）
			existing.Confidence = (existing.Confidence + p.Confidence) / 2
			existing.Frequency += p.Frequency
		} else {
			ml.patterns[p.ID] = &p
		}
	}
	
	// 古いパターンを削除
	ml.prunePatterns()
	
	// モデル作成
	allPatterns := []Pattern{}
	for _, p := range ml.patterns {
		allPatterns = append(allPatterns, *p)
	}
	
	// 信頼度でソート
	sort.Slice(allPatterns, func(i, j int) bool {
		return allPatterns[i].Confidence > allPatterns[j].Confidence
	})
	
	// 精度計算
	accuracy := ml.calculateAccuracy(allPatterns)
	
	return HeuristicModel{
		Patterns:    allPatterns,
		Accuracy:    accuracy,
		LastUpdated: time.Now(),
	}
}

// prunePatterns 古いパターンの削除
func (ml *MLPipeline) prunePatterns() {
	threshold := 0.1
	
	for id, pattern := range ml.patterns {
		if pattern.Confidence < threshold {
			delete(ml.patterns, id)
		}
	}
}

// calculateAccuracy 精度計算
func (ml *MLPipeline) calculateAccuracy(patterns []Pattern) float64 {
	if len(patterns) == 0 {
		return 0
	}
	
	totalConfidence := 0.0
	for _, p := range patterns {
		totalConfidence += p.Confidence
	}
	
	return totalConfidence / float64(len(patterns))
}

// GetCurrentModel 現在のモデル取得
func (ml *MLPipeline) GetCurrentModel() HeuristicModel {
	ml.mu.RLock()
	defer ml.mu.RUnlock()
	
	allPatterns := []Pattern{}
	for _, p := range ml.patterns {
		allPatterns = append(allPatterns, *p)
	}
	
	return HeuristicModel{
		Patterns:    allPatterns,
		Accuracy:    ml.calculateAccuracy(allPatterns),
		LastUpdated: ml.lastCycleTime,
	}
}

// GetMetrics メトリクス取得
func (ml *MLPipeline) GetMetrics() map[string]interface{} {
	ml.mu.RLock()
	defer ml.mu.RUnlock()
	
	return map[string]interface{}{
		"bufferSize":     len(ml.actionBuffer),
		"patternCount":   len(ml.patterns),
		"modelVersion":   ml.modelVersion,
		"lastCycleTime":  ml.lastCycleTime,
		"historyLength":  len(ml.patternHistory),
		"cycleInterval":  ml.cycleInterval.Seconds(),
	}
}

// PredictNext 次のアクション予測
func (ml *MLPipeline) PredictNext(recentActions []UserAction) []string {
	ml.mu.RLock()
	defer ml.mu.RUnlock()
	
	predictions := make(map[string]float64)
	
	// 最近のアクションからパターンマッチング
	for _, pattern := range ml.patterns {
		if matchesRecent(pattern.Name, recentActions) {
			nextAction := predictFromPattern(pattern.Name)
			if nextAction != "" {
				predictions[nextAction] += pattern.Confidence
			}
		}
	}
	
	// 予測をソート
	type prediction struct {
		action     string
		confidence float64
	}
	
	var sortedPredictions []prediction
	for action, conf := range predictions {
		sortedPredictions = append(sortedPredictions, prediction{action, conf})
	}
	
	sort.Slice(sortedPredictions, func(i, j int) bool {
		return sortedPredictions[i].confidence > sortedPredictions[j].confidence
	})
	
	// トップ3の予測を返す
	result := []string{}
	for i := 0; i < 3 && i < len(sortedPredictions); i++ {
		result = append(result, sortedPredictions[i].action)
	}
	
	return result
}

// HTTPハンドラー

// SyncHandler 同期エンドポイント
func (ml *MLPipeline) SyncHandler(c *gin.Context) {
	var req struct {
		Model     HeuristicModel `json:"model"`
		Version   int            `json:"version"`
		Timestamp int64          `json:"timestamp"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	// モデルをマージ
	ml.mu.Lock()
	for _, pattern := range req.Model.Patterns {
		if existing, ok := ml.patterns[pattern.ID]; ok {
			existing.Confidence = (existing.Confidence + pattern.Confidence) / 2
			existing.Frequency += pattern.Frequency
		} else {
			ml.patterns[pattern.ID] = &pattern
		}
	}
	ml.mu.Unlock()
	
	c.JSON(200, gin.H{
		"status":  "synced",
		"version": ml.modelVersion,
	})
}

// RecordHandler アクション記録エンドポイント
func (ml *MLPipeline) RecordHandler(c *gin.Context) {
	var action UserAction
	
	if err := c.ShouldBindJSON(&action); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	ml.RecordAction(action)
	
	c.JSON(200, gin.H{"status": "recorded"})
}

// ModelHandler モデル取得エンドポイント
func (ml *MLPipeline) ModelHandler(c *gin.Context) {
	model := ml.GetCurrentModel()
	c.JSON(200, model)
}

// MetricsHandler メトリクス取得エンドポイント
func (ml *MLPipeline) MetricsHandler(c *gin.Context) {
	metrics := ml.GetMetrics()
	c.JSON(200, metrics)
}

// PredictHandler 予測エンドポイント
func (ml *MLPipeline) PredictHandler(c *gin.Context) {
	var req struct {
		RecentActions []UserAction `json:"recentActions"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	predictions := ml.PredictNext(req.RecentActions)
	
	c.JSON(200, gin.H{
		"predictions": predictions,
	})
}

// ユーティリティ関数

func generatePatternID() string {
	return fmt.Sprintf("pattern-%d-%d", time.Now().UnixNano(), rand.Intn(1000))
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func splitSequence(seq string) []string {
	return strings.Split(seq, "->")
}

func allSame(items []string) bool {
	if len(items) == 0 {
		return true
	}
	first := items[0]
	for _, item := range items[1:] {
		if item != first {
			return false
		}
	}
	return true
}

func matchesRecent(pattern string, actions []UserAction) bool {
	parts := splitSequence(pattern)
	if len(parts) > len(actions) {
		return false
	}
	
	for i, part := range parts {
		if actions[len(actions)-len(parts)+i].ActionType != part {
			return false
		}
	}
	return true
}

func predictFromPattern(pattern string) string {
	parts := splitSequence(pattern)
	if len(parts) >= 2 {
		return parts[len(parts)-1]
	}
	return ""
}