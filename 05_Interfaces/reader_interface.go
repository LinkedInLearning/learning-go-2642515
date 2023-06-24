package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	const jsonBody = `{ "message": "Hello Reader!" }`

	// Reader erstellen
	r := strings.NewReader(jsonBody)

	var j map[string]string

	// JSON per Reader einlesen
	json.NewDecoder(r).Decode(&j)

	fmt.Println(j)
}
