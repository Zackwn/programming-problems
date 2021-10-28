package main

import "sync"

type Account struct {
	sync.Mutex
	balance int64
	open    bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	a := new(Account)
	a.balance = amount
	a.open = true
	return a
}

func (a *Account) Balance() (int64, bool) {
	if !a.open {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if !a.open {
		return 0, false
	}
	a.Lock()
	defer a.Unlock()
	if a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	if !a.open {
		return 0, false
	}
	a.open = false
	return a.balance, true
}
