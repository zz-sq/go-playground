package main

import "sync"

var (
	l       sync.Mutex
	balance int
)

func Deposit(amount int) {
	l.Lock()
	defer l.Unlock()
	balance += amount
}

func Balance() int {
	l.Lock()
	defer l.Unlock()
	return balance
}

func WithDraw(amount int) bool {
	l.Lock()
	defer l.Unlock()
	if balance < amount {
		return false
	}
	balance -= amount
	return true
}
