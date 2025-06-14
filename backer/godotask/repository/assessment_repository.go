package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
)

type AssessmentRepository struct {
	DB *gorm.DB
}

func (r *AssessmentRepository) Create(a *model.Assessment) error {
	return r.DB.Create(a).Error
}

func (r *AssessmentRepository) FindByID(id string) (*model.Assessment, error) {
	var assessment model.Assessment
	if err := r.DB.Where("id = ?", id).First(&assessment).Error; err != nil {
		return nil, err
	}
	return &assessment, nil
}

func (r *AssessmentRepository) FindAll() ([]model.Assessment, error) {
	var assessments []model.Assessment
	if err := r.DB.Find(&assessments).Error; err != nil {
		return nil, err
	}
	return assessments, nil
}

func (r *AssessmentRepository) Update(id string, a *model.Assessment) error {
	return r.DB.Model(&model.Assessment{}).Where("id = ?", id).Updates(a).Error
}

func (r *AssessmentRepository) Delete(id string) error {
	return r.DB.Delete(&model.Assessment{}, id).Error
}
