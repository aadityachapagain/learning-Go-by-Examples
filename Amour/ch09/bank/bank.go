// Package bank provides a concurrency-safe single-account bank.
package bank

import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

// Deposit func to deposit the money inside bank
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance = balance + amount
}

// Balance func Get the balace of bank account
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	b := balance
	return b
}
