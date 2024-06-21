package services

import (
	"Perico6255/task-go-rest/internal/models"
	"Perico6255/task-go-rest/internal/repository"
)

type UserService interface {
	Create(user models.UserModel) (models.UserModel, error)
	FindById(id uint) (models.UserModel, error)
	FindAll() ([]models.UserModel, error)
	Delete(id uint) (models.UserModel, error)
	Update(user models.UserModel) (models.UserModel, error)
}
type userService struct {
	r repository.UserRepository
}

// Delete implements UserService.
func (s *userService) Delete(id uint) (models.UserModel, error) {
	return s.r.Delete(id)
}

// FindAll implements UserService.
func (s *userService) FindAll() ([]models.UserModel, error) {
	return s.r.FindAll()
}

// Update implements UserService.
func (s *userService) Update(user models.UserModel) (models.UserModel, error) {
	return s.r.Update(user)
}

// FindById implements UserService.
func (s *userService) FindById(id uint) (models.UserModel, error) {
	return s.r.FindById(id)
}

// Create implements UserService.
func (s *userService) Create(user models.UserModel) (models.UserModel, error) {
	return s.r.Create(user)
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}
