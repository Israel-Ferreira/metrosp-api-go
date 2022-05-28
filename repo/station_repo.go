package repo

import "github.com/Israel-Ferreira/metrosp-api/models"

type StationRepo interface {
	FindAll() ([]models.Station, error)
	FindById(id uint64) (models.Station, error)
	Create(models.Station) (models.Station, error)
	Delete(id uint64) error
	Update(uint64, models.Station) error
}
