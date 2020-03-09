package utils

var been []*Room

var pathIndex int

func beenThere(v *Room) bool {
	for _, s := range been {
		if s == v {
			return true
		}
	}
	return false
}

// Implements Depth First Search on Graph
func DFS(r *Room) {
	if r.end {
		newPath := make([]*Room, len(been))
		copy(newPath, been)
		newPath = append(newPath, r)
		Paths = append(Paths, newPath)
		return
	}
	been = append(been, r)
	for _, v := range r.Conn {
		if !beenThere(v) {
			DFS(v)
		}
	}
	been = been[:len(been)-1]
	Visited[r.Name] = false
	return
}
