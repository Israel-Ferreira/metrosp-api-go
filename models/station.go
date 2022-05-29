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
	LineID           int  `json:"lineId"`
	ImgUrl           string `json:"imgUrl"`
	InaugurationDate string `json:"inaugurationDate"`
}

func NewStation(data data.StationDTO) Station {
	return Station{
		Name:             data.Name,
		Street:           data.Street,
		LineID:           int(data.LineNumber),
		Neighborhood:     data.Neighborhood,
		InaugurationDate: data.InaugurationDate,
		ImgUrl:           data.ImagesUrls[0],
	}
}
