package utils

// Checks if the node was visited before
func beenThere(v *Room) bool {
	for _, s := range Been {
		if s == v {
			return true
		}
	}
	return false
}

//var cache []*Room

// Implements Depth First Search on Graph
func DFS(r *Room) {
	if r.end {
		//cache = append(cache, Been...)
		newPath := make([]*Room, len(Been))
		copy(newPath, Been)
		newPath = append(newPath, r)
		newPath = newPath[1:]
		Paths = append(Paths, newPath)
		return
	}
	Been = append(Been, r)
	for _, v := range r.Conn {
		if !beenThere(v) {
			DFS(v)
		}
	}
	Been = Been[:len(Been)-1]
	return
}
