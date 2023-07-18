package main

import "fmt"

// Konstante für Package
const greetingPart2 = "Hello"

// Variable für Package
var aPackageVar string = "?"

func main() {
	// -- Variablen

	// Deklaration und Zuweisung in langer Form
	var s1 string = "a string"

	// Deklaration und Zuweisung in kurzer Form
	s2 := "another string"

	// Deklaration und Zuweisung trennen
	var s3 string          // Deklaration
	s3 = "and yet another" // Zuweisung des Werts

	fmt.Println("Variablen:", s1, s2, s3)

	// Vorbelegung mit "leeren" Werten
	var s0 string
	var i0 int
	var b0 bool
	fmt.Println("Leere Variablen", s0, i0, b0)

	// -- Konstanten

	// Lokale Konstante
	const greetingPart1 = "Hello"
	fmt.Println("Konstanten:", greetingPart1, greetingPart2)
}
