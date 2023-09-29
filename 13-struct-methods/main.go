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

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("The client %s has been disabled\n", c.Name)
}

func main() {
	jerry := Client{
		Name:   "Jerry Augusto",
		Age:    28,
		Active: true,
	}

	jerry.Disable()
}
