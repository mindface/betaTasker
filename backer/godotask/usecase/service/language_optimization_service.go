package service

import (
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/db/repository"
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
func (s *LanguageOptimizationService) ListLanguageOptimizations(userID uint) ([]model.LanguageOptimization, error) {
	return s.Repo.FindAll(userID)
}
func (s *LanguageOptimizationService) UpdateLanguageOptimization(id string, languageOptimization *model.LanguageOptimization) error {
	return s.Repo.Update(id, languageOptimization)
}
func (s *LanguageOptimizationService) DeleteLanguageOptimization(id string) error {
	return s.Repo.Delete(id)
}
