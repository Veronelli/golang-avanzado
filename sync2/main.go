package main

import (
	"fmt"
	"sync"
)

var (
	balance uint = 100
)

func Deposit(amount uint, wg *sync.WaitGroup, lock *sync.Mutex) {
	wg.Done()
	b := balance
	balance = b + amount

}

func Balance() uint {
	b := balance
	return b

}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := uint(0); i <= 5; i++ {
		wg.Add(1)
		go Deposit(i*100, &wg, &lock)

	}
	wg.Wait()
	fmt.Println(Balance())
}
