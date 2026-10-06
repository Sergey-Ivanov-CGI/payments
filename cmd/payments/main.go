package main

import (
	"fmt"

	"github.com/Sergey-Ivanov-CGI/payments/internal/domain"
)

func main() {
	amount := domain.Amount(150_00) // 150 рублей в копейках
	units := domain.MinorUnitsPerWhole(domain.RUB)

	fmt.Println("Amount:", amount, "minor units per RUB:", units)
	fmt.Println("Amount in rubles:", float64(amount)/float64(units))
}
