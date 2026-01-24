package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/infrastructure/db/model"
	"github.com/godotask/infrastructure/helper"
)

type TeachingFreeControlRepositoryImpl struct {
	DB *gorm.DB
}

func (r *TeachingFreeControlRepositoryImpl) Create(teachingFreeControl *model.TeachingFreeControl) error {
	return r.DB.Create(teachingFreeControl).Error
}

func (r *TeachingFreeControlRepositoryImpl) FindByID(id string) (*model.TeachingFreeControl, error) {
	var teachingFreeControl model.TeachingFreeControl
	if err := r.DB.Where("id = ?", id).First(&teachingFreeControl).Error; err != nil {
		return nil, err
	}
	return &teachingFreeControl, nil
}

func (r *TeachingFreeControlRepositoryImpl) FindAll(userID uint) ([]model.TeachingFreeControl, error) {
	var teachingFreeControls []model.TeachingFreeControl
	if err := r.DB.Scopes(helper.WithUserFilter(userID)).Order("created_at DESC, id DESC").Find(&teachingFreeControls).Error; err != nil {
		return nil, err
	}
	return teachingFreeControls, nil
}

func (r *TeachingFreeControlRepositoryImpl) Update(id string, teachingFreeControl *model.TeachingFreeControl) error {
	return r.DB.Model(&model.TeachingFreeControl{}).Where("id = ?", id).Updates(teachingFreeControl).Error
}

func (r *TeachingFreeControlRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.TeachingFreeControl{}, id).Error
}

// NewTeachingFreeControlRepository は TeachingFreeControlRepositoryInterface を返すコンストラクタ
func NewTeachingFreeControlRepository(db *gorm.DB) TeachingFreeControlRepositoryInterface {
	return &TeachingFreeControlRepositoryImpl{DB: db}
}