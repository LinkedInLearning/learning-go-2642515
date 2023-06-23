package main

import (
	"fmt"
)

// Einfaches Struct
type Congressman struct {
	salary int    // Private Property
	Name   string // Öffentliche Property
}

// Struct mit Embedding
type President struct {
	Congressman // Embedded Struct

	NuclearWeaponCode string
}

func main() {
	// -- Structs
	c := Congressman{3000, "Peter Russo"}
	c = Congressman{Name: "Peter Russo", salary: 3000}

	fmt.Println("Congressman:", c)

	p := President{NuclearWeaponCode: "**code123**"}
	p.Name = "Frank Underwood"

	fmt.Println("President:", p)

	var e Congressman
	fmt.Println("e:", e)
}
