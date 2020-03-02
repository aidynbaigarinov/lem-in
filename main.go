package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Graph struct {
	Rooms []*Room
}

type Room struct {
	Name   string
	Parent *Room
	start  bool
	end    bool
	busy   bool
	Conn   []*Room
}

type Path struct {
	antCounter int
	ID         int
	route      []*Room
}

type Ant struct {
	ID      int
	roomNum int
	finish  bool
	route   *Path
}

func New() *Graph {
	return &Graph{
		Rooms: []*Room{},
	}
}

func (g *Graph) AddNode(name string, start, end bool) {
	g.Rooms = append(g.Rooms, &Room{
		Name:  name,
		start: start,
		end:   end,
		Conn:  []*Room{},
	})
}

func (r *Room) AddConn(c *Room) {
	r.Conn = append(r.Conn, c)
}

type Queue struct {
	data []*Room
}

func NewQueue() *Queue {
	return &Queue{
		data: []*Room{},
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Enqueue(n *Room) *Queue {
	q.data = append(q.data, n)
	return q
}

func (q *Queue) Dequeue() (*Room, error) {
	if len(q.data) == 0 {
		return nil, fmt.Errorf("Queue is Empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}

/*
* * * Opens the file to get instructions for ant farm * * *
 */
func getInstructions(g *Graph, file string) []string {
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

/*
* * * Adds connections between rooms * * *
 */
func buildConn(g *Graph, arr []string) *Graph {

	for _, s := range arr {

		if len(s) > 2 {

			if c, ok := isConn(s); len(c) == 2 && ok {
				for _, v := range g.Rooms {

					if v.Name == c[0] {
						for _, k := range g.Rooms {
							if k.Name == c[1] {
								v.Conn = append(v.Conn, k)
							}
						}
					}

					if v.Name == c[1] {
						for _, k := range g.Rooms {
							if k.Name == c[0] {
								v.Conn = append(v.Conn, k)
							}
						}
					}

				}
			}
		}
	}
	return g
}

/*
* * * Checks is it a start room * * *
 */
func isStart(s string) bool {
	if s[:2] == "##" && s[2:] == "start" {
		return true
	}
	return false
}

/*
* * * Checks is it an end room* * *
 */
func isEnd(s string) bool {
	if s[:2] == "##" && s[2:] == "end" {
		return true
	}
	return false
}

/*
* * * Checks is it a room * * *
 */
func isRoom(s string) (string, bool) {
	if s[0] == '#' || s[0] == 'L' {
		return "", false
	}
	for i, v := range s {
		if v == ' ' && i > 0 {
			return s[:i], true
		}
	}
	return "", false
}

/*
* * * Check is it a connection between the rooms & returns the array of rooms* * *
 */
func isConn(s string) ([]string, bool) {
	var ret []string
	if s[0] == '#' || s[0] == 'L' {
		return nil, false
	}
	for i, v := range s {
		if v == '-' {
			ret = append(ret, s[:i], s[i+1:len(s)])
		}
	}
	return ret, true
}

func SavePath(r *Room, path []*Room) []*Room {
	/*
	* * * Recursively, through parents, add rooms to the path * * *
	 */
	if r.start == true {
		// path = append(path, r)
		return path
	}
	if r.start != true && r.end != true {
		vis = append(vis, r.Name)
	}
	path = SavePath(r.Parent, path)
	path = append(path, r)
	return path
}

var visited = make(map[string]bool)

func BFS(g *Graph) ([]*Room, bool) {
	var q = NewQueue()
	path := []*Room{}
	for _, r := range g.Rooms {
		if r.start == true {
			visited[r.Name] = true
			q.Enqueue(r)
		}
	}

	for !q.IsEmpty() {
		v, err := q.Dequeue()
		if err != nil {
			fmt.Println(err)
		}
		/*
		* * * Save path * * *
		 */
		if v.end == true {
			path = SavePath(v, path)
			return path, true
		}
		/*
		* * * Add neighbours rooms to queue * * *
		 */
		for _, a := range v.Conn {
			if visited[a.Name] == false {
				visited[a.Name] = true
				a.Parent = v
				q.Enqueue(a)
				// fmt.Println("Parent: ", v.Name, "a: ", a.Name)
			}
		}
	}
	return nil, false
}

// * Add connections between rooms
func antPath(p []*Path, aN int) {
	var ants = []Ant{}
	for i := 0; i < aN; i++ {
		ants = append(ants, Ant{ID: i + 1})
	}
	ants[0].route = p[0]
	p[0].antCounter++
	currPath := 0
	for i, l := 1, len(ants); i < l; i++ {
		if currPath == len(p)-1 {
			currPath = 0
		}
		// fmt.Println(currPath)
		if currPath < len(p)-1 {
			if (len(p[currPath].route) + p[currPath].antCounter) >
				(len(p[currPath+1].route) + p[currPath+1].antCounter) {
				currPath++
				ants[i].route = p[currPath]
				p[currPath].antCounter++
			} else {
				ants[i].route = p[currPath]
				p[currPath].antCounter++
			}
		} else {
			ants[i].route = p[currPath]
		}
	}

	z := len(p[0].route) - 1
	end := p[0].route[z].Name
	// fmt.Println("end:", end)
	printResult(ants, end)
}

func printResult(a []Ant, end string) {
	// m := make(map[int]bool)
	for !a[len(a)-1].finish {
		for i, l := 0, len(a); i < l; i++ {
			// ok := m[a[i].ID]
			// if !ok {
			if !a[i].finish {
				// fmt.Println(a[i].route)
				if a[i].route.route[a[i].roomNum].Name == end {
					fmt.Printf("L%d-%s ", a[i].ID, a[i].route.route[a[i].roomNum].Name)
					// m[a[i].ID] = true
					if a[i].roomNum > 0 {
						a[i].route.route[a[i].roomNum-1].busy = false
					} else {
						a[i].route.route[a[i].roomNum].busy = false
					}
					a[i].finish = true

				} else if !a[i].route.route[a[i].roomNum].busy {
					fmt.Printf("L%d-%s ", a[i].ID, a[i].route.route[a[i].roomNum].Name)
					a[i].route.route[a[i].roomNum].busy = true
					if a[i].roomNum > 0 {
						a[i].route.route[a[i].roomNum-1].busy = false
					}
					a[i].roomNum++
				}
			}
			// }
		}
		fmt.Println()
	}
}

func printInstructions(a []string) {
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println()
}

var vis []string

func TimeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("\n------------------------------------------------------------------------------\nTIME: %s took %s\n", name, elapsed)
}

func addRoom(g *Graph, a []string) {
	s, e := false, false
	for _, v := range a {
		// * Look for start & end rooms
		if len(v) == 7 && isStart(v) {
			s = true
			continue
		} else if len(v) == 5 && isEnd(v) {
			e = true
			continue
		}
		// * Add Room
		if len(v) > 0 {
			if r, ok := isRoom(v); ok {
				g.AddNode(r, s, e)
				s, e = false, false
			}
		}
	}
}

func makePath(g *Graph) []*Path {
	// * Get number of initial paths
	var num int
	for _, v := range g.Rooms {
		if v.start == true {
			num = len(v.Conn)
		}
	}
	p := make([]*Path, num)
	ok := false
	for i, _ := range p {
		p[i] = &Path{0, 1, nil}
	}

	for i := range p {
		p[i].route, ok = BFS(g)
		if !ok {
			continue
		}
		// * If there is a path, make path nodes visited, except start & end
		for _, j := range p[i].route {
			if j.start != true && j.end != true {
				vis = append(vis, j.Name)
			}
		}
		for k := range visited {
			visited[k] = false
		}
		for _, v := range vis {
			visited[v] = true
		}
	}
	pathTrue := []*Path{}

	for _, v := range p {
		if len(v.route) > 0 {
			pathTrue = append(pathTrue, v)
		}
	}
	return pathTrue
}

func main() {
	// TODO: refactor code, organize packages, handle errors and invalid inputs

	start := time.Now()

	graph := New()

	// * Get a filename
	farm := os.Args[1:]
	if len(farm) != 1 {
		fmt.Println("Please include a filename of an Antfarm")
		return
	}

	// * Get instructions
	arr := getInstructions(graph, farm[0])
	printInstructions(arr)

	// * Number of ants
	antsNum, _ := strconv.Atoi(arr[0])
	addRoom(graph, arr)

	// * Build connections between rooms
	graph = buildConn(graph, arr)

	// * BFS algo to find paths
	pathTrue := makePath(graph)

	// * deploy ants!!!
	antPath(pathTrue, antsNum)
	TimeTaken(start, "Lem-In")
}
