package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Number  int     `json:"number"`
	Balance float64 `json:"balance"`
}

func main() {
	account := Account{Number: 123, Balance: 100.50}
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}

	println(string(res))

	// Using econder
	encoder := json.NewEncoder(os.Stdout)
	err = encoder.Encode(account)

	if err != nil {
		panic(err)
	}

	// Json to struct
	jsonString := `{"number":2,"balance":100.50}`
	var account2 Account

	err = json.Unmarshal([]byte(jsonString), &account2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", account2)
}
