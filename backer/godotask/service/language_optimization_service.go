package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type LanguageOptimizationService struct {
  Repo repository.LanguageOptimizationRepositoryInterface
}

func (s *LanguageOptimizationService) CreateLanguageOptimization(languageOptimization *model.LanguageOptimization) error {
	return s.Repo.Create(languageOptimization)
}
func (s *LanguageOptimizationService) GetLanguageOptimizationByID(id string) (*model.LanguageOptimization, error) {
	return s.Repo.FindByID(id)
}
func (s *LanguageOptimizationService) ListLanguageOptimizations() ([]model.LanguageOptimization, error) {
	return s.Repo.FindAll()
}
func (s *LanguageOptimizationService) UpdateLanguageOptimization(id string, languageOptimization *model.LanguageOptimization) error {
	return s.Repo.Update(id, languageOptimization)
}
func (s *LanguageOptimizationService) DeleteLanguageOptimization(id string) error {
	return s.Repo.Delete(id)
}
