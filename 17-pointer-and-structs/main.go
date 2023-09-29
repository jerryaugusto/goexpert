package main

import "fmt"

type Account struct {
	balance int
}

func NewAccount() *Account {
	return &Account{balance: 0}
}

func (a *Account) simulate(value int) int {
	a.balance += value
	fmt.Println(a.balance)
	return a.balance
}

func main() {
	account := Account{balance: 5000}
	account.simulate(2000)

	fmt.Println(account.balance)
}
