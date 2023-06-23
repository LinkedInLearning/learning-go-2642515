package main

import "fmt"

func main() {
	// -- Arrays
	var finishedSeasons = [6]int{1, 2, 3, 4, 5, 99}
	finishedSeasons[5] = 6
	names := [2]string{"Zoe", "Frank"}

	fmt.Println(finishedSeasons)
	fmt.Println(names)

	// -- Slices
	mainCharacters := make([]string, 2)
	mainCharacters[0] = "Frank Underwood"
	mainCharacters[1] = "Clair Underwood"

	fmt.Println("Main Characters:", mainCharacters)

	supportingCharacters := []string{"Freddy Hayes", "Donald Blythe"}
	fmt.Printf("supportingCharacters = %v\n", supportingCharacters)
	fmt.Printf("length = %d\n", len(supportingCharacters))
	fmt.Printf("capacity = %d\n", cap(supportingCharacters))

	supportingCharacters = append(supportingCharacters, "Hector Mendoza", "Nancy Kaufberger")

	fmt.Println("Featured Parts:", supportingCharacters)

}
