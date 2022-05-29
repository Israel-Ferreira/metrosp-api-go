package models

type Line struct {
	ID        uint64  `json:"id"`
	Number    string  `json:"number"`
	Color     string  `json:"color"`
	Extension float64 `json:"extension"`
	MapImgUrl string  `json:"mapImgUrl"`
}
