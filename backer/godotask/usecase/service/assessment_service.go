package service

import (
	dtoquery "github.com/godotask/dto/query"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/db/repository"
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
func (s *AssessmentService) ListAssessments(userID uint) ([]model.Assessment, int64, error) {
	return s.Repo.FindAll(userID)
}
func (s *AssessmentService) ListAssessmentsPager(userID uint, pager dtoquery.PagerQuery) ([]model.Assessment, int64, error) {
  return s.Repo.ListAssessmentsPager(userID, pager.Offset, pager.Limit)
}
func (s *AssessmentService) ListAssessmentsForTaskUserPager(filter dtoquery.QueryFilter, pager dtoquery.PagerQuery) ([]model.Assessment, int64, error) {
  return s.Repo.ListAssessmentsForTaskUserPager(filter, pager.Offset, pager.Limit)
}
func (s *AssessmentService) UpdateAssessment(id string, task *model.Assessment) error {
	return s.Repo.Update(id, task)
}
func (s *AssessmentService) DeleteAssessment(id string) error {	
	return s.Repo.Delete(id)
}
