package main

import "fmt"

func main() {
	var myVar interface{} = "Jerry Augusto"
	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("O valor de res é %v e o valor de ok é %v", res, ok)
}
