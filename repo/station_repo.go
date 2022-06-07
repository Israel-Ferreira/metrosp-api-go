package repo

import (
	"github.com/Israel-Ferreira/metrosp-api/models"
	"gorm.io/gorm"
)

type StationRepo interface {
	FindAll() ([]models.Station, error)
	FindById(id uint64) (models.Station, error)
	Create(models.Station) (models.Station, error)
	Delete(id uint64) error
	Update(uint64, models.Station) error
}

type DbStationRepo struct {
	db *gorm.DB
}

func (dbRepo DbStationRepo) FindAll() ([]models.Station, error) {
	var stations []models.Station

	if txn := dbRepo.db.Find(&stations); txn.Error != nil {
		return nil, txn.Error
	}

	return stations, nil
}

func (dbRepo DbStationRepo) FindById(id uint64) (models.Station, error) {
	var station models.Station

	if txn := dbRepo.db.Find(&station, "id = ?", id); txn.Error != nil {
		return models.Station{}, txn.Error
	}

	return station, nil
}

func (dbRepo DbStationRepo) Create(station models.Station) (models.Station, error) {

	if txn := dbRepo.db.Save(&station); txn.Error != nil {
		return models.Station{}, txn.Error
	}

	return station, nil
}

func (dbRepo DbStationRepo) Delete(_ uint64) error {
	return nil
}

func (dbRepo DbStationRepo) Update(id uint64, station models.Station) error {
	txn := dbRepo.db.Save(&station)

	if txn.Error != nil {
		return txn.Error
	}

	return nil
}

func NewStationDbRepo(db *gorm.DB) StationRepo {
	return DbStationRepo{db: db}
}
