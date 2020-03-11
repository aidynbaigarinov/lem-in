package utils

// Global visited variable
var Been []*Room

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

var Ret [][][]*Room

var Paths [][]*Room
