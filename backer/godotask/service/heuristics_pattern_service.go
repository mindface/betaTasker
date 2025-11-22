package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type HeuristicsPatternService struct {
	Repo repository.HeuristicsPatternRepositoryInterface
}

func (s *HeuristicsPatternService) CreatePatternData(pattern *model.HeuristicsPattern) (*model.HeuristicsPattern, error) {
  if err := s.Repo.CreatePattern(pattern); err != nil {
      return nil, err
  }
  return pattern, nil
}

func (s *HeuristicsPatternService) GetPatternById(id string) (*model.HeuristicsPattern, error) {
  return s.Repo.GetPatternById(id)
}

func (s *HeuristicsPatternService) ListPattern() ([]model.HeuristicsPattern, error) {
  return s.Repo.ListPattern()
}

func (s *HeuristicsPatternService) UpdatePatternData(id string, pattern *model.HeuristicsPattern) error {
	return s.Repo.UpdatePattern(id, pattern)
}

func (s *HeuristicsPatternService) DeletePatternData(id string) error {
	return s.Repo.DeletePattern(id)
}
