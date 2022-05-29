package tests

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Israel-Ferreira/metrosp-api/server"
	"github.com/Israel-Ferreira/metrosp-api/services"
	"github.com/Israel-Ferreira/metrosp-api/tests/mocks"
)

func TestGetAllEndpoint(t *testing.T) {
	t.Run("Deve retornar todas as estações cadastradas e retornar 200", func(t *testing.T) {

		svc := services.NewStationService(mocks.NewMockStationRepo(false), &mocks.MockSuccessLineRepo{})
		router := server.SetupServer(svc, nil)

		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/api/stations", nil)

		router.ServeHTTP(w, req)

		if err != nil {
			t.Errorf("Error on Requisition:  %v \n", err)
		}

		if w.Result().StatusCode != 200 {
			t.Errorf("Expected 200, but returns %d \n", w.Result().StatusCode)
		}

	})

	t.Run("Deve retornar 500 ao falhar", func(t *testing.T) {
		svc := services.NewStationService(mocks.NewMockStationRepo(true), &mocks.MockSuccessLineRepo{})
		router := server.SetupServer(svc, nil)

		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/api/stations", nil)

		router.ServeHTTP(w, req)

		if err != nil {
			t.Errorf("Error on Requisition:  %v \n", err)
		}

		if w.Result().StatusCode != 500 {
			t.Errorf("Expected 500, but returns %d \n", w.Result().StatusCode)
		}
	})
}

func TestGetByIdEndpoint(t *testing.T) {
	t.Run("Deve retornar 200, ao encontrar o id na base", func(t *testing.T) {
		svc := services.NewStationService(mocks.NewMockStationRepo(false), &mocks.MockSuccessLineRepo{})
		router := server.SetupServer(svc, nil)

		w := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/api/stations/433", nil)

		router.ServeHTTP(w, req)

		if err != nil {
			t.Errorf("Error on Requisition:  %v \n", err)
		}

		if w.Result().StatusCode != 200 {
			t.Errorf("Expected 200, but returns %d \n", w.Result().StatusCode)
		}
	})

	t.Run("Deve retornar o status 404, ao não encontrar a estação", func(t *testing.T) {
		svc := services.NewStationService(mocks.NewMockStationRepo(false), &mocks.MockSuccessLineRepo{})
		router := server.SetupServer(svc, nil)

		w := httptest.NewRecorder()
		id := rand.Uint64()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/stations/%d", id), nil)

		router.ServeHTTP(w, req)

		if err != nil {
			t.Errorf("Error on Requisition:  %v \n", err)
		}

		if w.Result().StatusCode != 404 {
			t.Errorf("Expected 404, but returns %d \n", w.Result().StatusCode)
		}
	})

	t.Run("Deve retornar 404, se ocorrer um erro", func(t *testing.T) {
		svc := services.NewStationService(mocks.NewMockStationRepo(true), &mocks.MockSuccessLineRepo{})
		router := server.SetupServer(svc, nil)

		w := httptest.NewRecorder()
		id := rand.Uint64()

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/stations/%d", id), nil)

		router.ServeHTTP(w, req)

		if err != nil {
			t.Errorf("Error on Requisition:  %v \n", err)
		}

		if w.Result().StatusCode != 404 {
			t.Errorf("Expected 404, but returns %d \n", w.Result().StatusCode)
		}
	})
}

func TestCreate(t *testing.T) {

}

func TestDeleteEndpoint(t *testing.T) {

}
