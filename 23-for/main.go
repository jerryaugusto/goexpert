package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	numbers := []string{"one", "two", "three"}
	for _, v := range numbers {
		fmt.Println(v)
	}

	// Infinity loop
	// for {
	// 	fmt.Println("Hello, world!")
	// }
}
