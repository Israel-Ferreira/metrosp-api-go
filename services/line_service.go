package services

import (
	"github.com/Israel-Ferreira/metrosp-api/data"
	"github.com/Israel-Ferreira/metrosp-api/exceptions"
	"github.com/Israel-Ferreira/metrosp-api/models"
	"github.com/Israel-Ferreira/metrosp-api/repo"
)

type LineService interface {
	FindAll() ([]models.Line, error)
	FindById(id string) (models.Line, error)
	Create(data.LineDTO) (models.Line, error)
	Update(string, data.LineDTO) error
}

type lineService struct {
	lineRepo repo.LineRepo
}

func (ls *lineService) FindAll() ([]models.Line, error) {
	lines, err := ls.lineRepo.FindAll()

	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (ls *lineService) FindById(id string) (models.Line, error) {
	intId, err := parseIdToUint64(id)

	if err != nil {
		return models.Line{}, err
	}

	line, err := ls.lineRepo.FindById(intId)

	if err != nil {
		return models.Line{}, exceptions.ErrorNotFound
	}

	return line, nil

}

func (ls *lineService) Create(dto data.LineDTO) (models.Line, error) {

	if err := dto.IsValid(); err != nil {
		return models.Line{}, exceptions.ErrorValidation
	}

	lineNumber, err := parseIdToUint64(dto.Number)

	if err != nil {
		return models.Line{}, err
	}

	line, err := ls.lineRepo.FindByLineNumber(uint(lineNumber))

	if err != nil && err.Error() != "record not found" {
		return models.Line{}, err
	}

	if line.ID != 0 {
		return models.Line{}, exceptions.ErrorLineExists
	}

	newLine := models.NewLine(dto)

	newLine, err = ls.lineRepo.Create(newLine)

	if err != nil {
		return models.Line{}, err
	}

	return newLine, nil
}

func (ls *lineService) Update(id string, dto data.LineDTO) error {
	return nil
}

func NewLineService(repository repo.LineRepo) LineService {
	return &lineService{
		lineRepo: repository,
	}
}
