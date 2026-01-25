package repository

import "github.com/godotask/domain/entity"

// type UserRepository interface {
// 	Create(user *entity.User) error
// 	FindByEmail(email string) (*entity.User, error)
// }

type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
}