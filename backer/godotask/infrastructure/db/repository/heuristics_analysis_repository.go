package repository

import (
	"strconv"
	"github.com/godotask/infrastructure/db/model"
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

func (r *HeuristicsAnalysisRepositoryImpl) FindAllAnalyses() ([]model.HeuristicsAnalysis, error) {
    var analyses []model.HeuristicsAnalysis
    if err := r.DB.Find(&analyses).Error; err != nil {
        return nil, err
    }
    return analyses, nil
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


func (r *HeuristicsAnalysisRepositoryImpl) CreateTracking(tracking *model.HeuristicsTracking) error {
	return r.DB.Create(tracking).Error
}

func (r *HeuristicsAnalysisRepositoryImpl) GetTrackingByUserID(userID string) ([]model.HeuristicsTracking, error) {
	var trackings []model.HeuristicsTracking
	uid, _ := strconv.Atoi(userID)
	if err := r.DB.Where("user_id = ?", uid).Find(&trackings).Error; err != nil {
		return nil, err
	}
	return trackings, nil
}

func (r *HeuristicsAnalysisRepositoryImpl) DetectPatterns(userID, dataType, period string) ([]model.HeuristicsPattern, error) {
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

func (r *HeuristicsAnalysisRepositoryImpl) CreateModel(modelData *model.HeuristicsModel) error {
	return r.DB.Create(modelData).Error
}