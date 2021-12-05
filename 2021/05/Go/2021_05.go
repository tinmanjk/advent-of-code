package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)

	lineSegmentSlice := parseInput(lines)
	var result int
	result = task01(lineSegmentSlice)
	fmt.Println(result)

	// result = task02(numbers)
	// fmt.Println(result)
}

func parseInput(slicesOfLines []string) (sliceOfLineSegments []lineSegment) {
	sliceOfLineSegments = make([]lineSegment, len(slicesOfLines))
	for i := 0; i < len(slicesOfLines); i++ {
		lineSeg := lineSegment{}
		line := strings.Split(slicesOfLines[i], " -> ")
		x1y1 := strings.Split(line[0], ",")
		x2y2 := strings.Split(line[1], ",")

		lineSeg.x1, _ = strconv.Atoi(x1y1[0])
		lineSeg.y1, _ = strconv.Atoi(x1y1[1])
		lineSeg.x2, _ = strconv.Atoi(x2y2[0])
		lineSeg.y2, _ = strconv.Atoi(x2y2[1])

		sliceOfLineSegments[i] = lineSeg
	}
	return
}

// x1 = x2 or y1 = y2
// horizontal or vertical lines
func task01(sliceOfLineSegments []lineSegment) (result int) {

	mapPoints := make(map[point]int, 0)
	for i := 0; i < len(sliceOfLineSegments); i++ {
		// vertical or horizontal cases
		createMapPoitns(sliceOfLineSegments[i], mapPoints)
	}

	for _, v := range mapPoints {
		if v > 1 {
			result++
		}
	}
	return result
}

func createMapPoitns(lineSeg lineSegment, mapPoints map[point]int) {
	if lineSeg.x1 == lineSeg.x2 ||
		lineSeg.y1 == lineSeg.y2 {
		if lineSeg.x1 == lineSeg.x2 {
			var min, max int
			if lineSeg.y1 <= lineSeg.y2 {
				min = lineSeg.y1
				max = lineSeg.y2

			} else {
				min = lineSeg.y2
				max = lineSeg.y1
			}

			for i := min; i < max+1; i++ {
				p := point{lineSeg.x1, i}

				if _, ok := mapPoints[p]; ok {
					mapPoints[p]++
					continue
				}
				mapPoints[p] = 1
			}

		}

		if lineSeg.y1 == lineSeg.y2 {
			var min, max int
			if lineSeg.x1 <= lineSeg.x2 {
				min = lineSeg.x1
				max = lineSeg.x2

			} else {
				min = lineSeg.x2
				max = lineSeg.x1
			}

			for i := min; i < max+1; i++ {
				p := point{i, lineSeg.y1}

				if _, ok := mapPoints[p]; ok {
					mapPoints[p]++
					continue
				}
				mapPoints[p] = 1
			}

		}
	} else {
		// up/down and direction
		// dva tipa
		// tip 1
		var leftPoint, rightPoint point
		if lineSeg.x1 < lineSeg.x2 {
			leftPoint = point{lineSeg.x1, lineSeg.y1}
			rightPoint = point{lineSeg.x2, lineSeg.y2}
		} else {
			leftPoint = point{lineSeg.x2, lineSeg.y2}
			rightPoint = point{lineSeg.x1, lineSeg.y1}
		}

		// downslope
		if leftPoint.y > rightPoint.y {
			// generate downslope
			// x++ / y--
			y := leftPoint.y
			for i := leftPoint.x; i < rightPoint.x+1; i++ {
				p := point{i, y}
				y--
				if _, ok := mapPoints[p]; ok {
					mapPoints[p]++
					continue
				}
				mapPoints[p] = 1
			}
		} else {
			// generate downslope
			// x++ / y--
			y := leftPoint.y
			for i := leftPoint.x; i < rightPoint.x+1; i++ {
				p := point{i, y}
				y++
				if _, ok := mapPoints[p]; ok {
					mapPoints[p]++
					continue
				}
				mapPoints[p] = 1
			}
		}

	}
	return
}

// dictionary ot point obv
type point struct {
	x int
	y int
}
type lineSegment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

const finalSum int = 2020

// three numbers should sum to 2020 -> multiply
func task02(numbers []int) (result int) {

	return
}

func returnSliceOfIntsFromFile(filePath string) (sliceOfLines []int) {
	slicesOfLines := returnSliceOfLinesFromFile(filePath)

	lines := make([]int, 0, len(slicesOfLines))
	for i := 0; i < len(slicesOfLines); i++ {
		trimmed := strings.TrimRight(slicesOfLines[i], "\n ")
		if trimmed == "" {
			continue
		}
		// TODO better Error handling Atoi
		number, err := strconv.Atoi(trimmed)
		if err != nil {
			log.Panic(err)
		}
		lines = append(lines, number)
	}

	return lines
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

func splitLine(line string) (firstNumber int, secondNumber int,
	char rune, password string) {

	// Example line: 5-6 v: hvvgvrm
	lineSplit := strings.Split(line, " ") // should be 3
	numbers := strings.Split(lineSplit[0], "-")
	// TODO: Better Error handling
	firstNumber, err := strconv.Atoi(numbers[0])
	if err != nil {
		log.Panic(err)
	}
	secondNumber, err = strconv.Atoi(numbers[1])
	if err != nil {
		log.Panic(err)
	}

	// use if there are multi-byte unicode chars
	for _, r := range lineSplit[1] {
		char = r
		break
	}

	password = lineSplit[2]
	return
}
