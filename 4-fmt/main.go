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
	fmt.Printf("O tipo de f é %T", f)
}
