package main

import "fmt"

// Benannte Return Werte, aka "naked return"
func humanAgeV2(dogAge int) (humanAge int) {
	humanAge = dogAge * 7
	return
}

// Mehrere Rückgabewerte
func humanAges(dog1, dog2 int) (int, int) {
	return humanAgeV2(dog1), humanAgeV2(dog2)
}

func main() {
	hassoAge := 3
	fmt.Println("Human Age of Hasso:", humanAgev2(hassoAge))

	human1, human2 := humanAges(3, 5)
	fmt.Println("Human Ages:", human1, ",", human2)
}
