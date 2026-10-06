package main

import (
	"fmt"

	"github.com/Sergey-Ivanov-CGI/payments/internal/domain"
)

func main() {
	c, err := domain.ParseCurrency("RUB")
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("parsed:", c)

	_, err = domain.ParseCurrency("XYZ")
	if err != nil {
		fmt.Println("expected error:", err)
	}
}
