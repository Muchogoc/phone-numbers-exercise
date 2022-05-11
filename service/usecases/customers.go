package usecases

import (
	"context"
	"fmt"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
	"github.com/Muchogoc/phone-numbers-exercise/service/infrastructure"
)

// ICustomers method signatures for customer interaction
type ICustomers interface {
	ListCustomers(ctx context.Context, filters *domain.FilterInput, pagination domain.PaginationInput) ([]*domain.Customer, error)
}

// Usecases is the interactor exposed by the usecase layer
type Usecases interface {
	ICustomers
}

// UsecasesImpl is an implementation of usecase interface
type UsecasesImpl struct {
	infra infrastructure.Infrastructure
}

// NewUsecasesImpl a new Usecases interface implementation
func NewUsecasesImpl(infra infrastructure.Infrastructure) Usecases {
	return &UsecasesImpl{
		infra: infra,
	}
}

// ListCustomers lists the customers and their phone numbers
func (u UsecasesImpl) ListCustomers(ctx context.Context, filters *domain.FilterInput, pagination domain.PaginationInput) ([]*domain.Customer, error) {
	customers, err := u.infra.GetCustomers(ctx, filters, pagination)
	if err != nil {
		return nil, fmt.Errorf("error fetching customers: %w", err)
	}

	return customers, nil
}
