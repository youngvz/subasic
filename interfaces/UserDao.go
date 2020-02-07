package interfaces

import (
	"subasic/models"
)

type UserDao interface {
	Create(u *models.User) error
	GetAll() ([]models.User, error)
	Update(u *models.User) error
	GetById(i int) (models.User, error)
	Delete(i int) error
}