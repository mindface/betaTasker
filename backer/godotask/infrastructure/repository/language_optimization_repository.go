package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
)

type LanguageOptimizationRepositoryImpl struct {
	DB *gorm.DB
}

func (r *LanguageOptimizationRepositoryImpl) Create(languageOptimization *model.LanguageOptimization) error {
	return r.DB.Create(languageOptimization).Error
}

func (r *LanguageOptimizationRepositoryImpl) FindByID(id string) (*model.LanguageOptimization, error) {
	var languageOptimization model.LanguageOptimization
	if err := r.DB.Where("id = ?", id).First(&languageOptimization).Error; err != nil {
		return nil, err
	}
	return &languageOptimization, nil
}

func (r *LanguageOptimizationRepositoryImpl) FindAll() ([]model.LanguageOptimization, error) {
	var languageOptimizations []model.LanguageOptimization
	if err := r.DB.Find(&languageOptimizations).Error; err != nil {
		return nil, err
	}
	return languageOptimizations, nil
}

func (r *LanguageOptimizationRepositoryImpl) Update(id string, languageOptimization *model.LanguageOptimization) error {
	return r.DB.Model(&model.LanguageOptimization{}).Where("id = ?", id).Updates(languageOptimization).Error
}

func (r *LanguageOptimizationRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.LanguageOptimization{}, id).Error
}

// NewLanguageOptimizationRepository は LanguageOptimizationRepositoryInterface を返すコンストラクタ
func NewLanguageOptimizationRepository(db *gorm.DB) LanguageOptimizationRepositoryInterface {
	return &LanguageOptimizationRepositoryImpl{DB: db}
}