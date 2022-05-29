package repo

import "github.com/Israel-Ferreira/metrosp-api/models"

type LineRepo interface {
	FindAll() ([]models.Line, error)
	FindById(id uint64) (models.Line, error)
	Create(models.Line) (models.Line, error)
	Delete(uint64) error
	Update(uint64, models.Line) error
}
