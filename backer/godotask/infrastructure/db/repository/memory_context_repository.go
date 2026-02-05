package repository

import (
	"github.com/godotask/infrastructure/db/model"
	"gorm.io/gorm"
	"github.com/rs/zerolog/log"
)

type MemoryContextRepositoryImpl struct {
	DB *gorm.DB
}

// func (r *MemoryContextRepositoryImpl) FindByCode(code string, contexts *[]model.MemoryContext) error {
// 	 return r.DB.Where("work_target LIKE ?", "%"+code+"%").Find(contexts).Error
// }

func (r *MemoryContextRepositoryImpl) FindByCode(code string, contexts *[]model.MemoryContext) error {
	return r.DB.Where("work_target LIKE ?", "%"+code+"%").Find(contexts).Error
}

// FindWithAidsByCode: work_targetにcodeが含まれるもの＋リレーションを返す
// func (r *MemoryContextRepositoryImpl) FindWithAidsByCode(code string, contexts *[]model.MemoryContext) error {
// 	return r.DB.Preload("TechnicalFactors").Preload("KnowledgeTransformations").Where("work_target LIKE ?", "%"+code+"%").Find(contexts).Error
// }

func (r *MemoryContextRepositoryImpl) FindWithAidsByCode(code string, contexts *[]model.MemoryContext) error {
	log.Info().Msgf("Finding memory contexts with code: %s", code)
	return r.DB.
		Preload("TechnicalFactors").
		Preload("KnowledgeTransformations").
		Where("work_target LIKE ?", "%"+code+"%").
		Find(contexts).Error
}

func NewMemoryContextRepository(db *gorm.DB) MemoryContextRepositoryInterface {
	return &MemoryContextRepositoryImpl{DB: db}
}
