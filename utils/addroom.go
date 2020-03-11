package utils

// Adds valid room to the graph
func AddRoom(g *Graph, a []string) {
	s, e := false, false
	sCount, eCount := 0, 0
	for i, v := range a {
		// * Look for start & end rooms
		if sCount > 0 && IsStart(v) ||
			eCount > 0 && IsEnd(v) {
			ErrHandler()
		}
		if IsStart(v) && i < len(a)-1 {
			s = true
			sCount++
			continue
		} else if IsEnd(v) && i < len(a)-1 {
			e = true
			eCount++
			continue
		}
		// * Add Room
		g.AddNode(IsRoom(v), s, e)
		s, e = false, false

	}
	if sCount == 0 || eCount == 0 {
		ErrHandler()
	}
}
