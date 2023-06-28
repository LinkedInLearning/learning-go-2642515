package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function we'll run in every goroutine.
func worker(id int) {
	fmt.Printf("I, worker %d, am starting\n", id)

	// Schlafen als Simulation aufwändiger Verarbeitung
	time.Sleep(time.Second)

	fmt.Printf("I, worker %d, am done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Launch several goroutines and increment the WaitGroup
	// counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		// Obacht: Die Loop Variable i teilen sich alle Goroutinen, und das führt zu Fehlern,
		// daher neue Variable erstellen (siehe https://golang.org/doc/faq#closures_and_goroutines)
		i := i

		// Worker in eigener Goroutine starten
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Warten bis alle Goroutinen fertig sind
	wg.Wait()


	// INFO: Für eine WaitGroup mit expliziter Fehlerbehandlung
	// siehe https://pkg.go.dev/golang.org/x/sync/errgroup
}
