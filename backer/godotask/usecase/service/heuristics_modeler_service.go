package service

import (
	dtoquery "github.com/godotask/dto/query"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/db/repository"
)

type HeuristicsModelerService struct {
	Repo repository.HeuristicsModelerRepositoryInterface
}

func (s *HeuristicsModelerService) CreateModelerData(modeler *model.HeuristicsModeler) (*model.HeuristicsModeler, error) {
  if err := s.Repo.CreateModeler(modeler); err != nil {
    return nil, err
  }
  return modeler, nil
}

func (s *HeuristicsModelerService) GetModelerById(id string) (*model.HeuristicsModeler, error) {
  return s.Repo.GetModelerById(id)
}

func (s *HeuristicsModelerService) ListModeler(userID uint) ([]model.HeuristicsModeler, error) {
  return s.Repo.ListModeler(userID)
}

func (s *HeuristicsModelerService) ListModelerPager(filter dtoquery.QueryFilter, pager dtoquery.PagerQuery) ([]model.HeuristicsModeler, int64, error) {
  return s.Repo.ListModelerPager(filter, pager.Offset, pager.Limit)
}

func (s *HeuristicsModelerService) UpdateModelerData(id string, modeler *model.HeuristicsModeler) error {
	return s.Repo.UpdateModeler(id, modeler)
}

func (s *HeuristicsModelerService) DeleteModelerData(id string) error {
	return s.Repo.DeleteModeler(id)
}
