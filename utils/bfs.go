package utils

// Implements Breadth First Search on Graph
func BFS(start *Room) bool {
	var q = NewQueue()
	// path := []*Room{}
	// for _, r := range g.Rooms {
	// 	if r.start == true {
	Visited[start.Name] = true
	q.Enqueue(start)
	// 	}
	// }

	for !q.IsEmpty() {
		v, _ := q.Dequeue()
		// if err != nil {
		// 	ErrHandler()
		// }
		// * Save path
		if v.end {
			// path = SavePath(v, path)
			return true
		}
		// * Add neighbours rooms to queue
		for _, a := range v.Conn {
			if !Visited[a.Name] {
				Visited[a.Name] = true
				a.Parent = v
				q.Enqueue(a)
			}
		}
	}
	return false
}
