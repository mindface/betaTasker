package heuristics

import (
	"math"
	"time"
)

// CognitivePattern 認知パターン
type CognitivePattern struct {
	Type        string    `json:"type"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Strength    float64   `json:"strength"`
	DetectedAt  time.Time `json:"detectedAt"`
}

// CognitiveAnalyzer 認知分析器
type CognitiveAnalyzer struct {
	patterns []CognitivePattern
}

// NewCognitiveAnalyzer 認知分析器の初期化
func NewCognitiveAnalyzer() *CognitiveAnalyzer {
	return &CognitiveAnalyzer{
		patterns: []CognitivePattern{},
	}
}

// AnalyzeActions アクションから認知パターンを分析
func (ca *CognitiveAnalyzer) AnalyzeActions(actions []UserAction) []CognitivePattern {
	patterns := []CognitivePattern{}
	
	// 確証バイアスの検出
	if bias := ca.detectConfirmationBias(actions); bias != nil {
		patterns = append(patterns, *bias)
	}
	
	// アンカリング効果の検出
	if anchor := ca.detectAnchoringEffect(actions); anchor != nil {
		patterns = append(patterns, *anchor)
	}
	
	// 損失回避の検出
	if aversion := ca.detectLossAversion(actions); aversion != nil {
		patterns = append(patterns, *aversion)
	}
	
	// 利用可能性ヒューリスティクスの検出
	if availability := ca.detectAvailabilityHeuristic(actions); availability != nil {
		patterns = append(patterns, *availability)
	}
	
	// 認知負荷の検出
	if load := ca.detectCognitiveOverload(actions); load != nil {
		patterns = append(patterns, *load)
	}
	
	return patterns
}

// detectConfirmationBias 確証バイアスの検出
func (ca *CognitiveAnalyzer) detectConfirmationBias(actions []UserAction) *CognitivePattern {
	// 同じアクションの繰り返し回数をカウント
	repetitions := make(map[string]int)
	for _, action := range actions {
		repetitions[action.ActionType]++
	}
	
	// 最も頻繁なアクション
	maxCount := 0
	var maxAction string
	for action, count := range repetitions {
		if count > maxCount {
			maxCount = count
			maxAction = action
		}
	}

	// 全体の70%以上が同じアクションなら確証バイアス
	if float64(maxCount)/float64(len(actions)) > 0.7 {
		return &CognitivePattern{
			Type:        "confirmation_bias",
			Name:        "確証バイアス",
			Description: "ユーザーは「" + maxAction + "」を繰り返し選択しています",
			Strength:    float64(maxCount) / float64(len(actions)),
			DetectedAt:  time.Now(),
		}
	}

	return nil
}

// detectAnchoringEffect アンカリング効果の検出
func (ca *CognitiveAnalyzer) detectAnchoringEffect(actions []UserAction) *CognitivePattern {
	if len(actions) < 5 {
		return nil
	}
	
	// 最初のアクションと後続のアクションを比較
	firstAction := actions[0].ActionType
	relatedCount := 0
	
	for i := 1; i < len(actions); i++ {
		// 最初のアクションに関連するアクションをカウント
		if ca.isRelatedAction(firstAction, actions[i].ActionType) {
			relatedCount++
		}
	}
	
	// 50%以上が最初のアクションに関連していればアンカリング
	if float64(relatedCount)/float64(len(actions)-1) > 0.5 {
		return &CognitivePattern{
			Type:        "anchoring_effect",
			Name:        "アンカリング効果",
			Description: "最初の選択「" + firstAction + "」が後続の決定に影響しています",
			Strength:    float64(relatedCount) / float64(len(actions)-1),
			DetectedAt:  time.Now(),
		}
	}
	
	return nil
}

// detectLossAversion 損失回避の検出
func (ca *CognitiveAnalyzer) detectLossAversion(actions []UserAction) *CognitivePattern {
	cancelCount := 0
	backCount := 0
	undoCount := 0
	
	for _, action := range actions {
		switch action.ActionType {
		case "cancel", "close", "exit":
			cancelCount++
		case "back", "previous", "return":
			backCount++
		case "undo", "revert", "restore":
			undoCount++
		}
	}
	
	totalAvoidance := cancelCount + backCount + undoCount
	
	// 全体の30%以上が回避行動なら損失回避
	if float64(totalAvoidance)/float64(len(actions)) > 0.3 {
		return &CognitivePattern{
			Type:        "loss_aversion",
			Name:        "損失回避",
			Description: "ユーザーは頻繁に操作を取り消したり戻したりしています",
			Strength:    float64(totalAvoidance) / float64(len(actions)),
			DetectedAt:  time.Now(),
		}
	}
	
	return nil
}

// detectAvailabilityHeuristic 利用可能性ヒューリスティクスの検出
func (ca *CognitiveAnalyzer) detectAvailabilityHeuristic(actions []UserAction) *CognitivePattern {
	if len(actions) < 10 {
		return nil
	}
	
	// 最近5つのアクションを取得
	recentActions := actions[len(actions)-5:]
	recentTypes := make(map[string]bool)
	for _, action := range recentActions {
		recentTypes[action.ActionType] = true
	}
	
	// 次の5つのアクションで最近のものがどれだけ使われるか
	if len(actions) > 10 {
		nextActions := actions[len(actions)-10 : len(actions)-5]
		matchCount := 0
		for _, action := range nextActions {
			if recentTypes[action.ActionType] {
				matchCount++
			}
		}

		// 60%以上が最近のアクションと一致
		if float64(matchCount)/float64(len(nextActions)) > 0.6 {
			return &CognitivePattern{
				Type:        "availability_heuristic",
				Name:        "利用可能性ヒューリスティクス",
				Description: "ユーザーは最近使用した機能を優先的に選択しています",
				Strength:    float64(matchCount) / float64(len(nextActions)),
				DetectedAt:  time.Now(),
			}
		}
	}
	
	return nil
}

// detectCognitiveOverload 認知負荷の検出
func (ca *CognitiveAnalyzer) detectCognitiveOverload(actions []UserAction) *CognitivePattern {
	if len(actions) < 2 {
		return nil
	}
	
	// アクション間の時間間隔を計算
	totalInterval := int64(0)
	shortIntervals := 0
	
	for i := 1; i < len(actions); i++ {
		interval := actions[i].Timestamp - actions[i-1].Timestamp
		totalInterval += interval
		
		// 500ms未満の短い間隔
		if interval < 500 {
			shortIntervals++
		}
	}
	
	avgInterval := totalInterval / int64(len(actions)-1)
	
	// 平均間隔が1秒未満、または短い間隔が50%以上
	if avgInterval < 1000 || float64(shortIntervals)/float64(len(actions)-1) > 0.5 {
		load := 1.0 - (float64(avgInterval) / 5000.0) // 5秒を基準に負荷を計算
		if load > 1.0 {
			load = 1.0
		}
		
		return &CognitivePattern{
			Type:        "cognitive_overload",
			Name:        "認知負荷過多",
			Description: "ユーザーは急いで操作しており、認知負荷が高い状態です",
			Strength:    load,
			DetectedAt:  time.Now(),
		}
	}
	
	return nil
}

// isRelatedAction アクションが関連しているか判定
func (ca *CognitiveAnalyzer) isRelatedAction(action1, action2 string) bool {
	// 簡単な関連性チェック（実際はより高度な判定が必要）
	relatedGroups := [][]string{
		{"click", "select", "choose", "pick"},
		{"scroll", "swipe", "pan", "drag"},
		{"type", "input", "write", "edit"},
		{"save", "submit", "confirm", "apply"},
		{"cancel", "close", "exit", "quit"},
		{"back", "previous", "return", "undo"},
	}
	
	for _, group := range relatedGroups {
		inGroup1 := false
		inGroup2 := false
		
		for _, action := range group {
			if action == action1 {
				inGroup1 = true
			}
			if action == action2 {
				inGroup2 = true
			}
		}
		
		if inGroup1 && inGroup2 {
			return true
		}
	}
	
	return action1 == action2
}

// CalculateCognitiveLoad 認知負荷の計算
func CalculateCognitiveLoad(actions []UserAction) float64 {
	if len(actions) == 0 {
		return 0
	}
	
	// アクションの多様性
	uniqueTypes := make(map[string]bool)
	for _, action := range actions {
		uniqueTypes[action.ActionType] = true
	}
	diversity := float64(len(uniqueTypes)) / float64(len(actions))
	
	// アクション頻度
	if len(actions) < 2 {
		return diversity
	}
	
	timeRange := actions[len(actions)-1].Timestamp - actions[0].Timestamp
	if timeRange == 0 {
		return diversity
	}
	
	frequency := float64(len(actions)) / (float64(timeRange) / 1000.0) // actions per second
	
	// 認知負荷スコア（0-1）
	load := (diversity*0.3 + math.Min(frequency/10.0, 1.0)*0.7)
	
	return math.Min(load, 1.0)
}

// OptimizeUI UI最適化の提案
func OptimizeUI(patterns []CognitivePattern) map[string]interface{} {
	optimizations := map[string]interface{}{
		"reduceOptions":    false,
		"emphasizeFirst":   false,
		"prioritizeRecent": false,
		"simplifyLayout":   false,
		"addConfirmation":  false,
	}
	
	for _, pattern := range patterns {
		switch pattern.Type {
		case "cognitive_overload":
			if pattern.Strength > 0.7 {
				optimizations["reduceOptions"] = true
				optimizations["simplifyLayout"] = true
			}
		case "confirmation_bias":
			if pattern.Strength > 0.8 {
				optimizations["emphasizeFirst"] = false // 偏りを防ぐ
			}
		case "availability_heuristic":
			if pattern.Strength > 0.6 {
				optimizations["prioritizeRecent"] = true
			}
		case "loss_aversion":
			if pattern.Strength > 0.5 {
				optimizations["addConfirmation"] = true
			}
		}
	}
	
	return optimizations
}