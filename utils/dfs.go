package utils

import "fmt"

var array [][]*Room

// Implements Depth First Search on Graph
func DFS(g *Graph) []*Path {
	rooms := []*Room{}
	for _, r := range g.Rooms {
		if r.start == true {
			rec(r, rooms)
		}
	}

	for _, v := range array {
		fmt.Println("path:")
		for _, z := range v {
			fmt.Print("room:", z.Name, " ")
		}
		fmt.Println()
	}

	return nil
}

func rec(r *Room, path []*Room) [][]*Room {
	Visited[r.Name] = true
	path = append(path, r)
	if r.end {
		fmt.Println("end:", r.Name)
		// path = append(path, r)
		copy := make([]*Room, len(path))
		for i := range path {
			copy[i] = path[i]
		}
		array = append(array, path)
	} else {
		for _, v := range r.Conn {
			if !Visited[v.Name] {
				rec(v, path)
			}
		}
	}
	path = path[:len(path)-1]
	Visited[r.Name] = false
	return array
}
