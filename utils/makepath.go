package utils

import (
	"fmt"
)

var Paths [][]*Room

// Gets all available path via Breadth First Search (BFS)
func MakePath(g *Graph) []*Path {
	// TODO:
	// 1. MakePath(startNode, *Graph)

	// * Get number of initial paths
	// var num int
	var start *Room
	// var end *Room
	for _, v := range g.Rooms {
		if v.start {
			start = v
		}
	}
	DFS(start)
	comb := make([]*Path, len(Paths))
	for i := range comb {
		comb[i] = &Path{0, 1, nil}
	}
	for i, v := range Paths {
		fmt.Println(len(Paths))
		comb[i].route = v
		fmt.Println("I:", i)
		for _, r := range v {
			fmt.Printf("%s ", r.Name)
		}
		fmt.Println()
	}

	// DFS(start, Paths)
	// for i, v := range comb {
	// 	fmt.Println("PATH", i)
	// 	for _, k := range v.route {
	// 		fmt.Printf("ROOM: %s ", k.Name)
	// 	}
	// 	fmt.Println()
	// }
	// paths := []*Path{}
	// paths = DFS(g)
	// for _, v := range paths {
	// 	fmt.Println(v.ID)
	// 	for _, j := range v.route {
	// 		fmt.Printf("room:%s\n", j.Name)
	// 	}
	// }
	// for i := range p {
	// 	p[i].route, ok = BFS(g)
	// 	if !ok {
	// 		continue
	// 	}
	// 	// * If there is a path, make path nodes Visited, except start & end
	// 	for _, j := range p[i].route {
	// 		if j.start != true && j.end != true {
	// 			Vis = append(Vis, j.Name)
	// 		}
	// 	}
	// 	for k := range Visited {
	// 		Visited[k] = false
	// 	}
	// 	for _, v := range Vis {
	// 		Visited[v] = true
	// 	}
	// }

	// * Get only valid paths
	// pathTrue := []*Path{}
	// for _, v := range Paths {
	// 	if len(v.route) > 0 {
	// 		pathTrue = append(pathTrue, v)
	// 	}
	// }
	// if len(pathTrue) == 0 {
	// 	// ErrHandler()
	// 	fmt.Println("pe4al'")
	// 	os.Exit(0)
	// }
	return comb
}
