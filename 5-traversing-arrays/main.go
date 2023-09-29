package main

import "fmt"

const a = "Hello, world!"

type ID int

var (
	b bool    = true
	c int     = 1
	d string  = "Jerry Augusto"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3

	for i, v := range myArray {
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}
}
