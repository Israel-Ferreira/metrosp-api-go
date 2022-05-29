package mocks

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Israel-Ferreira/metrosp-api/models"
	"github.com/Israel-Ferreira/metrosp-api/repo"
)

type MockSuccessStationRepo struct {
	stations []models.Station
}

func (mckStation *MockSuccessStationRepo) FindAll() ([]models.Station, error) {
	return mckStation.stations, nil
}

func (mckStation *MockSuccessStationRepo) FindById(id uint64) (models.Station, error) {
	var stationIsFound bool
	var subwayStation models.Station

	for _, station := range mckStation.stations {
		if station.ID == id {
			stationIsFound = true
			subwayStation = station
		}
	}

	if !stationIsFound {
		return models.Station{}, errors.New("erro: Linha não encontrada")
	}

	return subwayStation, nil

}

func (mckStation *MockSuccessStationRepo) Create(line models.Station) (models.Station, error) {
	line.ID = rand.Uint64()
	return line, nil
}

func (mckStation *MockSuccessStationRepo) Delete(_ uint64) error {
	return nil
}

func (mckStation *MockSuccessStationRepo) Update(id uint64, station models.Station) error {
	estacao, err := mckStation.FindById(id)

	if err != nil {
		return err
	}

	fmt.Println(estacao)

	fmt.Println(station)

	return nil
}

func NewMockStationRepo(errorMock bool) repo.StationRepo {
	if errorMock {
		return &MockErrorStationRepo{}
	}

	return &MockSuccessStationRepo{
		stations: estacoes,
	}
}
