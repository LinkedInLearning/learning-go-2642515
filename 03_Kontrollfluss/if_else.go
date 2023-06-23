package main

import "fmt"

func main() {
	// Einfaches if
	president := "Walker"
	if "Frank" != president {
		fmt.Println("Frank is not President (yet).")
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
