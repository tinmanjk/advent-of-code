package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := returnSliceOfLinesFromFile(inputPath)
	var result int
	mapOfPoints, instructions := parseInput(lines)

	// part 1
	result = findResult(mapOfPoints, instructions, true)
	fmt.Println(result)

	// part 2
	result = findResult(mapOfPoints, instructions, false)
}

type point struct {
	x int
	y int
}

type inst struct {
	axis  string
	units int
}

func parseInput(slicesOfLines []string) (mapOfPoints map[point]point,
	instructions []inst) {

	// points
	mapOfPoints = map[point]point{}
	i := 0
	for ; i < len(slicesOfLines) && slicesOfLines[i] != ""; i++ {
		line := slicesOfLines[i]
		splitted := strings.Split(line, ",")
		x, _ := strconv.Atoi(splitted[0])
		y, _ := strconv.Atoi(splitted[1])
		p := point{x, y}
		mapOfPoints[p] = p
	}

	//instructions
	for j := i + 1; j < len(slicesOfLines) && slicesOfLines[j] != ""; j++ {
		line := slicesOfLines[j]
		splitted := strings.Split(line, " ")
		instruction := strings.Split(splitted[2], "=")
		axis := instruction[0]
		units, _ := strconv.Atoi(instruction[1])
		instr := inst{axis, units}
		instructions = append(instructions, instr)
	}

	return
}

const inputPath = "../input.txt"

func findResult(mapOfPoints map[point]point,
	instructions []inst, part1 bool) (result int) {

	for i := 0; i < len(instructions); i++ {
		mapOfPoints = fold(mapOfPoints, instructions[i].axis, instructions[i].units)

		if part1 {
			return len(mapOfPoints)
		}
	}
	printPoints(mapOfPoints)
	return len(mapOfPoints)
}

func printPoints(mapOfPoints map[point]point) {

	maxX := math.MinInt32
	maxY := math.MinInt32
	for _, p := range mapOfPoints {
		if p.x >= maxX {
			maxX = p.x
		}
		if p.y >= maxY {
			maxY = p.y
		}
	}

	printMatrix := make([][]rune, maxY+1)
	for i := 0; i < len(printMatrix); i++ {
		printMatrix[i] = make([]rune, maxX+1)
		for j := 0; j < len(printMatrix[i]); j++ {
			pointAtCoordinate := point{j, i}
			if _, ok := mapOfPoints[pointAtCoordinate]; ok {
				printMatrix[i][j] = '#'
			} else {
				printMatrix[i][j] = ' '
			}
		}
		fmt.Println(string(printMatrix[i]))
	}
}

func fold(mapOfPoints map[point]point, axis string, units int) (foldedMap map[point]point) {
	foldedMap = map[point]point{}
	for _, oldPoint := range mapOfPoints {
		switch axis {
		case "y": // down to up
			if oldPoint.y == units {
				break // not to be added to new one
			}
			if oldPoint.y < units {
				foldedMap[oldPoint] = oldPoint
				break
			}
			oldPoint.y = 2*units - oldPoint.y
			foldedMap[oldPoint] = oldPoint
		case "x": // right to left
			if oldPoint.x == units {
				break // not to be added to new one
			}
			if oldPoint.x < units {
				foldedMap[oldPoint] = oldPoint
				break
			}

			oldPoint.x = 2*units - oldPoint.x
			foldedMap[oldPoint] = oldPoint
		}
	}
	return
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
