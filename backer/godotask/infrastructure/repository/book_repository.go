package repository

import (
	"gorm.io/gorm"
	"github.com/godotask/model"
	"github.com/godotask/infrastructure/helper"
)

// BookRepositoryImpl
type BookRepositoryImpl struct {
	DB *gorm.DB
}

func (r *BookRepositoryImpl) Create(book *model.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepositoryImpl) FindByID(id string) (*model.Book, error) {
  var book model.Book
  if err := r.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepositoryImpl) FindAll(userID uint) ([]model.Book, error) {
	var books []model.Book
	if err := r.DB.Scopes(helper.WithUserFilter(userID)).Order("created_at DESC, id DESC").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepositoryImpl) FindAllPager(userID uint, offset, limit int) ([]model.Book, int64, error) {
    var books []model.Book
    var total int64

    q := r.DB.Model(&model.Book{}).Scopes(helper.WithUserFilter(userID))

    if err := q.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    if err := q.Order("created_at DESC, id DESC").Limit(limit).Offset(offset).Find(&books).Error; err != nil {
        return nil, 0, err
    }
    return books, total, nil
}

func (r *BookRepositoryImpl) Update(id string, book *model.Book) error {
	return r.DB.Model(&model.Book{}).Where("id = ?", id).Updates(book).Error
}

func (r *BookRepositoryImpl) Delete(id string) error {
	return r.DB.Delete(&model.Book{}, id).Error
}

// NewBookRepository は BookRepositoryInterface を返すコンストラクタ
func NewBookRepository(db *gorm.DB) BookRepositoryInterface {
	return &BookRepositoryImpl{DB: db}
}