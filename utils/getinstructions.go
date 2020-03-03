package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Open the file to get instructions for ant farm
func GetInstructions(g *Graph, file string) []string {
	var arr []string
	a, err := os.Open("maps/" + file)
	if err != nil {
		fmt.Println("There is no such file... :(")
		os.Exit(0)
	}
	defer a.Close()

	bf := bufio.NewReader(a)
	for {
		line, err := bf.ReadString('\n')
		if err == io.EOF {
			arr = append(arr, line)
			break
		}
		arr = append(arr, line[:len(line)-1])
	}
	return arr
}
