package repository

import (
	"github.com/godotask/domain/entity"
	domainRepo "github.com/godotask/domain/repository"
	"gorm.io/gorm"
)

type UserRepositoryGorm struct {
	db *gorm.DB
}

// compile-time check（強く推奨）
var _ domainRepo.UserRepository = (*UserRepositoryGorm)(nil)

func NewGormUserRepository(db *gorm.DB) domainRepo.UserRepository {
	return &UserRepositoryGorm{db: db}
}

func (r *UserRepositoryGorm) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepositoryGorm) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryGorm) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
