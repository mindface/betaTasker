package repository

import (
	"strconv"
	"github.com/godotask/model"
	"gorm.io/gorm"
)

type HeuristicsRepository interface {
	CreateAnalysis(analysis *model.HeuristicsAnalysis) error
	GetAnalysisById(id string) (*model.HeuristicsAnalysis, error)
	CreateTracking(tracking *model.HeuristicsTracking) error
	GetTrackingByUserID(userID string) ([]model.HeuristicsTracking, error)
	GetInsights(userID string, limit, offset int) ([]model.HeuristicsInsight, int, error)
	GetInsightById(id string) (*model.HeuristicsInsight, error)
	DetectPatterns(userID, dataType, period string) ([]model.HeuristicsPattern, error)
	CreateModel(model *model.HeuristicsModel) error
}

type HeuristicsRepositoryImpl struct {
	DB *gorm.DB
}

func (r *HeuristicsRepositoryImpl) CreateAnalysis(analysis *model.HeuristicsAnalysis) error {
	return r.DB.Create(analysis).Error
}

func (r *HeuristicsRepositoryImpl) GetAnalysisById(id string) (*model.HeuristicsAnalysis, error) {
	var analysis model.HeuristicsAnalysis
	if err := r.DB.First(&analysis, id).Error; err != nil {
		return nil, err
	}
	return &analysis, nil
}

func (r *HeuristicsRepositoryImpl) CreateTracking(tracking *model.HeuristicsTracking) error {
	return r.DB.Create(tracking).Error
}

func (r *HeuristicsRepositoryImpl) GetTrackingByUserID(userID string) ([]model.HeuristicsTracking, error) {
	var trackings []model.HeuristicsTracking
	uid, _ := strconv.Atoi(userID)
	if err := r.DB.Where("user_id = ?", uid).Find(&trackings).Error; err != nil {
		return nil, err
	}
	return trackings, nil
}

func (r *HeuristicsRepositoryImpl) GetInsights(userID string, limit, offset int) ([]model.HeuristicsInsight, int, error) {
	var insights []model.HeuristicsInsight
	var total int64
	
	query := r.DB.Model(&model.HeuristicsInsight{})
	if userID != "" {
		uid, _ := strconv.Atoi(userID)
		query = query.Where("user_id = ?", uid)
	}
	
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	
	if err := query.Limit(limit).Offset(offset).Find(&insights).Error; err != nil {
		return nil, 0, err
	}
	
	return insights, int(total), nil
}

func (r *HeuristicsRepositoryImpl) GetInsightById(id string) (*model.HeuristicsInsight, error) {
	var insight model.HeuristicsInsight
	if err := r.DB.First(&insight, id).Error; err != nil {
		return nil, err
	}
	return &insight, nil
}

func (r *HeuristicsRepositoryImpl) DetectPatterns(userID, dataType, period string) ([]model.HeuristicsPattern, error) {
	var patterns []model.HeuristicsPattern
	query := r.DB.Model(&model.HeuristicsPattern{})
	
	// 簡単なフィルタリング（実際のプロジェクトではより複雑なロジック）
	if dataType != "all" {
		query = query.Where("category = ?", dataType)
	}
	
	if err := query.Find(&patterns).Error; err != nil {
		return nil, err
	}
	
	return patterns, nil
}

func (r *HeuristicsRepositoryImpl) CreateModel(modelData *model.HeuristicsModel) error {
	return r.DB.Create(modelData).Error
}