package utils

import "fmt"

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

// New queue
type Queue struct {
	data []*Room
}

// Make new queue
func NewQueue() *Queue {
	return &Queue{
		data: []*Room{},
	}
}

// Check is a queue empty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Enqueue method for queue
func (q *Queue) Enqueue(n *Room) *Queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue method for queue
func (q *Queue) Dequeue() (*Room, error) {
	if len(q.data) == 0 {
		return nil, fmt.Errorf("Queue is Empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}
