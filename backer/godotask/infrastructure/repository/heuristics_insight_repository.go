package repository

import (
	"strconv"
	"github.com/godotask/model"
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

func (r *HeuristicsInsightRepositoryImpl) GetInsights(userID string, limit, offset int) ([]model.HeuristicsInsight, int, error) {
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

func (r *HeuristicsInsightRepositoryImpl) UpdateInsight(id string, insight *model.HeuristicsInsight) error {
    return r.DB.Model(&model.HeuristicsInsight{}).Where("id = ?", id).Updates(insight).Error
}

func (r *HeuristicsInsightRepositoryImpl) DeleteInsight(id string) error {
    return r.DB.Delete(&model.HeuristicsInsight{}, id).Error
}