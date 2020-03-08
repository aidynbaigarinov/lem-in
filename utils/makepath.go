package utils

import "fmt"

// Gets all available path via Breadth First Search (BFS)
func MakePath(g *Graph) []*Path {
	// TODO:
	// 1. MakePath(startNode, *Graph)

	// * Get number of initial paths
	var num int
	var start *Room
	for _, v := range g.Rooms {
		if v.end {
			num = len(v.Conn)
		} else if v.start {
			start = v
		}
	}
	p := make([]*Path, num)
	// ok := false
	for i, _ := range p {
		p[i] = &Path{0, 1, nil}
	}
	Paths := make([]*Path, 10)

	for i := 0; i < 10; i++ {
		Paths[i] = &Path{0, 1, nil}
	}

	DFS(start, Paths)
	for i, v := range Paths {
		fmt.Println("PATH", i)
		for _, k := range v.route {
			fmt.Printf("ROOM: %s ", k.Name)
		}
		fmt.Println()
	}
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
	pathTrue := []*Path{}
	for _, v := range Paths {
		if len(v.route) > 0 {
			pathTrue = append(pathTrue, v)
		}
	}
	if len(pathTrue) == 0 {
		ErrHandler()
		// fmt.Println("pe4al'")
	}
	return pathTrue
}
