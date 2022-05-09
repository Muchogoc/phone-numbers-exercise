package infrastructure

import (
	"context"
	"fmt"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
	"github.com/Muchogoc/phone-numbers-exercise/service/infrastructure/database"
)

// Infrastructure is a set of methods exposed by the infrastructure layer
type Infrastructure interface {
	database.Repository
}

type InfrastructureImpl struct {
	repository database.Repository
}

// NewInfrastructure returns an implementation of the infra interface
func NewInfrastructure() Infrastructure {
	repository := database.NewDatabaseImpl()

	return &InfrastructureImpl{
		repository: repository,
	}
}

func (d InfrastructureImpl) GetCustomers(ctx context.Context, filters *domain.FilterInput, pagination domain.PaginationInput) ([]*domain.Customer, error) {
	customers, err := d.repository.GetCustomers(ctx, filters, pagination)
	if err != nil {
		return nil, fmt.Errorf("infrastructure error fetching customers: %w", err)
	}

	return customers, nil
}
