package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	// -- Mit dem io.Reader Interface JSON verarbeiten
	const jsonBody = `{ "message": "Hello Reader!" }`

	// Reader erstellen
	r := strings.NewReader(jsonBody)

	var jsonMap map[string]string

	// JSON per Reader einlesen
	json.NewDecoder(r).Decode(&jsonMap)

	fmt.Println(jsonMap)
}
