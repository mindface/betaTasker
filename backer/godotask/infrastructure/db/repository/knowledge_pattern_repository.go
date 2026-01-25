package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/helper"
)

type KnowledgePatternRepositoryImpl struct {
	DB *gorm.DB
}

func (r *KnowledgePatternRepositoryImpl) Create(knowledgePattern *model.KnowledgePattern) error {
	return r.DB.Create(knowledgePattern).Error
}

func (r *KnowledgePatternRepositoryImpl) FindByID(id string) (*model.KnowledgePattern, error) {
	var knowledgePattern model.KnowledgePattern
	if err := r.DB.Where("id = ?", id).First(&knowledgePattern).Error; err != nil {
		return nil, err
	}
	return &knowledgePattern, nil
}

func (r *KnowledgePatternRepositoryImpl) FindAll(userID uint) ([]model.KnowledgePattern, error) {
	var knowledgePatterns []model.KnowledgePattern
	if err := r.DB.Scopes(helper.WithUserFilter(userID)).Order("created_at DESC, id DESC").Find(&knowledgePatterns).Error; err != nil {
		return nil, err
	}
	return knowledgePatterns, nil
}

func (r *KnowledgePatternRepositoryImpl) Update(id string, knowledgePattern *model.KnowledgePattern) error {
	return r.DB.Model(&model.KnowledgePattern{}).Where("id = ?", id).Updates(knowledgePattern).Error
}

func (r *KnowledgePatternRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.KnowledgePattern{}, id).Error
}

// NewKnowledgePatternRepository は KnowledgePatternRepositoryInterface を返すコンストラクタ
func NewKnowledgePatternRepository(db *gorm.DB) KnowledgePatternRepositoryInterface {
	return &KnowledgePatternRepositoryImpl{DB: db}
}