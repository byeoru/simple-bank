package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	KRW = "KRW"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, KRW:
		return true
	}
	return false
}
