package service

import (
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/repository"
)

type AssessmentService struct {
	Repo repository.AssessmentRepositoryInterface
}

func NewAssessmentService(repo repository.AssessmentRepositoryInterface) *AssessmentService {
	return &AssessmentService{Repo: repo}
}

func (s *AssessmentService) CreateAssessment(task *model.Assessment) error {
	return s.Repo.Create(task)
}
func (s *AssessmentService) GetAssessmentByID(id string) (*model.Assessment, error) {
	return s.Repo.FindByID(id)
}
func (s *AssessmentService) ListAssessments() ([]model.Assessment, error) {
	return s.Repo.FindAll()
}
func (s *AssessmentService) ListAssessmentsForTaskUser(userID int, taskID int) ([]model.Assessment, error) {
	return s.Repo.FindByTaskIDAndUserID(userID,taskID)
}
func (s *AssessmentService) ListAssessmentsTOPager(page, perPage, offset int) ([]model.Assessment, int64, error) {
  return s.Repo.ListAssessmentsPager(offset, perPage)
}
func (s *AssessmentService) ListAssessmentsForTaskUserPager(userID, taskID, page, perPage, offset int) ([]model.Assessment, int64, error) {
  return s.Repo.ListAssessmentsForTaskUserPager(userID, taskID, offset, perPage)
}
func (s *AssessmentService) UpdateAssessment(id string, task *model.Assessment) error {
	return s.Repo.Update(id, task)
}
func (s *AssessmentService) DeleteAssessment(id string) error {	
	return s.Repo.Delete(id)
}
