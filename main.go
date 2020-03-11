package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	utils "./utils"
)

func main() {
	startTime := time.Now() // benchmark

	graph := utils.New()
	// Get a filename
	farm := os.Args[1:]
	if len(farm) != 1 {
		utils.Usage()
	}
	// Get instructions
	arr := utils.GetInstructions(graph, farm[0])

	// Number of ants
	antsNum, err := strconv.Atoi(arr[0])
	if err != nil || antsNum <= 0 {
		utils.ErrHandler()
	}
	// Check instructions
	aR, aC := utils.Check(arr[1:])

	// Add Rooms
	utils.AddRoom(graph, aR)

	// Build connections between rooms
	start := utils.BuildConn(graph, aC)

	// DFS algorithm to find paths
	paths := utils.MakePath(start, antsNum)

	// deploy ants!!!
	utils.AntPath(paths, antsNum, arr)
	fmt.Println("\n", time.Now().Sub(startTime))
}
