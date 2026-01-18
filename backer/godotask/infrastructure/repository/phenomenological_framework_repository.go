package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/helper"
)

type PhenomenologicalFrameworkRepositoryImpl struct {
	DB *gorm.DB
}

func (r *PhenomenologicalFrameworkRepositoryImpl) Create(phenomenologicalFramework *model.PhenomenologicalFramework) error {
	return r.DB.Create(phenomenologicalFramework).Error
}

func (r *PhenomenologicalFrameworkRepositoryImpl) FindByID(id string) (*model.PhenomenologicalFramework, error) {
	var phenomenologicalFramework model.PhenomenologicalFramework
	if err := r.DB.Where("id = ?", id).First(&phenomenologicalFramework).Error; err != nil {
		return nil, err
	}
	return &phenomenologicalFramework, nil
}

func (r *PhenomenologicalFrameworkRepositoryImpl) FindAll(userID uint) ([]model.PhenomenologicalFramework, error) {
	var phenomenologicalFrameworks []model.PhenomenologicalFramework
	if err := r.DB.Scopes(helper.WithUserFilter(userID)).Order("created_at DESC, id DESC").Find(&phenomenologicalFrameworks).Error; err != nil {
		return nil, err
	}
	return phenomenologicalFrameworks, nil
}

func (r *PhenomenologicalFrameworkRepositoryImpl) Update(id string, phenomenologicalFramework *model.PhenomenologicalFramework) error {
	return r.DB.Model(&model.PhenomenologicalFramework{}).Where("id = ?", id).Updates(phenomenologicalFramework).Error
}

func (r *PhenomenologicalFrameworkRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.PhenomenologicalFramework{}, id).Error
}

// NewPhenomenologicalFrameworkRepository は PhenomenologicalFrameworkRepositoryInterface を返すコンストラクタ
func NewPhenomenologicalFrameworkRepository(db *gorm.DB) PhenomenologicalFrameworkRepositoryInterface {
	return &PhenomenologicalFrameworkRepositoryImpl{DB: db}
}