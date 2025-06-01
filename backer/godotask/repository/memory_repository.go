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
// 他CRUDも同様

// AssessmentRepository ...
type AssessmentRepository struct {
	DB *gorm.DB
}

func (r *AssessmentRepository) Create(a *model.Assessment) error {
	return r.DB.Create(a).Error
}
