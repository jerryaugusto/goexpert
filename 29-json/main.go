package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number int
	Balance int
}

func main() {
	account := Account{ Number: 1, Balance: 100 }
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(account)
}
