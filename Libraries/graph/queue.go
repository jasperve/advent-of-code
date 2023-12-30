package graph

import "sort"

//Queue comment
type Queue struct {
	nodes     map[int]*Node
	paths     map[int][]*Node
	distances map[int]int
	ids       []int
}

// Len is part of sort.Interface
func (q *Queue) Len() int {
	return len(q.ids)
}

// Swap is part of sort.Interface
func (q *Queue) Swap(i, j int) {
	q.ids[i], q.ids[j] = q.ids[j], q.ids[i]
}

// Less is part of sort.Interface
func (q *Queue) Less(i, j int) bool {
	return q.distances[q.ids[i]] < q.distances[q.ids[j]]
}

// Set updates or inserts a new id in the priority queue
func (q *Queue) Set(node *Node, path []*Node, distance int) {

	id := node.GetID()

	if _, ok := q.nodes[id]; !ok {
		q.ids = append(q.ids, id)
	}
	q.nodes[id] = node
	q.paths[id] = path
	q.distances[id] = distance

	// Sort queue so shortest distances get picked first by Next function
	sort.Sort(q)

}

// Next removes the element with the lowest distance from the queue and returns the node, path and distance
func (q *Queue) Next() (*Node, []*Node, int) {
	
	id := q.ids[0]
	q.ids = q.ids[1:]

	node := q.nodes[id]
	path := q.paths[id]
	distance := q.distances[id]
	delete(q.nodes, id)
	delete(q.paths, id)
	delete(q.distances, id)
	return node, path, distance

}

// GetDistance returns the ditsance to a node in the queue
func (q *Queue) GetDistance(id *Node) (int, bool) {
	distance, ok := q.distances[id.GetID()]
	return distance, ok
}

// IsEmpty returns true when the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.ids) == 0
}

// NewQueue creates a new empty priority queue
func NewQueue() *Queue {
	var q Queue
	q.nodes = make(map[int]*Node)
	q.paths = make(map[int][]*Node)
	q.distances = make(map[int]int)
	return &q
}
