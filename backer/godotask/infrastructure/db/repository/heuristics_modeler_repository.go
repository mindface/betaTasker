package repository

import (
	"github.com/godotask/infrastructure/db/model"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"gorm.io/gorm"
)

type HeuristicsModelerRepositoryImpl struct {
	DB *gorm.DB
}

func (r *HeuristicsModelerRepositoryImpl) CreateModeler(modeler *model.HeuristicsModeler) error {
  return r.DB.Create(modeler).Error
}

func (r *HeuristicsModelerRepositoryImpl) GetModelerById(id string) (*model.HeuristicsModeler, error) {
	var modeler model.HeuristicsModeler
	if err := r.DB.First(&modeler, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &modeler, nil
}

func (r *HeuristicsModelerRepositoryImpl) ListModeler(userID uint) ([]model.HeuristicsModeler, error) {
	var modelers []model.HeuristicsModeler

	if err := r.DB.Where("user_id = ?", userID).Find(&modelers).Error; err != nil {
		return nil, err
	}
	return modelers, nil
}

func (r *HeuristicsModelerRepositoryImpl) ListModelerPager(filter dtoquery.QueryFilter, offset int, limit int) ([]model.HeuristicsModeler, int64, error) {
	var modelers []model.HeuristicsModeler
	var total int64

	q := r.DB.Model(&model.HeuristicsModeler{}).Scopes(helperquery.WithDynamicFilters(filter))
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.Order("created_at DESC, id DESC").Limit(limit).Offset(offset).Find(&modelers).Error; err != nil {
		return nil, 0, err
	}

	return modelers, total, nil
}

func (r *HeuristicsModelerRepositoryImpl) UpdateModeler(id string, modeler *model.HeuristicsModeler) error {
  return r.DB.Model(&model.HeuristicsModeler{}).Where("id = ?", id).Updates(modeler).Error
}

func (r *HeuristicsModelerRepositoryImpl) DeleteModeler(id string) error {
  return r.DB.Delete(&model.HeuristicsModeler{}, id).Error
}
