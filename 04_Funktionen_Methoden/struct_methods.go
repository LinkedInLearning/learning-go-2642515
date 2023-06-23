package main

import "fmt"

type Congressman struct {
	Name string
}

// Receiver Function für Congressman
func (c Congressman) swearOathOfOffice() {
	fmt.Printf("I, %v, swear to serve the USA.\n", c.Name)
}

func main() {
	// Aufruf Receiver Function an Struct
	c := Congressman{Name: "Peter Russo"}
	c.swearOathOfOffice()

	// Aufruf Receiver Function an Pointer
	p := &Congressman{Name: "Garret Walker"}
	p.swearOathOfOffice()
}
