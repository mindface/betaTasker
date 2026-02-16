package repository

import (
	"strconv"
	"github.com/godotask/infrastructure/db/model"
	dtoquery "github.com/godotask/dto/query"
 "github.com/godotask/infrastructure/helper"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"gorm.io/gorm"
)

type HeuristicsPatternRepositoryImpl struct {
	DB *gorm.DB
}

func (r *HeuristicsPatternRepositoryImpl) CreatePattern(pattern *model.HeuristicsPattern) error {
    return r.DB.Create(pattern).Error
}

func (r *HeuristicsPatternRepositoryImpl) GetPatternById(id string) (*model.HeuristicsPattern, error) {
	var pattern model.HeuristicsPattern
	if err := r.DB.First(&pattern, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pattern, nil
}

func (r *HeuristicsPatternRepositoryImpl) ListPattern(userID uint) ([]model.HeuristicsPattern, error) {
	var patterns []model.HeuristicsPattern
	if err := r.DB.Scopes(helper.WithUserFilter(userID)).Order("created_at DESC, id DESC").Find(&patterns).Error; err != nil {
		return nil, err
	}
	return patterns, nil
}

func (r *HeuristicsPatternRepositoryImpl) ListPatternPager(filter dtoquery.QueryFilter, offset int, limit int) ([]model.HeuristicsPattern, int64, error) {
	var patterns []model.HeuristicsPattern
	var total int64

	q := r.DB.Model(&model.HeuristicsPattern{}).Scopes(helperquery.WithDynamicFilters(filter))
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.Order("created_at DESC, id DESC").Limit(limit).Offset(offset).Find(&patterns).Error; err != nil {
		return nil, 0, err
	}

	return patterns, total, nil
}

func (r *HeuristicsPatternRepositoryImpl) GetPatterns(userID string, limit, offset int) ([]model.HeuristicsPattern, int, error) {
	var patterns []model.HeuristicsPattern
	var total int64

	query := r.DB.Model(&model.HeuristicsPattern{})
	if userID != "" {
		uid, _ := strconv.Atoi(userID)
		query = query.Where("user_id = ?", uid)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(limit).Offset(offset).Find(&patterns).Error; err != nil {
		return nil, 0, err
	}

	return patterns, int(total), nil
}

func (r *HeuristicsPatternRepositoryImpl) UpdatePattern(id string, pattern *model.HeuristicsPattern) error {
  return r.DB.Model(&model.HeuristicsPattern{}).Where("id = ?", id).Updates(pattern).Error
}

func (r *HeuristicsPatternRepositoryImpl) DeletePattern(id string) error {
  return r.DB.Delete(&model.HeuristicsPattern{}, id).Error
}