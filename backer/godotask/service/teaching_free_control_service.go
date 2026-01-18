package service

import (
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/repository"
)

type TeachingFreeControlService struct {
  Repo repository.TeachingFreeControlRepositoryInterface
}

func (s *TeachingFreeControlService) CreateTeachingFreeControl(teachingFreeControl *model.TeachingFreeControl) error {
	return s.Repo.Create(teachingFreeControl)
}
func (s *TeachingFreeControlService) GetTeachingFreeControlByID(id string) (*model.TeachingFreeControl, error) {
	return s.Repo.FindByID(id)
}
func (s *TeachingFreeControlService) ListTeachingFreeControls(userID uint) ([]model.TeachingFreeControl, error) {
	return s.Repo.FindAll(userID)
}
func (s *TeachingFreeControlService) UpdateTeachingFreeControl(id string, teachingFreeControl *model.TeachingFreeControl) error {
	return s.Repo.Update(id, teachingFreeControl)
}
func (s *TeachingFreeControlService) DeleteTeachingFreeControl(id string) error {
	return s.Repo.Delete(id)
}
