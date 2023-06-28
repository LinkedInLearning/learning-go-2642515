package main

import "fmt"

func main() {
	// 1. Channel erzeugen
	money := make(chan int)

	go func() {
		// 3. Receive auf Channel money
		amount := <-money
		fmt.Println("Received", amount, "$!")
	}()

	// 3. Send an Channel money
	money <- 2000
}
