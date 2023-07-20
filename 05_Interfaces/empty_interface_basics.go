package main

import "fmt"

func main() {
	//-- Das leere Interface

	var i interface{} = "Hello"

	//  Type Assertion ohne Prüfung
	var s string = i.(string)
	fmt.Println(s)

	//  Type Assertion mit Prüfung
	s, ok := i.(string)
	fmt.Println(s, "is a string? ", ok)

	// Falsche Type Assertion ohne Prüfung erzeugt Panic
	var f float64 = i.(float64)
	fmt.Println(f)
}
