package utils

import "fmt"

// Print instructions
func PrintInstructions(a []string) {
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println()
}
