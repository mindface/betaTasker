package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
)

// MemoryRepositoryImpl
type MemoryRepositoryImpl struct {
	DB *gorm.DB
}

func (r *MemoryRepositoryImpl) Create(memory *model.Memory) error {
	return r.DB.Create(memory).Error
}

func (r *MemoryRepositoryImpl) FindByID(id string) (*model.Memory, error) {
	var m model.Memory
	if err := r.DB.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *MemoryRepositoryImpl) FindAll() ([]model.Memory, error) {
	var memories []model.Memory
	if err := r.DB.Find(&memories).Error; err != nil {
		return nil, err
	}
	return memories, nil
}

func (r *MemoryRepositoryImpl) Update(id string, memory *model.Memory) error {
	return r.DB.Model(&model.Memory{}).Where("id = ?", id).Updates(memory).Error
}

func (r *MemoryRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.Memory{}, id).Error
}

func NewMemoryRepository(db *gorm.DB) MemoryRepositoryInterface {
	return &MemoryRepositoryImpl{DB: db}
}
