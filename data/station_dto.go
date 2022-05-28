package data

import "errors"

type StationDTO struct {
	Name             string   `json:"name"`
	LineNumber       int      `json:"lineNumber"`
	InaugurationDate string   `json:"inaugurationDate"`
	Street           string   `json:"street"`
	Neighborhood     string   `json:"neighborhood"`
	ImagesUrls       []string `json:"imagesUrl"`
}

func (dto StationDTO) IsValid() error {
	if dto.Name == "" {
		return errors.New("erro: o nome da estação não pode estar em branco")
	}

	if dto.LineNumber <= 0 {
		return errors.New("erro: Número da Linha tem que ser maior que 0")
	}

	return nil
}
