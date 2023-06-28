package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Channel erzeugen
	money := make(chan int)

	go func() {
		// 3. Select auf 2 Channels
		select {
		case amount := <-money:
			fmt.Println("Received", amount, "$!")
		case <-time.After(1 * time.Second):
			fmt.Println("Got nothing ...!")
		}
	}()

	// Kein Send an Channel money

	// 2. Warten
	time.Sleep(2 * time.Second)

}
