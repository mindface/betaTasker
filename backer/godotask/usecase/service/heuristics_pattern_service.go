package service

import (
	dtoquery "github.com/godotask/dto/query"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/db/repository"
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

func (s *HeuristicsPatternService) ListPattern(userID uint) ([]model.HeuristicsPattern, error) {
  return s.Repo.ListPattern(userID)
}

func (s *HeuristicsPatternService) ListPatternPager(filter dtoquery.QueryFilter, pager dtoquery.PagerQuery) ([]model.HeuristicsPattern, int64, error) {
  return s.Repo.ListPatternPager(filter, pager.Offset, pager.Limit)
}

func (s *HeuristicsPatternService) UpdatePatternData(id string, pattern *model.HeuristicsPattern) error {
	return s.Repo.UpdatePattern(id, pattern)
}

func (s *HeuristicsPatternService) DeletePatternData(id string) error {
	return s.Repo.DeletePattern(id)
}
