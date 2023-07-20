package main

import "fmt"

func helloDefer() {
	defer fmt.Println("Gopher")
	fmt.Print("Hello ")
}

func helloDeferWithPanic() {
	defer fmt.Println("Gopher")

	var s = []string{}

	// Zugriff auf ungültigen Index
	fmt.Println("Hello ", s[42])
}

func main() {
	//-- Verzögerte Auführung mit dever
	helloDefer()
}
