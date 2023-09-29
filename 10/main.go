package main

import "fmt"

func main() {
	times := func() int {
		return sum(1, 2, 3, 4, 5, 6, 7, 8, 9) * 2
	}()

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9), times)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
