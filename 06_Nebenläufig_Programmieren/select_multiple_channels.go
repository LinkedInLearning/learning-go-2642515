package main

import (
	"fmt"
	"time"
)

func main() {
	account1 := make(chan int)
	account2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		account1 <- 1000
	}()

	go func() {
		time.Sleep(2 * time.Second)
		account2 <- 2000
	}()

	for i := 0; i < 2; i++ {
		select {
		case amount1 := <-account1:
			fmt.Println("Received", amount1, "$ on account 1!")
		case amount2 := <-account2:
			fmt.Println("Received", amount2, "$ on account 2!")
		case <-time.After(10 * time.Second):
			fmt.Println("Got nothing ...!")
		}
	}
}
