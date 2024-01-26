package banking

import (
	"errors"
	"strconv"
)

type BankAccount struct {
	owner   string
	balance int
}

// create constructor
func NewAccount(owner string) *BankAccount {
	account := BankAccount{owner: owner, balance: 0}
	return &account
}

// deposit x amount on your account
func (a *BankAccount) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a *BankAccount) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
func (a *BankAccount) Withdraw(amount int) error {
	if a.balance < amount {
		balance := strconv.Itoa(a.balance)
		message := "Cannot withdraw your balance is " + balance
		return errors.New(message)
	}

	a.balance -= amount
	return nil
}

var errNoMoney = errors.New("")
