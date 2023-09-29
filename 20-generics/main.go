package main

import "fmt"

type Number interface {
	~int | ~float64
}

func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

func Compare[T comparable](a, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Jerry": 5000, "Maria": 2000, "Pedro": 3000}
	fmt.Println(Sum(m))
	fmt.Println(Compare(10, 10.0))
}
