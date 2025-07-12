package repository

import (
	"github.com/godotask/model"
	"gorm.io/gorm"
)

type MemoryContextRepository struct {
	DB *gorm.DB
}

func (r *MemoryContextRepository) FindByCode(code string, contexts *[]model.MemoryContext) error {
	return r.DB.Where("work_target LIKE ?", "%"+code+"%").Find(contexts).Error
}

// FindWithAidsByCode: work_targetにcodeが含まれるもの＋リレーションを返す
func (r *MemoryContextRepository) FindWithAidsByCode(code string, contexts *[]model.MemoryContext) error {
	return r.DB.Preload("TechnicalFactors").Preload("KnowledgeTransformations").Where("work_target LIKE ?", "%"+code+"%").Find(contexts).Error
}
