package repository

import (
	"errors"
	"gorm.io/gorm"
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/helper"
)

type AssessmentRepositoryImpl struct {
	DB *gorm.DB
}

func (r *AssessmentRepositoryImpl) Create(a *model.Assessment) error {
	return r.DB.Create(a).Error
}

func (r *AssessmentRepositoryImpl) FindByID(id string) (*model.Assessment, error) {
	var assessment model.Assessment
	if err := r.DB.Where("id = ?", id).First(&assessment).Error; err != nil {
		return nil, err
	}
	return &assessment, nil
}

func (r *AssessmentRepositoryImpl) FindByTaskIDAndUserID(userID int, taskID int) ([]model.Assessment, error) {
	var assessments []model.Assessment
	if err := r.DB.
		Where("task_id = ? AND user_id = ?", taskID, userID).
		Find(&assessments).Error; err != nil {
			return nil, err
		}
	return assessments, nil
}

func (r *AssessmentRepositoryImpl) FindAll() ([]model.Assessment, error) {
	var assessments []model.Assessment
	if err := r.DB.Find(&assessments).Error; err != nil {
		return nil, err
	}
	return assessments, nil
}

func (r *AssessmentRepositoryImpl) ListAssessmentsPager(userID uint, offset int, limit int) ([]model.Assessment, int64, error) {
	var assessments []model.Assessment
	var total int64

	q := r.DB.Model(&model.Assessment{}).Scopes(helper.WithUserFilter(userID))

	if err := q.Count(&total).Error; err != nil {
			return nil, 0, err
	}

	if err := q.Order("created_at DESC, id DESC").Limit(limit).Offset(offset).Find(&assessments).Error; err != nil {
			return nil, 0, err
	}
	return assessments, total, nil
}

func (r *AssessmentRepositoryImpl) ListAssessmentsForTaskUserPager(userID, taskID, offset, limit int) ([]model.Assessment, int64, error) {
	var assessments []model.Assessment
	var total int64

	q := r.DB.Model(&model.Assessment{})

	// フィルタ条件
	if userID > 0 {
			q = q.Where("user_id = ?", userID)
	}
	if taskID > 0 {
			q = q.Where("task_id = ?", taskID)
	}

	if err := q.Count(&total).Error; err != nil {
			return nil, 0, err
	}

	if err := q.Order("created_at DESC, id DESC").Limit(limit).Offset(offset).Find(&assessments).Error; err != nil {
			return nil, 0, err
	}
	return assessments, total, nil
}

func (r *AssessmentRepositoryImpl) Update(id string, a *model.Assessment) error {
	return r.DB.Model(&model.Assessment{}).Where("id = ?", id).Updates(a).Error
}

func (r *AssessmentRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.Assessment{}, id).Error
}


// ErrorMockAssessmentRepository is a mock implementation of AssessmentRepositoryInterface that returns an error for FindByID
type ErrorMockAssessmentRepository struct{}

func (e *ErrorMockAssessmentRepository) Create(assessment *model.Assessment) error {
	return errors.New("not implemented")
}

func (e *ErrorMockAssessmentRepository) FindByID(id string) (*model.Assessment, error) {
	return nil, gorm.ErrRecordNotFound
}

func (e *ErrorMockAssessmentRepository) FindAll() ([]model.Assessment, error) {
	return nil, errors.New("not implemented")
}

func (e *ErrorMockAssessmentRepository) Update(id string, assessment *model.Assessment) error {
	return errors.New("not implemented")
}

func (e *ErrorMockAssessmentRepository) Delete(id string) error {
	return errors.New("not implemented")
}