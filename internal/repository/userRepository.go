package repository

import (
	"Perico6255/task-go-rest/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.UserModel) (models.UserModel, error)
	FindById(id uint) (models.UserModel, error)
}
type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) FindById(id uint) (models.UserModel, error) {
  var user models.UserModel
  err:= r.db.Find(&user,id ).Error
  return user, err
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user models.UserModel) (models.UserModel, error) {
	err := r.db.Create(&user).Error
	return user, err
}
