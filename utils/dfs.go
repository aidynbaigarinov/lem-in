package utils

import "fmt"

func Save(r *Room) []*Room {
	Visited[r.Name] = false
	if r.start {
		Visited[r.Name] = false
		return nil
	}
	Save(r.Parent)
	path = append(path, r)
	return path
}

var path []*Room

var pathIndex int

// Implements Depth First Search on Graph
func DFS(r *Room, Paths []*Path) []*Path {
	fmt.Println("R:::", r.Name)
	if r.end {
		newPath := make([]*Room, len(path))
		copy(newPath, path)
		newPath = append(newPath, r)
		Paths[pathIndex].route = newPath
		pathIndex++
		// Visited[r.Name] = false
		return Paths
	}
	if r.start {
		for k := range Visited {
			Visited[k] = false
		}
	}
	Visited[r.Name] = true
	path = append(path, r)
	for _, v := range r.Conn {
		fmt.Println("  V:::", v.Name)
		ok := Visited[v.Name]
		if !ok {
			Paths = DFS(v, Paths)
		}
	}
	path = path[:len(path)-1]
	return Paths
}
