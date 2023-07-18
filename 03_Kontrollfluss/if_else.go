package main

import "fmt"

func main() {
	// -- If/Else

	// Einfaches if
	president := "Walker"
	if "Underwood" != president {
		fmt.Println("Underwood is not President (yet).")
	}

	// If mit else
	proVoteCount := 193
	if proVoteCount < 300 {
		fmt.Println("Not enough.")
	} else if proVoteCount == 301 {
		fmt.Println("Draw.")
	} else {
		fmt.Println("We've got the votes.")
	}
}
