package main

import (
	"fmt"
)

func main() {
	for dogAge := 1; dogAge < 11; dogAge++ {
		humanAge := dogAge * 7
		fmt.Printf("Dog: %3d Human: %5d\n", dogAge, humanAge)
	}
}
