package main

import (
	"fmt"
	"log"

	"github.com/jyp90/practice/src/banking"
)

func main() {
	account := banking.NewAccount("jypark")
	account.Deposit(10000)
	account.Deposit(10000)
	account.Deposit(10000)
	fmt.Println(account.Balance())
	account.Withdraw(5000)
	err := account.Withdraw(40000)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
