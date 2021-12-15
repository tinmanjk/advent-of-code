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

func addPaddings(solutionsData [][]int, paddedValue int) (paddedSolutionsData [][]int) {

	paddedSolutionsData = make([][]int, len(solutionsData))
	for i := 0; i < len(solutionsData); i++ {
		paddedSolutionsData[i] = append([]int{paddedValue}, solutionsData[i]...)
		paddedSolutionsData[i] = append(paddedSolutionsData[i], paddedValue)
	}

	lenghtPaddedSingleLine := len(solutionsData[0]) + 2 // should be the same for all
	paddedBeginRow := make([]int, lenghtPaddedSingleLine)
	for i := 0; i < len(paddedBeginRow); i++ {
		paddedBeginRow[i] = paddedValue
	}
	paddedEndRow := make([]int, lenghtPaddedSingleLine)
	copy(paddedEndRow, paddedBeginRow)

	paddedMatrix := make([][]int, 0)
	paddedMatrix = append(paddedMatrix, paddedBeginRow)

	paddedSolutionsData = append(paddedMatrix, paddedSolutionsData...)
	paddedSolutionsData = append(paddedSolutionsData, paddedEndRow)

	return
}

func fiveXmatrix(matrixOfInts [][]int) (fiveXMatrix [][]int) {
	fiveXMatrix = make([][]int, 5*len(matrixOfInts))
	for i := 0; i < len(fiveXMatrix); i++ {
		fiveXMatrix[i] = make([]int, 5*len(matrixOfInts))
	}

	// nadqsno
	size := len(matrixOfInts)
	for k := 0; k < 5; k++ {
		for i := 0; i < len(matrixOfInts); i++ {
			for j := 0; j < len(matrixOfInts); j++ {
				// if k == 0 {
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
				// if k == 0 {
				fiveXMatrix[i+k*size][j] = matrixOfInts[i][j] + k
				if fiveXMatrix[i+k*size][j] > 9 {
					fiveXMatrix[i+k*size][j] = fiveXMatrix[i+k*size][j] % 9
				}
			}
		}
	}

	// pak nadqsno
	for k := 1; k < 5; k++ {
		for i := size; i < size*5; i++ {
			for j := 0; j < len(matrixOfInts); j++ {
				// if k == 0 {
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

	paddedMatrix := addPaddings(matrixOfInts, 99)

	// InputData
	var inputDataSlice = []AdjacencyList{}
	for i := 1; i < len(paddedMatrix)-1; i++ {
		for j := 1; j < len(paddedMatrix[i])-1; j++ {
			source := point{i, j}
			// destination nadqsno i nadolu samo
			// nadqsno
			if paddedMatrix[i][j+1] != 99 {
				inputData := AdjacencyList{}
				destination := point{}

				destination.i = i
				destination.j = j + 1

				inputData.Source = source
				inputData.Destination = destination
				inputData.Weight = paddedMatrix[i][j+1]
				inputDataSlice = append(inputDataSlice, inputData)
			}

			if paddedMatrix[i+1][j] != 99 {
				inputData := AdjacencyList{}
				destination := point{}

				destination.i = i + 1
				destination.j = j

				inputData.Source = source
				inputData.Destination = destination
				inputData.Weight = paddedMatrix[i+1][j]
				inputDataSlice = append(inputDataSlice, inputData)
			}
		}
	}

	fromPoint := point{1, 1}
	toPoint := point{len(paddedMatrix) - 2, len(paddedMatrix) - 2}
	inputGraph := InputGraph{
		inputDataSlice, fromPoint, toPoint,
	}

	startNode, endNode, itemGraph := CreateGraph(inputGraph)

	finalArr, result = getShortestPath(startNode, endNode, itemGraph)
	return
}

// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Vertex) {
	g.Vertexes = append(g.Vertexes, n)
}

// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Vertex, weight int) {
	if g.Edges == nil {
		g.Edges = make(map[Vertex][]*Edge)
	}
	ed1 := Edge{
		ToVertex: n2,
		Weight:   weight,
	}

	ed2 := Edge{
		ToVertex: n1,
		Weight:   weight,
	}
	g.Edges[*n1] = append(g.Edges[*n1], &ed1)
	g.Edges[*n2] = append(g.Edges[*n2], &ed2)
}

type AdjacencyList struct {
	Source      point
	Destination point
	Weight      int
}

type InputGraph struct {
	Graph []AdjacencyList
	From  point
	To    point
}

func CreateGraph(data InputGraph) (startVertex *Vertex, endVertex *Vertex, itemGraph *ItemGraph) {
	var g ItemGraph
	vertics := make(map[point]*Vertex)
	for _, v := range data.Graph {
		if _, found := vertics[v.Source]; !found {
			nA := Vertex{v.Source, nil}
			vertics[v.Source] = &nA
			if nA.Value == data.From {
				startVertex = &nA
			}

			g.AddNode(&nA)
		}
		if _, found := vertics[v.Destination]; !found {
			nA := Vertex{v.Destination, nil}
			vertics[v.Destination] = &nA

			if nA.Value == data.To {
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

	q := VertexDistanceQueue{}
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
		v := pq.Dequeue()
		if visited[v.Vertex.Value] {
			continue
		}
		visited[v.Vertex.Value] = true
		near := g.Edges[*v.Vertex]

		for _, val := range near {
			if val.ToVertex.Value.i < v.Vertex.Value.i {
				continue
			}

			if val.ToVertex.Value.j < v.Vertex.Value.j {
				continue
			}

			if !visited[val.ToVertex.Value] {
				if dist[v.Vertex.Value]+val.Weight < dist[val.ToVertex.Value] {
					store := VertexDistance{
						Vertex:   val.ToVertex,
						Distance: dist[v.Vertex.Value] + val.Weight,
					}
					dist[val.ToVertex.Value] = dist[v.Vertex.Value] + val.Weight
					prev[val.ToVertex.Value] = v.Vertex.Value
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

type point struct {
	i int
	j int
}

type Vertex struct {
	Value point
	Edges *[]*Vertex //->TODO
}

// node.a ima Edges
type Edge struct {
	ToVertex *Vertex
	Weight   int
}

// tova za priority queue.to
type VertexDistance struct {
	Vertex   *Vertex
	Distance int
}

type ItemGraph struct {
	Vertexes []*Vertex
	Edges    map[Vertex][]*Edge
}

type VertexDistanceQueue struct {
	Items []VertexDistance
}

// Enqueue adds an Node to the end of the queue
func (s *VertexDistanceQueue) Enqueue(t VertexDistance) {
	if len(s.Items) == 0 {
		s.Items = append(s.Items, t)
		return
	}
	var insertFlag bool
	for k, v := range s.Items {
		if t.Distance < v.Distance {
			if k > 0 {
				s.Items = append(s.Items[:k+1], s.Items[k:]...)
				s.Items[k] = t
				insertFlag = true
			} else {
				s.Items = append([]VertexDistance{t}, s.Items...)
				insertFlag = true
			}
		}
		if insertFlag {
			break
		}
	}
	if !insertFlag {
		s.Items = append(s.Items, t)
	}
}

// Dequeue removes an Node from the start of the queue
func (s *VertexDistanceQueue) Dequeue() *VertexDistance {
	item := s.Items[0]
	s.Items = s.Items[1:len(s.Items)]
	return &item
}

//NewQ Creates New Queue
func (s *VertexDistanceQueue) NewQ() *VertexDistanceQueue {
	s.Items = []VertexDistance{}
	return s
}

// IsEmpty returns true if the queue is empty
func (s *VertexDistanceQueue) IsEmpty() bool {
	return len(s.Items) == 0
}

// Size returns the number of Nodes in the queue
func (s *VertexDistanceQueue) Size() int {
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
