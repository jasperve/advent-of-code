package main

import (

	"fmt"
	"log"
	"slices"

	"github.com/jasperve/libraries/file"
	
)

type node struct {
	value byte
	number int
	ident int
}

func main() {

	nodes := map[int]map[int]*node{}

	lines, err := file.ReadLines("input-test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	for y := 0; y < len(lines); y++ {
		nodes[y] = map[int]*node{}
		for x := 0; x < len(lines[y]); x++ {
			nodes[y][x] = &node{value: lines[y][x]}
		}
	}

	ident := 0

	for y := 0; y < len(nodes); y++ {
		beginX, endX, number := -1, -1, 0
		for x := 0; x < len(nodes[y]); x++ {
			if nodes[y][x].value >= '0' && nodes[y][x].value <= '9' {
				if beginX == -1 { beginX = x } 
				endX = x
				number *= 10
				number += int(nodes[y][x].value) - 48
			} else {
				if beginX != -1 && endX != -1 && number != 0 {
					for i := beginX; i <= endX; i++ { nodes[y][i].number = number; nodes[y][i].ident = ident }
					beginX, endX, number = -1, -1, 0
					ident++
				}
			}
		}
	}

	total := 0

	neighbours := [][]int{[]int{-1, 0}, []int{-1, 1}, []int{0, 1}, []int{1, 1}, []int{1, 0}, []int{1, -1}, []int{0, -1}, []int{-1, -1}}

	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {

			number := 1

			if nodes[y][x].value == '*' {

				foundIdents := []int{}
				for _, neighbour := range neighbours {
					if v, ok := nodes[y+neighbour[0]][x+neighbour[1]]; ok && nodes[y+neighbour[0]][x+neighbour[1]].number > 0 { 
						if !slices.Contains(foundIdents, v.ident) {
							number *= v.number; found++ 
							foundIdents = append(foundIdents, v.ident)
						}
					} 
				}
				fmt.Println(foundIdents)

				if len(foundIdents) == 2 {
					fmt.Println(number)
					total += number
				}
		
			}
		}
	}
	fmt.Println(total)


}