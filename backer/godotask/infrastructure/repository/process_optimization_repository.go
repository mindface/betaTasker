package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
)

type ProcessOptimizationRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ProcessOptimizationRepositoryImpl) Create(processOptimization *model.ProcessOptimization) error {
	return r.DB.Create(processOptimization).Error
}

func (r *ProcessOptimizationRepositoryImpl) FindByID(id string) (*model.ProcessOptimization, error) {
	var processOptimization model.ProcessOptimization
	if err := r.DB.Where("id = ?", id).First(&processOptimization).Error; err != nil {
		return nil, err
	}
	return &processOptimization, nil
}

func (r *ProcessOptimizationRepositoryImpl) FindAll() ([]model.ProcessOptimization, error) {
	var processOptimizations []model.ProcessOptimization
	if err := r.DB.Find(&processOptimizations).Error; err != nil {
		return nil, err
	}
	return processOptimizations, nil
}

func (r *ProcessOptimizationRepositoryImpl) Update(id string, processOptimization *model.ProcessOptimization) error {
	return r.DB.Model(&model.ProcessOptimization{}).Where("id = ?", id).Updates(processOptimization).Error
}

func (r *ProcessOptimizationRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.ProcessOptimization{}, id).Error
}

// NewProcessOptimizationRepository は ProcessOptimizationRepositoryInterface を返すコンストラクタ
func NewProcessOptimizationRepository(db *gorm.DB) ProcessOptimizationRepositoryInterface {
	return &ProcessOptimizationRepositoryImpl{DB: db}
}