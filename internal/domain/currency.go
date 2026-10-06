package domain

import "fmt"

type Currency string

const (
	RUB Currency = "RUB"
	BYN Currency = "BYN"
	KZT Currency = "KZT"
	CNY Currency = "CNY"
	EUR Currency = "EUR"
	USD Currency = "USD"
	BTC Currency = "BTC"
)

// IsValid сообщает, поддерживается ли данная валюта системой.
func IsValid(c Currency) bool {
	switch c {
	case RUB, BYN, KZT, CNY, EUR, USD, BTC:
		return true
	default:
		return false
	}
}

/*ParseCurrency превращает строку в валюту, если она поддерживается. Возвращает ошибку, если валюта неизвестна*/
func ParseCurrency(s string) (Currency, error) {
	c := Currency(s)
	if !IsValid(c) {
		return "", fmt.Errorf("unknow currency: %d", s)
	}
	return c, nil
}
