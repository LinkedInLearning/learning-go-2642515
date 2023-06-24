package main

import (
	"log"
)

type Dog struct {
	Name string
}
type Cat struct {
	Name string
}

type Congressman[T any] struct {
	Name string
	Pet  T
}

func main() {
	p := Congressman[Cat]{Name: "Peter Russo", Pet: Cat{"Kitty"}}
	j := Congressman[Dog]{Name: "Jacky Sharp", Pet: Dog{"Hasso"}}

	log.Println(p)
	log.Println(j)
}
