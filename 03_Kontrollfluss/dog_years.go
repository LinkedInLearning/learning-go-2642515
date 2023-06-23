package main

import (
	"fmt"
)

func main() {
	fmt.Printf("| %3s | %3s |\n", "Dog", "Human")
	fmt.Printf("|-------------|\n")
	for dogAge := 1; dogAge < 10; dogAge++ {
		humanAge := dogAge * 7
		fmt.Printf("| %3d | %5d |\n", dogAge, humanAge)

	}
	fmt.Printf("|-------------|\n")
}
