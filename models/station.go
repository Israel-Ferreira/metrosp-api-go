package models

import "github.com/Israel-Ferreira/metrosp-api/data"

type Station struct {
	Name             string `json:"name"`
	Neighborhood     string `json:"neighborhood"`
	Street           string `json:"street"`
	LineID           int64  `json:"lineId"`
	ImgUrl           string `json:"imgUrl"`
	InaugurationDate string `json:"inaugurationDate"`
}

func NewStation(data data.StationDTO) Station {
	return Station{
		Name:             data.Name,
		Street:           data.Street,
		LineID:           int64(data.LineNumber),
		Neighborhood:     data.Neighborhood,
		InaugurationDate: data.InaugurationDate,
		ImgUrl:           data.ImagesUrls[0],
	}
}
