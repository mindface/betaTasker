package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type HeuristicsInsightService struct {
	Repo repository.HeuristicsInsightRepositoryInterface
}

func (s *HeuristicsInsightService) CreateInsightData(insight *model.HeuristicsInsight) (*model.HeuristicsInsight, error) {
	return s.Repo.CreateInsight(insight)
}

func (s *HeuristicsInsightService) GetInsighById(id string) (*model.HeuristicsInsight, error) {
  return s.Repo.GetInsightById(id)
}

func (s *HeuristicsInsightService) ListInsight() ([]model.HeuristicsInsight, error) {
  return s.Repo.ListInsights()
}

func (s *HeuristicsInsightService) UpdateInsightData(id string, insight *model.HeuristicsInsight) error {
	return s.Repo.UpdateInsight(id, insight)
}

func (s *HeuristicsInsightService) DeleteInsightData(id string) error {
	return s.Repo.DeleteInsight(id)
}
