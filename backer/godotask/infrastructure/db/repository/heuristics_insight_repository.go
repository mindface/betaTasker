package repository

import (
	"github.com/godotask/infrastructure/db/model"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"gorm.io/gorm"
)

type HeuristicsInsightRepositoryImpl struct {
	DB *gorm.DB
}

func (r *HeuristicsInsightRepositoryImpl) CreateInsight(insight *model.HeuristicsInsight) error {
  return r.DB.Create(insight).Error
}

func (r *HeuristicsInsightRepositoryImpl) GetInsightById(id string) (*model.HeuristicsInsight, error) {
	var insight model.HeuristicsInsight
	if err := r.DB.First(&insight, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &insight, nil
}

func (r *HeuristicsInsightRepositoryImpl) ListInsight() ([]model.HeuristicsInsight, error) {
	var insights []model.HeuristicsInsight

	if err := r.DB.Find(&insights).Error; err != nil {
		return nil, err
	}
	return insights, nil
}

func (r *HeuristicsInsightRepositoryImpl) ListInsightPager(filter dtoquery.QueryFilter, offset int, limit int) ([]model.HeuristicsInsight, int64, error) {
	var insights []model.HeuristicsInsight
	var total int64

	q := r.DB.Model(&model.HeuristicsInsight{}).Scopes(helperquery.WithDynamicFilters(filter))
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.Order("created_at DESC, id DESC").Limit(limit).Offset(offset).Find(&insights).Error; err != nil {
		return nil, 0, err
	}

	return insights, total, nil
}

func (r *HeuristicsInsightRepositoryImpl) UpdateInsight(id string, insight *model.HeuristicsInsight) error {
  return r.DB.Model(&model.HeuristicsInsight{}).Where("id = ?", id).Updates(insight).Error
}

func (r *HeuristicsInsightRepositoryImpl) DeleteInsight(id string) error {
  return r.DB.Delete(&model.HeuristicsInsight{}, id).Error
}