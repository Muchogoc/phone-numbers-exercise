package presentation

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Muchogoc/phone-numbers-exercise/service/infrastructure"
	"github.com/Muchogoc/phone-numbers-exercise/service/presentation/rest"
	"github.com/Muchogoc/phone-numbers-exercise/service/usecases"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func PrepareServer(port int) *http.Server {
	addr := fmt.Sprintf(":%d", port)
	router := mux.NewRouter()

	infra := infrastructure.NewInfrastructure()
	usecases := usecases.NewUsecasesImpl(infra)
	rest := rest.NewHandlers(usecases)

	api := router.PathPrefix("/api/v1").Subrouter()
	api.Path("/customers").Methods(
		http.MethodOptions,
		http.MethodGet,
	).Handler(rest.ListCustomers())

	// serves the SPA
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/build/")))

	h := handlers.CombinedLoggingHandler(os.Stdout, router)
	h = handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(true),
	)(h)
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	h = handlers.CORS(allowedOrigins)(h)

	srv := &http.Server{
		Handler:      h,
		Addr:         addr,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		IdleTimeout:  4 * time.Second,
	}

	return srv
}
