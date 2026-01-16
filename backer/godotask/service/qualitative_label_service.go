package service

import (
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/repository"
)

type QualitativeLabelService struct {
  Repo repository.QualitativeLabelRepositoryInterface
}

func (s *QualitativeLabelService) CreateQualitativeLabel(qualitativeLabel *model.QualitativeLabel) error {
	return s.Repo.Create(qualitativeLabel)
}
func (s *QualitativeLabelService) GetQualitativeLabelByID(id string) (*model.QualitativeLabel, error) {
	return s.Repo.FindByID(id)
}
func (s *QualitativeLabelService) ListQualitativeLabels() ([]model.QualitativeLabel, error) {
	return s.Repo.FindAll()
}
func (s *QualitativeLabelService) UpdateQualitativeLabel(id string, qualitativeLabel *model.QualitativeLabel) error {
	return s.Repo.Update(id, qualitativeLabel)
}
func (s *QualitativeLabelService) DeleteQualitativeLabel(id string) error {
	return s.Repo.Delete(id)
}
