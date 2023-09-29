package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	jerry := Client{
		Name:   "Jerry Augusto",
		Age:    28,
		Active: true,
	}

	fmt.Printf("O nome do cliente é %s", jerry.Name)
}
