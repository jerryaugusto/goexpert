package main

import (
	"fmt"
)

func main() {
	// salary := make(map[string]int)
	// salary := map[string]int{}
	salary := map[string]int{"Jerry": 1000, "Maria": 2000, "João": 3000}
	delete(salary, "Jerry")
	salary["Jerry Augusto"] = 5000
	fmt.Println(salary["Jerry Augusto"])

	for name, aSalary := range salary {
		fmt.Printf("The salary of %s is: %d\n", name, aSalary)
	}

	for _, aSalary := range salary {
		fmt.Printf("The salary is: %d\n", aSalary)
	}
}
