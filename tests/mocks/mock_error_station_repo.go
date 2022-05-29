package mocks

import (
	"errors"

	"github.com/Israel-Ferreira/metrosp-api/models"
)

type MockErrorStationRepo struct {
}

func (mckStation *MockErrorStationRepo) FindAll() ([]models.Station, error) {
	return nil, errors.New("error")
}

func (mckStation *MockErrorStationRepo) FindById(id uint64) (models.Station, error) {
	return models.Station{}, errors.New("error: estação não encontrada")
}

func (mckStation *MockErrorStationRepo) Create(line models.Station) (models.Station, error) {
	return models.Station{}, errors.New("error: erro ao salvar a estação")
}

func (mckStation *MockErrorStationRepo) Delete(_ uint64) error {
	return errors.New("erro: não implementado")
}

func (mckStation *MockErrorStationRepo) Update(id uint64, station models.Station) error {
	return errors.New("erro: falha ao atualizar a linha ")
}
