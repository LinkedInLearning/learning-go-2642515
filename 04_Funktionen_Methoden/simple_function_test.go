package main

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	dogAge := DogAge(20)
	if dogAge != 7 {
		t.Fatalf("No match")
	}
}
