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
	result = findResult(mapOfNodes)
	fmt.Println(result)

	// part 2
	// result = findResult(parsedInput)
	// fmt.Println(result)
}

type Vertex struct {
	node1 Node
	node2 Node
}

type Node struct {
	Name               string
	ConnectionsHashSet *map[*Node]*Node // undirected graph
	isSmall            bool
}

func parseInput(slicesOfLines []string) (mapOfNodes map[string]*Node) {
	// adjacency list
	// map of nodes
	mapOfNodes = map[string]*Node{}

	for i := 0; i < len(slicesOfLines); i++ {
		names := strings.Split(slicesOfLines[i], "-")
		// create empty node if not already
		for j := 0; j < 2; j++ {
			if _, ok := mapOfNodes[names[j]]; !ok {
				isSmall := 'a' <= names[j][0] && names[j][0] <= 'z'
				connections := &map[*Node]*Node{}
				node := Node{names[j], connections, isSmall}
				mapOfNodes[names[j]] = &node
			}
		}

		// connect the two
		map0Connections := (*mapOfNodes[names[0]]).ConnectionsHashSet
		node1 := mapOfNodes[names[1]]
		(*map0Connections)[node1] = node1

		map1Connections := (*mapOfNodes[names[1]]).ConnectionsHashSet
		node0 := mapOfNodes[names[0]]
		(*map1Connections)[node0] = node0
	}
	return
}

const inputPath = "../input0.txt"

// distinct paths
// dont visit small caves more than once in between
// big caves -> Uppercase -> ANY TIME
// small caves -> lowerCase -> AT MOST ONCE

func findResult(mapOfNodes map[string]*Node) (result int) {

	return
}

// func (n *Node) DepthFirstSearch(array []int) []int {
// 	array = append(array, n.Value)
// 	for _, child := range n.Children {
// 		array = child.DepthFirstSearch(array)
// 	}

// 	return array
// }

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
