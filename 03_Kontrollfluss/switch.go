package main

import (
	"fmt"
	"time"
)

func main() {
	// Switch über Typ
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Bloody Monday")
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Switch mit Expressions
	proVoteCount := 301
	switch {
	case proVoteCount < 300:
		fmt.Println("Not enough.")
	case proVoteCount == 301:
		fmt.Println("Draw.")
	default:
		fmt.Println("We've got the votes.")
	}

	// Switch mit Fallthrough
	// Fallthrough gibt die Kontrolle an nächsten Case, auch bei Match.
	s := 3000
	switch {
	case s < 1000:
		fmt.Println("Salary s is less than 10")
		fallthrough
	case s < 4000:
		fmt.Println("Salary s is less than 4000")
		fallthrough
	case s < 8000:
		fmt.Println("Salary s is less than 8000")
	}
}
