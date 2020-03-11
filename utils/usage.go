package utils

import (
	"fmt"
	"os"
)

// Print the valid way of using the program
func Usage() {
	fmt.Printf("\nPlease include a filename of an Antfarm:\n\n	go build  --->  ./lem-in <map.txt>\n\nYou can make your own map and place it in 'maps/' folder\n\n")
	os.Exit(0)
}
