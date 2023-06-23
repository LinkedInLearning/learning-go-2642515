package main

import "fmt"

// Private Funktion
func humanAge(dogAge int) int {
	return dogAge * 7
}

// Öffentliche Funktion
func DogAge(humanAge int) int {
	return humanAge / 7
}

func main() {
	hassoAge := 3
	fmt.Println("Human Age of Hasso:", humanAge(hassoAge))

	dougAge := 45
	fmt.Println("Dog Age of Doug:", DogAge(dougAge))
}
