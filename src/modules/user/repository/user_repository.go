package repository

import "crowdfunding-api/src/modules/user/model/entity"

type UserRepository interface {
	Create(user entity.User) error
	Update(user entity.User, userId uint) error
	UpdateRole(user entity.User, userId uint) error
	Delete(userId uint, deleteBy string) error
	FindAll() ([]entity.User, error)
	FindById(userId uint) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}
