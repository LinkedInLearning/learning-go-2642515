package main

import (
	"testing"
)

func TestDogAge(t *testing.T) {
	wantHumanAge := 3
	dogAge := DogAge(21)
	if dogAge != wantHumanAge {
		t.Fatalf("No match, expected %v got %v", wantHumanAge, dogAge)
	}
}
