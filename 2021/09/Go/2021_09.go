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
	solutionData := parseInput(lines)

	var result int

	result = findResult(solutionData, true)
	fmt.Println(result)

	result = findResult(solutionData, false)
	fmt.Println(result)
}

func parseInput(slicesOfLines []string) (matrixLocations [][]int) {

	length := len(slicesOfLines) + 2
	matrixLocations = make([][]int, length)
	otherLength := len(slicesOfLines[0]) + 2
	for i := 0; i < length; i++ {
		matrixLocations[i] = make([]int, otherLength)
		if i == 0 || i == length-1 {
			for j := 0; j < otherLength; j++ {
				matrixLocations[i][j] = blockValue
			}
			continue
		}
		line := slicesOfLines[i-1]
		matrixLocations[i][0] = blockValue
		matrixLocations[i][otherLength-1] = blockValue

		for j := 1; j < otherLength-1; j++ {
			matrixLocations[i][j] = int(line[j-1]) - 48
		}
	}

	return
}

func findResult(solutionData [][]int, partOne bool) (result int) {

	lengthTOtalLines := len(solutionData)
	lenghSingleLine := len(solutionData[0])

	mapOfPoints := make(map[pointCoord]*point, 0)
	mapLowPoints := make(map[pointCoord]*point, 0)
	for i := 1; i < lengthTOtalLines-1; i++ {
		for j := 1; j < lenghSingleLine-1; j++ {

			current := solutionData[i][j]
			if current == blockValue {
				continue
			}

			currentPoint := point{pointCoord{i, j}, current, nil}

			if partOne {
				left := solutionData[i][j-1]
				right := solutionData[i][j+1]
				up := solutionData[i-1][j]
				down := solutionData[i+1][j]
				if current < left && current < right && current < up && current < down {
					mapLowPoints[pointCoord{i, j}] = &currentPoint
				}
				continue
			}

			mapOfPoints[currentPoint.coord] = &currentPoint

			noLeft := (solutionData[i][j-1] == blockValue)
			noUp := (solutionData[i-1][j] == blockValue)
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

	if partOne {
		for _, v := range mapLowPoints {
			result += *&v.val + 1
		}
		return result
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
