package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Graph struct {
	Rooms []*Room
}

type Room struct {
	Name   string
	Parent *Room
	start  bool
	end    bool
	Conn   []*Room
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

func (q *Queue) Peek() (*Room, error) {
	if len(q.data) == 0 {
		return nil, fmt.Errorf("Queue is Empty")
	}
	return q.data[0], nil
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

func buildConn(g *Graph, arr []string) *Graph {
	for _, s := range arr {
		if len(s) > 0 {

			if c, ok := isConn(s); ok {
				if len(c) > 0 {
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
	}
	return g
}

func isStart(s string) bool {
	if s[:2] == "##" && s[2:] == "start" {
		return true
	}
	return false
}
func isEnd(s string) bool {
	if s[:2] == "##" && s[2:] == "end" {
		return true
	}
	return false
}

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
	path = append(path, r)
	// fmt.Println("r ", r.Name, "visited ", visited[r.Name])
	if r.start != true && r.end != true {
		visited[r.Name] = true
	}

	if r.start == true {
		return path
	}
	return SavePath(r.Parent, path)
}

var visited = make(map[string]bool)

func BFS(g *Graph) []*Room {
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
		// fmt.Println("v deq: ", v.Name)
		if err != nil {
			fmt.Println(err)
		}
		if v.end == true {
			for k := range visited {
				visited[k] = false
			}
			return SavePath(v, path)
		}
		for _, a := range v.Conn {
			if visited[a.Name] == false {
				visited[a.Name] = true
				a.Parent = v
				q.Enqueue(a)
				// fmt.Println("Parent: ", v.Name, "a: ", a.Name)
			}
		}
	}
	return nil
}

func main() {
	graph := New()
	arr := getInstructions(graph)
	s := false
	e := false
	// fmt.Println(arr)
	for _, v := range arr {

		if len(v) == 7 && isStart(v) {
			s = true
			continue
		} else if len(v) == 5 && isEnd(v) {
			e = true
			continue
		}
		if len(v) > 0 {
			if r, ok := isRoom(v); ok {
				graph.AddNode(r, s, e)
				s, e = false, false
			}
		}
	}
	graph = buildConn(graph, arr)
	// for _, v := range graph.Rooms {
	// 	fmt.Println("room name: ", v.Name, "start: ", v.start, "end: ", v.end)
	// 	for _, c := range v.Conn {
	// 		fmt.Println("  conn name: ", c.Name)
	// 	}
	// }
	var num int
	for _, v := range graph.Rooms {
		if v.start == true {
			num = len(v.Conn)
		}
	}
	fmt.Println(num)
	path := make([][]*Room, num)
	fmt.Println("path", path)
	for i := range path {
		path[i] = BFS(graph)
	}
	for i, v := range path {
		fmt.Println("path ", i)
		for j := len(v) - 1; j >= 0; j-- {
			fmt.Printf("%s ", v[j].Name)
		}
		fmt.Println()
	}
}
