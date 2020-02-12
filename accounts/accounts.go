package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// Construct
// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Methods
// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Owner of your account
func (a Account) Owner() string {
	return a.owner
}

// __str__
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account - $", a.Balance())
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount on your account
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// Change owner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}
