package database

import (
	"fmt"
	"strings"

	"github.com/Muchogoc/phone-numbers-exercise/service/domain"
)

func phoneCountry(code string) domain.Country {
	// Converts a country code to the standard i.e (254) to +254
	code = strings.Replace(code, "(", "", 1)
	code = strings.Replace(code, ")", "", 1)
	code = fmt.Sprintf("+%s", code)

	for _, country := range domain.AllCountries {
		if country.Code() == code {
			return country
		}
	}

	return ""
}
