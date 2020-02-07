package interfaces

import (
	"subasic/models"
)

type GroupDao interface {
	Create(u *models.Group) error
	GetAll() ([]models.Group, error)
	Update(u *models.Group) error
	GetById(i int)(models.Group, error)
	Delete(i int) error
}
