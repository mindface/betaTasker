package service

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/godotask/model"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ToolMatchingService struct {
	db *gorm.DB
}

func NewToolMatchingService(db *gorm.DB) *ToolMatchingService {
	return &ToolMatchingService{db: db}
}

func (s *ToolMatchingService) FindOptimalTools(req *model.ToolMatchingRequest) (*model.ToolMatchingResult, error) {
	// Get state evaluation context
	var stateEval model.StateEvaluation
	if err := s.db.Where("id = ?", req.StateEvaluationID).First(&stateEval).Error; err != nil {
		return nil, fmt.Errorf("state evaluation not found: %v", err)
	}

	// Get available robots and optimization models
	robots, err := s.getAvailableRobots()
	if err != nil {
		return nil, fmt.Errorf("failed to get robots: %v", err)
	}

	models, err := s.getOptimizationModels()
	if err != nil {
		return nil, fmt.Errorf("failed to get optimization models: %v", err)
	}

	// Calculate matching scores
	bestRobot, robotScore := s.findBestRobot(robots, &stateEval, req.Requirements, req.Constraints)
	bestModel, modelScore := s.findBestOptimizationModel(models, &stateEval, req.Requirements)

	// Generate recommendations
	recommendations := s.generateRecommendations(bestRobot, bestModel, &stateEval, req.Requirements)

	// Calculate combined matching score
	combinedScore := (robotScore + modelScore) / 2.0

	// Create tool matching result
	result := &model.ToolMatchingResult{
		ID:                uuid.New().String(),
		StateEvaluationID: req.StateEvaluationID,
		RobotID:          bestRobot.ID,
		OptimizationModelID: bestModel.ID,
		MatchingScore:    combinedScore,
		CreatedAt:        time.Now(),
	}

	// Marshal recommendations and parameters
	recommendationsJSON, err := json.Marshal(recommendations)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal recommendations: %v", err)
	}
	result.Recommendations = datatypes.JSON(recommendationsJSON)

	parameters := s.generateOptimalParameters(bestRobot, bestModel, &stateEval, req.Requirements)
	parametersJSON, err := json.Marshal(parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal parameters: %v", err)
	}
	result.Parameters = datatypes.JSON(parametersJSON)

	expectedPerformance := s.predictPerformance(bestRobot, bestModel, parameters, &stateEval)
	expectedPerformanceJSON, err := json.Marshal(expectedPerformance)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal expected performance: %v", err)
	}
	result.ExpectedPerformance = datatypes.JSON(expectedPerformanceJSON)

	// Save to database
	if err := s.db.Create(result).Error; err != nil {
		return nil, fmt.Errorf("failed to save tool matching result: %v", err)
	}

	return result, nil
}

func (s *ToolMatchingService) getAvailableRobots() ([]model.RobotSpecification, error) {
	var robots []model.RobotSpecification
	if err := s.db.Find(&robots).Error; err != nil {
		return nil, err
	}
	return robots, nil
}

func (s *ToolMatchingService) getOptimizationModels() ([]model.OptimizationModel, error) {
	var models []model.OptimizationModel
	if err := s.db.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

func (s *ToolMatchingService) findBestRobot(robots []model.RobotSpecification, stateEval *model.StateEvaluation, requirements, constraints map[string]interface{}) (*model.RobotSpecification, float64) {
	if len(robots) == 0 {
		return nil, 0.0
	}

	bestRobot := &robots[0]
	bestScore := 0.0

	for i := range robots {
		score := s.calculateRobotScore(&robots[i], stateEval, requirements, constraints)
		if score > bestScore {
			bestScore = score
			bestRobot = &robots[i]
		}
	}

	return bestRobot, bestScore
}

func (s *ToolMatchingService) calculateRobotScore(robot *model.RobotSpecification, stateEval *model.StateEvaluation, requirements, constraints map[string]interface{}) float64 {
	var totalScore float64 = 0.0
	var weights float64 = 0.0

	// Payload matching
	if reqPayload, exists := requirements["payload"]; exists {
		if payload, ok := reqPayload.(float64); ok {
			if robot.PayloadKg >= payload {
				payloadScore := math.Min(100.0, (robot.PayloadKg/payload)*50.0)
				totalScore += payloadScore * 0.25
			} else {
				totalScore += 0.0 // Cannot meet requirement
			}
			weights += 0.25
		}
	}

	// Reach matching
	if reqReach, exists := requirements["reach"]; exists {
		if reach, ok := reqReach.(float64); ok {
			if robot.ReachMm >= reach {
				reachScore := math.Min(100.0, (robot.ReachMm/reach)*50.0)
				totalScore += reachScore * 0.20
			}
			weights += 0.20
		}
	}

	// Precision matching
	if reqPrecision, exists := requirements["precision"]; exists {
		if precision, ok := reqPrecision.(float64); ok {
			// Lower repeat accuracy is better
			if robot.RepeatAccuracyMm <= precision {
				precisionScore := (precision / robot.RepeatAccuracyMm) * 50.0
				totalScore += math.Min(100.0, precisionScore) * 0.30
			}
			weights += 0.30
		}
	}

	// Speed matching
	if reqSpeed, exists := requirements["speed"]; exists {
		if speed, ok := reqSpeed.(float64); ok {
			if robot.MaxSpeedMmS >= speed {
				speedScore := math.Min(100.0, (robot.MaxSpeedMmS/speed)*50.0)
				totalScore += speedScore * 0.15
			}
			weights += 0.15
		}
	}

	// AI capability bonus
	if stateEval.Level >= 3 && robot.AICapability.Valid {
		aiCapabilities := strings.Split(robot.AICapability.String, "|")
		aiScore := float64(len(aiCapabilities)) * 10.0
		totalScore += math.Min(20.0, aiScore) * 0.10
		weights += 0.10
	}

	if weights == 0 {
		return 50.0 // Default score
	}

	return totalScore / weights
}

func (s *ToolMatchingService) findBestOptimizationModel(models []model.OptimizationModel, stateEval *model.StateEvaluation, requirements map[string]interface{}) (*model.OptimizationModel, float64) {
	if len(models) == 0 {
		return nil, 0.0
	}

	bestModel := &models[0]
	bestScore := 0.0

	for i := range models {
		score := s.calculateModelScore(&models[i], stateEval, requirements)
		if score > bestScore {
			bestScore = score
			bestModel = &models[i]
		}
	}

	return bestModel, bestScore
}

func (s *ToolMatchingService) calculateModelScore(model *model.OptimizationModel, stateEval *model.StateEvaluation, requirements map[string]interface{}) float64 {
	var totalScore float64 = 0.0

	// Domain matching
	workTargetDomain := extractDomainFromWorkTarget(stateEval.WorkTarget)
	if strings.Contains(strings.ToLower(model.Domain), workTargetDomain) {
		totalScore += 40.0
	}

	// Performance metrics matching
	if model.PerformanceMetric.Valid {
		metrics := strings.Split(model.PerformanceMetric.String, "|")
		for _, metric := range metrics {
			if strings.Contains(strings.ToLower(stateEval.WorkTarget), extractMetricKeyword(metric)) {
				totalScore += 15.0
			}
		}
	}

	// Convergence rate bonus
	if model.ConvergenceRate.Valid && model.ConvergenceRate.Float64 > 0.9 {
		totalScore += 20.0
	}

	// Iteration count efficiency
	if model.IterationCount.Valid {
		iterationEfficiency := math.Min(25.0, 10000.0/model.IterationCount.Float64*25.0)
		totalScore += iterationEfficiency
	}

	return math.Min(100.0, totalScore)
}

func extractDomainFromWorkTarget(workTarget string) string {
	workTarget = strings.ToLower(workTarget)
	
	if strings.Contains(workTarget, "切削") || strings.Contains(workTarget, "加工") {
		return "machining"
	}
	if strings.Contains(workTarget, "組立") || strings.Contains(workTarget, "assembly") {
		return "robot_assembly"
	}
	if strings.Contains(workTarget, "溶接") || strings.Contains(workTarget, "welding") {
		return "robot_welding"
	}
	if strings.Contains(workTarget, "画像") || strings.Contains(workTarget, "vision") {
		return "robot_vision"
	}
	
	return "robot_motion"
}

func extractMetricKeyword(metric string) string {
	metric = strings.ToLower(metric)
	if strings.Contains(metric, "accuracy") || strings.Contains(metric, "精度") {
		return "accuracy"
	}
	if strings.Contains(metric, "time") || strings.Contains(metric, "時間") {
		return "time"
	}
	if strings.Contains(metric, "energy") || strings.Contains(metric, "エネルギー") {
		return "energy"
	}
	if strings.Contains(metric, "force") || strings.Contains(metric, "力") {
		return "force"
	}
	return metric
}

func (s *ToolMatchingService) generateRecommendations(robot *model.RobotSpecification, optModel *model.OptimizationModel, stateEval *model.StateEvaluation, requirements map[string]interface{}) map[string]interface{} {
	recommendations := make(map[string]interface{})

	// Robot recommendations
	robotRec := map[string]interface{}{
		"model":              robot.ModelName,
		"dof":               robot.DOF,
		"payload_capacity":  robot.PayloadKg,
		"reach":             robot.ReachMm,
		"precision":         robot.RepeatAccuracyMm,
		"max_speed":         robot.MaxSpeedMmS,
		"recommended_use":   s.generateRobotUseCase(robot, stateEval),
	}

	// Add safety recommendations based on level
	if stateEval.Level <= 2 {
		robotRec["safety_level"] = "High supervision required"
		robotRec["training_time"] = "2-4 weeks"
	} else {
		robotRec["safety_level"] = "Standard safety protocols"
		robotRec["training_time"] = "1-2 weeks"
	}

	recommendations["robot"] = robotRec

	// Optimization model recommendations
	modelRec := map[string]interface{}{
		"model_name":        optModel.Name,
		"type":             optModel.Type,
		"objective":        optModel.ObjectiveFunction,
		"expected_improvement": s.estimateImprovement(optModel, stateEval),
		"implementation_complexity": s.assessComplexity(optModel, stateEval.Level),
	}

	recommendations["optimization"] = modelRec

	// Process recommendations
	processRec := s.generateProcessRecommendations(robot, optModel, stateEval, requirements)
	recommendations["process"] = processRec

	return recommendations
}

func (s *ToolMatchingService) generateRobotUseCase(robot *model.RobotSpecification, stateEval *model.StateEvaluation) string {
	useCases := []string{}

	if strings.Contains(stateEval.WorkTarget, "精度") || strings.Contains(stateEval.WorkTarget, "precision") {
		if robot.RepeatAccuracyMm <= 0.02 {
			useCases = append(useCases, "高精度作業に最適")
		}
	}

	if strings.Contains(stateEval.WorkTarget, "速度") || strings.Contains(stateEval.WorkTarget, "speed") {
		if robot.MaxSpeedMmS >= 2000 {
			useCases = append(useCases, "高速作業対応")
		}
	}

	if robot.AICapability.Valid && strings.Contains(robot.AICapability.String, "learning") {
		useCases = append(useCases, "AI学習機能搭載")
	}

	if len(useCases) == 0 {
		return "汎用作業向け"
	}

	return strings.Join(useCases, ", ")
}

func (s *ToolMatchingService) estimateImprovement(model *model.OptimizationModel, stateEval *model.StateEvaluation) string {
	if model.PerformanceMetric.Valid {
		metrics := strings.Split(model.PerformanceMetric.String, "|")
		for _, metric := range metrics {
			if improvement := extractImprovementValue(metric); improvement != "" {
				return improvement
			}
		}
	}

	// Default estimate based on current evaluation score
	if stateEval.EvaluationScore < 50 {
		return "30-50%の改善が期待できます"
	} else if stateEval.EvaluationScore < 80 {
		return "15-30%の改善が期待できます"
	} else {
		return "5-15%の改善が期待できます"
	}
}

func extractImprovementValue(metric string) string {
	// Extract percentage improvements from metric strings
	if strings.Contains(metric, "reduction:") {
		parts := strings.Split(metric, "reduction:")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1]) + "の削減"
		}
	}
	if strings.Contains(metric, "improvement:") {
		parts := strings.Split(metric, "improvement:")
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1]) + "の改善"
		}
	}
	return ""
}

func (s *ToolMatchingService) assessComplexity(model *model.OptimizationModel, level int) string {
	complexityLevels := map[string]int{
		"control_theory":    3,
		"ml_based":         4,
		"hybrid":           3,
		"heuristic":        2,
		"genetic_algorithm": 4,
		"constraint_programming": 4,
		"statistical_optimization": 3,
		"geometric_optimization": 2,
		"kalman_filter":    4,
		"queueing_theory":  3,
	}

	modelComplexity, exists := complexityLevels[model.Type]
	if !exists {
		modelComplexity = 3 // Default
	}

	if level >= modelComplexity {
		return "適切な複雑度"
	} else if level < modelComplexity-1 {
		return "高度な知識が必要"
	} else {
		return "段階的な学習が推奨"
	}
}

func (s *ToolMatchingService) generateProcessRecommendations(robot *model.RobotSpecification, model *model.OptimizationModel, stateEval *model.StateEvaluation, requirements map[string]interface{}) map[string]interface{} {
	process := make(map[string]interface{})

	// Setup recommendations
	setup := []string{}
	if robot.VisionSystem.Valid && robot.VisionSystem.String != "none" {
		setup = append(setup, "視覚システムのキャリブレーション")
	}
	if robot.ForceSensor.Valid && robot.ForceSensor.String != "none" {
		setup = append(setup, "力センサの初期設定")
	}
	if robot.AICapability.Valid {
		setup = append(setup, "AI学習データの準備")
	}
	process["setup_steps"] = setup

	// Monitoring parameters
	monitoring := map[string]interface{}{
		"key_metrics": s.extractKeyMetrics(model),
		"update_frequency": s.determineUpdateFrequency(model),
		"alert_thresholds": s.generateAlertThresholds(stateEval),
	}
	process["monitoring"] = monitoring

	// Success criteria
	process["success_criteria"] = s.generateSuccessCriteria(stateEval, requirements)

	return process
}

func (s *ToolMatchingService) extractKeyMetrics(model *model.OptimizationModel) []string {
	metrics := []string{}
	
	if model.PerformanceMetric.Valid {
		metricPairs := strings.Split(model.PerformanceMetric.String, "|")
		for _, pair := range metricPairs {
			if strings.Contains(pair, ":") {
				metric := strings.Split(pair, ":")[0]
				metrics = append(metrics, metric)
			}
		}
	}

	// Default metrics if none specified
	if len(metrics) == 0 {
		metrics = []string{"accuracy", "processing_time", "error_rate"}
	}

	return metrics
}

func (s *ToolMatchingService) determineUpdateFrequency(model *model.OptimizationModel) string {
	if model.IterationCount.Valid {
		if model.IterationCount.Float64 > 5000 {
			return "real_time"
		} else if model.IterationCount.Float64 > 1000 {
			return "every_10_seconds"
		} else {
			return "every_minute"
		}
	}
	return "every_30_seconds"
}

func (s *ToolMatchingService) generateAlertThresholds(stateEval *model.StateEvaluation) map[string]float64 {
	baseThreshold := stateEval.EvaluationScore * 0.8 // Alert if performance drops below 80% of current

	return map[string]float64{
		"performance_degradation": baseThreshold,
		"error_rate_increase":     0.05, // 5% error rate threshold
		"response_time_increase":  2.0,  // 2x response time threshold
	}
}

func (s *ToolMatchingService) generateSuccessCriteria(stateEval *model.StateEvaluation, requirements map[string]interface{}) map[string]interface{} {
	criteria := make(map[string]interface{})

	// Base improvement target
	targetImprovement := math.Min(95.0, stateEval.EvaluationScore*1.2) // 20% improvement, capped at 95%
	criteria["target_score"] = targetImprovement

	// Specific requirements
	if accuracy, exists := requirements["target_accuracy"]; exists {
		criteria["accuracy_threshold"] = accuracy
	}

	if efficiency, exists := requirements["target_efficiency"]; exists {
		criteria["efficiency_threshold"] = efficiency
	}

	// Time-based criteria
	criteria["evaluation_period"] = "1_week"
	criteria["minimum_samples"] = 10

	return criteria
}

func (s *ToolMatchingService) generateOptimalParameters(robot *model.RobotSpecification, model *model.OptimizationModel, stateEval *model.StateEvaluation, requirements map[string]interface{}) map[string]interface{} {
	parameters := make(map[string]interface{})

	// Robot parameters
	robotParams := map[string]interface{}{
		"max_payload":      robot.PayloadKg * 0.8,        // Use 80% of max payload for safety
		"working_speed":    robot.MaxSpeedMmS * 0.6,      // Use 60% of max speed initially
		"precision_mode":   robot.RepeatAccuracyMm < 0.02, // Enable precision mode for high-accuracy robots
	}

	// Adjust based on level
	if stateEval.Level <= 2 {
		robotParams["working_speed"] = robot.MaxSpeedMmS * 0.4 // Slower for beginners
		robotParams["safety_factor"] = 1.5
	}

	parameters["robot"] = robotParams

	// Optimization parameters
	if model.Parameters.Valid {
		modelParams := make(map[string]interface{})
		paramPairs := strings.Split(model.Parameters.String, "|")
		
		for _, pair := range paramPairs {
			if strings.Contains(pair, ":") {
				parts := strings.Split(pair, ":")
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					
					// Try to parse as number
					if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
						modelParams[key] = floatVal
					} else {
						modelParams[key] = value
					}
				}
			}
		}
		parameters["optimization"] = modelParams
	}

	// Process parameters based on requirements
	processParams := map[string]interface{}{
		"update_interval": s.determineUpdateFrequency(model),
		"batch_size":      32,
		"learning_rate":   0.01,
	}

	// Adjust learning rate based on level
	if stateEval.Level >= 4 {
		processParams["learning_rate"] = 0.001 // More conservative for advanced users
	}

	parameters["process"] = processParams

	return parameters
}

func (s *ToolMatchingService) predictPerformance(robot *model.RobotSpecification, model *model.OptimizationModel, parameters map[string]interface{}, stateEval *model.StateEvaluation) map[string]interface{} {
	performance := make(map[string]interface{})

	// Base performance prediction
	baseScore := stateEval.EvaluationScore

	// Robot capability factor
	robotFactor := 1.0
	if robot.RepeatAccuracyMm <= 0.01 {
		robotFactor += 0.15 // High precision bonus
	}
	if robot.AICapability.Valid && strings.Contains(robot.AICapability.String, "learning") {
		robotFactor += 0.10 // AI capability bonus
	}

	// Model effectiveness factor
	modelFactor := 1.0
	if model.ConvergenceRate.Valid && model.ConvergenceRate.Float64 > 0.95 {
		modelFactor += 0.20 // High convergence rate bonus
	}

	// Combined prediction
	predictedScore := baseScore * robotFactor * modelFactor
	predictedScore = math.Min(98.0, predictedScore) // Cap at 98% to be realistic

	performance["predicted_score"] = predictedScore
	performance["confidence_level"] = s.calculateConfidenceLevel(robot, model, stateEval)
	performance["estimated_timeline"] = s.estimateTimeline(stateEval.Level, predictedScore-baseScore)
	
	// Detailed metrics
	metrics := map[string]interface{}{
		"accuracy_improvement":   fmt.Sprintf("%.1f%%", (predictedScore-baseScore)*0.3),
		"efficiency_improvement": fmt.Sprintf("%.1f%%", (predictedScore-baseScore)*0.4),
		"consistency_improvement": fmt.Sprintf("%.1f%%", (predictedScore-baseScore)*0.3),
	}
	performance["detailed_metrics"] = metrics

	return performance
}

func (s *ToolMatchingService) calculateConfidenceLevel(robot *model.RobotSpecification, model *model.OptimizationModel, stateEval *model.StateEvaluation) string {
	confidence := 0.7 // Base confidence

	// Robot maturity factor
	if robot.MaintenanceIntervalHours >= 2000 {
		confidence += 0.1 // Mature, reliable robot
	}

	// Model convergence factor
	if model.ConvergenceRate.Valid && model.ConvergenceRate.Float64 > 0.9 {
		confidence += 0.15
	}

	// User level factor
	if stateEval.Level >= 3 {
		confidence += 0.05 // Experienced user
	}

	if confidence >= 0.9 {
		return "Very High (90%+)"
	} else if confidence >= 0.8 {
		return "High (80-90%)"
	} else if confidence >= 0.7 {
		return "Medium (70-80%)"
	} else {
		return "Low (<70%)"
	}
}

func (s *ToolMatchingService) estimateTimeline(level int, improvement float64) string {
	baseTime := []int{8, 6, 4, 3, 2}[level-1] // weeks based on level

	// Adjust based on expected improvement
	if improvement > 20 {
		baseTime += 2 // More complex improvements take longer
	} else if improvement < 5 {
		baseTime -= 1 // Minor improvements are quicker
	}

	if baseTime <= 0 {
		baseTime = 1
	}

	return fmt.Sprintf("%d weeks", baseTime)
}

func (s *ToolMatchingService) GetMatchingHistory(stateEvaluationID string) ([]model.ToolMatchingResult, error) {
	var results []model.ToolMatchingResult

	if err := s.db.Where("state_evaluation_id = ?", stateEvaluationID).Order("created_at DESC").Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to get matching history: %v", err)
	}

	return results, nil
}