package main

import (
	"fmt"
)

type Congressman struct {
	Name           string
	AccountBalance float64
}

// Fehler als Rückgabewert
func (c Congressman) bribe(amount float64) error {
	if c.Name != "Peter" {
		return fmt.Errorf("%v is not corrupt", c.Name)
	}

	c.AccountBalance += amount

	fmt.Println(c.Name, "has", c.AccountBalance)

	return nil
}

func main() {
	c := Congressman{Name: "Jackie", AccountBalance: 8000.0}

	// Fehler behandeln
	err := c.bribe(5000.0)
	if err != nil {
		fmt.Println("could not bribe:", err.Error())
	}
}
