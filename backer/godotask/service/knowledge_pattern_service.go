package service

import (
	"github.com/godotask/model"
	"github.com/godotask/repository"
)

type KnowledgePatternService struct {
  Repo repository.KnowledgePatternRepositoryInterface
}

func (s *KnowledgePatternService) CreateKnowledgePattern(knowledgePattern *model.KnowledgePattern) error {
	return s.Repo.Create(knowledgePattern)
}
func (s *KnowledgePatternService) GetKnowledgePatternByID(id string) (*model.KnowledgePattern, error) {
	return s.Repo.FindByID(id)
}
func (s *KnowledgePatternService) ListKnowledgePatterns() ([]model.KnowledgePattern, error) {
	return s.Repo.FindAll()
}
func (s *KnowledgePatternService) UpdateKnowledgePattern(id string, knowledgePattern *model.KnowledgePattern) error {
	return s.Repo.Update(id, knowledgePattern)
}
func (s *KnowledgePatternService) DeleteKnowledgePattern(id string) error {
	return s.Repo.Delete(id)
}
