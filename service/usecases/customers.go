package usecases

import (
	"context"
	"fmt"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
	"github.com/Muchogoc/phone-numbers-exercise/service/infrastructure"
)

type ICustomers interface {
	ListCustomers(ctx context.Context, filters *domain.FilterInput, pagination domain.PaginationInput) ([]*domain.Customer, error)
}

type Usecases interface {
	ICustomers
}

type UsecasesImpl struct {
	infra infrastructure.Infrastructure
}

func NewUsecasesImpl(infra infrastructure.Infrastructure) Usecases {
	return &UsecasesImpl{
		infra: infra,
	}
}

func (u UsecasesImpl) ListCustomers(ctx context.Context, filters *domain.FilterInput, pagination domain.PaginationInput) ([]*domain.Customer, error) {
	customers, err := u.infra.GetCustomers(ctx, filters, pagination)
	if err != nil {
		return nil, fmt.Errorf("error fetching customers: %w", err)
	}

	return customers, nil
}
