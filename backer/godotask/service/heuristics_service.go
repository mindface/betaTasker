package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type HeuristicsService struct {
	Repo repository.HeuristicsRepositoryInterface
}

func (s *HeuristicsService) AnalyzeData(request *model.HeuristicsAnalysisRequest) (*model.HeuristicsAnalysis, error) {
	analysis := &model.HeuristicsAnalysis{
		UserID:       request.UserID,
		TaskID:       request.TaskID,
		AnalysisType: request.AnalysisType,
		Result:       toJSONString(request.Data),
		Score:        calculateScore(request.Data),
		Status:       "completed",
	}

	if err := s.Repo.CreateAnalysis(analysis); err != nil {
		return nil, err
	}
	
	return analysis, nil
}

func (s *HeuristicsService) GetAnalysisById(id string) (*model.HeuristicsAnalysis, error) {
	return s.Repo.GetAnalysisById(id)
}

func (s *HeuristicsService) TrackUserBehavior(trackData *model.HeuristicsTrackingData) error {
	tracking := &model.HeuristicsTracking{
		UserID:    trackData.UserID,
		Action:    trackData.Action,
		Context:   toJSONString(trackData.Context),
		SessionID: trackData.SessionID,
		Duration:  trackData.Duration,
	}
	
	return s.Repo.CreateTracking(tracking)
}

func (s *HeuristicsService) GetTrackingDataByUserID(userID string) ([]model.HeuristicsTracking, error) {
	return s.Repo.GetTrackingByUserID(userID)
}

func (s *HeuristicsService) GetInsights(userID string, limit, offset int) ([]model.HeuristicsInsight, int, error) {
	return s.Repo.GetInsights(userID, limit, offset)
}

func (s *HeuristicsService) GetInsightById(id string) (*model.HeuristicsInsight, error) {
	return s.Repo.GetInsightById(id)
}

func (s *HeuristicsService) DetectPatterns(userID, dataType, period string) ([]model.HeuristicsPattern, error) {
	return s.Repo.DetectPatterns(userID, dataType, period)
}

func (s *HeuristicsService) TrainModel(request *model.HeuristicsTrainRequest) (*model.HeuristicsModel, error) {
	modelData := &model.HeuristicsModel{
		ModelType:   request.ModelType,
		Version:     "1.0.0",
		Parameters:  toJSONString(request.Parameters),
		Performance: toJSONString(map[string]interface{}{"status": "training"}),
		Status:      "training",
	}
	
	if err := s.Repo.CreateModel(modelData); err != nil {
		return nil, err
	}
	
	return modelData, nil
}

// Helper functions
func toJSONString(data interface{}) string {
	// 簡単なJSON変換（実際のプロジェクトではjson.Marshalを使用）
	return "{}"
}

func calculateScore(data map[string]interface{}) float64 {
	// 簡単なスコア計算（実際のプロジェクトでは複雑なロジック）
	return 85.0
}