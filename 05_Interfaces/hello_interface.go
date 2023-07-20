package main

import "fmt"

type Congressman struct {
	Name string
}

type Enemy struct{}

// Interface definieren
type Greeter interface {
	greet()
}

// Interface nutzen
func passBy(g1 Greeter, g2 Greeter) {
	g1.greet()
	g2.greet()
}

func main() {
	c := Congressman{Name: "Frank U."}
	e := Enemy{}
	passBy(c, e)
}

// Interface Implementierung 1
func (c Congressman) greet() {
	fmt.Println("Hello", c.Name)
}

// Interface Implementierung 2
func (e Enemy) greet() {
	fmt.Println("Go to hell!")
}
