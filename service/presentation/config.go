package presentation

import (
	"fmt"
	"io/fs"
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

	var frontend fs.FS = os.DirFS("frontend")
	httpFS := http.FS(frontend)
	fileServer := http.FileServer(httpFS)

	router.Path("/").Handler(fileServer)
	router.Path("/customers/").Methods(
		http.MethodOptions,
		http.MethodGet,
	).Handler(rest.ListCustomers())

	h := handlers.CombinedLoggingHandler(os.Stdout, router)
	h = handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(true),
	)(h)

	srv := &http.Server{
		Handler:      h,
		Addr:         addr,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
		IdleTimeout:  4 * time.Second,
	}

	return srv
}
