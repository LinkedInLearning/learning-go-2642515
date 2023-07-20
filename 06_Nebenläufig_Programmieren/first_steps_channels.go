package main

import "fmt"

func main() {
	//-- Goroutinen und Channels

	// 1. Channel erzeugen
	money := make(chan int)

	go func() {
		// 3. Receive auf Channel money
		amount := <-money
		fmt.Println("Received", amount, "$!")
	}()

	// 2. Send an Channel money
	money <- 2000
}
