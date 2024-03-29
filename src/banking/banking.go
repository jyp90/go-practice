package banking

import (
	"errors"
	"fmt"
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
		// return errNoMoney
	}

	a.balance -= amount
	return nil
}

// Change owner of the account
func (a *BankAccount) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a *BankAccount) Owner() string {
	return a.owner
}

func (a BankAccount) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas : ", a.Balance())
}

// Error
var errNoMoney = errors.New("")
