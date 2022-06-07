package mocks

import (
	"github.com/Israel-Ferreira/metrosp-api/models"
	"gorm.io/gorm"
)

var estacoes = []models.Station{
	{Model: gorm.Model{ID: 433}, Name: "Paulista", Neighborhood: "Consolação", LineNumber: 4},
	{Model: gorm.Model{ID: 234}, Name: "Brigadeiro", Neighborhood: "Bela Vista", LineNumber: 2},
	{Model: gorm.Model{ID: 124}, Name: "Santana", Neighborhood: "Santana", LineNumber: 1},
	{Model: gorm.Model{ID: 321}, Name: "Corinthians Itaquera", Neighborhood: "Itaquera", LineNumber: 3},
	{Model: gorm.Model{ID: 521}, Name: "Chacara Klabin", Neighborhood: "Vila Mariana", LineNumber: 5},
	{Model: gorm.Model{ID: 321}, Name: "São Mateus", Neighborhood: "São Mateus", LineNumber: 15},
}

var linhas = []models.Line{
	{Model: gorm.Model{ID: 3221}, Number: 5, Color: "Lilas", Extension: 20.0, MapImgUrl: ""},
	{Model: gorm.Model{ID: 3232}, Number: 15, Color: "Prata", Extension: 14.6, MapImgUrl: ""},
	{Model: gorm.Model{ID: 3234}, Number: 4, Color: "Amarela", Extension: 12.0, MapImgUrl: ""},
	{Model: gorm.Model{ID: 3621}, Number: 2, Color: "Verde", Extension: 14.7, MapImgUrl: ""},
	{Model: gorm.Model{ID: 3121}, Number: 1, Color: "Azul", Extension: 20.2, MapImgUrl: ""},
}
