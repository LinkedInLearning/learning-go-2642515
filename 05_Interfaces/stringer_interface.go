package main

import "fmt"

type Congressman struct {
	Name string
}

func (c Congressman) String() string {
	fmt.Println("Hello", c.Name)
}

func main() {
	c := Congressman{Name: "Frank U."}
	fmt.Println(c)
}
