package repository

import (
	"github.com/godotask/infrastructure/db/model"
	dtoquery "github.com/godotask/dto/query"
	helperquery "github.com/godotask/infrastructure/helper/query"
	"gorm.io/gorm"
)

type HeuristicsAnalysisRepositoryImpl struct {
	DB *gorm.DB
}

func (r *HeuristicsAnalysisRepositoryImpl) CreateAnalysis(analysis *model.HeuristicsAnalysis) error {
  return r.DB.Create(analysis).Error
}

func (r *HeuristicsAnalysisRepositoryImpl) GetAnalysisById(id string) (*model.HeuristicsAnalysis, error) {
	var analysis model.HeuristicsAnalysis
	if err := r.DB.First(&analysis, id).Error; err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *HeuristicsAnalysisRepositoryImpl) ListAnalyze() ([]model.HeuristicsAnalysis, error) {
	var analyses []model.HeuristicsAnalysis
	if err := r.DB.Find(&analyses).Error; err != nil {
			return nil, err
	}
	return analyses, nil
}

func (r *HeuristicsAnalysisRepositoryImpl) ListAnalysesPager(filter dtoquery.QueryFilter, offset int, limit int) ([]model.HeuristicsAnalysis, int64, error) {
	var analyses []model.HeuristicsAnalysis
	var total int64

	q := r.DB.Model(&model.HeuristicsAnalysis{}).Scopes(helperquery.WithDynamicFilters(filter))
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.Order("created_at DESC, id DESC").Limit(limit).Offset(offset).Find(&analyses).Error; err != nil {
		return nil, 0, err
	}

	return analyses, total, nil
}

func (r *HeuristicsAnalysisRepositoryImpl) UpdateAnalysis(id string, analysis *model.HeuristicsAnalysis) error {
    return r.DB.Model(&model.HeuristicsAnalysis{}).Where("id = ?", id).Updates(analysis).Error
}

func (r *HeuristicsAnalysisRepositoryImpl) DeleteAnalysis(id string) error {
    return r.DB.Delete(&model.HeuristicsAnalysis{}, id).Error
}

func (r *HeuristicsAnalysisRepositoryImpl) ListAnalyses() ([]model.HeuristicsAnalysis, error) {
    var analyses []model.HeuristicsAnalysis
    if err := r.DB.Find(&analyses).Error; err != nil {
        return nil, err
    }
    return analyses, nil
}
