package main

import (
	"fmt"
	"math/rand"
	"time"
)

func speaker(name string, debate chan int) {
	for {
		microphone := <-debate // Auf Mikro warten (Nachricht empfangen)

		fmt.Printf("Turn %v: %v says '%v'\n", microphone, name, randomAnswer())
		time.Sleep(400 * time.Millisecond)

		microphone++
		debate <- microphone // Mikro zurückgeben (Nachricht senden)
	}
}

func randomAnswer() string {
	answers := []string{"I'm right.", "Take this argument.", "But I've got that.", "You b***"}
	return answers[rand.Intn(len(answers)-1)]
}

func main() {
	// Channel deklarieren und initialisieren
	debate := make(chan int)

	go speaker("Jackie", debate)
	go speaker("Frank", debate)

	microphone := 1

	// Mikro geben und Diskussion starten
	debate <- microphone

	// Diskussion für diese Dauer laufen lassen
	time.Sleep(2 * time.Second)

	// Mikro nehmen und Diskussion beenden
	<-debate
}
