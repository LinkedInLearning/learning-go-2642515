package main

import (
	"fmt"
)

// Einfaches Struct
type Congressman struct {
	salary int    // Privates Feld
	Name   string // Öffentliches Feld
}

func main() {
	// -- Leerer Wert Nil

	// Pointer ist nil
	var p *Congressman
	fmt.Println("Pointer p:", p, " is nil: ", p == nil)

	// Slice ist nil
	var s []float64
	fmt.Println("Slice s:", s, " is nil: ", s == nil)

	// Map ist nil
	var m map[string]float64
	fmt.Println("Map m:", m, " is nil: ", m == nil)
}
