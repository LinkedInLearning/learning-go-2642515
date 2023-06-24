package main

import (
	"errors"
	"fmt"
)

type Congressman struct {
	Name           string
	AccountBalance float64
}

// Fehler als Rückgabewert
func (c Congressman) bribe(amount float64) error {
	if c.Name != "Peter Russo" {
		// Ohne Parameter
		// return errors.New("Not corrupt!")

		// Mit Parameter
		return fmt.Errorf("%v is not corrupt", c.Name)
	}

	// Fehler weitergeben
	err := transferMoney(amount)
	if err != nil {
		// Wrap error mit %w
		return fmt.Errorf("could not transfer: %w", err)
	}

	c.AccountBalance += amount
	return nil
}

func transferMoney(amount float64) error {
	return fmt.Errorf("could not transfer %v", amount)
}

func main() {
	c := Congressman{Name: "Jackie Sharp", AccountBalance: 8000.0}

	// Fehler behandeln
	err := c.bribe(5000.0)
	if err != nil {
		fmt.Println("could not bribe:", err.Error())
	}

	c2 := Congressman{Name: "Peter Russo", AccountBalance: -10.0}
	err = c2.bribe(5000.0)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Root error:", errors.Unwrap(err))
	}
}
