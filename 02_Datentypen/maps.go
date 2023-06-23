package main

import (
	"fmt"
)

func main() {
	// -- Map
	// Map bildet wie ein Wörterbuch einen Schlüssel (key) auf einen Wert (value) ab.

	// Map mit Werten initialisieren
	var lastNames = map[string]string{"Frank": "Underwood", "Zoe": "Barnes"}

	fmt.Println("Last Names:", lastNames)
	fmt.Println("Last Name of Frank:", lastNames["Frank"])

	// Map leer initialisieren
	var age = make(map[string]int)
	age["Frank"] = 52
	age["Zoe"] = 28

	fmt.Println("Age:", age)
	fmt.Println("Age Length", len((age)))
}
