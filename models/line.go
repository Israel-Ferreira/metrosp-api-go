package models

type Line struct {
	Number    string  `json:"number"`
	Color     string  `json:"color"`
	Extension float64 `json:"extension"`
	MapImgUrl string  `json:"mapImgUrl"`
}
