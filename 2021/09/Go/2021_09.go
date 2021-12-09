package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

const inputPath = "../input.txt"
const blockValue = 9

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	inputData := parseInput(lines)
	var result int

	// part 1
	lowPoints := findLowPoints(inputData)
	result = findSumRiskLevels(lowPoints)
	fmt.Println(result)

	// part 2
	result = findThreeLargestBasins(inputData)
	fmt.Println(result)
}

func parseInput(slicesOfLines []string) (solutionsData [][]int) {

	lengthTotalLines := len(slicesOfLines)
	lenghSingleLine := len(slicesOfLines[0]) // should be the same for all
	solutionsData = make([][]int, lengthTotalLines)
	for i := 0; i < lengthTotalLines; i++ {
		solutionsData[i] = make([]int, lenghSingleLine)
		for j := 0; j < lenghSingleLine; j++ {
			solutionsData[i][j] = int(slicesOfLines[i][j] - '0')
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

func findLowPoints(inputData [][]int) (mapLowPoints map[pointCoord]*point) {

	paddedMatrix := addPaddings(inputData, blockValue)
	lengthTotalLines := len(paddedMatrix)
	lenghSingleLine := len(paddedMatrix[0]) // should be the same for all
	mapLowPoints = make(map[pointCoord]*point, 0)

	for i := 1; i < lengthTotalLines-1; i++ {
		for j := 1; j < lenghSingleLine-1; j++ {

			current := paddedMatrix[i][j]
			if current == blockValue {
				continue
			}

			currentPoint := point{pointCoord{i, j}, current, nil}

			left := paddedMatrix[i][j-1]
			right := paddedMatrix[i][j+1]
			up := paddedMatrix[i-1][j]
			down := paddedMatrix[i+1][j]
			if current < left && current < right && current < up && current < down {
				mapLowPoints[pointCoord{i, j}] = &currentPoint
			}
		}
	}
	return
}

func findSumRiskLevels(mapLowPoints map[pointCoord]*point) (result int) {
	for _, v := range mapLowPoints {
		result += *&v.val + 1
	}
	return result
}

func findThreeLargestBasins(inputData [][]int) (result int) {

	paddedMatrix := addPaddings(inputData, blockValue)
	lengthTotalLines := len(paddedMatrix)
	lenghSingleLine := len(paddedMatrix[0])
	mapOfPoints := make(map[pointCoord]*point, 0)

	for i := 1; i < lengthTotalLines-1; i++ {
		for j := 1; j < lenghSingleLine-1; j++ {

			current := paddedMatrix[i][j]
			if current == blockValue {
				continue
			}

			currentPoint := point{pointCoord{i, j}, current, nil}
			mapOfPoints[currentPoint.coord] = &currentPoint

			noLeft := (paddedMatrix[i][j-1] == blockValue)
			noUp := (paddedMatrix[i-1][j] == blockValue)
			up := !noUp
			left := !noLeft

			var joinBasin *[]*point
			switch {
			case noLeft && noUp: // Start new basin
				newBasin := make([]*point, 0)
				joinBasin = &newBasin
			case noLeft && up: // join up
				upPoint := mapOfPoints[pointCoord{i - 1, j}]
				joinBasin = upPoint.basin
			case left && noUp: // join left
				leftPoint := mapOfPoints[pointCoord{i, j - 1}]
				joinBasin = leftPoint.basin
			case left && up: // MERGE
				leftPoint := mapOfPoints[pointCoord{i, j - 1}]
				upPoint := mapOfPoints[pointCoord{i - 1, j}]

				if len(*upPoint.basin) > len(*leftPoint.basin) {
					mergeBasins(leftPoint, upPoint)
				} else if len(*leftPoint.basin) > len(*upPoint.basin) {
					mergeBasins(upPoint, leftPoint)
				} else { // Equal number
					differentBasins := !(&(*leftPoint.basin) == &(*upPoint.basin))
					if differentBasins {
						mergeBasins(upPoint, leftPoint)
					}
				}
				// already merged doesn't matter which one (left or up point to the same basin)
				joinBasin = leftPoint.basin
			}

			*(joinBasin) = append(*(joinBasin), &currentPoint)
			currentPoint.basin = joinBasin
		}
	}

	basinsHashSet := make(map[*[]*point]bool, 0)
	basinSizes := make([]int, 0)
	for _, v := range mapOfPoints {
		if _, ok := basinsHashSet[v.basin]; !ok {
			basinsHashSet[v.basin] = true
			basinSizes = append(basinSizes, len(*v.basin))
		}
	}
	sort.Ints(basinSizes)
	result = 1
	for i := 1; i <= 3; i++ {
		basinSize := basinSizes[len(basinSizes)-i]
		result *= basinSize
	}

	return
}

func mergeBasins(srcBasin *point, destBasin *point) {
	for _, p := range *srcBasin.basin {
		*(destBasin.basin) = append(*(destBasin.basin), p)
		p.basin = destBasin.basin
	}
}

type pointCoord struct {
	i int
	j int
}
type point struct {
	coord pointCoord
	val   int
	basin *[]*point
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
