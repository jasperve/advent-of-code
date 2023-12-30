package graph

import (
	"sort"
	"fmt"

	//"blackmesa.nl/Jasper/advent-of-code-2023/Libraries/navigation"
)

type identifyTraversableNode func(*Node) bool
type prioritizeEdges func(map[int]*Edge) []int


type ListItem struct {
	Parent *ListItem
	Node *Node
	Direction string
	G int
	H int
}

type byF []ListItem
func (c byF) Len() int {
	return len(c)
}
func (c byF) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c byF) Less(i, j int) bool {
	return c[i].G + c[i].H < c[j].G + c[j].H
}


//Path comment
type Path struct {
	distance int
	nodes    []*Node
}

//GetDistance comment
func (p *Path) GetDistance() int {
	return p.distance
}

//GetNodes comment
func (p *Path) GetNodes() []*Node {
	return p.nodes
}

//ByWeight comment
func ByWeight(edges map[int]*Edge) []int {

	keys := make([]int, len(edges), len(edges))
	i := 0
	for k := range edges {
		keys[i] = k
		i++
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return edges[keys[i]].GetDetail("weight").(int) < edges[keys[j]].GetDetail("weight").(int)
	})

	return keys

}

//GetPathsToNodes comment
func GetPathsToNodes(start *Node, targets map[int]*Node, tnFn identifyTraversableNode, peFn prioritizeEdges) map[*Node]Path {

	paths := make(map[*Node]Path)

	openList := NewQueue()
	openList.Set(start, []*Node{}, 0)
	closedList := make(map[*Node]struct{})

	for openList.Len() > 0 {

		node, path, distance := openList.Next()

		closedList[node] = struct{}{}
		for _, target := range targets {
			if target == node {
				paths[target] = Path{nodes: path, distance: distance}
			}
		}

		//MOGELIJK NIET NODIG START TE CHECKEN
		//Check if we can traverse further via this node
		if !tnFn(node) && node != start {
			continue
		}

		edges := node.graph.GetEdges("from", node)
		edgeKeys := peFn(edges)

		for _, edgeKey := range edgeKeys {
		
			if _, ok := closedList[edges[edgeKey].GetTo()]; ok {
				continue
			}

			if existingDistance, ok := openList.GetDistance(edges[edgeKey].GetTo()); ok && existingDistance <= distance+edges[edgeKey].GetDetail("weight").(int) {
				continue
			}

			pathCopy := make([]*Node, len(path))
			copy(pathCopy, path)
			pathCopy = append(pathCopy, edges[edgeKey].GetTo())

			openList.Set(edges[edgeKey].GetTo(), pathCopy, distance+edges[edgeKey].GetDetail("weight").(int))

		}

	}

	return paths

}




func AStar(start *Node, target *Node) ListItem {

	startLocation := start.GetDetail("location").(navigation.Coordinate)
	targetLocation := target.GetDetail("location").(navigation.Coordinate)

	openList := []ListItem{ListItem{Parent: nil, Node: start, G: 0, H: abs(startLocation.Y, targetLocation.Y) + abs(startLocation.X, targetLocation.X)}}
	closedList := map[int]ListItem{}

	for len(openList) > 0 {

		// Sort the openlist and get current item from the openList
		sort.Sort(byF(openList))
		currentItem := openList[0]
		openList = openList[1:]

		// Add current item to the openList
		closedList[currentItem.Node.GetID()] = currentItem

		l := currentItem.Node.GetDetail("location").(navigation.Coordinate)

		if l.Y == 0 && l.X == 8 {
			fmt.Println(l, currentItem.G, currentItem.Direction)
			fmt.Println(currentItem.Parent.Node.GetDetail("location").(navigation.Coordinate), currentItem.Parent.Parent.Node.GetDetail("location").(navigation.Coordinate), currentItem.Parent.Parent.Parent.Node.GetDetail("location").(navigation.Coordinate), currentItem.Parent.Parent.Parent.Parent.Node.GetDetail("location").(navigation.Coordinate))
		}

		edges := currentItem.Node.graph.GetEdges("from", currentItem.Node)
		for _, edge := range edges {

			// Calculate F value
			neighbourLocation := edge.GetTo().GetDetail("location").(navigation.Coordinate)
			g := currentItem.G + edge.GetDetail("weight").(int)
			h := abs(neighbourLocation.Y, targetLocation.Y) + abs(neighbourLocation.X, targetLocation.X)

			if _, ok := closedList[edge.GetTo().GetID()]; ok && closedList[edge.GetTo().GetID()].G == g {
				continue
			}

			direction := edge.GetDetail("direction").(string)

			// Check for reverse direction
			if currentItem.Parent != nil { 
				switch direction {
					case "up": if currentItem.Parent.Direction == "down" { continue }
					case "right": if currentItem.Parent.Direction == "left" { continue }
					case "down": if currentItem.Parent.Direction == "up" { continue }
					case "left": if currentItem.Parent.Direction == "right" { continue }
				}
			}
			
			// If this is the third movement in this direction this edge is NOT valid so continue with the next edge
			
			if 	currentItem.Direction == direction &&
				currentItem.Parent != nil && currentItem.Parent.Direction == direction &&
				currentItem.Parent.Parent != nil && currentItem.Parent.Parent.Direction == direction {
				//	fmt.Println("Same direction, so continue with next edge")
				continue
			}
		
			// Target found
			if edge.GetTo() == target {
				fmt.Println("ROUTE TO TARGET FOUND!")
				return ListItem{Parent: &currentItem, Node: edge.GetTo(), Direction: direction, G: g, H: h}
			}

			openList = append(openList, ListItem{Parent: &currentItem, Node: edge.GetTo(), Direction: direction, G: g, H: h})

		}

	}

	return ListItem{}

}


func abs(a, b int) int {
	if a < b {
	   return b - a
	}
	return a - b
 }