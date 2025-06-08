package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
)

// MemoryRepository
type MemoryRepository struct {
	DB *gorm.DB
}

func (r *MemoryRepository) Create(memory *model.Memory) error {
	return r.DB.Create(memory).Error
}
func (r *MemoryRepository) FindByID(id string) (*model.Memory, error) {
	var m model.Memory
	if err := r.DB.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MemoryRepository) FindAll() ([]model.Memory, error) {
	var memories []model.Memory
	if err := r.DB.Find(&memories).Error; err != nil {
		return nil, err
	}
	return memories, nil
}

func (r *MemoryRepository) Update(id string, memory *model.Memory) error {
	return r.DB.Model(&model.Memory{}).Where("id = ?", id).Updates(memory).Error
}

func (r *MemoryRepository) Delete(id string) error {
	return r.DB.Delete(&model.Memory{}, id).Error
}

// TaskRepository ...
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
