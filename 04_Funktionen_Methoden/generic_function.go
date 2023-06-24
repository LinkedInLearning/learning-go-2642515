package main

import (
	"log"
)

func humanAge[T int | float64](dogAge T) T {
	return dogAge * 7
}

func main() {
	// Generische Funktion mit Typ int aufrufen
	var i int = humanAge[int](3)
	log.Println(i)

	// Generische Funktion mit Typ float aufrufen
	var f float64 = humanAge[float64](4.9)
	log.Println(f)
}
