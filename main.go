package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
func getInstructions(g *Graph) []string {
	var arr []string
	a, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
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
		path = append(path, r)
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

/*
* * * Adds connections between rooms * * *
 */
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
		}
	}

	// for _, v := range ants {
	// 	fmt.Println("ant ID", v.ID)
	// 	for _, k := range v.route.route {
	// 		fmt.Printf("%s ", k.Name)
	// 	}
	// 	fmt.Println()
	// }
	z := len(p[0].route) - 1
	end := p[0].route[z].Name
	printResult(ants, len(p), end)
}

func printResult(a []Ant, l int, end string) {
	// z := len(a) - l + 1
	m := make(map[int]int)
	counter := 0
	z := 0
	for z < 10 {
		for _, v := range a {
			// v.roomNum = v.roomNum + 1
			// k := 0
			ok := false
			if _, ok = m[v.ID]; !ok {
				m[v.ID] = 1
			} else {
				m[v.ID]++

				// fmt.Println("ant ID", v.ID, "room number", m[v.ID])
				if v.roomNum < len(v.route.route)-2 && !v.route.route[m[v.ID]].busy {
					fmt.Printf("L%d-%s ", v.ID, v.route.route[m[v.ID]].Name)
					v.route.route[m[v.ID]].busy = true
				} else {
					v.route.route[m[v.ID]].busy = false
				}
				// fmt.Println("name", v.route.route[m[v.ID]].Name)
				if v.route.route[m[v.ID]].Name == end {
					counter++
				}
			}
		}
		fmt.Println()
		if counter == len(a) {
			break
		}
		z++
	}
}

var vis []string

func main() {
	graph := New()
	/*
	* * * Get instructions * * *
	 */
	arr := getInstructions(graph)
	/*
	* * * Number of ants * * *
	 */
	antsNum, _ := strconv.Atoi(arr[0])
	s := false
	e := false
	// fmt.Println(arr)
	for _, v := range arr {
		/*
		* * * Look for start & end rooms * * *
		 */
		if len(v) == 7 && isStart(v) {
			s = true
			continue
		} else if len(v) == 5 && isEnd(v) {
			e = true
			continue
		}
		/*
		* * * Add Room * * *
		 */
		if len(v) > 0 {
			if r, ok := isRoom(v); ok {
				graph.AddNode(r, s, e)
				s, e = false, false
			}
		}
	}
	/*
	* * * Build connections between rooms * * *
	 */
	graph = buildConn(graph, arr)

	/*
	* * * Get number of initial routes, depending on number of connections from start room * * *
	 */
	var num int
	for _, v := range graph.Rooms {
		if v.start == true {
			num = len(v.Conn)
		}
	}

	p := make([]*Path, num)
	ok := false
	for i, _ := range p {
		p[i] = &Path{0, 1, nil}
	}
	/*
	* * * BFS algo to find paths * * *
	 */
	for i := range p {
		p[i].route, ok = BFS(graph)
		if !ok {
			continue
		}
		/*
		* * * If there is a path, make path nodes visited, except start & end * * *
		 */

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

	// for i, v := range pathTrue {
	// 	fmt.Println("path ", i)
	// 	for _, j := range v.route {
	// 		fmt.Printf(" " + j.Name)
	// 	}
	// 	fmt.Println()
	// }

	antPath(pathTrue, antsNum)
}
