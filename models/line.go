package models

import (
	"github.com/Israel-Ferreira/metrosp-api/data"
	"gorm.io/gorm"
)

type Line struct {
	gorm.Model
	Number    uint      `json:"number" gorm:"unique"`
	Color     string    `json:"color" gorm:"unique"`
	Extension float64   `json:"extension"`
	MapImgUrl string    `json:"mapImgUrl"`
	Stations  []Station `json:"stations" gorm:"foreignKey:LineNumber;references:Number"`
}

func NewLine(dto data.LineDTO) Line {
	return Line{
		Number:    uint(dto.Number),
		Extension: dto.Extension,
		MapImgUrl: dto.MapImageUrl,
		Color:     dto.Color,
	}
}
