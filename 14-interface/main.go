package main

import "fmt"

type Address struct {
	City     string
	State    string
	District string
	Number   int
}

type Person interface {
	Disable()
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

func Deactivation(person Person) {
	person.Disable()
}

func main() {
	jerry := Client{
		Name:   "Jerry Augusto",
		Age:    28,
		Active: true,
	}

	Deactivation(jerry)
}
