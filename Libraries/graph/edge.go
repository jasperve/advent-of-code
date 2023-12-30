package graph

//Edge comment
type Edge struct {
	graph   *Graph
	id      int
	from    *Node
	to      *Node
	details map[string]interface{}
}

//AddDetail function
func (e *Edge) AddDetail(key string, value interface{}) {
	if _, ok := e.details[key]; !ok {
		e.graph.edgeDict.add(key, value, e)
		e.details[key] = value
	}
}

//UpdateDetail function
func (e *Edge) UpdateDetail(key string, value interface{}) {
	if _, ok := e.details[key]; ok {
		if e.details[key] != value {
			e.graph.edgeDict.remove(key, e.details[key], e)
			e.graph.edgeDict.add(key, value, e)
			e.details[key] = value
		}
	}
}

//RemoveDetail function
func (e *Edge) RemoveDetail(key string) {
	if _, ok := e.details[key]; ok {
		e.graph.edgeDict.remove(key, e.details[key], e)
		delete(e.details, key)
	}
}

//GetDetail function
func (e *Edge) GetDetail(key string) interface{} {
	if _, ok := e.details[key]; ok {
		return e.details[key]
	}
	return nil
}

//GetID function
func (e *Edge) GetID() int {
	return e.id
}

//GetFrom function
func (e *Edge) GetFrom() *Node {
	return e.from
}

//GetTo function
func (e *Edge) GetTo() *Node {
	return e.to
}
