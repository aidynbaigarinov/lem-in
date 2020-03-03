package utils

import (
	"fmt"
	"os"
)

func ErrHandler() {
	fmt.Println("ERROR: invalid data format")
	os.Exit(0)
}
