package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	utils "./utils"
)

func main() {

	start := time.Now()

	graph := utils.New()

	// * Get a filename
	farm := os.Args[1:]
	if len(farm) != 1 {
		fmt.Println("Please include a filename of an Antfarm")
		return
	}

	// * Get instructions
	arr := utils.GetInstructions(graph, farm[0])

	// * Number of ants
	antsNum, err := strconv.Atoi(arr[0])
	if err != nil || antsNum <= 0 {
		utils.ErrHandler()
	}
	utils.AddRoom(graph, arr)

	// * Build connections between rooms
	graph = utils.BuildConn(graph, arr)

	// * BFS algo to find paths
	pathTrue := utils.MakePath(graph)
	// utils.PrintInstructions(arr)

	// * deploy ants!!!
	utils.AntPath(pathTrue, antsNum)
	utils.TimeTaken(start, "Lem-In")
}
