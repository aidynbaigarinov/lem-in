package utils

import "fmt"

// Implements Breadth First Search on Graph
func BFS(g *Graph) ([]*Room, bool) {
	var q = NewQueue()
	path := []*Room{}
	for _, r := range g.Rooms {
		if r.start == true {
			Visited[r.Name] = true
			q.Enqueue(r)
		}
	}

	for !q.IsEmpty() {
		v, err := q.Dequeue()
		if err != nil {
			fmt.Println(err)
		}
		// * Save path
		if v.end == true {
			path = SavePath(v, path)
			return path, true
		}
		// * Add neighbours rooms to queue
		for _, a := range v.Conn {
			if Visited[a.Name] == false {
				Visited[a.Name] = true
				a.Parent = v
				q.Enqueue(a)
			}
		}
	}
	return nil, false
}
