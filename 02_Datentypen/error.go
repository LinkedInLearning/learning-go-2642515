package main

import (
	"fmt"
)

func main() {
	// -- Error
	err := fmt.Errorf("user %v not found", "Frank")
	fmt.Println("Error err:", err)
}
