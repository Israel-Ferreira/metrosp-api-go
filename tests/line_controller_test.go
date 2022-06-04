package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Israel-Ferreira/metrosp-api/data"
	"github.com/Israel-Ferreira/metrosp-api/server"
	"github.com/Israel-Ferreira/metrosp-api/services"
	"github.com/Israel-Ferreira/metrosp-api/tests/mocks"
)

func TestCreateLineEndpoint(t *testing.T) {
	t.Run("Deve retornar 400, se o número da linha estiver vazio", func(t *testing.T) {
		lineSvc := services.NewLineService(mocks.NewMockLineRepo(true))
		stationSvc := services.NewEmptyStationService()

		w := httptest.NewRecorder()

		router := server.SetupServer(stationSvc, lineSvc)

		line := data.LineDTO{
			Color:       "Rosa",
			Number:      "",
			Extension:   28.8,
			MapImageUrl: "",
		}

		jsonVal, _ := json.Marshal(line)

		req, _ := http.NewRequest(http.MethodPost, "/api/lines", bytes.NewBuffer(jsonVal))

		router.ServeHTTP(w, req)

		if w.Result().StatusCode != 400 {
			t.Errorf("Expected 500, but returns %d \n", w.Result().StatusCode)
		}
	})

	t.Run("Deve retornar 404, se o número da linha estiver vazio", func(t *testing.T) {
		lineSvc := services.NewLineService(mocks.NewMockLineRepo(false))
		stationSvc := services.NewEmptyStationService()

		w := httptest.NewRecorder()

		router := server.SetupServer(stationSvc, lineSvc)

		line := data.LineDTO{
			Color:       "Rosa",
			Number:      "20",
			Extension:   28.8,
			MapImageUrl: "",
		}

		jsonVal, _ := json.Marshal(line)

		req, _ := http.NewRequest(http.MethodPost, "/api/lines", bytes.NewBuffer(jsonVal))

		router.ServeHTTP(w, req)

		if w.Result().StatusCode != 201 {
			t.Errorf("Expected 201, but returns %d \n", w.Result().StatusCode)
		}
	})

	t.Run("Deve retornar 404, se a linha for existente", func(t *testing.T) {
		lineSvc := services.NewLineService(mocks.NewMockLineRepo(false))
		stationSvc := services.NewEmptyStationService()

		w := httptest.NewRecorder()

		router := server.SetupServer(stationSvc, lineSvc)

		line := data.LineDTO{
			Color:       "Rosa",
			Number:      "15",
			Extension:   28.8,
			MapImageUrl: "",
		}

		jsonVal, _ := json.Marshal(line)

		req, _ := http.NewRequest(http.MethodPost, "/api/lines", bytes.NewBuffer(jsonVal))

		router.ServeHTTP(w, req)

		if w.Result().StatusCode != 404 {
			t.Errorf("Expected 404, but returns %d \n", w.Result().StatusCode)
		}
	})
}

func TestGetAllLineEndpoint(t *testing.T) {
	t.Run("Deve retornar o status 500, se ocorrer um erro", func(t *testing.T) {
		lineSvc := services.NewLineService(mocks.NewMockLineRepo(true))
		stationSvc := services.NewEmptyStationService()

		w := httptest.NewRecorder()

		router := server.SetupServer(stationSvc, lineSvc)

		req, _ := http.NewRequest(http.MethodGet, "/api/lines", nil)

		router.ServeHTTP(w, req)

		if w.Result().StatusCode != 500 {
			t.Errorf("Expected 500, but returns %d \n", w.Result().StatusCode)
		}

	})

	t.Run("Deve retornar o status 200", func(t *testing.T) {
		lineSvc := services.NewLineService(mocks.NewMockLineRepo(false))
		stationSvc := services.NewEmptyStationService()

		w := httptest.NewRecorder()

		router := server.SetupServer(stationSvc, lineSvc)

		req, _ := http.NewRequest(http.MethodGet, "/api/lines", nil)

		router.ServeHTTP(w, req)

		if w.Result().StatusCode != 200 {
			t.Errorf("Expected 200, but returns %d \n", w.Result().StatusCode)
		}

	})

}

func TestGetLineById(t *testing.T) {

}

func TestUpdateLine(t *testing.T) {

}
