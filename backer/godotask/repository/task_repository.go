package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
)

type TaskRepositoryImpl struct {
	DB *gorm.DB
}

func (r *TaskRepositoryImpl) Create(task *model.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepositoryImpl) FindByID(id string) (*model.Task, error) {
	var task model.Task
	if err := r.DB.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	// if err := r.DB.Find(&tasks).Error; err != nil {
	// 	return nil, err
	// }
	err := r.DB.Preload("QualitativeLabels").
		Preload("QuantificationLabels").
		Preload("MultimodalData").
		Preload("HeuristicsModel").
		Preload("HeuristicsTracking").
		Preload("HeuristicsInsight").
		Preload("KnowledgePatterns").
		Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) Update(id string, task *model.Task) error {
	return r.DB.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error
}

func (r *TaskRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.Task{}, id).Error
}
