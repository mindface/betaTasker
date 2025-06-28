package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
)

type TaskRepository struct {
	DB *gorm.DB
}

func (r *TaskRepository) Create(task *model.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) FindByID(id string) (*model.Task, error) {
	var t model.Task
	if err := r.DB.Where("id = ?", id).First(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TaskRepository) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepository) Update(id string, task *model.Task) error {
	return r.DB.Model(&model.Task{}).Where("id = ?", id).Updates(task).Error
}

func (r *TaskRepository) Delete(id string) error {
	return r.DB.Delete(&model.Task{}, id).Error
}
