package models

import (
	"github.com/Israel-Ferreira/metrosp-api/data"
	"gorm.io/gorm"
)

type Station struct {
	gorm.Model
	Name             string `json:"name"`
	Neighborhood     string `json:"neighborhood"`
	Street           string `json:"street"`
	ImgUrl           string `json:"imgUrl"`
	InaugurationDate string `json:"inaugurationDate"`
	LineNumber       uint   `json:"lineId"`
}

func NewStation(data data.StationDTO) Station {
	return Station{
		Name:             data.Name,
		Street:           data.Street,
		LineNumber:       uint(data.LineNumber),
		Neighborhood:     data.Neighborhood,
		InaugurationDate: data.InaugurationDate,
		ImgUrl:           "",
	}
}
