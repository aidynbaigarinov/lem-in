package utils

var Paths [][]*Room

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

	DFS(start)
	sort(Paths)
	comb := FindComb(Paths)
	ret := OptimalComb(comb, num)
	// fmt.Println(ret)
	allPaths := []*Path{}
	for _, v := range ret {
		tmp := &Path{route: v}
		allPaths = append(allPaths, tmp)
	}
	return allPaths
}
