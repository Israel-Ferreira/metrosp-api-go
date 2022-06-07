package mocks

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/Israel-Ferreira/metrosp-api/models"
	"github.com/Israel-Ferreira/metrosp-api/repo"
)

type MockSuccessLineRepo struct {
	lines []models.Line
}

func (mckLine *MockSuccessLineRepo) FindAll() ([]models.Line, error) {
	return mckLine.lines, nil
}

func (mckLine *MockSuccessLineRepo) FindById(id uint64) (models.Line, error) {
	var lineIsFound bool
	var subwayLine models.Line

	for _, line := range mckLine.lines {
		if line.ID == uint(id) {
			lineIsFound = true
			subwayLine = line
		}
	}

	if !lineIsFound {
		return models.Line{}, errors.New("erro: Linha não encontrada")
	}

	return subwayLine, nil

}

func (mckLine *MockSuccessLineRepo) Create(line models.Line) (models.Line, error) {
	line.ID = uint(rand.Uint64())
	return line, nil
}

func (mckLine *MockSuccessLineRepo) Delete(_ uint64) error {
	return nil
}

func (mckLine *MockSuccessLineRepo) FindByLineNumber(lineNumber uint) (models.Line, error) {
	var line models.Line

	for _, lin := range mckLine.lines {
		if lin.Number == lineNumber {
			line = lin
		}
	}

	return line, nil
}

func (mckLine *MockSuccessLineRepo) Update(id uint64, line models.Line) error {
	linha, err := mckLine.FindById(id)

	if err != nil {
		return err
	}

	fmt.Println(line)

	linha.MapImgUrl = line.MapImgUrl

	return nil
}

func NewMockLineRepo(errorMock bool) repo.LineRepo {
	if errorMock {
		return &MockErrorLineRepo{}
	}

	return &MockSuccessLineRepo{
		lines: linhas,
	}
}
