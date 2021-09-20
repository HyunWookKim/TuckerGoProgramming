package main

import (
	"fmt"

	"exB1/bankaccount"
)

func main() {
	account := bankaccount.NewAccount()
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
