package mocks

import "github.com/Israel-Ferreira/metrosp-api/models"

var estacoes = []models.Station{
	{ID: 433, Name: "Paulista", Neighborhood: "Consolação", LineID: 4},
	{ID: 234, Name: "Brigadeiro", Neighborhood: "Bela Vista", LineID: 2},
	{ID: 124, Name: "Santana", Neighborhood: "Santana", LineID: 1},
	{ID: 321, Name: "Corinthians Itaquera", Neighborhood: "Itaquera", LineID: 3},
	{ID: 521, Name: "Chacara Klabin", Neighborhood: "Vila Mariana", LineID: 5},
	{ID: 321, Name: "São Mateus", Neighborhood: "São Mateus", LineID: 15},
}

var linhas = []models.Line{
	{ID: 3221, Number: "5", Color: "Lilas", Extension: 20.0, MapImgUrl: ""},
	{ID: 3232, Number: "15", Color: "Prata", Extension: 14.6, MapImgUrl: ""},
	{ID: 3234, Number: "4", Color: "Amarela", Extension: 12.0, MapImgUrl: ""},
	{ID: 3621, Number: "2", Color: "Verde", Extension: 14.7, MapImgUrl: ""},
	{ID: 3121, Number: "1", Color: "Azul", Extension: 20.2, MapImgUrl: ""},
}