package main

import "fmt"

func main() {
	// -- Arrays
	names := [2]string{"Zoe", "?"}
	names[1] = "Frank"

	fmt.Println(names)

	// -- Slices

	// Slice mit make intialisieren
	mainCharacters := make([]string, 2)
	mainCharacters[0] = "Frank"
	mainCharacters[1] = "Clair"

	fmt.Println("Main Characters:", mainCharacters)

	// Slice mit Werten initialisieren
	supportingCharacters := []string{"Freddy", "Donald"}
	fmt.Printf("supportingCharacters = %v\n", supportingCharacters)
	fmt.Printf("length = %d\n", len(supportingCharacters))
	fmt.Printf("capacity = %d\n", cap(supportingCharacters))

	// Neues Element an Slice anhängen
	supportingCharacters = append(supportingCharacters, "Hector Mendoza", "Nancy Kaufberger")

	fmt.Println("Featured Parts:", supportingCharacters)

}
