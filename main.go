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
	Name string
	Conn []*Room
}

func New() *Graph {
	return &Graph{
		Rooms: []*Room{},
	}
}

func (g *Graph) AddNode(name string) {
	g.Rooms = append(g.Rooms, &Room{
		Name: name,
		Conn: []*Room{},
	})
}

func (r *Room) AddConn(c *Room) {
	r.Conn = append(r.Conn, c)
}

// func (g *Graph) Neighbors(id int) []int {
// 	neighbors := []int{}
// 	for _, node := range g.nodes {
// 		for edge := range node.edges {
// 			if node.id == id {
// 				neighbors = append(neighbors, edge)
// 			}
// 			if edge == id {
// 				neighbors = append(neighbors, node.id)
// 			}
// 		}
// 	}
// 	return neighbors
// }

// func (g *Graph) Nodes() []int {
// 	nodes := make([]int, len(g.nodes))
// 	for i := range g.nodes {
// 		nodes[i] = i
// 	}
// 	return nodes
// }

// func (g *Graph) Edges() [][3]int {
// 	edges := make([][3]int, 0, len(g.nodes))
// 	for i := range g.nodes {
// 		for k, v := range g.nodes[i].edges {
// 			edges = append(edges, [3]int{i, k, int(v)})
// 		}
// 	}
// 	return edges
// }

func getRooms(g *Graph) []string {
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

// if r, ok := isRoom(line); ok {
// 	g.AddNode(r)
// }
// if c, ok := isConn(line); ok {
// 	if len(c) > 0 {

// 		for _, v := range g.Rooms {
// 			if v.Name == c[0] {
// 				for _, k := range g.Rooms {
// 					if k.Name == c[1] {
// 						v.Conn = append(v.Conn, k)
// 					}
// 				}

// 			}
// 			if v.Name == c[1] {
// 				for _, k := range g.Rooms {
// 					if k.Name == c[0] {
// 						v.Conn = append(v.Conn, k)
// 					}
// 				}

// 			}
// 		}

// 	}
// }

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
			ret = append(ret, s[:i], s[i+1:len(s)-1])
		}
	}
	return ret, true
}

func main() {
	graph := New()
	arr := getRooms(graph)
	fmt.Println(arr)

}
