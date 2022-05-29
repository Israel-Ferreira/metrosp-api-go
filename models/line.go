package models

import (
	"github.com/Israel-Ferreira/metrosp-api/data"
	"gorm.io/gorm"
)

type Line struct {
	gorm.Model
	Number    string    `json:"number"`
	Color     string    `json:"color"`
	Extension float64   `json:"extension"`
	MapImgUrl string    `json:"mapImgUrl"`
	Stations  []Station `json:"stations"`
}

func NewLine(dto data.LineDTO) Line {
	return Line{
		Number:    dto.Number,
		Extension: dto.Extension,
		MapImgUrl: dto.MapImageUrl,
		Color:     dto.Color,
	}
}
