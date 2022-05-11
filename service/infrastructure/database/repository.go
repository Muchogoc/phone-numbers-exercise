package database

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
	"github.com/Muchogoc/phone-numbers-exercise/service/infrastructure/database/sqlite"
)

type Repository interface {
	GetCustomers(ctx context.Context, filters *domain.FilterInput, pagination domain.PaginationInput) ([]*domain.Customer, error)
}

type DatabaseImpl struct {
	sqlite *sqlite.DBImpl
}

func NewDatabaseImpl() Repository {
	sq := sqlite.NewSqliteDBImpl()
	return &DatabaseImpl{
		sqlite: sq,
	}
}

func (d DatabaseImpl) GetCustomers(ctx context.Context, filters *domain.FilterInput, pagination domain.PaginationInput) ([]*domain.Customer, error) {
	mapped := []*domain.Customer{}

	customers, err := d.sqlite.GetCustomers(ctx, filters, pagination)
	if err != nil {
		return nil, fmt.Errorf("repository error fetching customers: %w", err)
	}

	for _, customer := range customers {
		m := &domain.Customer{
			ID:   customer.ID,
			Name: customer.Name,
		}

		// split the phone number to obtain the code and number
		// "(212) 698054317" -> ["(212)","698054317"]
		split := strings.Split(customer.Phone, " ")
		country := phoneCountry(split[0])
		if !country.IsValid() {
			continue
		}

		m.Country = country
		m.Phone = country.Code() + split[1]

		r := regexp.MustCompile(country.RegexPattern())

		status := r.MatchString(customer.Phone)
		m.IsValid = status

		// check if there is a state filter
		if filters.State != nil {
			// use filter to add the matching record
			if m.IsValid == *filters.State {
				mapped = append(mapped, m)
			}
		} else {
			mapped = append(mapped, m)
		}

	}

	return mapped, nil
}
