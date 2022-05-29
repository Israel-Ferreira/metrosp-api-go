package mocks

import (
	"errors"

	"github.com/Israel-Ferreira/metrosp-api/exceptions"
	"github.com/Israel-Ferreira/metrosp-api/models"
)

type MockErrorLineRepo struct {
}

func (mckLine *MockErrorLineRepo) FindAll() ([]models.Line, error) {
	return nil, errors.New("erro")
}

func (mckLine *MockErrorLineRepo) FindById(id uint64) (models.Line, error) {
	return models.Line{}, exceptions.ErrorNotFound
}

func (mckLine *MockErrorLineRepo) Create(line models.Line) (models.Line, error) {
	return models.Line{}, errors.New("erro")
}

func (mckLine *MockErrorLineRepo) FindByLineNumber(lineNumber uint) (models.Line, error) {
	var line models.Line

	return line, errors.New("erro ao fazer a consulta")
}

func (mckLine *MockErrorLineRepo) Update(_ uint64, _ models.Line) error {
	return errors.New("erro ao fazer o update")
}
