package utils

import (
	"fmt"
	"os"
)

// Handles invalid data format error
func ErrHandler() {
	fmt.Println("\nERROR: invalid data format\n")
	os.Exit(0)
}
