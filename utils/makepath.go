package utils

// implements bubble sort
func sort(p [][]*Room) {
	var (
		n = len(Paths)
	)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if len(Paths[j]) > len(Paths[j+1]) {
				Paths[j], Paths[j+1] = Paths[j+1], Paths[j]
			}
		}
	}
}

// Gets all available path via Breadth First Search (BFS)
func MakePath(start *Room, num int) []*Path {

	// Uses Depth First Search to traverse the graph
	DFS(start)

	// Sorts all found paths by its lengths
	sort(Paths)

	// Find all possible combinations of paths that don't intersect
	comb := FindComb(Paths)

	// Find one optimal combination of paths, depending on number of ants
	ret := OptimalComb(comb, num)
	allPaths := []*Path{}
	for _, v := range ret {
		tmp := &Path{route: v}
		allPaths = append(allPaths, tmp)
	}
	return allPaths
}
