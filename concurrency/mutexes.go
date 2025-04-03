package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank_Account struct {
	balance int
	mutex   sync.Mutex
	rwMutex sync.RWMutex
}

func New_account(initial_balance int) *Bank_Account {
	return &Bank_Account{
		balance: initial_balance,
	}
}

// regular mutex
// Exclusive access for both read and write
func (a *Bank_Account) Transfer(amt int) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	// simulations
	time.Sleep(100 * time.Millisecond)
	a.balance += amt
}

func (a *Bank_Account) GetBalance() int {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return a.balance
}

// RWMutex
// Allows multiple read but a single write
func (a *Bank_Account) GetBalanceRW() int {
	a.rwMutex.Lock()
	defer a.rwMutex.Unlock()
	return a.balance
}

func (a *Bank_Account) TransferRW(amount int) {
	a.rwMutex.Lock() // Exclusive lock for writing
	defer a.rwMutex.Unlock()

	time.Sleep(100 * time.Millisecond)
	a.balance += amount
}

func main() {
	// Example 1: Regular Mutex
	fmt.Println("Regular Mutex Example:")
	account := New_account(1000)

	// Using waitgroup
	var wg sync.WaitGroup
	start_time := time.Now()

	// 5 concurrent transfers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			new_balance := account.GetBalance()
			fmt.Printf("Balance: %d\n", new_balance)
		}()
	}

	wg.Wait()
	fmt.Printf("Regular Mutex took: %v\n\n", time.Since(start_time))

	// RWMutex
	fmt.Println("RWMutex Example:")
	accountRW := New_account(1000)
	start_time = time.Now()

	// Start 5 concurrent transfers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			accountRW.TransferRW(100)
		}()
	}

	// Start 20 concurrent balance checks (more readers to demonstrate RWMutex advantage)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			balance := accountRW.GetBalanceRW()
			fmt.Printf("Balance (RW): %d\n", balance)
		}()
	}

	wg.Wait()
	fmt.Printf("RWMutex took: %v\n", time.Since(start_time))
}
