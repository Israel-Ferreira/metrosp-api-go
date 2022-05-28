package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Israel-Ferreira/metrosp-api/server"
)

func TestGetAllEndpoint(t *testing.T) {
	t.Run("Deve todas as estações cadastradas e retornar 200", func(t *testing.T) {
		router := server.SetupServer(nil, nil)

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
}

func TestGetByIdEndpoint(t *testing.T) {

}

func TestCreate(t *testing.T) {

}

func TestDeleteEndpoint(t *testing.T) {

}
