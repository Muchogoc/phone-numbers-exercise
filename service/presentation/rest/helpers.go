package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
)

func paginationParams(r *http.Request) (*domain.PaginationInput, error) {
	// with a values as -1 we'll get a request without limit or offset as default
	pagination := &domain.PaginationInput{
		Offset: -1,
		Limit:  -1,
	}

	strLimit := r.URL.Query().Get("limit")
	if strLimit != "" {
		limit, err := strconv.Atoi(strLimit)
		if err != nil || limit < -1 {
			return nil, fmt.Errorf("limit query parameter is not a valid number")
		}
		pagination.Limit = limit
	}

	strOffset := r.URL.Query().Get("offset")
	if strOffset != "" {
		offset, err := strconv.Atoi(strOffset)
		if err != nil || offset < -1 {
			return nil, fmt.Errorf("offset query parameter is not a valid number")
		}
		pagination.Offset = offset
	}

	return pagination, nil

}

func filterParams(r *http.Request) (*domain.FilterInput, error) {
	filter := &domain.FilterInput{
		Country: nil,
		State:   nil,
	}

	countryValue := r.URL.Query().Get("country")
	if countryValue != "" {
		country := domain.Country(countryValue)
		if !country.IsValid() {
			return nil, fmt.Errorf("country query parameter is not a valid country")
		}

		filter.Country = &country
	}

	stateValue := r.URL.Query().Get("state")
	if stateValue != "" {
		state, err := strconv.ParseBool(stateValue)
		if err != nil {
			return nil, fmt.Errorf("state query parameter is not a valid boolean")
		}
		filter.State = &state
	}

	return filter, nil

}
