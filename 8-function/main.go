package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(5, 10)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}

// Funções em Go podem retornar mais do que um valor.
func sum(a, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("the sum is greater than 50")
	}
	return a + b, nil
}
