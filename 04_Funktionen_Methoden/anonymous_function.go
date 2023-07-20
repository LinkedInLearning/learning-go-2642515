package main

import (
	"fmt"
	"net/http"
)

func runIt(f func()) {
	f()
}

func main() {
	//-- Anonyme Funktionen

	// Anonyme Funktion definieren + ausführen
	func() {
		fmt.Print("Hello LinkedIn!")
	}()

	// Anonyme Funktion als Variable/Parameter
	f := func() {
		fmt.Println("Gopher!")
	}
	runIt(f)

	// Anonyme HTTP Handler Function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Gopher Web!"))
	})
	http.ListenAndServe(":8000", nil)
}
