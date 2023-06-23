package main

import "fmt"

func main() {
	// -- Schleifen

	// Klassische for Schleife
	fmt.Println("Classic for Loop")
	for i := 1; i < 7; i++ {
		fmt.Println("Season", i)
	}
	fmt.Println("")

	// While Schleife
	fmt.Println("While Loop")
	j := 1
	for j < 7 {
		fmt.Println("Season", j)
		j++
	}

	// Endlos Schleife
	fmt.Println("Infinite Loop")
	k := 1
	for {
		if k == 10 {
			k++
			fmt.Println("No Season 10")
			continue
		}

		if k == 15 {
			break
		}

		fmt.Println("Season", k)
		k++
	}

	// For-Each Schleife mit range
	fmt.Println("For Each Loop")
	supportingCharacters := []string{"Freddy Hayes", "Donald Blythe"}

	for index, character := range supportingCharacters {
		fmt.Println("Character", character, "at", index)
	}
	// Ignoriere index
	for _, character := range supportingCharacters {
		fmt.Println("Character", character)
	}

}
