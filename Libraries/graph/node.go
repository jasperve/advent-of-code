package graph

//Node struct
type Node struct {
	graph   *Graph
	id      int
	details map[string]interface{}
}

//AddDetail function
func (n *Node) AddDetail(key string, value interface{}) {
	if _, ok := n.details[key]; !ok {
		n.graph.nodeDict.add(key, value, n)
		n.details[key] = value
	}
}

//UpdateDetail function
func (n *Node) UpdateDetail(key string, value interface{}) {
	if _, ok := n.details[key]; ok {
		if n.details[key] != value {
			n.graph.nodeDict.remove(key, n.details[key], n)
			n.graph.nodeDict.add(key, value, n)
			n.details[key] = value
		}
	}
}

//RemoveDetail function
func (n *Node) RemoveDetail(key string) {
	if _, ok := n.details[key]; ok {
		n.graph.nodeDict.remove(key, n.details[key], n)
		delete(n.details, key)
	}
}

//GetDetail function
func (n *Node) GetDetail(key string) interface{} {
	if _, ok := n.details[key]; ok {
		return n.details[key]
	}
	return nil
}

//GetID function
func (n *Node) GetID() int {
	return n.id
}

//GetGraph function 
func (n *Node) GetGraph() *Graph {
	return n.graph
}
