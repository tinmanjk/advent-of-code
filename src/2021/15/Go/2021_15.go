package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

func main() {

	lines := returnSliceOfLinesFromFile(inputPath)
	var result int
	var finalArr []point
	matrixOfInts := parseInput(lines)

	// part 1
	finalArr, result = findResult(matrixOfInts, false)
	fmt.Println(result)
	fmt.Println(finalArr)

	// part 2
	finalArr, result = findResult(matrixOfInts, true)
	fmt.Println(result)
	fmt.Println(finalArr)

}

func parseInput(slicesOfLines []string) (matrixOfInts [][]int) {

	matrixOfInts = make([][]int, len(slicesOfLines))
	for i := 0; i < len(matrixOfInts); i++ {
		matrixOfInts[i] = make([]int, len(slicesOfLines[i]))
		line := slicesOfLines[i]
		for j := 0; j < len(matrixOfInts[i]); j++ {
			matrixOfInts[i][j] = int(line[j]) - 48
		}
	}

	return
}

// TODO: Refactor possibly away
func fiveXmatrix(matrixOfInts [][]int) (fiveXMatrix [][]int) {
	fiveXMatrix = make([][]int, 5*len(matrixOfInts))
	for i := 0; i < len(fiveXMatrix); i++ {
		fiveXMatrix[i] = make([]int, 5*len(matrixOfInts))
	}

	// to the right
	size := len(matrixOfInts)
	for k := 0; k < 5; k++ {
		for i := 0; i < len(matrixOfInts); i++ {
			for j := 0; j < len(matrixOfInts); j++ {
				fiveXMatrix[i][j+k*size] = matrixOfInts[i][j] + k
				if fiveXMatrix[i][j+k*size] > 9 {
					fiveXMatrix[i][j+k*size] = fiveXMatrix[i][j+k*size] % 9
				}
			}
		}
	}
	// to the bottom
	for k := 1; k < 5; k++ {
		for i := 0; i < len(matrixOfInts); i++ {
			for j := 0; j < len(matrixOfInts); j++ {
				fiveXMatrix[i+k*size][j] = matrixOfInts[i][j] + k
				if fiveXMatrix[i+k*size][j] > 9 {
					fiveXMatrix[i+k*size][j] = fiveXMatrix[i+k*size][j] % 9
				}
			}
		}
	}

	// left
	for k := 1; k < 5; k++ {
		for i := size; i < size*5; i++ {
			for j := 0; j < len(matrixOfInts); j++ {
				fiveXMatrix[i][j+k*size] = fiveXMatrix[i][j+(k-1)*size] + 1
				if fiveXMatrix[i][j+k*size] > 9 {
					fiveXMatrix[i][j+k*size] = fiveXMatrix[i][j+k*size] % 9
				}
			}
		}
	}

	return
}

const inputPath = "../input.txt"

func findResult(matrixOfInts [][]int, secondPart bool) (finalArr []point, result int) {

	if secondPart {
		matrixOfInts = fiveXmatrix(matrixOfInts)
	}

	lastIndexInMatrix := len(matrixOfInts) - 1

	// Create EdgeList Representation ofthe Graph
	var edgeList = []FullEdge{}
	for i := 0; i < len(matrixOfInts); i++ {
		for j := 0; j < len(matrixOfInts); j++ {
			source := point{i, j}
			// to the right
			if j != lastIndexInMatrix {
				fullEdge := FullEdge{}

				destination := point{}
				destination.i = i
				destination.j = j + 1

				fullEdge.Source = source
				fullEdge.Destination = destination
				fullEdge.Weight = matrixOfInts[i][j+1]
				edgeList = append(edgeList, fullEdge)

				// and to the left back
				fullEdge.Destination = source
				fullEdge.Source = destination
				fullEdge.Weight = matrixOfInts[i][j]
				edgeList = append(edgeList, fullEdge)
			}

			// to the bottom
			if i != lastIndexInMatrix {
				fullEdge := FullEdge{}
				destination := point{}

				destination.i = i + 1
				destination.j = j

				fullEdge.Source = source
				fullEdge.Destination = destination
				fullEdge.Weight = matrixOfInts[i+1][j]
				edgeList = append(edgeList, fullEdge)

				// and to the up back
				fullEdge.Destination = source
				fullEdge.Source = destination
				fullEdge.Weight = matrixOfInts[i][j]
				edgeList = append(edgeList, fullEdge)
			}
		}
	}

	fromPoint := point{0, 0}
	toPoint := point{lastIndexInMatrix, lastIndexInMatrix}

	startNode, endNode, itemGraph := CreateGraph(edgeList, fromPoint, toPoint)

	finalArr, result = getShortestPath(startNode, endNode, itemGraph)

	acc := 0
	for i := 1; i < len(finalArr); i++ {
		acc += matrixOfInts[finalArr[i].i][finalArr[i].j]
	}

	// fmt.Println(acc)
	return
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Vertex) {
	g.Vertexes = append(g.Vertexes, n)
}

// AddEdge adds a partial edge to the graph
// works as expected
func (g *ItemGraph) AddEdge(n1, n2 *Vertex, weight int) {
	if g.Edges == nil {
		g.Edges = make(map[*Vertex][]*PartialEdge)
	}
	ed1 := PartialEdge{
		ToVertex: n2,
		Weight:   weight,
	}

	g.Edges[n1] = append(g.Edges[n1], &ed1)
}

type FullEdge struct {
	Source      point
	Destination point
	Weight      int
}

// node.a ima Edges
type PartialEdge struct {
	ToVertex *Vertex
	Weight   int
}

type point struct {
	i int
	j int
}

type Vertex struct {
	Value point
}

type ItemGraph struct {
	Vertexes []*Vertex                  // can have partialEdges
	Edges    map[*Vertex][]*PartialEdge // needs full edges
}

// adjancy list
func CreateGraph(edgeListGraph []FullEdge, fromPoint point, toPoint point) (startVertex *Vertex, endVertex *Vertex, itemGraph *ItemGraph) {
	var g ItemGraph
	vertics := make(map[point]*Vertex)
	for _, v := range edgeListGraph {
		if _, found := vertics[v.Source]; !found {
			nA := Vertex{v.Source}
			vertics[v.Source] = &nA
			if nA.Value == fromPoint {
				startVertex = &nA
			}

			g.AddNode(&nA)
		}
		if _, found := vertics[v.Destination]; !found {
			nA := Vertex{v.Destination}
			vertics[v.Destination] = &nA

			if nA.Value == toPoint {
				endVertex = &nA
			}
			g.AddNode(&nA)
		}
		g.AddEdge(vertics[v.Source], vertics[v.Destination], v.Weight)
	}
	return startVertex, endVertex, &g
}

func getShortestPath(startVertex *Vertex, endVertex *Vertex, g *ItemGraph) ([]point, int) {
	visited := make(map[point]bool) // already visited
	dist := make(map[point]int)     // from Start TO
	prev := make(map[point]point)   //???

	q := PriorityQueue{}
	pq := q.NewQ() // priority Queue
	start := VertexDistance{
		Vertex:   startVertex,
		Distance: 0,
	}
	for _, nval := range g.Vertexes {
		dist[nval.Value] = math.MaxInt64
	}
	dist[startVertex.Value] = start.Distance
	pq.Enqueue(start)
	for !pq.IsEmpty() {
		current := pq.Dequeue()
		if visited[current.Vertex.Value] {
			continue
		}
		edges := g.Edges[current.Vertex]
		visited[current.Vertex.Value] = true

		for _, edge := range edges {
			if !visited[edge.ToVertex.Value] {
				if dist[current.Vertex.Value]+edge.Weight < dist[edge.ToVertex.Value] {
					store := VertexDistance{
						Vertex:   edge.ToVertex,
						Distance: dist[current.Vertex.Value] + edge.Weight,
					}
					dist[edge.ToVertex.Value] = dist[current.Vertex.Value] + edge.Weight
					prev[edge.ToVertex.Value] = current.Vertex.Value
					pq.Enqueue(store)
				}
			}
		}

	}
	pathval := prev[endVertex.Value]
	var finalArr []point
	finalArr = append(finalArr, endVertex.Value)
	for pathval != startVertex.Value {
		finalArr = append(finalArr, pathval)
		pathval = prev[pathval]
	}
	finalArr = append(finalArr, pathval)
	finalArrReversed := make([]point, len(finalArr))
	for i := 0; i < len(finalArr); i++ {
		finalArrReversed[i] = finalArr[len(finalArr)-1-i]
	}
	return finalArrReversed, dist[endVertex.Value]

}

// tova za priority queue.to
type VertexDistance struct {
	Vertex   *Vertex
	Distance int
}

type PriorityQueue struct {
	Items []VertexDistance
}

// Enqueue adds an Node to the end of the queue
func (s *PriorityQueue) Enqueue(t VertexDistance) {
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		return
	}

	if t.Distance < s.Items[0].Distance {
		s.Items = append([]VertexDistance{t}, s.Items...)
		return
	}

	var insertFlag bool
	for i := 1; i < len(s.Items); i++ {
		if t.Distance < s.Items[i].Distance {
			s.Items = append(s.Items[:i+1], s.Items[i:]...)
			s.Items[i] = t
			insertFlag = true
			break
		}
	}

	if !insertFlag {
		s.Items = append(s.Items, t)
	}
}

// Dequeue removes an Node from the start of the queue
func (s *PriorityQueue) Dequeue() *VertexDistance {
	item := s.Items[0]
	s.Items = s.Items[1:]
	return &item
}

//NewQ Creates New Queue
func (s *PriorityQueue) NewQ() *PriorityQueue {
	s.Items = []VertexDistance{}
	return s
}

// IsEmpty returns true if the queue is empty
func (s *PriorityQueue) IsEmpty() bool {
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *PriorityQueue) Size() int {
	return len(s.Items)
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
