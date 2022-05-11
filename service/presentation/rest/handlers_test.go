package rest

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Muchogoc/phone-numbers-exercise/service/infrastructure"
	"github.com/Muchogoc/phone-numbers-exercise/service/infrastructure/database/sqlite"
	"github.com/Muchogoc/phone-numbers-exercise/service/usecases"
)

var handlers Handlers

func TestMain(m *testing.M) {
	initialValue := os.Getenv(sqlite.RunningTestsEnvName)
	os.Setenv(sqlite.RunningTestsEnvName, "true")

	infra := infrastructure.NewInfrastructure()
	usecases := usecases.NewUsecasesImpl(infra)
	handlers = NewHandlers(usecases)

	exitVal := m.Run()

	os.Setenv(sqlite.RunningTestsEnvName, initialValue)

	os.Exit(exitVal)
}

func TestListCustomersHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/api/v1/customers", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(handlers.ListCustomers())

	// a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	// call the ServeHTTP method
	handler.ServeHTTP(rr, request)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
