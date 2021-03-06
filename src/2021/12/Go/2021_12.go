package main

import (
	"aoc/libs/go/inputParse"

	"fmt"
	"strings"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	var result int
	graph := parseInput(lines)

	// part 1
	result = findResult(&graph, false)
	fmt.Println(result)

	// part 2
	result = findResult(&graph, true)
	fmt.Println(result)
}

type Vertex struct {
	Key      string
	Vertices map[string]*Vertex
	isSmall  bool
}

type Graph struct {
	Vertices map[string]*Vertex
}

func parseInput(slicesOfLines []string) (graph Graph) {
	graph = Graph{map[string]*Vertex{}}

	for i := 0; i < len(slicesOfLines); i++ {
		names := strings.Split(slicesOfLines[i], "-")
		// create empty node if not already
		for j := 0; j < 2; j++ {
			if _, ok := graph.Vertices[names[j]]; !ok {
				isSmall := 'a' <= names[j][0] && names[j][0] <= 'z'
				vertices := map[string]*Vertex{}
				vertex := Vertex{names[j], vertices, isSmall}
				graph.Vertices[names[j]] = &vertex
			}
		}

		// connect the two
		graph.Vertices[names[0]].Vertices[names[1]] = graph.Vertices[names[1]]
		graph.Vertices[names[1]].Vertices[names[0]] = graph.Vertices[names[0]]
	}
	return
}

const inputPath = "../input.txt"

func countPaths(g *Graph, current *Vertex, end *Vertex,
	visitedTimes map[string]int, oneSmallCaveTwiceOption bool) (count int) {

	visitedTimes[current.Key]++
	if current.Key == end.Key {
		return 1
	}

	if oneSmallCaveTwiceOption {
		for k, v := range visitedTimes {
			if g.Vertices[k].isSmall && v > 1 {
				oneSmallCaveTwiceOption = false
				break
			}
		}
	}

	for _, v := range current.Vertices {
		if numberVisits, ok := visitedTimes[v.Key]; ok && v.isSmall && numberVisits > 0 {
			if !oneSmallCaveTwiceOption {
				continue
			}

			if oneSmallCaveTwiceOption && v.Key == "start" {
				continue
			}
		}
		count += countPaths(g, v, end, visitedTimes, oneSmallCaveTwiceOption)
		visitedTimes[v.Key]--
	}

	return count
}

func findResult(graph *Graph, oneSmallCaveTwiceOption bool) (result int) {

	start := graph.Vertices["start"]
	end := graph.Vertices["end"]
	visited := map[string]int{}

	result = countPaths(graph, start, end, visited, oneSmallCaveTwiceOption)

	return
}
