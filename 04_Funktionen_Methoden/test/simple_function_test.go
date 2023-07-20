package main

import (
	"testing"
)

func TestDogAge(t *testing.T) {
	wantDogAge := 3
	dogAge := DogAge(21)
	if dogAge != wantDogAge {
		t.Fatalf("No match, expected %v got %v", wantDogAge, dogAge)
	}
}
