package utils

// Adds valid room to the graph
func AddRoom(g *Graph, a []string) {
	s, e := false, false
	for _, v := range a {
		// * Look for start & end rooms
		if len(v) == 7 && IsStart(v) {
			s = true
			continue
		} else if len(v) == 5 && IsEnd(v) {
			e = true
			continue
		}
		// * Add Room
		if len(v) > 0 {
			if r, ok := IsRoom(v); ok {
				g.AddNode(r, s, e)
				s, e = false, false
			}
		}
	}
}
