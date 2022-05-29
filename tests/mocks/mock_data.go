package mocks

import (
	"github.com/Israel-Ferreira/metrosp-api/models"
	"gorm.io/gorm"
)

var estacoes = []models.Station{
	{Model: gorm.Model{ID: 433}, Name: "Paulista", Neighborhood: "Consolação", LineID: 4},
	{Model: gorm.Model{ID: 234}, Name: "Brigadeiro", Neighborhood: "Bela Vista", LineID: 2},
	{Model: gorm.Model{ID: 124}, Name: "Santana", Neighborhood: "Santana", LineID: 1},
	{Model: gorm.Model{ID: 321}, Name: "Corinthians Itaquera", Neighborhood: "Itaquera", LineID: 3},
	{Model: gorm.Model{ID: 521}, Name: "Chacara Klabin", Neighborhood: "Vila Mariana", LineID: 5},
	{Model: gorm.Model{ID: 321}, Name: "São Mateus", Neighborhood: "São Mateus", LineID: 15},
}

var linhas = []models.Line{
	{Model: gorm.Model{ID: 3221}, Number: "5", Color: "Lilas", Extension: 20.0, MapImgUrl: ""},
	{Model: gorm.Model{ID: 3232}, Number: "15", Color: "Prata", Extension: 14.6, MapImgUrl: ""},
	{Model: gorm.Model{ID: 3234}, Number: "4", Color: "Amarela", Extension: 12.0, MapImgUrl: ""},
	{Model: gorm.Model{ID: 3621}, Number: "2", Color: "Verde", Extension: 14.7, MapImgUrl: ""},
	{Model: gorm.Model{ID: 3121}, Number: "1", Color: "Azul", Extension: 20.2, MapImgUrl: ""},
}
