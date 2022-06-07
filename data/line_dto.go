package data

import "errors"

type LineDTO struct {
	Color       string  `json:"color"`
	Number      int     `json:"number"`
	MapImageUrl string  `json:"mapImageUrl"`
	Extension   float64 `json:"extension"`
}

func (dto LineDTO) IsValid() error {
	if dto.Color == "" {
		return errors.New("erro: a cor da linha não pode estar vazia")
	}

	if dto.Number <= 0 {
		return errors.New("erro:  O número não pode estar vazio")
	}

	if dto.Extension <= 0.0 {
		return errors.New("erro: a extensão não pode ser menor ou igual a 0 KM")
	}

	return nil
}
