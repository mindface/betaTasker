package service

import (
	dtoquery "github.com/godotask/dto/query"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/db/repository"
)

type HeuristicsAnalysisService struct {
	Repo repository.HeuristicsAnalysisRepositoryInterface
}

func (s *HeuristicsAnalysisService) CreateAnalyzeData(request *model.HeuristicsAnalysisRequest) (*model.HeuristicsAnalysis, error) {
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

func (s *HeuristicsAnalysisService) GetAnalysisById(id string) (*model.HeuristicsAnalysis, error) {
  return s.Repo.GetAnalysisById(id)
}

func (s *HeuristicsAnalysisService) ListAnalyze() ([]model.HeuristicsAnalysis, error) {
  return s.Repo.ListAnalyze()
}

func (s *HeuristicsAnalysisService) ListAnalysesPager(filter dtoquery.QueryFilter, pager dtoquery.PagerQuery) ([]model.HeuristicsAnalysis, int64, error) {
  return s.Repo.ListAnalysesPager(filter, pager.Offset, pager.Limit)
}

func (s *HeuristicsAnalysisService) UpdateAnalyzeData(id string, analyze *model.HeuristicsAnalysis) error {
	return s.Repo.UpdateAnalysis(id, analyze)
}

func (s *HeuristicsAnalysisService) DeleteAnalyzeData(id string) error {
	return s.Repo.DeleteAnalysis(id)
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