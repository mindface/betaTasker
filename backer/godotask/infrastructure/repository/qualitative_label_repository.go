package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/helper"
)

type QualitativeLabelRepositoryImpl struct {
	DB *gorm.DB
}

func (r *QualitativeLabelRepositoryImpl) Create(qualitativeLabel *model.QualitativeLabel) error {
	return r.DB.Create(qualitativeLabel).Error
}

func (r *QualitativeLabelRepositoryImpl) FindByID(id string) (*model.QualitativeLabel, error) {
	var qualitativeLabel model.QualitativeLabel
	if err := r.DB.Where("id = ?", id).First(&qualitativeLabel).Error; err != nil {
		return nil, err
	}
	return &qualitativeLabel, nil
}

func (r *QualitativeLabelRepositoryImpl) FindAll(userID uint) ([]model.QualitativeLabel, error) {
	var qualitativeLabels []model.QualitativeLabel
	if err := r.DB.Scopes(helper.WithUserFilter(userID)).Order("created_at DESC, id DESC").Find(&qualitativeLabels).Error; err != nil {
		return nil, err
	}
	return qualitativeLabels, nil
}

func (r *QualitativeLabelRepositoryImpl) Update(id string, qualitativeLabel *model.QualitativeLabel) error {
	return r.DB.Model(&model.QualitativeLabel{}).Where("id = ?", id).Updates(qualitativeLabel).Error
}

func (r *QualitativeLabelRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.QualitativeLabel{}, id).Error
}

// NewQualitativeLabelRepository は QualitativeLabelRepositoryInterface を返すコンストラクタ
func NewQualitativeLabelRepository(db *gorm.DB) QualitativeLabelRepositoryInterface {
	return &QualitativeLabelRepositoryImpl{DB: db}
}