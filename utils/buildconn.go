package utils

// Add connections between rooms
func BuildConn(g *Graph, a []string) *Room {
	var start *Room

	for _, s := range a {
		c := IsConn(s)
		for _, v := range g.Rooms {
			if v.start {
				start = v
			}
			if v.Name == c[0] {
				for _, k := range g.Rooms {
					if k.Name == c[1] {
						v.Conn = append(v.Conn, k)
					}
				}
			}

			if v.Name == c[1] {
				for _, k := range g.Rooms {
					if k.Name == c[0] {
						v.Conn = append(v.Conn, k)
					}
				}
			}

		}

	}
	return start
}
