package domain

// Amount - сумма в наименьших денежных единицах конкретной валюты.
// Сам по себе Amount бессмысленен без валюты, к которой он относится.
type Amount int64

func MinorUnitsPerWhole(c Currency) Amount {
	switch c {
	case CNY:
		return 10
	case RUB, USD, BYN, KZT, EUR:
		return 100
	case BTC:
		return 100_000_000
	default:
		return 0
	}
}
