package interfaces

import (
	"subasic/models"
)

type LocationDao interface {
	Create(l *models.Location) error
	GetAll() ([]models.Location, error)
	Update(l *models.Location) error
	GetById(i int)(models.Location, error)
	Delete(i int) error
}
