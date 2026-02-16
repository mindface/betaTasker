package service

import (
	dtoquery "github.com/godotask/dto/query"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/db/repository"
)

type HeuristicsService struct {
	Repo repository.HeuristicsRepositoryInterface
}

func (s *HeuristicsService) CreateAnalyzeData(request *model.HeuristicsAnalysisRequest) (*model.HeuristicsAnalysis, error) {
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

func (s *HeuristicsService) ListAnalyses() ([]model.HeuristicsAnalysis, error) {
  return s.Repo.FindAllAnalyses()
}

func (s *HeuristicsService) ListAnalysesPager(filter dtoquery.QueryFilter, pager dtoquery.PagerQuery) ([]model.HeuristicsAnalysis, int64, error) {
  return s.Repo.ListAnalysesPager(filter, pager.Offset, pager.Limit)
}

func (s *HeuristicsService) UpdateAnalyzeData(id string, analyze *model.HeuristicsAnalysis) error {
	return s.Repo.UpdateAnalysis(id, analyze)
}

func (s *HeuristicsService) DeleteAnalyzeData(id string) error {
	return s.Repo.DeleteAnalysis(id)
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

func (s *HeuristicsService) DetectPatterns(userID, dataType, period string) ([]model.HeuristicsPattern, error) {
	return s.Repo.DetectPatterns(userID, dataType, period)
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