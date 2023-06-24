package main

import "fmt"

// Interface definieren
type Greeter interface {
	greet()
}

// Interface nutzen
func passBy(c1 Greeter, c2 Greeter) {
	c1.greet()
	c2.greet()
}

func main() {
	c := Congressman{Name: "Frank U."}
	e := Enemy{}
	passBy(c, e)
}

type Congressman struct {
	Name string
}

// Interface Implementierung 1
func (c Congressman) greet() {
	fmt.Println("Hello", c.Name)
}

type Enemy struct{}

// Interface Implementierung 2
func (e Enemy) greet() {
	fmt.Println("Go to hell!")
}
