package main

import "fmt"

type Address struct {
	City     string
	State    string
	District string
	Number   int
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	jerry := Client{
		Name:   "Jerry Augusto",
		Age:    28,
		Active: true,
	}

	fmt.Printf("O nome do cliente é %s", jerry.Name)
}
