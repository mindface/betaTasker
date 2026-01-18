package repository

import (
	"strconv"
	"github.com/godotask/model"
	"gorm.io/gorm"
)

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

func (r *HeuristicsRepositoryImpl) FindAllAnalyses() ([]model.HeuristicsAnalysis, error) {
    var analyses []model.HeuristicsAnalysis
    if err := r.DB.Find(&analyses).Error; err != nil {
        return nil, err
    }
    return analyses, nil
}

func (r *HeuristicsRepositoryImpl) UpdateAnalysis(id string, analysis *model.HeuristicsAnalysis) error {
    return r.DB.Model(&model.HeuristicsAnalysis{}).Where("id = ?", id).Updates(analysis).Error
}

func (r *HeuristicsRepositoryImpl) DeleteAnalysis(id string) error {
    return r.DB.Delete(&model.HeuristicsAnalysis{}, id).Error
}

func (r *HeuristicsRepositoryImpl) ListAnalyses() ([]model.HeuristicsAnalysis, error) {
    var analyses []model.HeuristicsAnalysis
    if err := r.DB.Find(&analyses).Error; err != nil {
        return nil, err
    }
    return analyses, nil
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