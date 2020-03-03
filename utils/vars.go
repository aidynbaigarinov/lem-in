package utils

// Global visited variable
var Visited = make(map[string]bool)

// variable after the path is found
var Vis []string

// New graph
type Graph struct {
	Rooms []*Room
}

// New room structure
type Room struct {
	Name   string
	Parent *Room
	start  bool
	end    bool
	busy   bool
	Conn   []*Room
}

// Path structure
type Path struct {
	antCounter int
	ID         int
	route      []*Room
}

// Ant structure
type Ant struct {
	ID      int
	roomNum int
	finish  bool
	route   *Path
}
