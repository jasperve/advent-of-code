package graphbuilder

import (
//	"blackmesa.nl/Jasper/advent-of-code-2023/Libraries/graph"
//	"blackmesa.nl/Jasper/advent-of-code-2023/Libraries/navigation"
)

type identifyNode func(byte) bool
type identifyEdgeWeight func(*graph.Node, *graph.Node) int

//Dot function Identifies nodes by dot '.'
func Dot(b byte) bool {
	if b == '.' {
		return true
	}
	return false
}

//Letter function Identifies nodes by type letter 'A' - 'Z' or letter 'a' - 'z'
func Letter(b byte) bool {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return true
	}
	return false
}

//UppercaseLetter function Identifies nodes by type letter 'A' - 'Z' or letter 'a' - 'z'
func UppercaseLetter(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

//LowercaseLetter function Identifies nodes by type letter 'A' - 'Z' or letter 'a' - 'z'
func LowercaseLetter(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}

//Number function Identifies nodes by number '0' - '9'
func Number(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}


//None returns false
func None(b byte) bool {
	return false
}

//GridToWeightedGraph generated a weighted graph specified by functions
func GridToWeightedGraph(input []string, inFn identifyNode, iewFn identifyEdgeWeight) graph.WeightedGraph {

	// Coordinate format y, x
	directions := map[string][]int {
		"up": []int{-1, 0},
		"right": []int{0, 1},
		"down": []int{1, 0},
		"left": []int{0, -1},
	}

	opposites := map[string]string {
		"up": "down",
		"right": "left",
		"down": "up",
		"left": "right",
	}

	g := graph.InitWeightedGraph()

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {

			if inFn(input[y][x]) {

				node := g.AddNode()
				node.AddDetail("location", navigation.Coordinate{X: x, Y: y})
				node.AddDetail("type", input[y][x])

				for direction, coordinateChange := range directions {
					neighbourNodes := g.GetNodes("location", navigation.Coordinate{Y: y+coordinateChange[0], X: x+coordinateChange[1]}) 
					for _, neighbourNode := range neighbourNodes {

						edge := g.AddEdge(node, neighbourNode)
						edge.AddDetail("weight", iewFn(node, neighbourNode))
						edge.AddDetail("direction", direction)

						edge = g.AddEdge(neighbourNode, node)
						edge.AddDetail("weight", iewFn(neighbourNode, node))
						edge.AddDetail("direction", opposites[direction])

					}
				}

			}

		}
	}

	return g

}