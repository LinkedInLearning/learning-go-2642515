package main

import "fmt"

// Benannte Return Werte, aka "naked return"
func humanAge(dogAge int) (humanAge int) {
	humanAge = dogAge * 7
	return
}

// Mehrere Parameter und mehrere Rückgabewerte
func humanAges(dog1, dog2 int) (int, int) {
	return humanAge(dog1), humanAge(dog2)
}

func main() {
	hassoAge := 3
	fmt.Println("Human Age of Hasso:", humanAge(hassoAge))

	human1, human2 := humanAges(3, 5)
	fmt.Println("Human Ages:", human1, ",", human2)
}
