package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/godotask/model"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StateEvaluationService struct {
	db *gorm.DB
}

func NewStateEvaluationService(db *gorm.DB) *StateEvaluationService {
	return &StateEvaluationService{db: db}
}

func (s *StateEvaluationService) EvaluateState(req *model.EvaluationRequest) (*model.StateEvaluation, error) {
	// Create new state evaluation
	evaluation := &model.StateEvaluation{
		ID:         uuid.New().String(),
		UserID:     req.UserID,
		TaskID:     req.TaskID,
		Level:      req.Level,
		WorkTarget: req.WorkTarget,
		Framework:  req.Framework,
		Status:     "evaluating",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// Convert state maps to JSON
	currentStateJSON, err := json.Marshal(req.CurrentState)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal current state: %v", err)
	}
	evaluation.CurrentState = datatypes.JSON(currentStateJSON)

	targetStateJSON, err := json.Marshal(req.TargetState)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal target state: %v", err)
	}
	evaluation.TargetState = datatypes.JSON(targetStateJSON)

	// Calculate evaluation score
	score, err := s.calculateEvaluationScore(req.CurrentState, req.TargetState, req.Level)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate evaluation score: %v", err)
	}
	evaluation.EvaluationScore = score

	// Get appropriate phenomenological framework
	framework, err := s.selectFramework(req.TaskID, req.Level, req.WorkTarget)
	if err != nil {
		return nil, fmt.Errorf("failed to select framework: %v", err)
	}
	evaluation.Framework = framework

	// Save to database
	if err := s.db.Create(evaluation).Error; err != nil {
		return nil, fmt.Errorf("failed to save evaluation: %v", err)
	}

	return evaluation, nil
}

func (s *StateEvaluationService) calculateEvaluationScore(currentState, targetState map[string]interface{}, level int) (float64, error) {
	var totalScore float64 = 0.0
	var totalWeight float64 = 0.0

	// Define weights based on level (L1-L5)
	levelWeights := map[string]float64{
		"accuracy":    []float64{0.4, 0.3, 0.2, 0.15, 0.1}[level-1],
		"efficiency":  []float64{0.2, 0.25, 0.3, 0.35, 0.4}[level-1],
		"consistency": []float64{0.3, 0.3, 0.3, 0.25, 0.2}[level-1],
		"innovation":  []float64{0.1, 0.15, 0.2, 0.25, 0.3}[level-1],
	}

	// Calculate scores for each dimension
	for dimension, weight := range levelWeights {
		currentVal, currentExists := currentState[dimension]
		targetVal, targetExists := targetState[dimension]

		if !currentExists || !targetExists {
			continue
		}

		// Convert to float64 for calculation
		current, ok1 := currentVal.(float64)
		target, ok2 := targetVal.(float64)

		if !ok1 || !ok2 {
			if str, ok := currentVal.(string); ok {
				if f, err := parseNumericValue(str); err == nil {
					current = f
					ok1 = true
				}
			}
			if str, ok := targetVal.(string); ok {
				if f, err := parseNumericValue(str); err == nil {
					target = f
					ok2 = true
				}
			}
		}

		if ok1 && ok2 && target != 0 {
			// Calculate achievement ratio
			ratio := current / target
			if ratio > 1.0 {
				ratio = 1.0 // Cap at 100%
			}
			
			dimensionScore := ratio * 100 // Convert to percentage
			totalScore += dimensionScore * weight
			totalWeight += weight
		}
	}

	if totalWeight == 0 {
		return 50.0, nil // Default score when no comparable dimensions
	}

	return totalScore / totalWeight, nil
}

func parseNumericValue(s string) (float64, error) {
	// Remove common units and parse
	s = strings.ReplaceAll(s, "%", "")
	s = strings.ReplaceAll(s, "mm", "")
	s = strings.ReplaceAll(s, "ms", "")
	s = strings.ReplaceAll(s, "N", "")
	s = strings.TrimSpace(s)
	
	var val float64
	_, err := fmt.Sscanf(s, "%f", &val)
	return val, err
}

func (s *StateEvaluationService) selectFramework(taskID, level int, workTarget string) (string, error) {
	// Get relevant phenomenological frameworks from database
	var frameworks []model.PhenomenologicalFramework
	
	// Determine domain from work target
	domain := s.extractDomainFromWorkTarget(workTarget)
	
	query := s.db.Where("domain = ? OR domain = 'general'", domain)
	
	// Filter by abstract level (L0-L3 mapping from L1-L5)
	abstractLevel := fmt.Sprintf("L%d", min(level-1, 3))
	query = query.Where("abstract_level <= ?", abstractLevel)
	
	if err := query.Find(&frameworks).Error; err != nil {
		return "default_framework", nil // Fallback to default
	}

	if len(frameworks) == 0 {
		return "default_framework", nil
	}

	// Select best matching framework based on work target similarity
	bestFramework := frameworks[0]
	bestScore := 0.0

	for _, framework := range frameworks {
		score := s.calculateFrameworkSimilarity(workTarget, framework.Description)
		if score > bestScore {
			bestScore = score
			bestFramework = framework
		}
	}

	return bestFramework.ID, nil
}

func (s *StateEvaluationService) extractDomainFromWorkTarget(workTarget string) string {
	workTarget = strings.ToLower(workTarget)
	
	domainKeywords := map[string][]string{
		"robot_control":    {"制御", "control", "力制御", "position"},
		"robot_assembly":   {"組立", "assembly", "取付", "組み立て"},
		"robot_welding":    {"溶接", "welding", "weld", "アーク"},
		"robot_vision":     {"画像", "vision", "認識", "カメラ"},
		"robot_motion":     {"軌道", "motion", "移動", "経路"},
		"machining":        {"切削", "加工", "旋盤", "フライス"},
		"robot_maintenance": {"保全", "maintenance", "保守", "点検"},
		"robot_safety":     {"安全", "safety", "協働", "collaborative"},
	}

	for domain, keywords := range domainKeywords {
		for _, keyword := range keywords {
			if strings.Contains(workTarget, keyword) {
				return domain
			}
		}
	}

	return "general"
}

func (s *StateEvaluationService) calculateFrameworkSimilarity(workTarget, description string) float64 {
	workWords := strings.Fields(strings.ToLower(workTarget))
	descWords := strings.Fields(strings.ToLower(description))

	matches := 0
	for _, workWord := range workWords {
		for _, descWord := range descWords {
			if strings.Contains(descWord, workWord) || strings.Contains(workWord, descWord) {
				matches++
				break
			}
		}
	}

	if len(workWords) == 0 {
		return 0.0
	}

	return float64(matches) / float64(len(workWords))
}

func (s *StateEvaluationService) GetEvaluationHistory(userID string, limit int) ([]model.StateEvaluation, error) {
	var evaluations []model.StateEvaluation

	query := s.db.Where("user_id = ?", userID).Order("created_at DESC")
	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&evaluations).Error; err != nil {
		return nil, fmt.Errorf("failed to get evaluation history: %v", err)
	}

	return evaluations, nil
}

func (s *StateEvaluationService) GetEvaluationByID(id string) (*model.StateEvaluation, error) {
	var evaluation model.StateEvaluation

	if err := s.db.Where("id = ?", id).First(&evaluation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("evaluation not found")
		}
		return nil, fmt.Errorf("failed to get evaluation: %v", err)
	}

	return &evaluation, nil
}

func (s *StateEvaluationService) UpdateEvaluationResults(id string, results map[string]interface{}, learnedKnowledge string) error {
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return fmt.Errorf("failed to marshal results: %v", err)
	}

	updates := map[string]interface{}{
		"results":          datatypes.JSON(resultsJSON),
		"learned_knowledge": learnedKnowledge,
		"status":           "completed",
		"updated_at":       time.Now(),
	}

	if err := s.db.Model(&model.StateEvaluation{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return fmt.Errorf("failed to update evaluation: %v", err)
	}

	return nil
}

func (s *StateEvaluationService) GetLevelProgression(userID string) (map[int]float64, error) {
	var results []struct {
		Level int
		AvgScore float64
	}

	err := s.db.Model(&model.StateEvaluation{}).
		Select("level, AVG(evaluation_score) as avg_score").
		Where("user_id = ? AND status = 'completed'", userID).
		Group("level").
		Order("level").
		Scan(&results).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get level progression: %v", err)
	}

	progression := make(map[int]float64)
	for _, result := range results {
		progression[result.Level] = result.AvgScore
	}

	return progression, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}