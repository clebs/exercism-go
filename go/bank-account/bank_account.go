// Package account provides types and methods to work with bank accounts.
package account

import "sync"

const testVersion = 1

// Account defines the behaviour any bank account should have.
type Account interface {
	Close() (payout int64, ok bool)
	Balance() (balance int64, ok bool)
	Deposit(amount int64) (newBalance int64, ok bool)
}

type basicAccount struct {
	balance int64
	active  bool
	mutex   sync.Mutex
}

// Open creates a new bank account. If negative initial balance is given, no account is opened.
func Open(initial int64) Account {
	if initial < 0 {
		return nil
	}
	return &basicAccount{balance: initial, active: true}
}

func (a *basicAccount) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.active {
		return 0, false
	}

	a.active = false
	return a.balance, true
}

func (a *basicAccount) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.active {
		return 0, false
	}
	return a.balance, true
}

func (a *basicAccount) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.active || a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}
