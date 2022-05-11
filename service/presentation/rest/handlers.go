package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Muchogoc/phone-numbers-exercise/service/usecases"
)

// PresentationHandlers represents all the handler logic
type Handlers interface {
	ListCustomers() http.HandlerFunc
}

// HandlersImpl represents the usecase implementation object
type HandlersImpl struct {
	usecases usecases.Usecases
}

// NewHandlers initializes a new rest handlers usecase
func NewHandlers(usecases usecases.Usecases) Handlers {
	return &HandlersImpl{
		usecases: usecases,
	}
}

func (h HandlersImpl) ListCustomers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		pagination, err := paginationParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		filters, err := filterParams(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		customers, err := h.usecases.ListCustomers(ctx, filters, *pagination)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(customers)
		if err != nil {
			http.Error(w, "error encoding customers", http.StatusInternalServerError)
			return
		}
	}
}
