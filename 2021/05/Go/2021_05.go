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
	result = findResult(lineSegmentSlice, false)
	fmt.Println(result)

	result = findResult(lineSegmentSlice, true)
	fmt.Println(result)
}

type point struct {
	x int
	y int
}
type lineSegment struct {
	p1 point
	p2 point
}

func parseInput(slicesOfLines []string) (sliceOfLineSegments []lineSegment) {

	sliceOfLineSegments = make([]lineSegment, len(slicesOfLines))
	for i := 0; i < len(slicesOfLines); i++ {
		lineSeg := lineSegment{point{}, point{}}
		line := strings.Split(slicesOfLines[i], " -> ")
		x1y1 := strings.Split(line[0], ",")
		x2y2 := strings.Split(line[1], ",")
		lineSeg.p1.x, _ = strconv.Atoi(x1y1[0])
		lineSeg.p1.y, _ = strconv.Atoi(x1y1[1])
		lineSeg.p2.x, _ = strconv.Atoi(x2y2[0])
		lineSeg.p2.y, _ = strconv.Atoi(x2y2[1])

		sliceOfLineSegments[i] = lineSeg
	}
	return
}

func findResult(sliceOfLineSegments []lineSegment, includeDiagonal bool) (result int) {

	mapPoints := make(map[point]int, 0)
	for i := 0; i < len(sliceOfLineSegments); i++ {
		createMapPoints(sliceOfLineSegments[i], mapPoints, includeDiagonal)
	}

	for _, v := range mapPoints {
		if v > 1 {
			result++
		}
	}
	return result
}

func createMapPoints(lineSeg lineSegment, mapPoints map[point]int, includeDiagonal bool) {
	switch {
	case lineSeg.p1.x == lineSeg.p2.x: // vertical
		x := lineSeg.p1.x
		min, max := findMinMax(lineSeg.p1.y, lineSeg.p2.y)
		for y := min; y <= max; y++ {
			upsertMapPointsCounts(x, y, mapPoints)
		}
	case lineSeg.p1.y == lineSeg.p2.y: // horizontal
		y := lineSeg.p1.y
		min, max := findMinMax(lineSeg.p1.x, lineSeg.p2.x)
		for x := min; x <= max; x++ {
			upsertMapPointsCounts(x, y, mapPoints)
		}
	case includeDiagonal:
		var leftPoint, rightPoint point
		if lineSeg.p1.x < lineSeg.p2.x {
			leftPoint = point{lineSeg.p1.x, lineSeg.p1.y}
			rightPoint = point{lineSeg.p2.x, lineSeg.p2.y}
		} else {
			leftPoint = point{lineSeg.p2.x, lineSeg.p2.y}
			rightPoint = point{lineSeg.p1.x, lineSeg.p1.y}
		}

		for x, y := leftPoint.x, leftPoint.y; x <= rightPoint.x; x++ {
			upsertMapPointsCounts(x, y, mapPoints)
			if leftPoint.y > rightPoint.y {
				y--
			} else {
				y++
			}
		}
	}

	return
}

func findMinMax(number1 int, number2 int) (min int, max int) {
	if number1 <= number2 {
		min = number1
		max = number2

	} else {
		min = number2
		max = number1
	}
	return
}

func upsertMapPointsCounts(x int, y int, mapPoints map[point]int) {
	p := point{x, y}
	if _, ok := mapPoints[p]; ok {
		mapPoints[p]++
		return
	}
	mapPoints[p] = 1
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
