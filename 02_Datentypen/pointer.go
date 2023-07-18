package main

import (
	"fmt"
)

// Einfaches Struct
type Congressman struct {
	salary int    // Private Property
	Name   string // Öffentliche Property
}

func main() {
	// -- Pointer
	c := Congressman{3000, "Peter Russo"}
	d := c

	// Pointer p auf Struct c
	var p *Congressman = &c

	// Pointer mit * auflösen und Wert zuweisen
	(*p).Name = "Bob"
	p.Name = "Bob" // kurze Schreibweise für Structs

	fmt.Println("c:", c)
	fmt.Println("Pointer p:", p)
	fmt.Println("Ohne Pointer d:", d)
}
