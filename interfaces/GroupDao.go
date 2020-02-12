package interfaces

import (
	"subasic/models"
)

type GroupDao interface {
	Create(g *models.Group) error
	GetAll() ([]models.Group, error)
	Update(g *models.Group) error
	GetById(i int)(models.Group, error)
	Delete(i int) error
}
