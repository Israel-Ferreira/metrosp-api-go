package services

import (
	"strconv"

	"github.com/Israel-Ferreira/metrosp-api/data"
	"github.com/Israel-Ferreira/metrosp-api/models"
	"github.com/Israel-Ferreira/metrosp-api/repo"
)

func parseIdToUint64(id string) (intId uint64, err error) {
	intId, err = strconv.ParseUint(id, 10, 64)

	if err != nil {
		return 0, err
	}

	return intId, nil
}

type StationService interface {
	FindAll() ([]models.Station, error)
	FindById(id string) (models.Station, error)
	Create(data.StationDTO) (models.Station, error)
	Delete(id string) error
	Update(string, data.StationDTO) error
}

type stationService struct {
	repository     repo.StationRepo
	lineRepository repo.LineRepo
}

func (svc *stationService) FindAll() ([]models.Station, error) {
	stations, err := svc.repository.FindAll()

	if err != nil {
		return nil, err
	}

	return stations, nil
}

func (svc *stationService) FindById(id string) (models.Station, error) {
	intId, err := parseIdToUint64(id)

	if err != nil {
		return *new(models.Station), err
	}

	station, err := svc.repository.FindById(intId)

	if err != nil {
		return *new(models.Station), err
	}

	return station, nil
}

func (svc *stationService) Create(dto data.StationDTO) (models.Station, error) {

	if err := dto.IsValid(); err != nil {
		return models.Station{}, err
	}

	line, err := svc.lineRepository.FindByLineNumber(uint(dto.LineNumber))

	if err != nil {
		return models.Station{}, err
	}

	station := models.NewStation(dto, line.ID)

	station, err = svc.repository.Create(station)

	if err != nil {
		return models.Station{}, err
	}

	return station, nil
}

func (svc *stationService) Delete(id string) error {
	intId, err := parseIdToUint64(id)

	if err != nil {
		return err
	}

	if err := svc.repository.Delete(intId); err != nil {
		return err
	}

	return nil
}

func (svc *stationService) Update(id string, dto data.StationDTO) error {
	intId, err := parseIdToUint64(id)

	if err != nil {
		return err
	}

	station, err := svc.repository.FindById(intId)

	if err != nil {
		return err
	}

	if err := dto.IsValid(); err != nil {
		return err
	}

	station.Name = dto.Name
	station.ImgUrl = dto.ImagesUrls[0]
	station.InaugurationDate = dto.InaugurationDate

	if err := svc.repository.Update(intId, station); err != nil {
		return err
	}

	return nil
}

func NewStationService(stationRepo repo.StationRepo, lineRepo repo.LineRepo) StationService {
	return &stationService{
		repository:     stationRepo,
		lineRepository: lineRepo,
	}
}

func NewEmptyStationService() StationService {
	return &stationService{}
}
