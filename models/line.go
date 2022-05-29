package models

import "gorm.io/gorm"

type Line struct {
	gorm.Model
	Number    string    `json:"number"`
	Color     string    `json:"color"`
	Extension float64   `json:"extension"`
	MapImgUrl string    `json:"mapImgUrl"`
	Stations  []Station `json:"stations"`
}
