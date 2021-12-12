package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int
	mapOfNodes := parseInput(lines)

	// part 1
	result = findResult(&mapOfNodes)
	fmt.Println(result)

	// part 2
	// result = findResult(parsedInput)
	// fmt.Println(result)
}

type Vertex struct {
	Key      string
	Vertices map[string]*Vertex // undirected graph
	isSmall  bool
}

type Graph struct {
	Vertices map[string]*Vertex
}

func parseInput(slicesOfLines []string) (graph Graph) {
	// adjacency list
	// map of nodes
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
		map0Connections := graph.Vertices[names[0]].Vertices
		node1 := graph.Vertices[names[1]]
		map0Connections[names[1]] = node1

		map1Connections := graph.Vertices[names[1]].Vertices
		node0 := graph.Vertices[names[0]]
		map1Connections[names[0]] = node0
	}
	return
}

const inputPath = "../input.txt"

// distinct paths
// dont visit small caves more than once in between
// big caves -> Uppercase -> ANY TIME
// small caves -> lowerCase -> AT MOST ONCE
// all possible paths
func findPath(g *Graph, current *Vertex, end *Vertex,
	visited map[string]bool, thisPath []string, allPaths *[][]string) {

	// if current.isSmall {
	// 	visited[current.Key] = true
	// }
	// final destination
	thisPath = append(thisPath, current.Key)

	if current.Key == end.Key {
		*allPaths = append(*allPaths, thisPath)
		return
	}

outer:
	for _, v := range current.Vertices {
		alreadyVisitedSmallTwice := false
		// duplicate mi trqbva
		for _, valueInThisPath := range thisPath {
			isSmall := 'a' <= valueInThisPath[0] && valueInThisPath[0] <= 'z'
			if !isSmall {
				continue
			}
			// we have small guaranteed
			counter := 0
			for _, valueInThisPathToCheck := range thisPath {
				if valueInThisPath == valueInThisPathToCheck {
					counter++
				}
			}
			if counter > 1 {
				alreadyVisitedSmallTwice = true
				break
			}
		}

		if v.Key == "start" {
			continue
		}
		for _, vis := range thisPath {
			// alreadyVisitedSmallTwice
			// old...just once
			if thisPath[len(thisPath)-1] == v.Key {
				continue outer
			}

			if alreadyVisitedSmallTwice {
				if v.isSmall && v.Key == vis {
					continue outer
				}
			}

		}

		findPath(g, v, end, visited, thisPath, allPaths)
	}
}

func findResult(graph *Graph) (result int) {

	// traverse distinct paths
	// start at start - end at end
	start := graph.Vertices["start"]
	end := graph.Vertices["end"]
	visited := map[string]bool{}

	for _, v := range (*graph).Vertices {
		visited[v.Key] = false
	}
	visited["start"] = true
	pathSoFar := []string{}
	allPaths := [][]string{}
	findPath(graph, start, end, visited, pathSoFar, &allPaths)

	// traverse -> i da namerq "end"
	for _, v := range allPaths {
		fmt.Println(v)

	}
	return len(allPaths)
}

func returnSliceOfLinesFromFile(filePath string) (sliceOfLines []string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	rawBytes, err := io.ReadAll(file)
	if err != nil {
		log.Panic(err)
	}

	lines := strings.Split(string(rawBytes), "\n")

	return lines
}
