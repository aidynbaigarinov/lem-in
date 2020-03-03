package utils

import (
	"log"
	"time"
)

// Counts execution time
func TimeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("\n------------------------------------------------------------------------------\nTIME: %s took %s\n", name, elapsed)
}
