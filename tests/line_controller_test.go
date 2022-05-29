package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Israel-Ferreira/metrosp-api/server"
	"github.com/Israel-Ferreira/metrosp-api/services"
	"github.com/Israel-Ferreira/metrosp-api/tests/mocks"
)

func TestCreateLineEndpoint(t *testing.T) {

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
			t.Errorf("Expected 200, but returns %d \n", w.Result().StatusCode)
		}

	})

}

func TestGetLineById(t *testing.T) {

}

func TestUpdateLine(t *testing.T) {

}
