package utils

// Gets all available path via Breadth First Search (BFS)
func MakePath(g *Graph) []*Path {
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
		// * If there is a path, make path nodes Visited, except start & end
		for _, j := range p[i].route {
			if j.start != true && j.end != true {
				Vis = append(Vis, j.Name)
			}
		}
		for k := range Visited {
			Visited[k] = false
		}
		for _, v := range Vis {
			Visited[v] = true
		}
	}

	// * Get only valid paths
	pathTrue := []*Path{}
	for _, v := range p {
		if len(v.route) > 0 {
			pathTrue = append(pathTrue, v)
		}
	}
	if len(pathTrue) == 0 {
		ErrHandler()
	}
	return pathTrue
}
