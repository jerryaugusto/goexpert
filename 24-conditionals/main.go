package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3

	//if a < b || c > a {
	//	fmt.Println("'a' é menor do que 'b' e 'c' é maior do que 'a'")
	//}

	switch a {
	case 1:
		fmt.Println(a)
	case 2:
		fmt.Println(b)
	case 3:
		fmt.Println(c)
	default:
		fmt.Println("Invalid")
	}
}
