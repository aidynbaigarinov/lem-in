package utils

// Returns new graph
func New() *Graph {
	return &Graph{
		Rooms: []*Room{},
	}
}

// Graph method to add room to graph
func (g *Graph) AddNode(name string, start, end bool) {
	g.Rooms = append(g.Rooms, &Room{
		Name:  name,
		start: start,
		end:   end,
		Conn:  []*Room{},
	})
}

// Room method to add connection
func (r *Room) AddConn(c *Room) {
	r.Conn = append(r.Conn, c)
}
