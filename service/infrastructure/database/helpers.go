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

	switch code {
	case "+237":
		return domain.CountryCameroon
	case "+251":
		return domain.CountryEthiopia
	case "+212":
		return domain.CountryMorocco
	case "+258":
		return domain.CountryMozambique
	case "+256":
		return domain.CountryUganda
	default:
		return ""
	}
}
