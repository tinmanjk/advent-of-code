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
	result = findResult(mapOfPoints, instructions, false)
	fmt.Println(result)

	// // part 2
	// result = findResult(&graph, true)
	// fmt.Println(result)
}

type point struct {
	x int
	y int
}

type inst struct {
	axis  string
	units int
}

// list of points

func parseInput(slicesOfLines []string) (mapOfPoints map[point]point,
	instructions []inst) {

	// points
	mapOfPoints = map[point]point{}

	i := 0
	for ; i < len(slicesOfLines); i++ {
		line := slicesOfLines[i]
		if line == "" {
			i++
			break
		}
		splitted := strings.Split(line, ",")
		x, _ := strconv.Atoi(splitted[0])
		y, _ := strconv.Atoi(splitted[1])
		p := point{x, y}
		mapOfPoints[p] = p
	}
	// we have i

	//instructions
	for j := i; j < len(slicesOfLines); j++ {
		line := slicesOfLines[j]
		if line == "" {
			break
		}
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
	instructions []inst, justFirstInstr bool) (result int) {

	for i := 0; i < len(instructions); i++ {
		mapOfPoints = foldDirection(mapOfPoints, instructions[i].axis, instructions[i].units)

		if justFirstInstr {
			break
		}
	}
	printReadyPoints(mapOfPoints)
	return len(mapOfPoints)
}

func printReadyPoints(mapOfPoints map[point]point) {

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

	// y-x mi trqbva matrix
	//if point '#' otherwise '.'
	printMatrix := make([][]rune, maxY+1)
	for i := 0; i < len(printMatrix); i++ {
		printMatrix[i] = make([]rune, maxX+1)
		for j := 0; j < len(printMatrix[i]); j++ {
			pointAtCoordinate := point{j, i}
			if _, ok := mapOfPoints[pointAtCoordinate]; ok {
				printMatrix[i][j] = '#'
			} else {
				printMatrix[i][j] = '.'
			}
		}
		fmt.Println(string(printMatrix[i]))
	}
}

func foldDirection(mapOfPoints map[point]point, axis string, units int) (foldedMap map[point]point) {
	foldedMap = map[point]point{}
	for _, oldPoint := range mapOfPoints {
		switch axis {
		case "y":
			if oldPoint.y == units {
				continue // not to be added to new one
			}
			if oldPoint.y < units {
				foldedMap[oldPoint] = oldPoint
				continue
			}
			oldPoint.y = 2*units - oldPoint.y
			foldedMap[oldPoint] = oldPoint
		case "x":
			if oldPoint.x == units {
				continue // not to be added to new one
			}
			if oldPoint.x < units {
				foldedMap[oldPoint] = oldPoint
				continue
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
