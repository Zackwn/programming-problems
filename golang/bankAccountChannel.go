package main

type Account struct {
	balance int64
	open    bool
	deposit chan int64
}

func (a *Account) Start() {
	for {
		amount, ok := <-a.deposit
		if !ok {
			return
		}
		a.balance += amount
	}
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	a := new(Account)
	a.balance = amount
	a.open = true
	a.deposit = make(chan int64)
	go a.Start()
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
	if a.balance+amount < 0 {
		return 0, false
	}
	a.deposit <- amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	if !a.open {
		return 0, false
	}
	a.open = false
	close(a.deposit)
	return a.balance, true
}
