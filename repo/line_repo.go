package repo

import (
	"github.com/Israel-Ferreira/metrosp-api/models"
	"gorm.io/gorm"
)

type LineRepo interface {
	FindAll() ([]models.Line, error)
	FindById(id uint64) (models.Line, error)
	FindByLineNumber(uint) (models.Line, error)
	Create(models.Line) (models.Line, error)
	Update(uint64, models.Line) error
}

type DbLineRepo struct {
	db *gorm.DB
}

func (repo *DbLineRepo) FindAll() ([]models.Line, error) {

	var lines []models.Line

	txn := repo.db.Find(&lines)

	if txn.Error != nil {
		return nil, txn.Error
	}

	return lines, nil
}

func (repo *DbLineRepo) FindById(id uint64) (models.Line, error) {

	subwayLine := models.Line{Model: gorm.Model{ID: uint(id)}}

	txn := repo.db.Preload("Stations").First(&subwayLine)

	if txn.Error != nil {
		return models.Line{}, txn.Error
	}

	return subwayLine, nil
}

func (repo *DbLineRepo) Create(line models.Line) (models.Line, error) {
	txn := repo.db.Save(&line)

	if txn.Error != nil {
		return models.Line{}, txn.Error
	}

	return line, nil
}

func (repo *DbLineRepo) Delete(_ uint64) error {
	return nil
}

func (repo *DbLineRepo) Update(id uint64, line models.Line) error {
	// linha, err := repo.FindById(id)

	return nil

}

func (repo *DbLineRepo) FindByLineNumber(lineNumber uint) (models.Line, error) {
	var line models.Line

	result := repo.db.Where(&models.Line{Number: line.Number}).First(&line)

	if result.Error != nil {
		return models.Line{}, result.Error
	}

	return line, nil
}

func NewLineRepo(db *gorm.DB) LineRepo {
	return &DbLineRepo{
		db: db,
	}
}
