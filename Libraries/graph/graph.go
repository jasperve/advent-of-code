package graph

import "math/rand"

//Graph comment
type Graph struct {
	nodeDict dictType
	edgeDict dictType
}

//InitGraph function
func InitGraph() Graph {
	g := Graph{}
	g.nodeDict = make(dictType)
	g.edgeDict = make(dictType)
	return g
}

//WeightedGraph comment
type WeightedGraph struct {
	Graph
}

//InitWeightedGraph function
func InitWeightedGraph() WeightedGraph {
	g := WeightedGraph{}
	g.nodeDict = make(dictType)
	g.edgeDict = make(dictType)
	return g
}

//AddEdge comment
func (g *Graph) AddEdge(from *Node, to *Node) *Edge {
	id := rand.Int()
	edge := &Edge{graph: g, id: id, from: from, to: to, details: make(map[string]interface{})}
	g.edgeDict.add("id", id, edge)
	g.edgeDict.add("from", from, edge)
	g.edgeDict.add("to", to, edge)
	return edge
}

//GetEdges comment
func (g *Graph) GetEdges(key string, value interface{}) map[int]*Edge {
	edges := make(map[int]*Edge)
	for k, v := range g.edgeDict.get(key, value) {
		edges[k] = v.(*Edge)
	}
	return edges
}

//GetEdge function
func (g *Graph) GetEdge(id int) *Edge {
	for _, v := range g.GetEdges("id", id) {
		return v
	}
	return nil
}

//AddNode function will add a new node to several maps based on its details for quick recovery of the nodes.
func (g *Graph) AddNode() *Node {
	id := rand.Int()
	node := &Node{graph: g, id: id, details: make(map[string]interface{})}
	g.nodeDict.add("id", id, node)
	return node
}

//GetNodes function
func (g *Graph) GetNodes(key string, value interface{}) map[int]*Node {
	nodes := make(map[int]*Node)
	for k, v := range g.nodeDict.get(key, value) {
		nodes[k] = v.(*Node)
	}
	return nodes
}

//GetNode function
func (g *Graph) GetNode(id int) *Node {
	for _, v := range g.GetNodes("id", id) {
		return v
	}
	return nil
}
