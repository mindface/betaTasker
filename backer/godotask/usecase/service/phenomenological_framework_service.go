package service

import (
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/db/repository"
)

type PhenomenologicalFrameworkService struct {
  Repo repository.PhenomenologicalFrameworkRepositoryInterface
}

func (s *PhenomenologicalFrameworkService) CreatePhenomenologicalFramework(phenomenologicalFramework *model.PhenomenologicalFramework) error {
	return s.Repo.Create(phenomenologicalFramework)
}
func (s *PhenomenologicalFrameworkService) GetPhenomenologicalFrameworkByID(id string) (*model.PhenomenologicalFramework, error) {
	return s.Repo.FindByID(id)
}
func (s *PhenomenologicalFrameworkService) ListPhenomenologicalFrameworks(userID uint) ([]model.PhenomenologicalFramework, error) {
	return s.Repo.FindAll(userID)
}
func (s *PhenomenologicalFrameworkService) UpdatePhenomenologicalFramework(id string, phenomenologicalFramework *model.PhenomenologicalFramework) error {
	return s.Repo.Update(id, phenomenologicalFramework)
}
func (s *PhenomenologicalFrameworkService) DeletePhenomenologicalFramework(id string) error {
	return s.Repo.Delete(id)
}
