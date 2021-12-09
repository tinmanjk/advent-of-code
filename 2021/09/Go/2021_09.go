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

	listOfPoints := make([]*point, 0)
	for i := 1; i < lengthTOtalLines-1; i++ {
		for j := 1; j < lenghSingleLine-1; j++ {

			current := solutionData[i][j]
			if current == blockValue {
				continue
			}

			if partOne {
				left := solutionData[i][j-1]
				right := solutionData[i][j+1]
				up := solutionData[i-1][j]
				down := solutionData[i+1][j]
				if current < left && current < right && current < up && current < down {
					result += current + 1
				}
				continue
			}

			newPoint := point{current, i, j, nil}
			listOfPoints = append(listOfPoints, &newPoint)

			noLeft := (solutionData[i][j-1] == blockValue)
			noUp := (solutionData[i-1][j] == blockValue)
			up := !noUp
			left := !noLeft

			// from left to right
			// top to bottom
			// traversing
			switch {
			case noLeft && noUp: // Start new basin
				newBasin := make([]*point, 0)
				newBasin = append(newBasin, &newPoint)
				newPoint.basin = &newBasin
			case noLeft && up: // join up
				upPoint := findPoint(i-1, j, listOfPoints)
				*(upPoint.basin) = append(*(upPoint.basin), &newPoint)
				newPoint.basin = upPoint.basin
			case left && noUp: // join left
				leftPoint := findPoint(i, j-1, listOfPoints)
				*(leftPoint.basin) = append(*(leftPoint.basin), &newPoint)
				newPoint.basin = leftPoint.basin
			case left && up: // MERGE
				leftPoint := findPoint(i, j-1, listOfPoints)
				upPoint := findPoint(i-1, j, listOfPoints)

				switch {
				case len(*upPoint.basin) > len(*leftPoint.basin):
					if len(*leftPoint.basin) > 1 {
						for _, p := range *leftPoint.basin {
							// bez left
							if p != leftPoint {
								*(upPoint.basin) = append(*(upPoint.basin), p)
								p.basin = upPoint.basin
							}
						}
					}
					*(upPoint.basin) = append(*(upPoint.basin), leftPoint)
					*(upPoint.basin) = append(*(upPoint.basin), &newPoint)
					leftPoint.basin = upPoint.basin
					newPoint.basin = upPoint.basin
				case len(*leftPoint.basin) > len(*upPoint.basin):
					if len(*upPoint.basin) > 1 {
						for _, p := range *upPoint.basin {
							// bez left
							if p != upPoint {
								*(leftPoint.basin) = append(*(leftPoint.basin), p)
								p.basin = leftPoint.basin
							}
						}
					}
					*(leftPoint.basin) = append(*(leftPoint.basin), upPoint)
					*(leftPoint.basin) = append(*(leftPoint.basin), &newPoint)
					upPoint.basin = leftPoint.basin
					newPoint.basin = leftPoint.basin
				default: //TODO: fix logic to be complete
					equal := &(*leftPoint.basin) == &(*upPoint.basin)
					if !equal {
					}
					if len(*upPoint.basin) == 1 {
						*(leftPoint.basin) = append(*(leftPoint.basin), upPoint)
						*(leftPoint.basin) = append(*(leftPoint.basin), &newPoint)
						upPoint.basin = leftPoint.basin
						newPoint.basin = leftPoint.basin
						break
					}
					*(leftPoint.basin) = append(*(leftPoint.basin), &newPoint)
					newPoint.basin = leftPoint.basin
				}
			}

		}
	}

	if partOne {
		return result
	}

	basins := make(map[*[]*point]int, 0)
	for _, v := range listOfPoints {
		basins[v.basin] = len(*v.basin)
	}

	basinSizes := make([]int, 0)
	for _, v := range basins {
		basinSizes = append(basinSizes, v)
	}
	sort.Ints(basinSizes)

	result = 1
	for i := 1; i <= 3; i++ {
		basinSize := basinSizes[len(basinSizes)-i]
		result *= basinSize
	}

	return
}

func findPoint(i int, j int, listOfPoints []*point) *point {
	for _, v := range listOfPoints {
		if v.i == i && v.j == j {
			return v
		}
	}
	return nil
}

type point struct {
	val   int
	i     int
	j     int
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
