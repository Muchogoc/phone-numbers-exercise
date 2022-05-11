package domain

type Country string

var (
	CountryCameroon   Country = "Cameroon"
	CountryEthiopia   Country = "Ethiopia"
	CountryMorocco    Country = "Morocco"
	CountryMozambique Country = "Mozambique"
	CountryUganda     Country = "Uganda"
)

var AllCountries = []Country{
	CountryCameroon,
	CountryEthiopia,
	CountryMorocco,
	CountryMozambique,
	CountryUganda,
}

func (c Country) RegexPattern() string {
	switch c {
	case CountryCameroon:
		return `\(237\)\ ?[2368]\d{7,8}$`
	case CountryEthiopia:
		return `\(251\)\ ?[1-59]\d{8}$`
	case CountryMorocco:
		return `\(212\)\ ?[5-9]\d{8}$`
	case CountryMozambique:
		return `\(258\)\ ?[28]\d{7,8}$`
	case CountryUganda:
		return `\(256\)\ ?\d{9}$`
	default:
		return ""
	}
}

func (c Country) Code() string {
	switch c {
	case CountryCameroon:
		return "+237"
	case CountryEthiopia:
		return "+251"
	case CountryMorocco:
		return "+212"
	case CountryMozambique:
		return "+258"
	case CountryUganda:
		return "+256"
	default:
		return ""
	}
}

func (c Country) IsValid() bool {
	switch c {
	case CountryCameroon, CountryEthiopia, CountryMorocco, CountryMozambique, CountryUganda:
		return true
	default:
		return false
	}
}
