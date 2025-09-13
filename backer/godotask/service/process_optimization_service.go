package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type ProcessOptimizationService struct {
  Repo repository.ProcessOptimizationRepositoryInterface
}

func (s *ProcessOptimizationService) CreateProcessOptimization(processOptimization *model.ProcessOptimization) error {
	return s.Repo.Create(processOptimization)
}
func (s *ProcessOptimizationService) GetProcessOptimizationByID(id string) (*model.ProcessOptimization, error) {
	return s.Repo.FindByID(id)
}
func (s *ProcessOptimizationService) ListProcessOptimizations() ([]model.ProcessOptimization, error) {
	return s.Repo.FindAll()
}
func (s *ProcessOptimizationService) UpdateProcessOptimization(id string, processOptimization *model.ProcessOptimization) error {
	return s.Repo.Update(id, processOptimization)
}
func (s *ProcessOptimizationService) DeleteProcessOptimization(id string) error {
	return s.Repo.Delete(id)
}
